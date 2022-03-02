package trafficlight

import "encoding/json"

//	"fmt"

type MOM struct {
	m    TrafficLight
	data []byte
}

func NewMOM() (*MOM, error) {
	mom := &MOM{}
	var err error
	mom.m, err = New(mom, nil)
	if err != nil {
		return nil, err
	}
	return mom, nil
}

//func (mom *MOM) Start(w http.ResponseWriter, r *http.Request) {
func (mom *MOM) Start() {
	mom.m.Start()
	mom.data, _ = json.Marshal(mom.m)
	mom.data = nil
}

func (mom *MOM) Stop() {
	mom.m.Stop()
}

func (mom *MOM) Tick() {

	json.Unmarshal(mom.data, mom.m)
	mom.m.Tick()
	mom.data, _ = json.Marshal(mom.m)
	mom.data = nil
}
