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
	MOMState_End
)

type MOM interface {
	Start()
	Stop()
	Tick()
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
