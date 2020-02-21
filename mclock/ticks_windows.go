package mclock

import "time"

func ticksNano() uint64 {
	return uint64(time.Now().UnixNano())
}
