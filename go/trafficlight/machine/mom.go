package machine

//	"fmt"

type MOM struct {
	m TrafficLight
}

func NewMOM() *MOM {
	mom := &MOM{}
	mom.m = New()
	return mom
}

//func (mom *MOM) Start(w http.ResponseWriter, r *http.Request) {
func (mom *MOM) Start() {
	mom.m.Start(mom)
}

func (mom *MOM) Stop() {
	mom.m.Stop()
}
