package main

import (
	"fmt"
	"strconv"

	"github.com/frame-lang/frame-demos/compartmentparams/framelang"
)

func NewCompartmentParams(state_param int, enter_param int) CompartmentParams {
	m := &compartmentParamsStruct{}

	// Validate interfaces
	var _ CompartmentParams = m
	var _ CompartmentParams_actions = m
	m._compartment_ = NewCompartmentParamsCompartment(CompartmentParamsState_S0)
	m._compartment_.StateArgs["state_param"] = state_param
	m._compartment_.StateVars["state_var"] = 100

	// Initialize domain

	// Send system start event
	params := make(map[string]interface{})
	params["enter_param"] = enter_param
	e := framelang.FrameEvent{Msg: ">", Params: params}
	m._mux_(&e)
	return m
}

type CompartmentParamsState uint

const (
	CompartmentParamsState_S0 CompartmentParamsState = iota
	CompartmentParamsState_S1
)

type CompartmentParams interface {
}

type CompartmentParams_actions interface {
	print(s string)
}

type compartmentParamsStruct struct {
	_compartment_     *CompartmentParamsCompartment
	_nextCompartment_ *CompartmentParamsCompartment
}

//====================== Multiplexer ====================//

func (m *compartmentParamsStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case CompartmentParamsState_S0:
		m._CompartmentParamsState_S0_(e)
	case CompartmentParamsState_S1:
		m._CompartmentParamsState_S1_(e)
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

func (m *compartmentParamsStruct) _CompartmentParamsState_S0_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(strconv.Itoa((m._compartment_.StateArgs["state_param"].(int))) + " " + strconv.Itoa((m._compartment_.StateVars["state_var"].(int))) + " " + strconv.Itoa(e.Params["enter_param"].(int)))
		compartment := NewCompartmentParamsCompartment(CompartmentParamsState_S1)
		compartment._forwardEvent_ = e
		compartment.StateArgs["state_param"] = m._compartment_.StateArgs["state_param"].(int) + 20

		compartment.StateVars["state_var"] = 200

		m._transition_(compartment)
		return
	}
}

func (m *compartmentParamsStruct) _CompartmentParamsState_S1_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(strconv.Itoa((m._compartment_.StateArgs["state_param"].(int))) + " " + strconv.Itoa((m._compartment_.StateVars["state_var"].(int))) + " " + strconv.Itoa(e.Params["enter_param"].(int)))
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *compartmentParamsStruct) _transition_(compartment *CompartmentParamsCompartment) {
	m._nextCompartment_ = compartment
}

func (m *compartmentParamsStruct) _do_transition_(nextCompartment *CompartmentParamsCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//

func (m *compartmentParamsStruct) print(s string) {
	fmt.Println(s)
}

/********************************************************

// Unimplemented Actions


********************************************************/

//=============== Compartment ==============//

type CompartmentParamsCompartment struct {
	State          CompartmentParamsState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewCompartmentParamsCompartment(state CompartmentParamsState) *CompartmentParamsCompartment {
	c := &CompartmentParamsCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
