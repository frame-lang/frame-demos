package trafficlight

import "encoding/json"

//	"fmt"

type MOM struct {
	m    TrafficLight
	data []byte
}

func NewMOM() (*MOM, error) {
	mom := &MOM{}
	return mom, nil
}

//func (mom *MOM) Start(w http.ResponseWriter, r *http.Request) {
func (mom *MOM) Start() {
	var err error
	mom.m, err = New(mom, nil)
	if err != nil {
		// TODO
		return
	}
	mom.m.Start()
	mom.data, _ = json.Marshal(mom.m)
	mom.m = nil
}

func (mom *MOM) Stop() {
	var err error
	mom.m, err = New(mom, mom.data)
	if err != nil {
		// TODO
		return
	}
	mom.m.Stop()
}

func (mom *MOM) Tick() {
	var err error
	mom.m, err = New(mom, mom.data)
	if err != nil {
		// TODO
		return
	}
	//	json.Unmarshal(mom.data, mom.m)
	mom.m.Tick()
	mom.data, _ = json.Marshal(mom.m)
	mom.m = nil
}
