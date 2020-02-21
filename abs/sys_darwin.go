// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abs

import "unsafe"

// Call fn with arg as its argument. Return what fn returns.
// fn is the raw pc value of the entry point of the desired function.
// Switches to the system stack, if not already there.
// Preserves the calling point as the location where a profiler traceback will begin.
//func libcCall(fn, arg unsafe.Pointer) int32 {
//	// Leave caller's PC/SP/G around for traceback.
//	gp := getg()
//	var mp *m
//	if gp != nil {
//		mp = gp.m
//	}
//	if mp != nil && mp.libcallsp == 0 {
//		mp.libcallg.set(gp)
//		mp.libcallpc = getcallerpc()
//		// sp must be the last, because once async cpu profiler finds
//		// all three values to be non-zero, it will use them
//		mp.libcallsp = getcallersp()
//	} else {
//		// Make sure we don't reset libcallsp. This makes
//		// libcCall reentrant; We remember the g/pc/sp for the
//		// first call on an M, until that libcCall instance
//		// returns.  Reentrance only matters for signals, as
//		// libc never calls back into Go.  The tricky case is
//		// where we call libcX from an M and record g/pc/sp.
//		// Before that call returns, a signal arrives on the
//		// same M and the signal handling code calls another
//		// libc function.  We don't want that second libcCall
//		// from within the handler to be recorded, and we
//		// don't want that call's completion to zero
//		// libcallsp.
//		// We don't need to set libcall* while we're in a sighandler
//		// (even if we're not currently in libc) because we block all
//		// signals while we're handling a signal. That includes the
//		// profile signal, which is the one that uses the libcall* info.
//		mp = nil
//	}
//	res := asmcgocall(fn, arg)
//	if mp != nil {
//		mp.libcallsp = 0
//	}
//	return res
//}

//go:noescape
//go:linkname libcCall runtime.libcCall
func libcCall(fn, arg unsafe.Pointer) int32

//go:noescape
//go:linkname funcPC runtime.funcPC
func funcPC(f interface{}) uintptr

//go:nosplit
//go:cgo_unsafe_args
func Nanotime1() int64 {
	var r struct {
		t            int64  // raw timer
		numer, denom uint32 // conversion factors. nanoseconds = t * numer / denom.
	}
	libcCall(unsafe.Pointer(funcPC(nanotime_trampoline)), unsafe.Pointer(&r))
	// Note: Apple seems unconcerned about overflow here. See
	// https://developer.apple.com/library/content/qa/qa1398/_index.html
	// Note also, numer == denom == 1 is common.
	t := r.t
	if r.numer != 1 {
		t *= int64(r.numer)
	}
	if r.denom != 1 {
		t /= int64(r.denom)
	}
	return t
}
func nanotime_trampoline()

// Tell the linker that the libc_* functions are to be found
// in a system library, with the libc_ prefix missing.

//go:cgo_import_dynamic libc_mach_timebase_info mach_timebase_info "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_mach_absolute_time mach_absolute_time "/usr/lib/libSystem.B.dylib"

// Magic incantation to get libSystem actually dynamically linked.
// TODO: Why does the code require this?  See cmd/link/internal/ld/go.go
//go:cgo_import_dynamic _ _ "/usr/lib/libSystem.B.dylib"
