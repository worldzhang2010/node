package main

import (
	"fmt"
	"time"

	// "unsafe"

	"github.com/mysteriumnetwork/node/mclock"
)

// #include <t.h>
import "C"

func main() {
	// cs := C.CString("Hello from stdio")
	for {
		t1 := mclock.Now()
		time.Sleep(1 * time.Second)
		t2 := mclock.Now()

		fmt.Printf("%d - %d = %v\n", t2, t1, t2.Sub(t1))
	}
	// fmt.Println(time.Now().Nanosecond())
	// C.free(unsafe.Pointer(cs))
}
