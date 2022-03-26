package main

import (
	"strconv"

	"github.com/frame-lang/frame-demos/transitioneventforwarding/framelang"
)

type TransitionEventForwardingState uint

const (
	TransitionEventForwardingState_One TransitionEventForwardingState = iota
	TransitionEventForwardingState_Two
	TransitionEventForwardingState_Three
	TransitionEventForwardingState_Stop
)

type TransitionEventForwarding interface {
}

type TransitionEventForwarding_actions interface {
	print(msg string)
}

type transitionEventForwardingStruct struct {
	_compartment_     *TransitionEventForwardingCompartment
	_nextCompartment_ *TransitionEventForwardingCompartment
}

func NewTransitionEventForwarding(cycles int) TransitionEventForwarding {
	m := &transitionEventForwardingStruct{}

	// Validate interfaces
	var _ TransitionEventForwarding = m
	var _ TransitionEventForwarding_actions = m
	m._compartment_ = NewTransitionEventForwardingCompartment(TransitionEventForwardingState_One)

	// Initialize domain

	// Send system start event
	params := make(map[string]interface{})
	params["cycles"] = cycles
	e := framelang.FrameEvent{Msg: ">", Params: params}
	m._mux_(&e)
	return m
}

//====================== Multiplexer ====================//

func (m *transitionEventForwardingStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case TransitionEventForwardingState_One:
		m._TransitionEventForwardingState_One_(e)
	case TransitionEventForwardingState_Two:
		m._TransitionEventForwardingState_Two_(e)
	case TransitionEventForwardingState_Three:
		m._TransitionEventForwardingState_Three_(e)
	case TransitionEventForwardingState_Stop:
		m._TransitionEventForwardingState_Stop_(e)
	}

	// if m._nextCompartment_ != nil {
	// 	nextCompartment := m._nextCompartment_
	// 	m._nextCompartment_ = nil
	// 	if nextCompartment._forwardEvent_ != nil {
	// 		if nextCompartment._forwardEvent_.Msg == ">" {
	// 			m._compartment_ = nextCompartment
	// 			m._mux_(m._compartment_._forwardEvent_)
	// 		} else {
	// 			m._do_transition_(nextCompartment)
	// 			if m._compartment_._forwardEvent_ != nil {
	// 				m._mux_(m._compartment_._forwardEvent_)
	// 				m._compartment_._forwardEvent_ = nil
	// 			}
	// 		}

	// 	} else {
	// 		m._do_transition_(nextCompartment)
	// 		if m._compartment_._forwardEvent_ != nil {
	// 			m._mux_(m._compartment_._forwardEvent_)
	// 			m._compartment_._forwardEvent_ = nil
	// 		}
	// 	}
	// 	m._compartment_._forwardEvent_ = nil
	// }

	// if m._nextCompartment_ != nil {
	// 	nextCompartment := m._nextCompartment_
	// 	m._nextCompartment_ = nil
	// 	if nextCompartment._forwardEvent_ != nil &&
	// 		nextCompartment._forwardEvent_.Msg == ">" {
	// 		m._compartment_ = nextCompartment
	// 		m._mux_(nextCompartment._forwardEvent_)
	// 	} else {
	// 		m._do_transition_(nextCompartment)
	// 		if nextCompartment._forwardEvent_ != nil {
	// 			m._mux_(nextCompartment._forwardEvent_)
	// 		}
	// 	}
	// 	nextCompartment._forwardEvent_ = nil
	// }

	if m._nextCompartment_ != nil {
		nextCompartment := m._nextCompartment_
		m._nextCompartment_ = nil
		if nextCompartment._forwardEvent_ != nil &&
			nextCompartment._forwardEvent_.Msg == ">" {
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

func (m *transitionEventForwardingStruct) _TransitionEventForwardingState_One_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		if e.Params["cycles"].(int) == 0 {

			compartment := NewTransitionEventForwardingCompartment(TransitionEventForwardingState_Stop)
			compartment._forwardEvent_ = e
			m._transition_(compartment)
			return
		} else {

			compartment := NewTransitionEventForwardingCompartment(TransitionEventForwardingState_Two)
			compartment._forwardEvent_ = e
			m._transition_(compartment)
		}
		return
	}
}

func (m *transitionEventForwardingStruct) _TransitionEventForwardingState_Two_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":

		compartment := NewTransitionEventForwardingCompartment(TransitionEventForwardingState_Three)
		compartment._forwardEvent_ = e
		m._transition_(compartment)
		return
	}
}

func (m *transitionEventForwardingStruct) _TransitionEventForwardingState_Three_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(strconv.Itoa(e.Params["cycles"].(int)))

		compartment := NewTransitionEventForwardingCompartment(TransitionEventForwardingState_One)
		compartment.AddEnterArg("cycles", e.Params["cycles"].(int)-1)
		m._transition_(compartment)
		return
	}
}

func (m *transitionEventForwardingStruct) _TransitionEventForwardingState_Stop_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(strconv.Itoa(e.Params["cycles"].(int)))
		m.print("done")
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *transitionEventForwardingStruct) _transition_(compartment *TransitionEventForwardingCompartment) {
	m._nextCompartment_ = compartment
}

func (m *transitionEventForwardingStruct) _do_transition_(nextCompartment *TransitionEventForwardingCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.GetExitArgs(), Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.GetEnterArgs(), Ret: nil})
}

/********************
// Sample Actions Implementation
package transitioneventforwarding

func (m *transitionEventForwardingStruct) print(msg string)  {}
********************/

//=============== Compartment ==============//

type TransitionEventForwardingCompartment struct {
	State          TransitionEventForwardingState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewTransitionEventForwardingCompartment(state TransitionEventForwardingState) *TransitionEventForwardingCompartment {
	c := &TransitionEventForwardingCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}

func (c *TransitionEventForwardingCompartment) AddStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *TransitionEventForwardingCompartment) SetStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *TransitionEventForwardingCompartment) GetStateArg(name string) interface{} {
	return c.StateArgs[name]
}

func (c *TransitionEventForwardingCompartment) AddStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *TransitionEventForwardingCompartment) SetStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *TransitionEventForwardingCompartment) GetStateVar(name string) interface{} {
	return c.StateVars[name]
}

func (c *TransitionEventForwardingCompartment) AddEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *TransitionEventForwardingCompartment) SetEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *TransitionEventForwardingCompartment) GetEnterArg(name string) interface{} {
	return c.EnterArgs[name]
}

func (c *TransitionEventForwardingCompartment) GetEnterArgs() map[string]interface{} {
	return c.EnterArgs
}

func (c *TransitionEventForwardingCompartment) AddExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *TransitionEventForwardingCompartment) SetExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *TransitionEventForwardingCompartment) GetExitArg(name string) interface{} {
	return c.ExitArgs[name]
}

func (c *TransitionEventForwardingCompartment) GetExitArgs() map[string]interface{} {
	return c.ExitArgs
}
