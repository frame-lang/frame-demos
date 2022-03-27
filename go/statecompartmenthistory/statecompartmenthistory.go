package main

import (
	"container/list"

	"github.com/frame-lang/frame-demos/statecompartmenthistory/framelang"
)

type HistoryStateContextState uint

const (
	HistoryStateContextState_Start HistoryStateContextState = iota
	HistoryStateContextState_S0
	HistoryStateContextState_S1
	HistoryStateContextState_DeadEnd
)

type HistoryStateContext interface {
	Start()
	SwitchState()
	GotoDeadEnd()
	Back()
}

type HistoryStateContext_actions interface {
	print(msg string)
}

type historyStateContextStruct struct {
	_compartment_     *HistoryStateContextCompartment
	_nextCompartment_ *HistoryStateContextCompartment
	_stateStack_      *Stack
}

func NewHistoryStateContext() HistoryStateContext {
	m := &historyStateContextStruct{}

	// Validate interfaces
	var _ HistoryStateContext = m
	var _ HistoryStateContext_actions = m

	m._stateStack_ = &Stack{stack: list.New()}
	m._compartment_ = NewHistoryStateContextCompartment(HistoryStateContextState_Start)

	// Initialize domain

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

//===================== Interface Block ===================//

func (m *historyStateContextStruct) Start() {
	e := framelang.FrameEvent{Msg: ">>"}
	m._mux_(&e)
}

func (m *historyStateContextStruct) SwitchState() {
	e := framelang.FrameEvent{Msg: "switchState"}
	m._mux_(&e)
}

func (m *historyStateContextStruct) GotoDeadEnd() {
	e := framelang.FrameEvent{Msg: "gotoDeadEnd"}
	m._mux_(&e)
}

func (m *historyStateContextStruct) Back() {
	e := framelang.FrameEvent{Msg: "back"}
	m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *historyStateContextStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case HistoryStateContextState_Start:
		m._HistoryStateContextState_Start_(e)
	case HistoryStateContextState_S0:
		m._HistoryStateContextState_S0_(e)
	case HistoryStateContextState_S1:
		m._HistoryStateContextState_S1_(e)
	case HistoryStateContextState_DeadEnd:
		m._HistoryStateContextState_DeadEnd_(e)
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

func (m *historyStateContextStruct) _HistoryStateContextState_Start_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":

		compartment := NewHistoryStateContextCompartment(HistoryStateContextState_S0)
		compartment.AddStateVar("enterMsg", "Enter S0")

		m._transition_(compartment)
		return
	}
}

func (m *historyStateContextStruct) _HistoryStateContextState_S0_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print((m._compartment_.GetStateVar("enterMsg").(string)))
		return
	case "switchState":

		// Switch\nState
		compartment := NewHistoryStateContextCompartment(HistoryStateContextState_S1)
		m._transition_(compartment)
		return
	case "gotoDeadEnd":
		m._stateStack_push_(m._compartment_)

		// Goto\nDead End
		compartment := NewHistoryStateContextCompartment(HistoryStateContextState_DeadEnd)
		m._transition_(compartment)
		return
	}
}

func (m *historyStateContextStruct) _HistoryStateContextState_S1_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print("Enter S1")
		return
	case "switchState":

		// Switch\nState
		compartment := NewHistoryStateContextCompartment(HistoryStateContextState_S0)
		compartment.AddStateVar("enterMsg", "Enter S0")

		m._transition_(compartment)
		return
	case "gotoDeadEnd":
		m._stateStack_push_(m._compartment_)

		// Goto\nDead End
		compartment := NewHistoryStateContextCompartment(HistoryStateContextState_DeadEnd)
		m._transition_(compartment)
		return
	}
}

func (m *historyStateContextStruct) _HistoryStateContextState_DeadEnd_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print("Enter $DeadEnd")
		return
	case "back":
		// Go Back
		compartment := m._stateStack_pop_()
		m._transition_(compartment)
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *historyStateContextStruct) _transition_(compartment *HistoryStateContextCompartment) {
	m._nextCompartment_ = compartment
}

func (m *historyStateContextStruct) _do_transition_(nextCompartment *HistoryStateContextCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.GetExitArgs(), Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.GetEnterArgs(), Ret: nil})
}

func (m *historyStateContextStruct) _stateStack_push_(compartment *HistoryStateContextCompartment) {
	m._stateStack_.Push(compartment)
}

func (m *historyStateContextStruct) _stateStack_pop_() *HistoryStateContextCompartment {
	compartment, _ := m._stateStack_.Front()
	return compartment.(*HistoryStateContextCompartment)
}

/********************
// Sample Actions Implementation
package historystatecontext

func (m *historyStateContextStruct) print(msg string)  {}
********************/

//=============== Compartment ==============//

type HistoryStateContextCompartment struct {
	State          HistoryStateContextState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewHistoryStateContextCompartment(state HistoryStateContextState) *HistoryStateContextCompartment {
	c := &HistoryStateContextCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}

func (c *HistoryStateContextCompartment) AddStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *HistoryStateContextCompartment) SetStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *HistoryStateContextCompartment) GetStateArg(name string) interface{} {
	return c.StateArgs[name]
}

func (c *HistoryStateContextCompartment) AddStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *HistoryStateContextCompartment) SetStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *HistoryStateContextCompartment) GetStateVar(name string) interface{} {
	return c.StateVars[name]
}

func (c *HistoryStateContextCompartment) AddEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *HistoryStateContextCompartment) SetEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *HistoryStateContextCompartment) GetEnterArg(name string) interface{} {
	return c.EnterArgs[name]
}

func (c *HistoryStateContextCompartment) GetEnterArgs() map[string]interface{} {
	return c.EnterArgs
}

func (c *HistoryStateContextCompartment) AddExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *HistoryStateContextCompartment) SetExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *HistoryStateContextCompartment) GetExitArg(name string) interface{} {
	return c.ExitArgs[name]
}

func (c *HistoryStateContextCompartment) GetExitArgs() map[string]interface{} {
	return c.ExitArgs
}
