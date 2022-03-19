package main

import (
	"container/list"

	"github.com/frame-lang/frame-demos/statecompartmenthistory/framelang"
)

type HistoryBasicState uint

const (
	HistoryBasicState_Start HistoryBasicState = iota
	HistoryBasicState_S0
	HistoryBasicState_S1
	HistoryBasicState_DeadEnd
)

type HistoryBasic interface {
	Start()
	SwitchState()
	GotoDeadEnd()
	Back()
}

type HistoryBasic_actions interface {
	print(msg string)
}

type historyBasicStruct struct {
	_state_      HistoryBasicState
	_stateStack_ *Stack
}

func NewHistoryBasic() HistoryBasic {
	m := &historyBasicStruct{}

	// Validate interfaces
	var _ HistoryBasic = m
	var _ HistoryBasic_actions = m

	m._stateStack_ = &Stack{stack: list.New()}

	// Initialize domain

	return m
}

//===================== Interface Block ===================//

func (m *historyBasicStruct) Start() {
	e := framelang.FrameEvent{Msg: ">>"}
	m._mux_(&e)
}

func (m *historyBasicStruct) SwitchState() {
	e := framelang.FrameEvent{Msg: "switchState"}
	m._mux_(&e)
}

func (m *historyBasicStruct) GotoDeadEnd() {
	e := framelang.FrameEvent{Msg: "gotoDeadEnd"}
	m._mux_(&e)
}

func (m *historyBasicStruct) Back() {
	e := framelang.FrameEvent{Msg: "back"}
	m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *historyBasicStruct) _mux_(e *framelang.FrameEvent) {
	switch m._state_ {
	case HistoryBasicState_Start:
		m._HistoryBasicState_Start_(e)
	case HistoryBasicState_S0:
		m._HistoryBasicState_S0_(e)
	case HistoryBasicState_S1:
		m._HistoryBasicState_S1_(e)
	case HistoryBasicState_DeadEnd:
		m._HistoryBasicState_DeadEnd_(e)
	}
}

//===================== Machine Block ===================//

func (m *historyBasicStruct) _HistoryBasicState_Start_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m._transition_(HistoryBasicState_S0)
		return
	}
}

func (m *historyBasicStruct) _HistoryBasicState_S0_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print("Enter S0")
		return
	case "switchState":
		// Switch\nState
		m._transition_(HistoryBasicState_S1)
		return
	case "gotoDeadEnd":
		m._stateStack_push_(m._state_)
		// Goto\nDead End
		m._transition_(HistoryBasicState_DeadEnd)
		return
	}
}

func (m *historyBasicStruct) _HistoryBasicState_S1_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print("Enter S1")
		return
	case "switchState":
		// Switch\nState
		m._transition_(HistoryBasicState_S0)
		return
	case "gotoDeadEnd":
		m._stateStack_push_(m._state_)
		// Goto\nDead End
		m._transition_(HistoryBasicState_DeadEnd)
		return
	}
}

func (m *historyBasicStruct) _HistoryBasicState_DeadEnd_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print("Enter $DeadEnd")
		return
	case "back":
		// Go Back
		state := m._stateStack_pop_()
		m._transition_(state.(HistoryBasicState))
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *historyBasicStruct) _transition_(newState HistoryBasicState) {
	m._mux_(&framelang.FrameEvent{Msg: "<"})
	m._state_ = newState
	m._mux_(&framelang.FrameEvent{Msg: ">"})
}

func (m *historyBasicStruct) _stateStack_push_(state HistoryBasicState) {
	m._stateStack_.Push(state)
}

func (m *historyBasicStruct) _stateStack_pop_() interface{} {
	val, _ := m._stateStack_.Front()
	m._stateStack_.Pop()
	return val
}

type HistoryBasicCompartment struct {
	State     HistoryBasicState
	StateArgs map[string]interface{}
	StateVars map[string]interface{}
	EnterArgs map[string]interface{}
}

func NewHistoryBasicCompartment(state HistoryBasicState) *HistoryBasicCompartment {
	c := &HistoryBasicCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	return c
}

func (c *HistoryBasicCompartment) AddStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *HistoryBasicCompartment) SetStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *HistoryBasicCompartment) GetStateArg(name string) interface{} {
	return c.StateArgs[name]
}

func (c *HistoryBasicCompartment) AddStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *HistoryBasicCompartment) SetStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *HistoryBasicCompartment) GetStateVar(name string) interface{} {
	return c.StateVars[name]
}

func (c *HistoryBasicCompartment) AddEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *HistoryBasicCompartment) SetEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *HistoryBasicCompartment) GetEnterArg(name string) interface{} {
	return c.EnterArgs[name]
}

func (c *HistoryBasicCompartment) GetEnterArgs() map[string]interface{} {
	return c.EnterArgs
}

/********************
// Sample Actions Implementation
package historybasic

func (m *historyBasicStruct) print(msg string)  {}
********************/
