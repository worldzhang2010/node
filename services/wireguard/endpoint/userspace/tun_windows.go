// +build windows

/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
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

package userspace

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/mysteriumnetwork/node/utils/netutil"
	"github.com/pkg/errors"

	// "github.com/songgao/water"
	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun"
)

type nativeTun struct {
	//tun    *water.Interface
	tun    tun.Device
	events chan tun.Event
}

// CreateTUN creates native TUN device for wireguard.
func CreateTUN(name string, subnet net.IPNet) (tun.Device, error) {
	//tunDevice, err := water.New(water.Config{
	//	DeviceType: water.TUN,
	//	PlatformSpecificParams: water.PlatformSpecificParams{
	//		ComponentID: "tap0901",
	//		Network:     subnet.String(),
	//	},
	//})
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to create new TUN device")
	//}

	reqGUID, err := windows.GenerateGUID()
	if err != nil {
		return nil, fmt.Errorf("could not generate win GUID: %w", err)
	}
	tunDevice, err := tun.CreateTUNWithRequestedGUID(name, &reqGUID, 0)
	if err != nil {
		return nil, fmt.Errorf("could not create wintun: %w", err)
	}
	nativeTun_ := tunDevice.(*tun.NativeTun)
	wintunVersion, ndisVersion, err := nativeTun_.Version()
	if err != nil {
		log.Printf("Warning: unable to determine Wintun version: %v", err)
	} else {
		log.Printf("Using Wintun/%s (NDIS %s)", wintunVersion, ndisVersion)
	}

	deviceName, err := tunDevice.Name()
	if err != nil {
		return nil, err
	}

	if err := netutil.AssignIP(deviceName, subnet); err != nil {
		return nil, errors.Wrap(err, "failed to assign IP address")
	}

	if deviceName != name {
		if err := renameInterface(deviceName, name); err != nil {
			return nil, errors.Wrap(err, "failed to rename network interface")
		}
	}

	return &nativeTun{
		tun:    tunDevice,
		events: make(chan tun.Event, 10),
	}, nil
}

func (tun *nativeTun) Name() (string, error) {
	return tun.tun.Name()
}

func (tun *nativeTun) File() *os.File {
	return nil
}

func (tun *nativeTun) Events() chan tun.Event {
	return tun.events
}

func (tun *nativeTun) Read(buff []byte, offset int) (int, error) {
	return tun.tun.Read(buff, offset)
}

func (tun *nativeTun) Write(buff []byte, offset int) (int, error) {
	return tun.tun.Write(buff, offset)
}

func (tun *nativeTun) Close() error {
	close(tun.events)
	return tun.tun.Close()
}

func (tun *nativeTun) Flush() error {
	return nil
}

func (tun *nativeTun) MTU() (int, error) {
	return device.DefaultMTU, nil
}

func renameInterface(name, newname string) error {
	out, err := exec.Command("powershell", "-Command", "netsh interface set interface name=\""+name+"\" newname=\""+newname+"\"").CombinedOutput()
	return errors.Wrap(err, string(out))
}

func destroyDevice(name string) error {
	// Windows implementation is using single device that are reused for the future needs.
	// Nothing to destroy here.
	return nil
}
