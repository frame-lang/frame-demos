package trafficlight

import (
	"github.com/frame-lang/frame-demos/persistenttrafficlight/framelang"
)

type MOMState uint

const (
	MOMState_new MOMState = iota
	MOMState_waiting
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
	// Verify MOMStruct implements actions interface
	//	var _ actions = m
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
	case MOMState_waiting:
		m._waiting_(e)
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
		m.trafficLight, _ = New(m, m.data)
		m.trafficLight.Start()
		m.data = m.trafficLight.Save()
		m._transition_(MOMState_waiting)
		return
	case "<":
		m.data = m.trafficLight.Save()
		m.trafficLight = nil
		return
	}
}

func (m *mOMStruct) _waiting_(e *framelang.FrameEvent) {
	switch e.Msg {
	case "tick":
		m._transition_(MOMState_working)
		return
	case "stop":
		m._transition_(MOMState_end)
		return
	}
}

func (m *mOMStruct) _working_(e *framelang.FrameEvent) {
	switch e.Msg {
	case ">":
		m.trafficLight, _ = New(m, m.data)
		m.trafficLight.Tick()
		m._transition_(MOMState_waiting)
		return
	case "<":
		m.data = m.trafficLight.Save()
		m.trafficLight = nil
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

// package trafficlight

// //	"fmt"

// type MOM struct {
// 	m    TrafficLight
// 	data []byte
// }

// func NewMOM() (*MOM, error) {
// 	mom := &MOM{}
// 	return mom, nil
// }

// //func (mom *MOM) Start(w http.ResponseWriter, r *http.Request) {
// func (mom *MOM) Start() {
// 	var err error
// 	mom.m, err = New(mom, nil)
// 	if err != nil {
// 		// TODO
// 		return
// 	}

// 	mom.m.Start()
// 	mom.data = mom.m.Save()
// 	mom.m = nil
// }

// func (mom *MOM) Stop() {
// 	var err error
// 	mom.m, err = New(mom, mom.data)
// 	if err != nil {
// 		// TODO
// 		return
// 	}
// 	mom.m.Stop()
// 	mom.m = nil
// }

// func (mom *MOM) Tick() {
// 	var err error
// 	mom.m, err = New(mom, mom.data)
// 	if err != nil {
// 		// TODO
// 		return
// 	}

// 	mom.m.Tick()
// 	mom.data = mom.m.Save()
// 	mom.m = nil
// }
