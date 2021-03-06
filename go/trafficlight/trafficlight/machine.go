package trafficlight

import (
	"time"

	"github.com/frame-lang/frame-demos/trafficlight/framelang"
)

const (
	begin framelang.FrameState = iota
	red
	green
	yellow
	flashingRed
	end
	working
)

type TrafficLight interface {
	Start(mom *MOM)
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
	_state_    framelang.FrameState
	flashColor string
	mom        *MOM
	ticker     *time.Ticker
}

func New() TrafficLight {
	m := new(trafficLightStruct)
	// Verify TrafficLightStruct implements actions interface
	var _ actions = m
	m.flashColor = ""
	m.mom = nil
	m.ticker = nil
	return m
}

//===================== Interface Block ===================//

func (m *trafficLightStruct) Start(mom *MOM) {
	params := make(map[string]interface{})
	params["mom"] = mom
	e := framelang.FrameEvent{Msg: ">>", Params: params}
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
	case end:
		m._end_(e)
	case working:
		m._working_(e)
	}
}

//===================== Machine Block ===================//

func (m *trafficLightStruct) _begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.startWorkingTimer()
		m._transition_(red)
		return
	}
}

func (m *trafficLightStruct) _red_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterRed()
		return
	case "tick":
		m._transition_(green)
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
		m._transition_(yellow)
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
		m._transition_(red)
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
		m._transition_(red)
		return
	case "stop":
		m._transition_(end)
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
		m._transition_(end)
		return
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
