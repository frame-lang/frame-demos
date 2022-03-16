package trafficlight

import (
	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

type TrafficLightMomState uint

const (
	TrafficLightMomState_New TrafficLightMomState = iota
	TrafficLightMomState_Saving
	TrafficLightMomState_Persisted
	TrafficLightMomState_Working
	TrafficLightMomState_TrafficLightApi
	TrafficLightMomState_End
)

type TrafficLightMom interface {
	Start()
	Stop()
	Tick()
	EnterRed()
	EnterGreen()
	EnterYellow()
	EnterFlashingRed()
	ExitFlashingRed()
	StartWorkingTimer()
	StopWorkingTimer()
	StartFlashingTimer()
	StopFlashingTimer()
	ChangeColor(color string)
	StartFlashing()
	StopFlashing()
	ChangeFlashingAnimation()
	Log(msg string)
}

type TrafficLightMom_actions interface {
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

type trafficLightMomStruct struct {
	_state_      TrafficLightMomState
	trafficLight TrafficLight
	data         []byte
}

func NewTrafficLightMom() TrafficLightMom {
	m := &trafficLightMomStruct{}

	// Validate interfaces
	var _ TrafficLightMom = m
	var _ TrafficLightMom_actions = m

	// Initialize domain
	m.trafficLight = nil
	m.data = nil

	return m
}

//===================== Interface Block ===================//

func (m *trafficLightMomStruct) Start() {
	e := framelang.FrameEvent{Msg: ">>"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) Stop() {
	e := framelang.FrameEvent{Msg: "<<"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) Tick() {
	e := framelang.FrameEvent{Msg: "tick"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterRed() {
	e := framelang.FrameEvent{Msg: "enterRed"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterGreen() {
	e := framelang.FrameEvent{Msg: "enterGreen"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterYellow() {
	e := framelang.FrameEvent{Msg: "enterYellow"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterFlashingRed() {
	e := framelang.FrameEvent{Msg: "enterFlashingRed"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) ExitFlashingRed() {
	e := framelang.FrameEvent{Msg: "exitFlashingRed"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) StartWorkingTimer() {
	e := framelang.FrameEvent{Msg: "startWorkingTimer"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) StopWorkingTimer() {
	e := framelang.FrameEvent{Msg: "stopWorkingTimer"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) StartFlashingTimer() {
	e := framelang.FrameEvent{Msg: "startFlashingTimer"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) StopFlashingTimer() {
	e := framelang.FrameEvent{Msg: "stopFlashingTimer"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) ChangeColor(color string) {
	params := make(map[string]interface{})
	params["color"] = color
	e := framelang.FrameEvent{Msg: "changeColor", Params: params}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) StartFlashing() {
	e := framelang.FrameEvent{Msg: "startFlashing"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) StopFlashing() {
	e := framelang.FrameEvent{Msg: "stopFlashing"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) ChangeFlashingAnimation() {
	e := framelang.FrameEvent{Msg: "changeFlashingAnimation"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) Log(msg string) {
	params := make(map[string]interface{})
	params["msg"] = msg
	e := framelang.FrameEvent{Msg: "log", Params: params}
	m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *trafficLightMomStruct) _mux_(e *framelang.FrameEvent) {
	switch m._state_ {
	case TrafficLightMomState_New:
		m._TrafficLightMomState_New_(e)
	case TrafficLightMomState_Saving:
		m._TrafficLightMomState_Saving_(e)
	case TrafficLightMomState_Persisted:
		m._TrafficLightMomState_Persisted_(e)
	case TrafficLightMomState_Working:
		m._TrafficLightMomState_Working_(e)
	case TrafficLightMomState_TrafficLightApi:
		m._TrafficLightMomState_TrafficLightApi_(e)
	case TrafficLightMomState_End:
		m._TrafficLightMomState_End_(e)
	}
}

//===================== Machine Block ===================//

func (m *trafficLightMomStruct) _TrafficLightMomState_New_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.trafficLight = NewTrafficLight(m)
		m.trafficLight.Start()
		// Traffic Light\nStarted
		m._transition_(TrafficLightMomState_Saving)
		return
	}
}

func (m *trafficLightMomStruct) _TrafficLightMomState_Saving_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.data = m.trafficLight.Marshal()
		m.trafficLight = nil
		// Saved
		m._transition_(TrafficLightMomState_Persisted)
		return
	}
}

func (m *trafficLightMomStruct) _TrafficLightMomState_Persisted_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "tick":
		// Tick
		m._transition_(TrafficLightMomState_Working)
		return
	case "<<":
		// Stop
		m._transition_(TrafficLightMomState_End)
		return
	}
}

func (m *trafficLightMomStruct) _TrafficLightMomState_Working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.trafficLight = LoadTrafficLight(m, m.data)
		m.trafficLight.Tick()
		// Done
		m._transition_(TrafficLightMomState_Saving)
		return
	}
	m._TrafficLightMomState_TrafficLightApi_(e)

}

func (m *trafficLightMomStruct) _TrafficLightMomState_TrafficLightApi_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "enterRed":
		m.enterRed()
		return
	case "enterGreen":
		m.enterGreen()
		return
	case "enterYellow":
		m.enterYellow()
		return
	case "enterFlashingRed":
		m.enterFlashingRed()
		return
	case "exitFlashingRed":
		m.exitFlashingRed()
		return
	case "startWorkingTimer":
		m.startWorkingTimer()
		return
	case "stopWorkingTimer":
		m.stopWorkingTimer()
		return
	case "startFlashingTimer":
		m.startFlashingTimer()
		return
	case "stopFlashingTimer":
		m.stopFlashingTimer()
		return
	case "changeColor":
		m.changeColor(e.Params["color"].(string))
		return
	case "startFlashing":
		m.startFlashing()
		return
	case "stopFlashing":
		m.stopFlashing()
		return
	case "changeFlashingAnimation":
		m.changeFlashingAnimation()
		return
	case "log":
		m.log(e.Params["msg"].(string))
		return
	}
}

func (m *trafficLightMomStruct) _TrafficLightMomState_End_(e *framelang.FrameEvent) {
	switch e.Msg {
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *trafficLightMomStruct) _transition_(newState TrafficLightMomState) {
	m._mux_(&framelang.FrameEvent{Msg: "<"})
	m._state_ = newState
	m._mux_(&framelang.FrameEvent{Msg: ">"})
}

/********************
// Sample Actions Implementation
package trafficlightmom

func (m *trafficLightMomStruct) enterRed()  {}
func (m *trafficLightMomStruct) enterGreen()  {}
func (m *trafficLightMomStruct) enterYellow()  {}
func (m *trafficLightMomStruct) enterFlashingRed()  {}
func (m *trafficLightMomStruct) exitFlashingRed()  {}
func (m *trafficLightMomStruct) startWorkingTimer()  {}
func (m *trafficLightMomStruct) stopWorkingTimer()  {}
func (m *trafficLightMomStruct) startFlashingTimer()  {}
func (m *trafficLightMomStruct) stopFlashingTimer()  {}
func (m *trafficLightMomStruct) changeColor(color string)  {}
func (m *trafficLightMomStruct) startFlashing()  {}
func (m *trafficLightMomStruct) stopFlashing()  {}
func (m *trafficLightMomStruct) changeFlashingAnimation()  {}
func (m *trafficLightMomStruct) log(msg string)  {}
********************/
