package trafficlight

import (
	"encoding/json"
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

func (m *trafficLightStruct) Save() []byte {
	data, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return data
}

func (m *trafficLightStruct) MarshalJSON() ([]byte, error) {
	data := marshalStruct{
		FrameState: m._state_,
		FlashColor: m.flashColor,
	}
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (m *trafficLightStruct) load(data []byte) error {
	var marshal marshalStruct

	err := json.Unmarshal(data, &marshal)
	if err != nil {
		return err
	}
	m._state_ = marshal.FrameState
	m.flashColor = marshal.FlashColor
	return nil
}

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
