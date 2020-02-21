/*
author: jbenet
os x, compile with: gcc -o testo test.c
linux, compile with: gcc -o testo test.c -lrt
*/

//    #include <time.h>
//    #include <sys/time.h>
//    #include <stdio.h>
//    #include <stdint.h>
//
//    void current_utc_time(struct timespec *ts) {
//        clock_gettime(CLOCK_BOOTTIME, ts);
//    }
//
//    static uint64_t ticks_nano() {
//        struct timespec ts;
//        clock_gettime(CLOCK_BOOTTIME, &ts);
//        return (uint64_t)(ts.tv_sec * 1e9 + ts.tv_nsec);
//    }

int main(int argc, char **argv) {

  uint64_t start = ticks_nano();
  sleep(2);
  uint64_t stop = ticks_nano();

  double result = stop - start;
  printf("Thread slept for: %f\n", result);
  return 0;
}
