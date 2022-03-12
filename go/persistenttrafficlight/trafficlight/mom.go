package trafficlight

import (
	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

type MOMState uint

const (
	MOMState_new MOMState = iota
	MOMState_saving
	MOMState_persisted
	MOMState_working
	MOMState_end
)

type mOMStruct struct {
	_state_      MOMState
	trafficLight TrafficLight
	data         []byte
}

func NewMOM() *mOMStruct {
	m := &mOMStruct{}
	// Verify MOMStruct implements system interface
	var _ MOM = m
	m.trafficLight = nil
	m.data = nil
	return m
}

//===================== Interface Block ===================//

type MOM interface {
	Start()
	Stop()
	Tick()
}

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
	case MOMState_new:
		m._new_(e)
	case MOMState_saving:
		m._saving_(e)
	case MOMState_persisted:
		m._persisted_(e)
	case MOMState_working:
		m._working_(e)
	case MOMState_end:
		m._end_(e)
	}
}

//===================== Machine Block ===================//

func (m *mOMStruct) _new_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">>":
		m.trafficLight = New(m)
		m.trafficLight.Start()
		m._transition_(MOMState_saving)
		return
	}
}

func (m *mOMStruct) _saving_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.data = m.trafficLight.Marshal()
		m.trafficLight = nil
		m._transition_(MOMState_persisted)
		return
	}
}

func (m *mOMStruct) _persisted_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "tick":
		m._transition_(MOMState_working)
		return
	case "<<":
		m._transition_(MOMState_end)
		return
	}
}

func (m *mOMStruct) _working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.trafficLight = Load(m, m.data)
		m.trafficLight.Tick()
		m._transition_(MOMState_saving)
		return
	}
}

func (m *mOMStruct) _end_(e *framelang.FrameEvent) {
	switch e.Msg {
	}
}

//=============== Machinery and Mechanisms ==============//

func (m *mOMStruct) _transition_(newState MOMState) {
	m._mux_(&framelang.FrameEvent{Msg: "<"})
	m._state_ = newState
	m._mux_(&framelang.FrameEvent{Msg: ">"})
}
