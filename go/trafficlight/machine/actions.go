package machine

import (
	"fmt"
	"time"
)

func (m *trafficLightStruct) enterRed() {
	fmt.Println("enterRed()")
}
func (m *trafficLightStruct) enterGreen() {
	fmt.Println("enterGreen()")
}
func (m *trafficLightStruct) enterYellow() {
	fmt.Println("enterYellow()")
}
func (m *trafficLightStruct) enterFlashingRed() {}
func (m *trafficLightStruct) exitFlashingRed()  {}
func (m *trafficLightStruct) startWorkingTimer() {
	m.ticker = time.NewTicker(1000 * time.Millisecond)

	go func() {
		for range m.ticker.C {
			m.Timer()
		}
	}()

}
func (m *trafficLightStruct) stopWorkingTimer() {
	m.ticker.Stop()
}
func (m *trafficLightStruct) startFlashingTimer()      {}
func (m *trafficLightStruct) stopFlashingTimer()       {}
func (m *trafficLightStruct) changeColor(color string) {}
func (m *trafficLightStruct) startFlashing()           {}
func (m *trafficLightStruct) stopFlashing()            {}
func (m *trafficLightStruct) changeFlashingAnimation() {}
func (m *trafficLightStruct) log(msg string)           {}
