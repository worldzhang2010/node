/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package wginterface

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/Microsoft/go-winio"
	"github.com/mysteriumnetwork/node/utils/netutil"
	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/ipc/winpipe"
	"golang.zx2c4.com/wireguard/tun"
)

// New creates new WgInterface instance.
func New(interfaceName string, uid string, subnet net.IPNet) (*WgInterface, error) {
	log.Println("Creating Wintun interface")

	reqGUID, err := windows.GenerateGUID()
	if err != nil {
		return nil, fmt.Errorf("could not generate win GUID: %w", err)
	}
	wintun, err := tun.CreateTUNWithRequestedGUID(interfaceName, &reqGUID, 0)
	if err != nil {
		return nil, fmt.Errorf("could not create wintun: %w", err)
	}
	nativeTun := wintun.(*tun.NativeTun)
	wintunVersion, ndisVersion, err := nativeTun.Version()
	if err != nil {
		log.Printf("Warning: unable to determine Wintun version: %v", err)
	} else {
		log.Printf("Using Wintun/%s (NDIS %s)", wintunVersion, ndisVersion)
	}

	tunDeviceName, err := wintun.Name()
	if err != nil {
		return nil, err
	}

	if tunDeviceName != interfaceName {
		if err := renameInterface(tunDeviceName, interfaceName); err != nil {
			return nil, fmt.Errorf("failed to rename network interface: %w", err)
		}
	}

	if err := netutil.AssignIP(interfaceName, subnet); err != nil {
		return nil, fmt.Errorf("could not assign IP: %w", err)
	}

	log.Println("Creating interface instance")
	logFile, err := os.OpenFile("myst_supervisor_wg.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return nil, fmt.Errorf("could not open wg log file: %w", err)
	}
	logger := NewLogger(device.LogLevelDebug, fmt.Sprintf("(%s) ", interfaceName), logFile)
	logger.Info.Println("Starting wireguard-go version", device.WireGuardGoVersion)

	wgDevice := device.NewDevice(wintun, logger)

	log.Println("Setting interface configuration")
	uapi, err := UAPIListen(interfaceName)
	if err != nil {
		return nil, fmt.Errorf("could not listen for user API wg configuration: %w", err)
	}

	wg := &WgInterface{
		Name:      interfaceName,
		device:    wgDevice,
		uapi:      uapi,
		logWriter: logFile,
	}
	go wg.handleUAPI()

	return wg, nil
}

// handleUAPI listens for WireGuard configuration changes via user space socket.
func (a *WgInterface) handleUAPI() {
	for {
		conn, err := a.uapi.Accept()
		if err != nil {
			log.Println("Closing UAPI listener, err:", err)
			return
		}
		go a.device.IpcHandle(conn)
	}
}

// Down closes device and user space api socket.
func (a *WgInterface) Down() {
	if err := a.uapi.Close(); err != nil {
		log.Printf("could not close uapi socket: %w", err)
	}
	a.device.Close()
	a.logWriter.Close()
}

func renameInterface(name, newname string) error {
	_, err := exec.Command("powershell", "-Command", "netsh interface set interface name=\""+name+"\" newname=\""+newname+"\"").CombinedOutput()
	return err
}

const (
	LogLevelSilent = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
)

func NewLogger(level int, prepend string, output io.Writer) *device.Logger {
	logger := new(device.Logger)

	logErr, logInfo, logDebug := func() (io.Writer, io.Writer, io.Writer) {
		if level >= LogLevelDebug {
			return output, output, output
		}
		if level >= LogLevelInfo {
			return output, output, ioutil.Discard
		}
		if level >= LogLevelError {
			return output, ioutil.Discard, ioutil.Discard
		}
		return ioutil.Discard, ioutil.Discard, ioutil.Discard
	}()

	logger.Debug = log.New(logDebug,
		"DEBUG: "+prepend,
		log.Ldate|log.Ltime,
	)

	logger.Info = log.New(logInfo,
		"INFO: "+prepend,
		log.Ldate|log.Ltime,
	)
	logger.Error = log.New(logErr,
		"ERROR: "+prepend,
		log.Ldate|log.Ltime,
	)
	return logger
}

func UAPIListen(name string) (net.Listener, error) {
	socketGroup := "Users"
	sddl := "D:P(A;;GA;;;BA)(A;;GA;;;SY)"
	sid, err := winio.LookupSidByName(socketGroup)
	if err != nil {
		return nil, err
	}
	sddl += fmt.Sprintf("(A;;GRGW;;;%s)", sid)
	securityDescriptor, err := windows.SecurityDescriptorFromString(sddl)
	if err != nil {
		return nil, err
	}

	config := winpipe.PipeConfig{
		SecurityDescriptor: securityDescriptor,
	}
	listener, err := winpipe.ListenPipe(`\\.\pipe\ProtectedPrefix\Administrators\WireGuard\`+name, &config)
	if err != nil {
		return nil, err
	}

	uapi := &UAPIListener{
		listener: listener,
		connNew:  make(chan net.Conn, 1),
		connErr:  make(chan error, 1),
	}

	go func(l *UAPIListener) {
		for {
			conn, err := l.listener.Accept()
			if err != nil {
				l.connErr <- err
				break
			}
			l.connNew <- conn
		}
	}(uapi)

	return uapi, nil
}

type UAPIListener struct {
	listener net.Listener // unix socket listener
	connNew  chan net.Conn
	connErr  chan error
	kqueueFd int
	keventFd int
}

func (l *UAPIListener) Accept() (net.Conn, error) {
	for {
		select {
		case conn := <-l.connNew:
			return conn, nil

		case err := <-l.connErr:
			return nil, err
		}
	}
}

func (l *UAPIListener) Close() error {
	return l.listener.Close()
}

func (l *UAPIListener) Addr() net.Addr {
	return l.listener.Addr()
}
