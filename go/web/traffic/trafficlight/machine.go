package trafficlight

import "github.com/frame-demos/go/web/traffic/framelang"

const (
	begin framelang.FrameState = iota
	red
	green
	yellow
	flashingRed
	working
)

type TrafficLight interface {
	Start()
	Timer()
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
	_state_    framelang.FrameState
	actions    actions
	flashColor string
}

func New() TrafficLight {
	m := new(trafficLightStruct)
	m.actions = &trafficLightActions{}
	m.flashColor = ""
	return m
}

//===================== Interface Block ===================//

func (m *trafficLightStruct) Start() {
	e := framelang.FrameEvent{Msg: ">>"}
	m._mux_(&e)
}

func (m *trafficLightStruct) Timer() {
	e := framelang.FrameEvent{Msg: "Timer"}
	m._mux_(&e)
}

func (m *trafficLightStruct) SystemError() {
	e := framelang.FrameEvent{Msg: "SystemError"}
	m._mux_(&e)
}

func (m *trafficLightStruct) SystemRestart() {
	e := framelang.FrameEvent{Msg: "SystemRestart"}
	m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *trafficLightStruct) _mux_(e *framelang.FrameEvent) {
	switch m._state_ {
	case begin:
		m._begin_(e)
	case red:
		m._red_(e)
	case green:
		m._green_(e)
	case yellow:
		m._yellow_(e)
	case flashingRed:
		m._flashingRed_(e)
	case working:
		m._working_(e)
	}
}

//===================== Machine Block ===================//

func (m *trafficLightStruct) _begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.actions.startWorkingTimer()
		m._transition_(red)
		return
	}
}

func (m *trafficLightStruct) _red_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.actions.enterRed()
		return
	case "timer":
		m._transition_(green)
		return
	}
	m._working_(e)

}

func (m *trafficLightStruct) _green_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.actions.enterGreen()
		return
	case "timer":
		m._transition_(yellow)
		return
	}
	m._working_(e)

}

func (m *trafficLightStruct) _yellow_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.actions.enterYellow()
		return
	case "timer":
		m._transition_(red)
		return
	}
	m._working_(e)

}

func (m *trafficLightStruct) _flashingRed_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.actions.enterFlashingRed()
		m.actions.stopWorkingTimer()
		m.actions.startFlashingTimer()
		return
	case "<":
		m.actions.exitFlashingRed()
		m.actions.stopFlashingTimer()
		m.actions.startWorkingTimer()
		return
	case "timer":
		m.actions.changeFlashingAnimation()
		return
	case "systemRestart":
		m._transition_(red)
		return
	}
}

func (m *trafficLightStruct) _working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "systemError":
		m._transition_(flashingRed)
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *trafficLightStruct) _transition_(newState framelang.FrameState) {
	m._mux_(&framelang.FrameEvent{Msg: "<"})
	m._state_ = newState
	m._mux_(&framelang.FrameEvent{Msg: ">"})
}

/********************
// Sample Actions Implementation
package TrafficLight

type trafficLightActions struct{}

func (m *trafficLightActions) enterRed()  {}
func (m *trafficLightActions) enterGreen()  {}
func (m *trafficLightActions) enterYellow()  {}
func (m *trafficLightActions) enterFlashingRed()  {}
func (m *trafficLightActions) exitFlashingRed()  {}
func (m *trafficLightActions) startWorkingTimer()  {}
func (m *trafficLightActions) stopWorkingTimer()  {}
func (m *trafficLightActions) startFlashingTimer()  {}
func (m *trafficLightActions) stopFlashingTimer()  {}
func (m *trafficLightActions) changeColor(color string)  {}
func (m *trafficLightActions) startFlashing()  {}
func (m *trafficLightActions) stopFlashing()  {}
func (m *trafficLightActions) changeFlashingAnimation()  {}
func (m *trafficLightActions) log(msg string)  {}
********************/
