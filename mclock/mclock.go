package mclock

import "time"

// Now returns new monotonic clock which holds current value of a clock that
// increments monotonically in tick units (starting at an arbitrary point)
// including while the system is asleep.
func Now() Clock {
	ticks := ticksNano()
	return Clock{ticks: ticks}
}

// A Clock represents an instant in time witch holds tick units represented
// in nano seconds. It should be used only for elapsed time measurements
// when it is important to have monotonic time which also includes time while
// the system is a asleep.
//
// Depending on platform it uses C bindings to call specific OS APIs and
// fetches underlying clock ticks.
type Clock struct {
	ticks uint64
}

func (t Clock) Sub(u Clock) time.Duration {
	return time.Duration(t.ticks - u.ticks)
}
