package DriveElevator

import (
	"time"
)

var timerFlag int
var t time.Time

func timer(chan_timer chan bool){
	for{
		if timerFlag == 1 {
			t = time.Now()
			timerFlag = 0
		}

		if int(time.Now().Second()) - int(t.Second()) > 3{
			chan_timer <- true
		}
	}
}

func timer_start() {
	timerFlag = 1
}