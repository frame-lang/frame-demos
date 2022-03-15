package trafficlight

import (
	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

type MOMState uint

const (
	MOMState_New MOMState = iota
	MOMState_Saving
	MOMState_Persisted
	MOMState_Working
	MOMState_TrafficLightApi
	MOMState_End
)

type MOM interface {
	Start()
	Stop()
	Tick()
	EnterRed()
	EnterGreen()
	EnterYellow()
	EnterFlashingRed()
	ExitFlashingRed()
	StartWorkingTimer()
	StopWorkingTimer()
	StartFlashingTimer()
	StopFlashingTimer()
	ChangeColor(color string)
	StartFlashing()
	StopFlashing()
	ChangeFlashingAnimation()
	Log(msg string)
}

type MOM_actions interface {
	enterRed()
	enterGreen()
	enterYellow()
	enterFlashingRed()
	exitFlashingRed()
	startWorkingTimer()
	stopWorkingTimer()
	startFlashingTimer()
	stopFlashingTimer()
	changeColor(color string)
	startFlashing()
	stopFlashing()
	changeFlashingAnimation()
	log(msg string)
}

type mOMStruct struct {
	_state_      MOMState
	trafficLight TrafficLight
	data         []byte
}

func NewMOM() MOM {
	m := &mOMStruct{}

	// Validate interfaces
	var _ MOM = m
	var _ MOM_actions = m

	// Initialize domain
	m.trafficLight = nil
	m.data = nil

	return m
}

//===================== Interface Block ===================//

func (m *mOMStruct) Start() {
	e := framelang.FrameEvent{Msg: ">>"}
	m._mux_(&e)
}

func (m *mOMStruct) Stop() {
	e := framelang.FrameEvent{Msg: "<<"}
	m._mux_(&e)
}

func (m *mOMStruct) Tick() {
	e := framelang.FrameEvent{Msg: "tick"}
	m._mux_(&e)
}

func (m *mOMStruct) EnterRed() {
	e := framelang.FrameEvent{Msg: "enterRed"}
	m._mux_(&e)
}

func (m *mOMStruct) EnterGreen() {
	e := framelang.FrameEvent{Msg: "enterGreen"}
	m._mux_(&e)
}

func (m *mOMStruct) EnterYellow() {
	e := framelang.FrameEvent{Msg: "enterYellow"}
	m._mux_(&e)
}

func (m *mOMStruct) EnterFlashingRed() {
	e := framelang.FrameEvent{Msg: "enterFlashingRed"}
	m._mux_(&e)
}

func (m *mOMStruct) ExitFlashingRed() {
	e := framelang.FrameEvent{Msg: "exitFlashingRed"}
	m._mux_(&e)
}

func (m *mOMStruct) StartWorkingTimer() {
	e := framelang.FrameEvent{Msg: "startWorkingTimer"}
	m._mux_(&e)
}

func (m *mOMStruct) StopWorkingTimer() {
	e := framelang.FrameEvent{Msg: "stopWorkingTimer"}
	m._mux_(&e)
}

func (m *mOMStruct) StartFlashingTimer() {
	e := framelang.FrameEvent{Msg: "startFlashingTimer"}
	m._mux_(&e)
}

func (m *mOMStruct) StopFlashingTimer() {
	e := framelang.FrameEvent{Msg: "stopFlashingTimer"}
	m._mux_(&e)
}

func (m *mOMStruct) ChangeColor(color string) {
	params := make(map[string]interface{})
	params["color"] = color
	e := framelang.FrameEvent{Msg: "changeColor", Params: params}
	m._mux_(&e)
}

func (m *mOMStruct) StartFlashing() {
	e := framelang.FrameEvent{Msg: "startFlashing"}
	m._mux_(&e)
}

func (m *mOMStruct) StopFlashing() {
	e := framelang.FrameEvent{Msg: "stopFlashing"}
	m._mux_(&e)
}

func (m *mOMStruct) ChangeFlashingAnimation() {
	e := framelang.FrameEvent{Msg: "changeFlashingAnimation"}
	m._mux_(&e)
}

func (m *mOMStruct) Log(msg string) {
	params := make(map[string]interface{})
	params["msg"] = msg
	e := framelang.FrameEvent{Msg: "log", Params: params}
	m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *mOMStruct) _mux_(e *framelang.FrameEvent) {
	switch m._state_ {
	case MOMState_New:
		m._MOMState_New_(e)
	case MOMState_Saving:
		m._MOMState_Saving_(e)
	case MOMState_Persisted:
		m._MOMState_Persisted_(e)
	case MOMState_Working:
		m._MOMState_Working_(e)
	case MOMState_TrafficLightApi:
		m._MOMState_TrafficLightApi_(e)
	case MOMState_End:
		m._MOMState_End_(e)
	}
}

//===================== Machine Block ===================//

func (m *mOMStruct) _MOMState_New_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.trafficLight = NewTrafficLight(m)
		m.trafficLight.Start()
		// Traffic Light\nStarted
		m._transition_(MOMState_Saving)
		return
	}
}

func (m *mOMStruct) _MOMState_Saving_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.data = m.trafficLight.Marshal()
		m.trafficLight = nil
		// Saved
		m._transition_(MOMState_Persisted)
		return
	}
}

func (m *mOMStruct) _MOMState_Persisted_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "tick":
		// Tick
		m._transition_(MOMState_Working)
		return
	case "<<":
		// Stop
		m._transition_(MOMState_End)
		return
	}
}

func (m *mOMStruct) _MOMState_Working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.trafficLight = LoadTrafficLight(m, m.data)
		m.trafficLight.Tick()
		// Done
		m._transition_(MOMState_Saving)
		return
	}
	m._MOMState_TrafficLightApi_(e)

}

func (m *mOMStruct) _MOMState_TrafficLightApi_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "enterRed":
		m.enterRed()
		return
	case "enterGreen":
		m.enterGreen()
		return
	case "enterYellow":
		m.enterYellow()
		return
	case "enterFlashingRed":
		m.enterFlashingRed()
		return
	case "exitFlashingRed":
		m.exitFlashingRed()
		return
	case "startWorkingTimer":
		m.startWorkingTimer()
		return
	case "stopWorkingTimer":
		m.stopWorkingTimer()
		return
	case "startFlashingTimer":
		m.startFlashingTimer()
		return
	case "stopFlashingTimer":
		m.stopFlashingTimer()
		return
	case "changeColor":
		m.changeColor(e.Params["color"].(string))
		return
	case "startFlashing":
		m.startFlashing()
		return
	case "stopFlashing":
		m.stopFlashing()
		return
	case "changeFlashingAnimation":
		m.changeFlashingAnimation()
		return
	case "log":
		m.log(e.Params["msg"].(string))
		return
	}
}

func (m *mOMStruct) _MOMState_End_(e *framelang.FrameEvent) {
	switch e.Msg {
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *mOMStruct) _transition_(newState MOMState) {
	m._mux_(&framelang.FrameEvent{Msg: "<"})
	m._state_ = newState
	m._mux_(&framelang.FrameEvent{Msg: ">"})
}

/********************
// Sample Actions Implementation
package mom

func (m *mOMStruct) enterRed()  {}
func (m *mOMStruct) enterGreen()  {}
func (m *mOMStruct) enterYellow()  {}
func (m *mOMStruct) enterFlashingRed()  {}
func (m *mOMStruct) exitFlashingRed()  {}
func (m *mOMStruct) startWorkingTimer()  {}
func (m *mOMStruct) stopWorkingTimer()  {}
func (m *mOMStruct) startFlashingTimer()  {}
func (m *mOMStruct) stopFlashingTimer()  {}
func (m *mOMStruct) changeColor(color string)  {}
func (m *mOMStruct) startFlashing()  {}
func (m *mOMStruct) stopFlashing()  {}
func (m *mOMStruct) changeFlashingAnimation()  {}
func (m *mOMStruct) log(msg string)  {}
********************/
