package main

import (
	"fmt"
	"strconv"

	"github.com/frame-lang/frame-demos/changestate/framelang"
)

type DebugBugState uint

const (
	DebugBugState_S0 DebugBugState = iota
	DebugBugState_S1
)

type DebugBug interface {
}

type DebugBug_actions interface {
	print(s string)
	test()
}

type debugBugStruct struct {
	_compartment_     *DebugBugCompartment
	_nextCompartment_ *DebugBugCompartment
}

func NewDebugBug(state_param int, enter_param int) DebugBug {
	m := &debugBugStruct{}

	// Validate interfaces
	var _ DebugBug = m
	var _ DebugBug_actions = m
	m._compartment_ = NewDebugBugCompartment(DebugBugState_S0)
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

//====================== Multiplexer ====================//

func (m *debugBugStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case DebugBugState_S0:
		m._DebugBugState_S0_(e)
	case DebugBugState_S1:
		m._DebugBugState_S1_(e)
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

func (m *debugBugStruct) _DebugBugState_S0_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(strconv.Itoa((m._compartment_.StateArgs["state_param"].(int))) + " " + strconv.Itoa((m._compartment_.StateVars["state_var"].(int))) + " " + strconv.Itoa(e.Params["enter_param"].(int)))
		compartment := NewDebugBugCompartment(DebugBugState_S1)
		compartment._forwardEvent_ = e
		compartment.StateArgs["state_param"] = m._compartment_.StateArgs["state_param"].(int) + 20

		compartment.StateVars["state_var"] = 200

		m._transition_(compartment)
		return
	}
}

func (m *debugBugStruct) _DebugBugState_S1_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(strconv.Itoa((m._compartment_.StateArgs["state_param"].(int))) + " " + strconv.Itoa((m._compartment_.StateVars["state_var"].(int))) + " " + strconv.Itoa(e.Params["enter_param"].(int)))
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *debugBugStruct) _transition_(compartment *DebugBugCompartment) {
	m._nextCompartment_ = compartment
}

func (m *debugBugStruct) _do_transition_(nextCompartment *DebugBugCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//

func (m *debugBugStruct) print(s string) {
	fmt.Println(s)
}

//********************************************************//

// Unimplemented Actions

func (m *debugBugStruct) test() {}

//********************************************************/

//=============== Compartment ==============//

type DebugBugCompartment struct {
	State          DebugBugState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewDebugBugCompartment(state DebugBugState) *DebugBugCompartment {
	c := &DebugBugCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
