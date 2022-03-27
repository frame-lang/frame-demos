package main

import "github.com/frame-lang/frame-demos/systemparams/framelang"

type SystemParamsState uint

const (
	SystemParamsState_Begin SystemParamsState = iota
)

type SystemParams interface {
}

type SystemParams_actions interface {
	print(msg string)
}

type systemParamsStruct struct {
	_compartment_     *SystemParamsCompartment
	_nextCompartment_ *SystemParamsCompartment
}

func NewSystemParams(msg string) SystemParams {
	m := &systemParamsStruct{}

	// Validate interfaces
	var _ SystemParams = m
	var _ SystemParams_actions = m
	m._compartment_ = NewSystemParamsCompartment(SystemParamsState_Begin)

	// Initialize domain

	// Send system start event
	params := make(map[string]interface{})
	params["msg"] = msg
	e := framelang.FrameEvent{Msg: ">", Params: params}
	m._mux_(&e)
	return m
}

//====================== Multiplexer ====================//

func (m *systemParamsStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case SystemParamsState_Begin:
		m._SystemParamsState_Begin_(e)
	}

	if m._nextCompartment_ != nil {
		nextCompartment := m._nextCompartment_
		m._nextCompartment_ = nil
		if nextCompartment._forwardEvent_ != nil &&
			nextCompartment._forwardEvent_.Msg == ">" {
			m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.GetExitArgs(), Ret: nil})
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

func (m *systemParamsStruct) _SystemParamsState_Begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(e.Params["msg"].(string))
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *systemParamsStruct) _transition_(compartment *SystemParamsCompartment) {
	m._nextCompartment_ = compartment
}

func (m *systemParamsStruct) _do_transition_(nextCompartment *SystemParamsCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.GetExitArgs(), Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.GetEnterArgs(), Ret: nil})
}

/********************
// Sample Actions Implementation
package systemparams

func (m *systemParamsStruct) print(msg string)  {}
********************/

//=============== Compartment ==============//

type SystemParamsCompartment struct {
	State          SystemParamsState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewSystemParamsCompartment(state SystemParamsState) *SystemParamsCompartment {
	c := &SystemParamsCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}

func (c *SystemParamsCompartment) AddStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *SystemParamsCompartment) SetStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *SystemParamsCompartment) GetStateArg(name string) interface{} {
	return c.StateArgs[name]
}

func (c *SystemParamsCompartment) AddStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *SystemParamsCompartment) SetStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *SystemParamsCompartment) GetStateVar(name string) interface{} {
	return c.StateVars[name]
}

func (c *SystemParamsCompartment) AddEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *SystemParamsCompartment) SetEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *SystemParamsCompartment) GetEnterArg(name string) interface{} {
	return c.EnterArgs[name]
}

func (c *SystemParamsCompartment) GetEnterArgs() map[string]interface{} {
	return c.EnterArgs
}

func (c *SystemParamsCompartment) AddExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *SystemParamsCompartment) SetExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *SystemParamsCompartment) GetExitArg(name string) interface{} {
	return c.ExitArgs[name]
}

func (c *SystemParamsCompartment) GetExitArgs() map[string]interface{} {
	return c.ExitArgs
}
