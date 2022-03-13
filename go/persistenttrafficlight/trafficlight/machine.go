package trafficlight

import (
	"encoding/json"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

type TrafficLightFrameState uint

const (
	TrafficLightFrameState_Begin TrafficLightFrameState = iota
	TrafficLightFrameState_Red
	TrafficLightFrameState_Green
	TrafficLightFrameState_Yellow
	TrafficLightFrameState_FlashingRed
	TrafficLightFrameState_End
	TrafficLightFrameState_Working
)

type Marshal interface {
	Marshal() []byte
}

type TrafficLight interface {
	Marshal
	Start()
	Stop()
	Tick()
	SystemError()
	SystemRestart()
}

type actions interface {
	enterRed()
	enterGreen()
	enterYellow()
	enterFlashingRed()
	exitFlashingRed()
	startWorkingTimer()
	stopWorkingTimer()
	startFlashingTimer()
	stopFlashingTimer()
	changeColor(color string)
	startFlashing()
	stopFlashing()
	changeFlashingAnimation()
	log(msg string)
}

type trafficLightStruct struct {
	mom        *mOMStruct
	_state_    TrafficLightFrameState
	flashColor string
}

type marshalStruct struct {
	TrafficLightState TrafficLightFrameState
	FlashColor        string
}

func NewTrafficLight(mom *mOMStruct) TrafficLight {
	m := &trafficLightStruct{}
	m.mom = mom

	// Validate interfaces
	var _ TrafficLight = m
	var _ actions = m

	// Initialize domain
	m.flashColor = ""

	return m
}

func LoadTrafficLight(mom *mOMStruct, data []byte) TrafficLight {
	m := &trafficLightStruct{}
	m.mom = mom

	// Validate interfaces
	var _ TrafficLight = m
	var _ actions = m

	// Unmarshal
	var marshal marshalStruct
	err := json.Unmarshal(data, &marshal)
	if err != nil {
		return nil
	}

	// Initialize machine
	m._state_ = marshal.TrafficLightState
	m.flashColor = marshal.FlashColor

	return m

}

func (m *trafficLightStruct) MarshalJSON() ([]byte, error) {
	data := marshalStruct{
		TrafficLightState: m._state_,
		FlashColor:        m.flashColor,
	}
	return json.Marshal(data)
}

func (m *trafficLightStruct) Marshal() []byte {
	data, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return data

}

//===================== Interface Block ===================//

func (m *trafficLightStruct) Start() {
	e := framelang.FrameEvent{Msg: ">>"}
	m._mux_(&e)
}

func (m *trafficLightStruct) Stop() {
	e := framelang.FrameEvent{Msg: "stop"}
	m._mux_(&e)
}

func (m *trafficLightStruct) Tick() {
	e := framelang.FrameEvent{Msg: "tick"}
	m._mux_(&e)
}

func (m *trafficLightStruct) SystemError() {
	e := framelang.FrameEvent{Msg: "systemError"}
	m._mux_(&e)
}

func (m *trafficLightStruct) SystemRestart() {
	e := framelang.FrameEvent{Msg: "systemRestart"}
	m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *trafficLightStruct) _mux_(e *framelang.FrameEvent) {
	switch m._state_ {
	case TrafficLightFrameState_Begin:
		m._TrafficLightFrameState_Begin_(e)
	case TrafficLightFrameState_Red:
		m._TrafficLightFrameState_Red_(e)
	case TrafficLightFrameState_Green:
		m._TrafficLightFrameState_Green_(e)
	case TrafficLightFrameState_Yellow:
		m._TrafficLightFrameState_Yellow_(e)
	case TrafficLightFrameState_FlashingRed:
		m._TrafficLightFrameState_FlashingRed_(e)
	case TrafficLightFrameState_End:
		m._TrafficLightFrameState_End_(e)
	case TrafficLightFrameState_Working:
		m._TrafficLightFrameState_Working_(e)
	}
}

//===================== Machine Block ===================//

func (m *trafficLightStruct) _TrafficLightFrameState_Begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.startWorkingTimer()
		m._transition_(TrafficLightFrameState_Red)
		return
	}
}

func (m *trafficLightStruct) _TrafficLightFrameState_Red_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterRed()
		return
	case "tick":
		m._transition_(TrafficLightFrameState_Green)
		return
	}
	m._TrafficLightFrameState_Working_(e)

}

func (m *trafficLightStruct) _TrafficLightFrameState_Green_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterGreen()
		return
	case "tick":
		m._transition_(TrafficLightFrameState_Yellow)
		return
	}
	m._TrafficLightFrameState_Working_(e)

}

func (m *trafficLightStruct) _TrafficLightFrameState_Yellow_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterYellow()
		return
	case "tick":
		m._transition_(TrafficLightFrameState_Red)
		return
	}
	m._TrafficLightFrameState_Working_(e)

}

func (m *trafficLightStruct) _TrafficLightFrameState_FlashingRed_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterFlashingRed()
		m.stopWorkingTimer()
		m.startFlashingTimer()
		return
	case "<":
		m.exitFlashingRed()
		m.stopFlashingTimer()
		m.startWorkingTimer()
		return
	case "tick":
		m.changeFlashingAnimation()
		return
	case "systemRestart":
		m._transition_(TrafficLightFrameState_Red)
		return
	case "stop":
		m._transition_(TrafficLightFrameState_End)
		return
	}
}

func (m *trafficLightStruct) _TrafficLightFrameState_End_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.stopWorkingTimer()
		return
	}
}

func (m *trafficLightStruct) _TrafficLightFrameState_Working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "stop":
		m._transition_(TrafficLightFrameState_End)
		return
	case "systemError":
		m._transition_(TrafficLightFrameState_FlashingRed)
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *trafficLightStruct) _transition_(newState TrafficLightFrameState) {
	m._mux_(&framelang.FrameEvent{Msg: "<"})
	m._state_ = newState
	m._mux_(&framelang.FrameEvent{Msg: ">"})
}

/********************
// Sample Actions Implementation
package trafficlight

func (m *trafficLightStruct) enterRed()  {}
func (m *trafficLightStruct) enterGreen()  {}
func (m *trafficLightStruct) enterYellow()  {}
func (m *trafficLightStruct) enterFlashingRed()  {}
func (m *trafficLightStruct) exitFlashingRed()  {}
func (m *trafficLightStruct) startWorkingTimer()  {}
func (m *trafficLightStruct) stopWorkingTimer()  {}
func (m *trafficLightStruct) startFlashingTimer()  {}
func (m *trafficLightStruct) stopFlashingTimer()  {}
func (m *trafficLightStruct) changeColor(color string)  {}
func (m *trafficLightStruct) startFlashing()  {}
func (m *trafficLightStruct) stopFlashing()  {}
func (m *trafficLightStruct) changeFlashingAnimation()  {}
func (m *trafficLightStruct) log(msg string)  {}
********************/
