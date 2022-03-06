package trafficlight

import (
	"encoding/json"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

type TrafficLightFrameState uint

const (
	TrafficLightFrameState_begin TrafficLightFrameState = iota
	TrafficLightFrameState_red
	TrafficLightFrameState_green
	TrafficLightFrameState_yellow
	TrafficLightFrameState_flashingRed
	TrafficLightFrameState_end
	TrafficLightFrameState_working
)

type Marshal interface {
	Save() []byte
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
	//	ticker     *time.Ticker
}

type marshalStruct struct {
	TrafficLightState TrafficLightFrameState
	FlashColor        string
}

func New(mom *mOMStruct) (TrafficLight, error) {
	m := &trafficLightStruct{}
	m.mom = mom

	// Validate interfaces
	var _ TrafficLight = m
	var _ actions = m

	m.flashColor = ""
	return m, nil
}

func Load(mom *mOMStruct, data []byte) error {
	m := &trafficLightStruct{}
	m.mom = mom
	m.mom.trafficLight = m

	// Validate interfaces
	var _ TrafficLight = m
	var _ actions = m

	var marshal marshalStruct

	err := json.Unmarshal(data, &marshal)
	if err != nil {
		return err
	}
	m._state_ = marshal.TrafficLightState
	m.flashColor = marshal.FlashColor
	return nil
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
	case TrafficLightFrameState_begin:
		m._begin_(e)
	case TrafficLightFrameState_red:
		m._red_(e)
	case TrafficLightFrameState_green:
		m._green_(e)
	case TrafficLightFrameState_yellow:
		m._yellow_(e)
	case TrafficLightFrameState_flashingRed:
		m._flashingRed_(e)
	case TrafficLightFrameState_end:
		m._end_(e)
	case TrafficLightFrameState_working:
		m._working_(e)
	}
}

//===================== Machine Block ===================//

func (m *trafficLightStruct) _begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.startWorkingTimer()
		m._transition_(TrafficLightFrameState_red)
		return
	}
}

func (m *trafficLightStruct) _red_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterRed()
		return
	case "tick":
		m._transition_(TrafficLightFrameState_green)
		return
	}
	m._working_(e)

}

func (m *trafficLightStruct) _green_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterGreen()
		return
	case "tick":
		m._transition_(TrafficLightFrameState_yellow)
		return
	}
	m._working_(e)

}

func (m *trafficLightStruct) _yellow_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterYellow()
		return
	case "tick":
		m._transition_(TrafficLightFrameState_red)
		return
	}
	m._working_(e)

}

func (m *trafficLightStruct) _flashingRed_(e *framelang.FrameEvent) {
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
		m._transition_(TrafficLightFrameState_red)
		return
	case "stop":
		m._transition_(TrafficLightFrameState_end)
		return
	}
}

func (m *trafficLightStruct) _end_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.stopWorkingTimer()
		return
	}
}

func (m *trafficLightStruct) _working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "stop":
		m._transition_(TrafficLightFrameState_end)
		return
	case "systemError":
		m._transition_(TrafficLightFrameState_flashingRed)
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
