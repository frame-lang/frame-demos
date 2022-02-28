package main

import (
	"time"

	"github.com/frame-lang/frame-demos/go/web/traffic/trafficlight"
)

func main() {

	mom := trafficlight.NewMOM()
	mom.Start()

	time.Sleep(600000 * time.Millisecond)
	mom.Stop()
}
