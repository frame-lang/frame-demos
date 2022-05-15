package main

import (
	"container/list"

	"github.com/frame-lang/frame-demos/statecompartmenthistory/framelang"
)

func NewHistoryBasic() HistoryBasic {
	m := &historyBasicStruct{}

	// Validate interfaces
	var _ HistoryBasic = m
	var _ HistoryBasic_actions = m
	// History mechanism used in spec. Create state stack.
	m._stateStack_ = &Stack{stack: list.New()}

	// Create and intialize start state compartment.
	m._compartment_ = NewHistoryBasicCompartment(HistoryBasicState_Start)

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

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
	_compartment_     *HistoryBasicCompartment
	_nextCompartment_ *HistoryBasicCompartment
	_stateStack_      *Stack
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
	switch m._compartment_.State {
	case HistoryBasicState_Start:
		m._HistoryBasicState_Start_(e)
	case HistoryBasicState_S0:
		m._HistoryBasicState_S0_(e)
	case HistoryBasicState_S1:
		m._HistoryBasicState_S1_(e)
	case HistoryBasicState_DeadEnd:
		m._HistoryBasicState_DeadEnd_(e)
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

func (m *historyBasicStruct) _HistoryBasicState_Start_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		compartment := NewHistoryBasicCompartment(HistoryBasicState_S0)
		m._transition_(compartment)
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
		compartment := NewHistoryBasicCompartment(HistoryBasicState_S1)
		m._transition_(compartment)
		return
	case "gotoDeadEnd":
		m._stateStack_push_(m._compartment_)
		// Goto\nDead End
		compartment := NewHistoryBasicCompartment(HistoryBasicState_DeadEnd)
		m._transition_(compartment)
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
		compartment := NewHistoryBasicCompartment(HistoryBasicState_S0)
		m._transition_(compartment)
		return
	case "gotoDeadEnd":
		m._stateStack_push_(m._compartment_)
		// Goto\nDead End
		compartment := NewHistoryBasicCompartment(HistoryBasicState_DeadEnd)
		m._transition_(compartment)
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
		compartment := m._stateStack_pop_()
		m._transition_(compartment)
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *historyBasicStruct) _transition_(compartment *HistoryBasicCompartment) {
	m._nextCompartment_ = compartment
}

func (m *historyBasicStruct) _do_transition_(nextCompartment *HistoryBasicCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

func (m *historyBasicStruct) _stateStack_push_(compartment *HistoryBasicCompartment) {
	m._stateStack_.Push(compartment)
}

func (m *historyBasicStruct) _stateStack_pop_() *HistoryBasicCompartment {
	compartment, _ := m._stateStack_.Front()
	return compartment.(*HistoryBasicCompartment)
}

//===================== Actions Block ===================//

/********************************************************

// Unimplemented Actions

func (m *historyBasicStruct) print(msg string)  {}

********************************************************/

//=============== Compartment ==============//

type HistoryBasicCompartment struct {
	State          HistoryBasicState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewHistoryBasicCompartment(state HistoryBasicState) *HistoryBasicCompartment {
	c := &HistoryBasicCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
