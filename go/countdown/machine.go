package main

import (
	"fmt"
	"strconv"

	"github.com/frame-lang/frame-demos/changestate/framelang"
)

func NewChangeState(i int) ChangeState {
	m := &changeStateStruct{}

	// Validate interfaces
	var _ ChangeState = m
	var _ ChangeState_actions = m
	m._compartment_ = NewChangeStateCompartment(ChangeStateState_S0)
	m._compartment_.StateArgs["i"] = i
	m._compartment_.StateVars["dec"] = 1

	// Initialize domain

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

type ChangeStateState uint

const (
	ChangeStateState_S0 ChangeStateState = iota
	ChangeStateState_S1
	ChangeStateState_Stop
)

type ChangeState interface {
}

type ChangeState_actions interface {
	print(s string)
}

type changeStateStruct struct {
	_compartment_     *ChangeStateCompartment
	_nextCompartment_ *ChangeStateCompartment
}

//====================== Multiplexer ====================//

func (m *changeStateStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case ChangeStateState_S0:
		m._ChangeStateState_S0_(e)
	case ChangeStateState_S1:
		m._ChangeStateState_S1_(e)
	case ChangeStateState_Stop:
		m._ChangeStateState_Stop_(e)
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

func (m *changeStateStruct) _ChangeStateState_S0_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m._compartment_.StateArgs["i"] = m._compartment_.StateArgs["i"].(int) - m._compartment_.StateVars["dec"].(int)
		m.print(strconv.Itoa((m._compartment_.StateArgs["i"].(int))))
		if (m._compartment_.StateArgs["i"].(int)) == 0 {
			compartment := NewChangeStateCompartment(ChangeStateState_Stop)
			m._transition_(compartment)
			return
		}
		compartment := NewChangeStateCompartment(ChangeStateState_S1)
		compartment.EnterArgs["i"] = m._compartment_.StateArgs["i"].(int)
		m._transition_(compartment)
		return
	}
}

func (m *changeStateStruct) _ChangeStateState_S1_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		compartment := NewChangeStateCompartment(ChangeStateState_S0)
		compartment.StateArgs["i"] = e.Params["i"].(int)

		compartment.StateVars["dec"] = 1

		m._transition_(compartment)
		return
	}
}

func (m *changeStateStruct) _ChangeStateState_Stop_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print("done")
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *changeStateStruct) _transition_(compartment *ChangeStateCompartment) {
	m._nextCompartment_ = compartment
}

func (m *changeStateStruct) _do_transition_(nextCompartment *ChangeStateCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//

func (m *changeStateStruct) print(s string) {
	fmt.Println(s)
}

//********************************************************//

// Unimplemented Actions

//********************************************************/

//=============== Compartment ==============//

type ChangeStateCompartment struct {
	State          ChangeStateState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewChangeStateCompartment(state ChangeStateState) *ChangeStateCompartment {
	c := &ChangeStateCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
