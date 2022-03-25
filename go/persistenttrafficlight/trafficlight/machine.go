package trafficlight

import (
	"encoding/json"

	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

type TrafficLightState uint

const (
	TrafficLightState_Begin TrafficLightState = iota
	TrafficLightState_Red
	TrafficLightState_Green
	TrafficLightState_Yellow
	TrafficLightState_FlashingRed
	TrafficLightState_End
	TrafficLightState_Working
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

type TrafficLight_actions interface {
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
	mom               TrafficLightMom
	_compartment_     *TrafficLightCompartment
	_nextCompartment_ *TrafficLightCompartment
	flashColor        string
}

type marshalStruct struct {
	TrafficLightCompartment
	FlashColor string
}

func NewTrafficLight(mom TrafficLightMom) TrafficLight {
	m := &trafficLightStruct{}
	m.mom = mom

	// Validate interfaces
	var _ TrafficLight = m
	var _ TrafficLight_actions = m
	m._compartment_ = NewTrafficLightCompartment(TrafficLightState_Begin)

	// Initialize domain
	m.flashColor = ""

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

func LoadTrafficLight(mom TrafficLightMom, data []byte) TrafficLight {
	m := &trafficLightStruct{}
	m.mom = mom

	// Validate interfaces
	var _ TrafficLight = m
	var _ TrafficLight_actions = m

	// Unmarshal
	var marshal marshalStruct
	err := json.Unmarshal(data, &marshal)
	if err != nil {
		return nil
	}

	// Initialize machine
	m._compartment_ = &marshal.TrafficLightCompartment

	m.flashColor = marshal.FlashColor

	return m

}

func (m *trafficLightStruct) MarshalJSON() ([]byte, error) {
	data := marshalStruct{
		TrafficLightCompartment: *m._compartment_,
		FlashColor:              m.flashColor,
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
	switch m._compartment_.State {
	case TrafficLightState_Begin:
		m._TrafficLightState_Begin_(e)
	case TrafficLightState_Red:
		m._TrafficLightState_Red_(e)
	case TrafficLightState_Green:
		m._TrafficLightState_Green_(e)
	case TrafficLightState_Yellow:
		m._TrafficLightState_Yellow_(e)
	case TrafficLightState_FlashingRed:
		m._TrafficLightState_FlashingRed_(e)
	case TrafficLightState_End:
		m._TrafficLightState_End_(e)
	case TrafficLightState_Working:
		m._TrafficLightState_Working_(e)
	}

	for m._nextCompartment_ != nil {
		nextCompartment := m._nextCompartment_
		m._nextCompartment_ = nil
		m._do_transition_(nextCompartment)
		if m._compartment_._forwardEvent_ != nil {
			m._mux_(m._compartment_._forwardEvent_)
			m._compartment_._forwardEvent_ = nil
		}
	}
}

//===================== Machine Block ===================//

func (m *trafficLightStruct) _TrafficLightState_Begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.startWorkingTimer()

		compartment := NewTrafficLightCompartment(TrafficLightState_Red)
		m._transition_(compartment)
		return
	}
}

func (m *trafficLightStruct) _TrafficLightState_Red_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterRed()
		return
	case "tick":

		compartment := NewTrafficLightCompartment(TrafficLightState_Green)
		m._transition_(compartment)
		return
	}
	m._TrafficLightState_Working_(e)

}

func (m *trafficLightStruct) _TrafficLightState_Green_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterGreen()
		return
	case "tick":

		compartment := NewTrafficLightCompartment(TrafficLightState_Yellow)
		m._transition_(compartment)
		return
	}
	m._TrafficLightState_Working_(e)

}

func (m *trafficLightStruct) _TrafficLightState_Yellow_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.enterYellow()
		return
	case "tick":

		compartment := NewTrafficLightCompartment(TrafficLightState_Red)
		m._transition_(compartment)
		return
	}
	m._TrafficLightState_Working_(e)

}

func (m *trafficLightStruct) _TrafficLightState_FlashingRed_(e *framelang.FrameEvent) {
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

		compartment := NewTrafficLightCompartment(TrafficLightState_Red)
		m._transition_(compartment)
		return
	case "stop":

		compartment := NewTrafficLightCompartment(TrafficLightState_End)
		m._transition_(compartment)
		return
	}
}

func (m *trafficLightStruct) _TrafficLightState_End_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.stopWorkingTimer()
		return
	}
}

func (m *trafficLightStruct) _TrafficLightState_Working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "stop":

		compartment := NewTrafficLightCompartment(TrafficLightState_End)
		m._transition_(compartment)
		return
	case "systemError":

		compartment := NewTrafficLightCompartment(TrafficLightState_FlashingRed)
		m._transition_(compartment)
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *trafficLightStruct) _transition_(compartment *TrafficLightCompartment) {
	m._nextCompartment_ = compartment
}

func (m *trafficLightStruct) _do_transition_(nextCompartment *TrafficLightCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.GetExitArgs(), Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.GetEnterArgs(), Ret: nil})
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

//=============== Compartment ==============//

type TrafficLightCompartment struct {
	State          TrafficLightState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewTrafficLightCompartment(state TrafficLightState) *TrafficLightCompartment {
	c := &TrafficLightCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}

func (c *TrafficLightCompartment) AddStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *TrafficLightCompartment) SetStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *TrafficLightCompartment) GetStateArg(name string) interface{} {
	return c.StateArgs[name]
}

func (c *TrafficLightCompartment) AddStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *TrafficLightCompartment) SetStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *TrafficLightCompartment) GetStateVar(name string) interface{} {
	return c.StateVars[name]
}

func (c *TrafficLightCompartment) AddEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *TrafficLightCompartment) SetEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *TrafficLightCompartment) GetEnterArg(name string) interface{} {
	return c.EnterArgs[name]
}

func (c *TrafficLightCompartment) GetEnterArgs() map[string]interface{} {
	return c.EnterArgs
}

func (c *TrafficLightCompartment) AddExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *TrafficLightCompartment) SetExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *TrafficLightCompartment) GetExitArg(name string) interface{} {
	return c.ExitArgs[name]
}

func (c *TrafficLightCompartment) GetExitArgs() map[string]interface{} {
	return c.ExitArgs
}
