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

func (m *trafficLightMomStruct) enterRed()                { fmt.Println("enterRed()") }
func (m *trafficLightMomStruct) enterGreen()              { fmt.Println("enterGreen()") }
func (m *trafficLightMomStruct) enterYellow()             { fmt.Println("enterYellow()") }
func (m *trafficLightMomStruct) enterFlashingRed()        {}
func (m *trafficLightMomStruct) exitFlashingRed()         {}
func (m *trafficLightMomStruct) startWorkingTimer()       {}
func (m *trafficLightMomStruct) stopWorkingTimer()        {}
func (m *trafficLightMomStruct) startFlashingTimer()      {}
func (m *trafficLightMomStruct) stopFlashingTimer()       {}
func (m *trafficLightMomStruct) changeColor(color string) {}
func (m *trafficLightMomStruct) startFlashing()           {}
func (m *trafficLightMomStruct) stopFlashing()            {}
func (m *trafficLightMomStruct) changeFlashingAnimation() {}
func (m *trafficLightMomStruct) log(msg string)           {}
