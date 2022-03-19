package main

/// This demo exercises the history mechanism with both basic state stack as well
/// as a compartment stack.

func main() {
	m_basic := NewHistoryBasic()
	m_basic.Start()
	m_basic.SwitchState()
	m_basic.GotoDeadEnd()
	m_basic.Back()

	m_compartment := NewHistoryStateContext()
	m_compartment.Start()
	m_compartment.SwitchState()
	m_compartment.GotoDeadEnd()
	m_compartment.Back()
}
