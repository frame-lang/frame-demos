package main

import (
	"fmt"

	"github.com/frame-lang/frame-demos/go/web/traffic/trafficlight"
)

func main() {

	m := trafficlight.New()
	m.Start()
	fmt.Println("Hello")
}
