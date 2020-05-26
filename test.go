package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("supervisor.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	for {
		log.Println("tick")
		time.Sleep(time.Second)
	}
}
