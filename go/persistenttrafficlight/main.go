package main

import (
	"fmt"
	"time"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/trafficlight"
)

func main() {

	stop := make(chan bool)
	finished := make(chan bool)
	mom := trafficlight.NewMOM()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	ticker := time.NewTicker(1000 * time.Millisecond)
	mom.Start()
	go func() {
		for {
			select {
			case <-stop:
				ticker.Stop()
				mom.Stop()
				finished <- true
				return
			case <-ticker.C:
				mom.Tick()
				fmt.Println("tick")
			}

		}
	}()

	time.Sleep(5000 * time.Millisecond)
	stop <- true
	<-finished
	fmt.Println("finished")
}
