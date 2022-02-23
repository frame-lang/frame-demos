package trafficlight

type trafficLightActions struct{}

func (m *trafficLightActions) enterRed()                {}
func (m *trafficLightActions) enterGreen()              {}
func (m *trafficLightActions) enterYellow()             {}
func (m *trafficLightActions) enterFlashingRed()        {}
func (m *trafficLightActions) exitFlashingRed()         {}
func (m *trafficLightActions) startWorkingTimer()       {}
func (m *trafficLightActions) stopWorkingTimer()        {}
func (m *trafficLightActions) startFlashingTimer()      {}
func (m *trafficLightActions) stopFlashingTimer()       {}
func (m *trafficLightActions) changeColor(color string) {}
func (m *trafficLightActions) startFlashing()           {}
func (m *trafficLightActions) stopFlashing()            {}
func (m *trafficLightActions) changeFlashingAnimation() {}
func (m *trafficLightActions) log(msg string)           {}
