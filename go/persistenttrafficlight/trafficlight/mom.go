package trafficlight

import (
	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

func NewTrafficLightMom() TrafficLightMom {
	m := &trafficLightMomStruct{}

	// Validate interfaces
	var _ TrafficLightMom = m
	var _ TrafficLightMom_actions = m

	// Create and intialize start state compartment.
	m._compartment_ = NewTrafficLightMomCompartment(TrafficLightMomState_New)

	// Override domain variables.
	m.trafficLight = nil
	m.data = nil

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

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
	SystemError()
	SystemRestart()
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
	systemError()
	systemRestart()
	log(msg string)
}

type trafficLightMomStruct struct {
	_compartment_     *TrafficLightMomCompartment
	_nextCompartment_ *TrafficLightMomCompartment
	trafficLight      TrafficLight
	data              []byte
}

//===================== Interface Block ===================//

func (m *trafficLightMomStruct) Stop() {
	e := framelang.FrameEvent{Msg: "stop"}
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

func (m *trafficLightMomStruct) SystemError() {
	e := framelang.FrameEvent{Msg: "systemError"}
	m._mux_(&e)
}

func (m *trafficLightMomStruct) SystemRestart() {
	e := framelang.FrameEvent{Msg: "systemRestart"}
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
	switch m._compartment_.State {
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

	if m._nextCompartment_ != nil {
		nextCompartment := m._nextCompartment_
		m._nextCompartment_ = nil
		if nextCompartment._forwardEvent_ != nil &&
			nextCompartment._forwardEvent_.Msg == ">" {
			m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
			m._compartment_ = nextCompartment
			m._mux_(nextCompartment._forwardEvent_)
		} else {
			m._do_transition_(nextCompartment)
			if nextCompartment._forwardEvent_ != nil {
				m._mux_(nextCompartment._forwardEvent_)
			}
		}
		nextCompartment._forwardEvent_ = nil
	}
}

//===================== Machine Block ===================//

func (m *trafficLightMomStruct) _TrafficLightMomState_New_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.trafficLight = NewTrafficLight(m)
		// Traffic Light\nStarted
		compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Saving)
		m._transition_(compartment)
		return
	}
	m._TrafficLightMomState_TrafficLightApi_(e)

}

func (m *trafficLightMomStruct) _TrafficLightMomState_Saving_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.data = m.trafficLight.Marshal()
		m.trafficLight = nil
		// Saved
		compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Persisted)
		m._transition_(compartment)
		return
	}
}

func (m *trafficLightMomStruct) _TrafficLightMomState_Persisted_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "tick":
		// Tick
		compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Working)
		compartment._forwardEvent_ = e
		m._transition_(compartment)
		return
	case "systemError":
		// System Error
		compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Working)
		compartment._forwardEvent_ = e
		m._transition_(compartment)
		return
	case "stop":
		// Stop
		compartment := NewTrafficLightMomCompartment(TrafficLightMomState_End)
		m._transition_(compartment)
		return
	}
}

func (m *trafficLightMomStruct) _TrafficLightMomState_Working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.trafficLight = LoadTrafficLight(m, m.data)
		return
	case "tick":
		m.trafficLight.Tick()
		// Done
		compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Saving)
		m._transition_(compartment)
		return
	case "systemError":
		m.trafficLight.SystemError()
		// Done
		compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Saving)
		m._transition_(compartment)
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
	case "systemError":
		m.systemError()
		return
	case "systemRestart":
		m.systemRestart()
		return
	case "log":
		m.log(e.Params["msg"].(string))
		return
	}
}

func (m *trafficLightMomStruct) _TrafficLightMomState_End_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.trafficLight = LoadTrafficLight(m, m.data)
		m.trafficLight.Stop()
		m.trafficLight = nil
		return
	}
	m._TrafficLightMomState_TrafficLightApi_(e)

}

//=============== Machinery and Mechanisms ==============//

func (m *trafficLightMomStruct) _transition_(compartment *TrafficLightMomCompartment) {
	m._nextCompartment_ = compartment
}

func (m *trafficLightMomStruct) _do_transition_(nextCompartment *TrafficLightMomCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//

/********************************************************

// Unimplemented Actions

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
func (m *trafficLightMomStruct) systemError()  {}
func (m *trafficLightMomStruct) systemRestart()  {}
func (m *trafficLightMomStruct) log(msg string)  {}

********************************************************/

//=============== Compartment ==============//

type TrafficLightMomCompartment struct {
	State          TrafficLightMomState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewTrafficLightMomCompartment(state TrafficLightMomState) *TrafficLightMomCompartment {
	c := &TrafficLightMomCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
