package mclock

//    #include <time.h>
//    #include <sys/time.h>
//    #include <stdio.h>
//    #include <stdint.h>
//
//    static uint64_t ticks_nano() {
//        struct timespec ts;
//        clock_gettime(CLOCK_BOOTTIME, &ts);
//        return (uint64_t)(ts.tv_sec * 1e9 + ts.tv_nsec);
//    }
import "C"

func ticksNano() uint64 {
	return uint64(C.ticks_nano())
}
