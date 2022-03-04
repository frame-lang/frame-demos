package trafficlight

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
	mom.data = mom.m.Save()
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
	mom.m = nil
}

func (mom *MOM) Tick() {
	var err error
	mom.m, err = New(mom, mom.data)
	if err != nil {
		// TODO
		return
	}

	mom.m.Tick()
	mom.data = mom.m.Save()
	mom.m = nil
}
