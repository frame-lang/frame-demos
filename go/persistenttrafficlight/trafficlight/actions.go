package trafficlight

import (
	"fmt"
)

func (m *trafficLightStruct) enterRed() {
	m.mom.EnterRed()
}
func (m *trafficLightStruct) enterGreen() {
	m.mom.EnterGreen()
}
func (m *trafficLightStruct) enterYellow() {
	m.mom.EnterYellow()
}
func (m *trafficLightStruct) enterFlashingRed() {}
func (m *trafficLightStruct) exitFlashingRed()  {}
func (m *trafficLightStruct) startWorkingTimer() {

}
func (m *trafficLightStruct) stopWorkingTimer() {
}
func (m *trafficLightStruct) startFlashingTimer()      {}
func (m *trafficLightStruct) stopFlashingTimer()       {}
func (m *trafficLightStruct) changeColor(color string) {}
func (m *trafficLightStruct) startFlashing()           {}
func (m *trafficLightStruct) stopFlashing()            {}
func (m *trafficLightStruct) changeFlashingAnimation() {}
func (m *trafficLightStruct) log(msg string)           {}

func (m *mOMStruct) enterRed()                { fmt.Println("enterRed()") }
func (m *mOMStruct) enterGreen()              { fmt.Println("enterGreen()") }
func (m *mOMStruct) enterYellow()             { fmt.Println("enterYellow()") }
func (m *mOMStruct) enterFlashingRed()        {}
func (m *mOMStruct) exitFlashingRed()         {}
func (m *mOMStruct) startWorkingTimer()       {}
func (m *mOMStruct) stopWorkingTimer()        {}
func (m *mOMStruct) startFlashingTimer()      {}
func (m *mOMStruct) stopFlashingTimer()       {}
func (m *mOMStruct) changeColor(color string) {}
func (m *mOMStruct) startFlashing()           {}
func (m *mOMStruct) stopFlashing()            {}
func (m *mOMStruct) changeFlashingAnimation() {}
func (m *mOMStruct) log(msg string)           {}
