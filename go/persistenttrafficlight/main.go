package main

import (
	"fmt"
	"time"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/trafficlight"
)

func main() {

	stop := make(chan bool)
	finished := make(chan bool)
	// if err != nil {
	// 	log.Fatal(err)	// }
	ticker := time.NewTicker(1000 * time.Millisecond)
	mom := trafficlight.NewTrafficLightMom()
	go func() {
		for {
			select {
			case <-stop:
				ticker.Stop()
				mom.Stop()
				finished <- true
				return
			case <-ticker.C:
				fmt.Println("tick")
				mom.Tick()
			}

		}
	}()

	time.Sleep(5000 * time.Millisecond)
	stop <- true
	<-finished
	fmt.Println("finished")
}
