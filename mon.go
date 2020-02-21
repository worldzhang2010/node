package main

import (
	"fmt"

	// "unsafe"
	// "github.com/mysteriumnetwork/node/mclock"
	t "time"

	"github.com/mysteriumnetwork/node/time"
)

func main() {
	// cs := C.CString("Hello from stdio")
	//for {
	//	t1 := mclock.Now()
	//	time.Sleep(1 * time.Second)
	//	t2 := mclock.Now()
	//
	//	fmt.Printf("%d - %d = %v\n", t2, t1, t2.Sub(t1))
	//}
	t1 := time.New()
	t.Sleep(2 * t.Second)
	t2 := time.New()
	fmt.Println(t2 - t1)
	// fmt.Println(time.Now().Nanosecond())
	// C.free(unsafe.Pointer(cs))
}
