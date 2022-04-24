// emitted from framec_v0.8.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
package main

import "github.com/frame-lang/frame-demos/stringtools/framelang"

func NewStringTools() StringTools {
	m := &stringToolsStruct{}

	// Validate interfaces
	var _ StringTools = m

	m._compartment_ = NewStringToolsCompartment(StringToolsState_Router)

	// Initialize domain

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

type StringToolsState uint

const (
	StringToolsState_Router StringToolsState = iota
	StringToolsState_Reverse
	StringToolsState_MakePalindrome
)

type StringTools interface {
	Reverse(str string) string
	MakePalindrome(str string) string
}

type stringToolsStruct struct {
	_compartment_     *StringToolsCompartment
	_nextCompartment_ *StringToolsCompartment
}

//===================== Interface Block ===================//

func (m *stringToolsStruct) Reverse(str string) string {
	params := make(map[string]interface{})
	params["str"] = str
	e := framelang.FrameEvent{Msg: "reverse", Params: params}
	m._mux_(&e)
	return e.Ret.(string)
}

func (m *stringToolsStruct) MakePalindrome(str string) string {
	params := make(map[string]interface{})
	params["str"] = str
	e := framelang.FrameEvent{Msg: "makePalindrome", Params: params}
	m._mux_(&e)
	return e.Ret.(string)
}

//====================== Multiplexer ====================//

func (m *stringToolsStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case StringToolsState_Router:
		m._StringToolsState_Router_(e)
	case StringToolsState_Reverse:
		m._StringToolsState_Reverse_(e)
	case StringToolsState_MakePalindrome:
		m._StringToolsState_MakePalindrome_(e)
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

func (m *stringToolsStruct) _StringToolsState_Router_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "makePalindrome":
		// make\npalindrome
		compartment := NewStringToolsCompartment(StringToolsState_MakePalindrome)
		compartment._forwardEvent_ = e
		m._transition_(compartment)
		return
	case "reverse":
		// reverse
		compartment := NewStringToolsCompartment(StringToolsState_Reverse)
		compartment._forwardEvent_ = e
		m._transition_(compartment)
		return
	}
}

func (m *stringToolsStruct) _StringToolsState_Reverse_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "reverse":
		e.Ret = reverse_str(e.Params["str"].(string))
		// ready
		compartment := NewStringToolsCompartment(StringToolsState_Router)
		m._transition_(compartment)
		return
	}
}

func (m *stringToolsStruct) _StringToolsState_MakePalindrome_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "makePalindrome":
		e.Ret = e.Params["str"].(string) + reverse_str(e.Params["str"].(string))
		// ready
		compartment := NewStringToolsCompartment(StringToolsState_Router)
		m._transition_(compartment)
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *stringToolsStruct) _transition_(compartment *StringToolsCompartment) {
	m._nextCompartment_ = compartment
}

func (m *stringToolsStruct) _do_transition_(nextCompartment *StringToolsCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//=============== Compartment ==============//

type StringToolsCompartment struct {
	State          StringToolsState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewStringToolsCompartment(state StringToolsState) *StringToolsCompartment {
	c := &StringToolsCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
