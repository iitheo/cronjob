package main

import (
	"fmt"
	"time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR_TO_TICK int = 11
const MINUTE_TO_TICK int = 58
const SECOND_TO_TICK int = 03

type jobTicker struct {
	timer *time.Timer
}

func runningRoutine() {
	jobTicker := &jobTicker{}
	jobTicker.updateTimer()
	for {
		<-jobTicker.timer.C
		fmt.Println(time.Now(), "- just ticked")
		jobTicker.updateTimer()
	}
}

func (t *jobTicker) updateTimer() {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick, "- next tick")
	diff := nextTick.Sub(time.Now())
	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}

func main(){
	runningRoutine()
}
