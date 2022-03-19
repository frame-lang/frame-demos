package main

func main() {
	m := NewHistoryStateContext()
	m.Start()
	m.SwitchState()
	m.GotoDeadEnd()
	m.Back()
}
