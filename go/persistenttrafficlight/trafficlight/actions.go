package trafficlight

import (
	"fmt"
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

// func (m *trafficLightStruct) UnmarshalJSON(data []byte) error {
// 	unmarshalleddata := struct {
// 		FrameState framelang.FrameState
// 		FlashColor string
// 	}{}

// 	err := json.Unmarshal(data, unmarshalleddata)
// 	if err != nil {
// 		return err
// 	}

// 	m._state_ = unmarshalleddata.FrameState
// 	m.flashColor = unmarshalleddata.FlashColor

// 	return nil
// }
