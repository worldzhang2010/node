package mclock

//    #include <time.h>
//    #include <sys/time.h>
//    #include <stdio.h>
//    #include <mach/clock.h>
//    #include <mach/mach.h>
//    #include <mach/mach_time.h>
//
//    static uint64_t ticks_nano() {
//        uint64_t t = mach_continuous_time();
//        static mach_timebase_info_data_t timebase;
//        mach_timebase_info(&timebase);
//        double clockTicksToNanosecons = (double)timebase.numer / timebase.denom;
//        return (uint64_t)(t * clockTicksToNanosecons);
//    }
import "C"

func ticksNano() uint64 {
	return uint64(C.ticks_nano())
}
