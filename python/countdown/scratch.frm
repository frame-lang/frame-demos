      i = i - 1
      i = i*5 - 1
      i = (i*5) - 1
      i = ((i*5) - 1)
      i = ((i*5) - i*1) + f(i)

#DebugBug $[i:int] >[name:"hi"]

        $S0[state_param:int]
    var state_var:int = 0

    |>| [enter_param:int]
      $[state_param] = 1
      state_param = 2
      $.state_var = 3
      state_var = 4
      ||[enter_param] = 5
      enter_param = 6
      (enter_param) -> (state_param) $S0 
      ^

--------

      Bugs

      $.missing_var = 1

--------

```
package main

import (

	"github.com/frame-lang/frame-demos/changestate/framelang"
)
```
#DebugBug $[state_param:int] >[enter_param:int]

  -machine-

   $S0[state_param:int]
    var state_var:int = 100

    |>| [enter_param:int]
      $[state_param] = 1
      state_param = 2
      $.state_var = 3
      state_var = 4
      ||[enter_param] = 5
      enter_param = 6
 
      ^

##

---------------

package main

import (
	"fmt"
	"strconv"

	"github.com/frame-lang/frame-demos/countdown/framelang"
)

func NewCountdown(i int) Countdown {
	m := &countdownStruct{}

	// Validate interfaces
	var _ Countdown = m
	var _ Countdown_actions = m
	// Create and intialize start state compartment.
	m._compartment_ = NewCountdownCompartment(CountdownState_S0)
	m._compartment_.StateArgs["i"] = i
	m._compartment_.StateVars["dec"] = 1

	// Send system start event
	e := framelang.FrameEvent{Msg: ">"}
	m._mux_(&e)
	return m
}

type CountdownState uint

const (
	CountdownState_S0 CountdownState = iota
	CountdownState_S1
	CountdownState_Stop
)

type Countdown interface {
}

type Countdown_actions interface {
	print(s *string)
}

type countdownStruct struct {
	_compartment_     *CountdownCompartment
	_nextCompartment_ *CountdownCompartment
}

//====================== Multiplexer ====================//

func (m *countdownStruct) _mux_(e *framelang.FrameEvent) {
	switch m._compartment_.State {
	case CountdownState_S0:
		m._CountdownState_S0_(e)
	case CountdownState_S1:
		m._CountdownState_S1_(e)
	case CountdownState_Stop:
		m._CountdownState_Stop_(e)
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

func (m *countdownStruct) _CountdownState_S0_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m._compartment_.StateArgs["i"] = m._compartment_.StateArgs["i"].(int) - m._compartment_.StateVars["dec"].(int)
		m.print(strconv.Itoa((m._compartment_.StateArgs["i"].(int))))
		if (m._compartment_.StateArgs["i"].(int)) == 0 {
			compartment := NewCountdownCompartment(CountdownState_Stop)
			m._transition_(compartment)
			return
		}
		compartment := NewCountdownCompartment(CountdownState_S1)
		compartment.EnterArgs["i"] = m._compartment_.StateArgs["i"].(int)
		m._transition_(compartment)
		return
	}
}

func (m *countdownStruct) _CountdownState_S1_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		compartment := NewCountdownCompartment(CountdownState_S0)
		compartment.StateArgs["i"] = e.Params["i"].(int)

		compartment.StateVars["dec"] = 1

		m._transition_(compartment)
		return
	}
}

func (m *countdownStruct) _CountdownState_Stop_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.print("done")
		return
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *countdownStruct) _transition_(compartment *CountdownCompartment) {
	m._nextCompartment_ = compartment
}

func (m *countdownStruct) _do_transition_(nextCompartment *CountdownCompartment) {
	m._mux_(&framelang.FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
	m._compartment_ = nextCompartment
	m._mux_(&framelang.FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//

func (m *countdownStruct) print(s *string) {
	fmt.Println(s)
}

/********************************************************

// Unimplemented Actions


********************************************************/

//=============== Compartment ==============//

type CountdownCompartment struct {
	State          CountdownState
	StateArgs      map[string]interface{}
	StateVars      map[string]interface{}
	EnterArgs      map[string]interface{}
	ExitArgs       map[string]interface{}
	_forwardEvent_ *framelang.FrameEvent
}

func NewCountdownCompartment(state CountdownState) *CountdownCompartment {
	c := &CountdownCompartment{State: state}
	c.StateArgs = make(map[string]interface{})
	c.StateVars = make(map[string]interface{})
	c.EnterArgs = make(map[string]interface{})
	c.ExitArgs = make(map[string]interface{})
	return c
}
