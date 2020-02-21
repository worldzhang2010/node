#include <stdio.h>
#include <inttypes.h>
#include <mach/mach_time.h>

static uint64_t get_time() {
    return mach_absolute_time();
}

int main ( ) {
    uint64_t clockTicksSinceSystemBoot = mach_continuous_time();
    printf("Clock ticks since system boot: %"PRIu64"\n",
        clockTicksSinceSystemBoot
    );

    static mach_timebase_info_data_t timebase;
    mach_timebase_info(&timebase);
    // Cast to double is required to make this a floating point devision,
    // otherwise it would be an interger division and only the result would
    // be converted to floating point!
    double clockTicksToNanosecons = (double)timebase.numer / timebase.denom;
    printf("Ticks in nano: %f\n", clockTicksToNanosecons);

//    uint64_t systemUptimeNanoseconds = (uint64_t)(
//        clockTicksToNanosecons * clockTicksSinceSystemBoot
//    );
    // uint64_t systemUptimeSeconds = systemUptimeNanoseconds / (1000 * 1000 * 1000);
    // printf("System uptime: %"PRIu64" seconds\n", systemUptimeSeconds);

    uint64_t machTimeBegin = mach_absolute_time();
    sleep(2);
    uint64_t machTimeEnd = mach_absolute_time();
    uint64_t machTimePassed = machTimeEnd - machTimeBegin;
    uint64_t timePassedNS = (uint64_t)(
       machTimePassed * clockTicksToNanosecons
    );
    printf("Thread slept for: %"PRIu64" ns\n", timePassedNS);
    printf("Thread slept for: %"PRIu64" ns\n", machTimePassed);
}
