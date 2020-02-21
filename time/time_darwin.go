package time

import "unsafe"

//go:noescape
//go:linkname libcCall runtime.libcCall
func libcCall(fn, arg unsafe.Pointer) int32

//go:noescape
//go:linkname funcPC runtime.funcPC
func funcPC(f interface{}) uintptr

type machTimebaseInfo struct {
	numer uint32
	denom uint32
}

//go:nosplit
//go:go:nocheckptr
func nanotime1() int64 {
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

func New() int64 {
	return nanotime1()
}

//go:cgo_import_dynamic libc_mach_timebase_info mach_timebase_info "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_mach_continuous_time mach_continuous_time "/usr/lib/libSystem.B.dylib"
