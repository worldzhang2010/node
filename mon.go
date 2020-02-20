package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/prashantgupta24/mac-sleep-notifier/notifier"
)

type timeTracker struct {
	sync.Mutex
	startTime time.Time
	sleepTime time.Time
	awakeTime time.Time
}

func (t *timeTracker) start() {
	t.startTime = time.Now()

	for activity := range notifier.GetInstance().Start() {
		t.handleStateChange(activity)
	}
}

func (t *timeTracker) elapsed() time.Duration {
	now := time.Now()
	return now.Sub(t.startTime)
}

func (t *timeTracker) elapsedWall() time.Duration {
	now := time.Now()
	return subWall(now, t.startTime)
}

func subWall(a, b time.Time) time.Duration  {
	return time.Duration(a.Unix()-b.Unix())*time.Second
}

func (t *timeTracker) handleStateChange(activity *notifier.Activity) {
	t.Lock()
	defer t.Unlock()

	if activity.Type == notifier.Awake {
		log.Println("machine awake")
		t.awakeTime = time.Now()
	} else if activity.Type == notifier.Sleep {
		log.Println("machine sleeping")
		t.sleepTime = time.Now()
	}

	if !t.awakeTime.IsZero() && !t.sleepTime.IsZero() {
		log.Println("sleep duration", subWall(t.awakeTime, t.sleepTime))
	}

	t.awakeTime.Nanosecond()
}

func main() {
	var tracker timeTracker
	go tracker.start()

	for {
		time.Sleep(3 * time.Second)
		fmt.Println(time.Now(), tracker.elapsedWall(), tracker.elapsed())
	}
}
