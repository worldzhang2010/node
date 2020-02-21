// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// System calls and other sys.stuff for AMD64, Darwin
// System calls are implemented in libSystem, this file contains
// trampolines that convert from Go to C calling convention.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

GLOBL timebase<>(SB),NOPTR,$(machTimebaseInfo__size)

TEXT runtime·nanotime_trampoline(SB),NOSPLIT,$0
	PUSHQ	BP
	MOVQ	SP, BP
	MOVQ	DI, BX
	CALL	libc_mach_absolute_time(SB)
	MOVQ	AX, 0(BX)
	MOVL	timebase<>+machTimebaseInfo_numer(SB), SI
	MOVL	timebase<>+machTimebaseInfo_denom(SB), DI // atomic read
	TESTL	DI, DI
	JNE	initialized

	SUBQ	$(machTimebaseInfo__size+15)/16*16, SP
	MOVQ	SP, DI
	CALL	libc_mach_timebase_info(SB)
	MOVL	machTimebaseInfo_numer(SP), SI
	MOVL	machTimebaseInfo_denom(SP), DI
	ADDQ	$(machTimebaseInfo__size+15)/16*16, SP

	MOVL	SI, timebase<>+machTimebaseInfo_numer(SB)
	MOVL	DI, AX
	XCHGL	AX, timebase<>+machTimebaseInfo_denom(SB) // atomic write

initialized:
	MOVL	SI, 8(BX)
	MOVL	DI, 12(BX)
	MOVQ	BP, SP
	POPQ	BP
	RET


ok:
	XORL	AX, AX        // no error (it's ignored anyway)
	MOVQ	BP, SP
	POPQ	BP
	RET
