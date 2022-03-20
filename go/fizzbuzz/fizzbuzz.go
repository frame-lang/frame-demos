package main

import (
	"strconv"

	"github.com/frame-lang/frame-demos/fizzbuzz/framelang"
)

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
	//	_state_       FizzBuzzState
	_compartment_      *FizzBuzzCompartment
	_next_compartment_ *FizzBuzzCompartment
}

func NewFizzBuzz() FizzBuzz {
	m := &fizzBuzzStruct{}

	// Validate interfaces
	var _ FizzBuzz = m
	var _ FizzBuzz_actions = m
	m._compartment_ = NewFizzBuzzCompartment(FizzBuzzState_Begin)

	// Initialize domain

	return m
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

	for m._next_compartment_ != nil {
		next_compartment := m._next_compartment_
		m._next_compartment_ = nil
		m._do_transition_(next_compartment)
	}
}

//===================== Machine Block ===================//  //  try on https://frame-lang.org

func (m *fizzBuzzStruct) _FizzBuzzState_Begin_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		// start
		compartment := NewFizzBuzzCompartment(FizzBuzzState_Fizz)
		compartment.AddEnterArg("i", 1)
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
			compartment.AddEnterArg("i", e.Params["i"].(int))
			compartment.AddEnterArg("fizzed", true)
			m._transition_(compartment)
		} else {
			// i % 3 != 0
			compartment := NewFizzBuzzCompartment(FizzBuzzState_Buzz)
			compartment.AddEnterArg("i", e.Params["i"].(int))
			compartment.AddEnterArg("fizzed", false)
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
			compartment := NewFizzBuzzCompartment(FizzBuzzState_Fizz)
			compartment.AddEnterArg("i", m.plus_1(e.Params["i"].(int)))
			compartment.AddExitArg("output", " ")
			m._transition_(compartment)
			return
		}
		if e.Params["fizzed"].(bool) {
			// fizzed
			compartment := NewFizzBuzzCompartment(FizzBuzzState_Fizz)
			compartment.AddEnterArg("i", m.plus_1(e.Params["i"].(int)))
			compartment.AddExitArg("output", " ")
			m._transition_(compartment)
			return
		}
		// ! mod3 | mod5
		compartment := NewFizzBuzzCompartment(FizzBuzzState_Digit)
		compartment.AddEnterArg("i", e.Params["i"].(int))
		compartment.AddExitArg("output", "")
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
		compartment.AddEnterArg("i", m.plus_1(e.Params["i"].(int)))
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
	m._next_compartment_ = compartment
}

func (m *fizzBuzzStruct) _do_transition_(compartment *FizzBuzzCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: compartment.GetExitArgs(), Ret: nil})
	m._compartment_ = compartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.GetEnterArgs(), Ret: nil})
}

type FizzBuzzCompartment struct {
	State     FizzBuzzState
	StateArgs map[string]interface{}
	StateVars map[string]interface{}
	EnterArgs map[string]interface{}
	ExitArgs  map[string]interface{}
}

func NewFizzBuzzCompartment(state FizzBuzzState) *FizzBuzzCompartment {
	c := &FizzBuzzCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}

func (c *FizzBuzzCompartment) AddStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *FizzBuzzCompartment) SetStateArg(name string, value interface{}) {
	c.StateArgs[name] = value
}

func (c *FizzBuzzCompartment) GetStateArg(name string) interface{} {
	return c.StateArgs[name]
}

func (c *FizzBuzzCompartment) AddStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *FizzBuzzCompartment) SetStateVar(name string, value interface{}) {
	c.StateVars[name] = value
}

func (c *FizzBuzzCompartment) GetStateVar(name string) interface{} {
	return c.StateVars[name]
}

func (c *FizzBuzzCompartment) AddEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *FizzBuzzCompartment) SetEnterArg(name string, value interface{}) {
	c.EnterArgs[name] = value
}

func (c *FizzBuzzCompartment) GetEnterArg(name string) interface{} {
	return c.EnterArgs[name]
}

func (c *FizzBuzzCompartment) GetEnterArgs() map[string]interface{} {
	return c.EnterArgs
}

func (c *FizzBuzzCompartment) AddExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *FizzBuzzCompartment) SetExitArg(name string, value interface{}) {
	c.ExitArgs[name] = value
}

func (c *FizzBuzzCompartment) GetExitArg(name string) interface{} {
	return c.ExitArgs[name]
}

func (c *FizzBuzzCompartment) GetExitArgs() map[string]interface{} {
	return c.ExitArgs
}

/********************
// Sample Actions Implementation
package fizzbuzz

func (m *fizzBuzzStruct) print(msg string)  {}
func (m *fizzBuzzStruct) gt_100(i int) bool {}
func (m *fizzBuzzStruct) mod3_eq0(i int) bool {}
func (m *fizzBuzzStruct) mod5_eq0(i int) bool {}
func (m *fizzBuzzStruct) plus_1(i int) int {}
********************/
