package main

import (
	"strconv"

	"github.com/frame-lang/frame-demos/fizzbuzz/framelang"
)

func NewFizzBuzz() FizzBuzz {
	m := &fizzBuzzStruct{}

	// Validate interfaces
	var _ FizzBuzz = m
	var _ FizzBuzz_actions = m

	// Create and intialize start state compartment.
	m._compartment_ = NewFizzBuzzCompartment(FizzBuzzState_Begin)

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

type FizzBuzzState uint

const (
	FizzBuzzState_Begin FizzBuzzState = iota
	FizzBuzzState_Fizz
	FizzBuzzState_Buzz
	FizzBuzzState_Digit
	FizzBuzzState_End
)

type FizzBuzz interface {
	Start()
}

type FizzBuzz_actions interface {
	print(msg string)
	gt_100(i int) bool
	mod3_eq0(i int) bool
	mod5_eq0(i int) bool
	plus_1(i int) int
}

type fizzBuzzStruct struct {
	_compartment_     *FizzBuzzCompartment
	_nextCompartment_ *FizzBuzzCompartment
}

//===================== Interface Block ===================//

func (m *fizzBuzzStruct) Start() {
	e := framelang.FrameEvent{Msg: ">>"}
	m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *fizzBuzzStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case FizzBuzzState_Begin:
		m._FizzBuzzState_Begin_(e)
	case FizzBuzzState_Fizz:
		m._FizzBuzzState_Fizz_(e)
	case FizzBuzzState_Buzz:
		m._FizzBuzzState_Buzz_(e)
	case FizzBuzzState_Digit:
		m._FizzBuzzState_Digit_(e)
	case FizzBuzzState_End:
		m._FizzBuzzState_End_(e)
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

func (m *fizzBuzzStruct) _FizzBuzzState_Begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		// start
		compartment := NewFizzBuzzCompartment(FizzBuzzState_Fizz)
		compartment.EnterArgs["i"] = 1
		m._transition_(compartment)
		return
	}
}

func (m *fizzBuzzStruct) _FizzBuzzState_Fizz_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		if m.gt_100(e.Params["i"].(int)) {
			// i > 100
			compartment := NewFizzBuzzCompartment(FizzBuzzState_End)
			m._transition_(compartment)
			return
		}
		if m.mod3_eq0(e.Params["i"].(int)) {
			m.print("Fizz")
			// i % 3 == 0
			compartment := NewFizzBuzzCompartment(FizzBuzzState_Buzz)
			compartment.EnterArgs["i"] = e.Params["i"].(int)
			compartment.EnterArgs["fizzed"] = true
			m._transition_(compartment)
		} else {
			// i % 3 != 0
			compartment := NewFizzBuzzCompartment(FizzBuzzState_Buzz)
			compartment.EnterArgs["i"] = e.Params["i"].(int)
			compartment.EnterArgs["fizzed"] = false
			m._transition_(compartment)
		}
		return
	}
}

func (m *fizzBuzzStruct) _FizzBuzzState_Buzz_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		if m.mod5_eq0(e.Params["i"].(int)) {
			m.print("Buzz")
			// i % 5 == 0
			m._compartment_.ExitArgs["output"] = " "
			compartment := NewFizzBuzzCompartment(FizzBuzzState_Fizz)
			compartment.EnterArgs["i"] = m.plus_1(e.Params["i"].(int))
			m._transition_(compartment)
			return
		}
		if e.Params["fizzed"].(bool) {
			// fizzed
			m._compartment_.ExitArgs["output"] = " "
			compartment := NewFizzBuzzCompartment(FizzBuzzState_Fizz)
			compartment.EnterArgs["i"] = m.plus_1(e.Params["i"].(int))
			m._transition_(compartment)
			return
		}
		// ! mod3 | mod5
		m._compartment_.ExitArgs["output"] = ""
		compartment := NewFizzBuzzCompartment(FizzBuzzState_Digit)
		compartment.EnterArgs["i"] = e.Params["i"].(int)
		m._transition_(compartment)
		return
	case "<":
		m.print(e.Params["output"].(string))
		return
	}
}

func (m *fizzBuzzStruct) _FizzBuzzState_Digit_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print(strconv.Itoa(e.Params["i"].(int)))
		m.print(" ")
		// loop
		compartment := NewFizzBuzzCompartment(FizzBuzzState_Fizz)
		compartment.EnterArgs["i"] = m.plus_1(e.Params["i"].(int))
		m._transition_(compartment)
		return
	}
}

func (m *fizzBuzzStruct) _FizzBuzzState_End_(e *framelang.FrameEvent) {
	switch e.Msg {
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *fizzBuzzStruct) _transition_(compartment *FizzBuzzCompartment) {
	m._nextCompartment_ = compartment
}

func (m *fizzBuzzStruct) _do_transition_(nextCompartment *FizzBuzzCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//

/********************************************************

// Unimplemented Actions

func (m *fizzBuzzStruct) print(msg string)  {}
func (m *fizzBuzzStruct) gt_100(i int) bool {}
func (m *fizzBuzzStruct) mod3_eq0(i int) bool {}
func (m *fizzBuzzStruct) mod5_eq0(i int) bool {}
func (m *fizzBuzzStruct) plus_1(i int) int {}

********************************************************/

//=============== Compartment ==============//

type FizzBuzzCompartment struct {
	State          FizzBuzzState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewFizzBuzzCompartment(state FizzBuzzState) *FizzBuzzCompartment {
	c := &FizzBuzzCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
