package main

import (
	"time"

	trafficlight "github.com/frame-lang/frame-demos/trafficlight/trafficlight"
)

func main() {

	mom := trafficlight.NewMOM()
	mom.Start()

	time.Sleep(5000 * time.Millisecond)
	mom.Stop()
}
