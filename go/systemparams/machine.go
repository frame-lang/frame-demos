package main

import (
	"fmt"

	"github.com/frame-lang/frame-demos/systemparams/framelang"
)

func NewSystemParams(stateMsg string, enterMsg string) SystemParams {
	m := &systemParamsStruct{}

	// Validate interfaces
	var _ SystemParams = m
	var _ SystemParams_actions = m
	m._compartment_ = NewSystemParamsCompartment(SystemParamsState_Begin)
	m._compartment_.StateArgs["stateMsg"] = stateMsg

	// Initialize domain

	// Send system start event
	params := make(map[string]interface{})
	params["enterMsg"] = enterMsg
	e := framelang.FrameEvent{Msg: ">", Params: params}
	m._mux_(&e)
	return m
}

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

func (m *systemParamsStruct) _SystemParamsState_Begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print((m._compartment_.StateArgs["stateMsg"].(string)) + " " + e.Params["enterMsg"].(string))
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *systemParamsStruct) _transition_(compartment *SystemParamsCompartment) {
	m._nextCompartment_ = compartment
}

func (m *systemParamsStruct) _do_transition_(nextCompartment *SystemParamsCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//

func (m *systemParamsStruct) print(msg string) {

	fmt.Println(msg)

}

/********************************************************

// Unimplemented Actions


********************************************************/

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
