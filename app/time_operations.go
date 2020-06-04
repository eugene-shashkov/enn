package app

import "time"

func OneSecondPause() {
	time.Sleep(time.Second * 1)
}

func FiveMillisecondsPause() {
	time.Sleep(time.Millisecond * 5)
}

func HalfOfMillisecondPause() {
	time.Sleep(time.Microsecond * 500)
}

func NanosecondPause(x int64) {
	for i := int64(0); i < x; i++ {
		time.Sleep(time.Microsecond)
	}
}
