package timer

var timerFlag int
var t time.Time

func timer(chan_timer chan bool){
	for{
		if timerFlag = 1 {
			t = time.now()
			timerFlag = 0
		}

		if time.now() - t > 3{
			chan_timer <- true
		}
	}
}

func start() {
	timerFlag = 1
}