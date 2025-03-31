package fsm

type state struct {
	transitions map[Event]State
}

type machine struct {
	name    string
	current State
	initial State
	states  map[State]state
	err     error
}

func (m *machine) CurrentState() State {
	return m.current
}

func (m *machine) IsInState(state State) bool {
	return m.current == state
}

func (m *machine) Transition(event Event) error {
	if m.err != nil {
		return m.err
	}
	next := m.getNextState(event)
	m.current.OnExit()
	m.current = next
	m.current.OnEnter()
	return nil
}

func (m *machine) HasTransition(event Event) bool {
	if _, ok := m.states[m.current].transitions[event]; ok {
		return true
	}
	return false
}

func (m *machine) Err() error {
	return m.err
}

func (m *machine) getNextState(event Event) State {
	if m.err != nil {
		return nil
	}
	next, ok := m.states[m.current].transitions[event]
	if !ok {
		m.err = ErrEventDeclined
		return nil
	}
	return next
}
