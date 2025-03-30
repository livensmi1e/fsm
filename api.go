package fsm

type Event interface {
}

type State interface {
	OnEnter()
	OnExit()
	Execute()
}

type Machine interface {
	CurrentState() State
	IsInState(state State) bool
	Transition(event Event) error
	HasTransition(event Event) bool
	Err() error
}

type Transition struct {
	Event Event
	Dst   State
}

type Transitions []Transition

type StateConfig struct {
	State       State
	Transitions Transitions
}

type States []StateConfig

func NewMachine(name string, initial State, states States) Machine {
	mStates := make(map[State]state)
	for _, s := range states {
		state := state{
			transitions: make(map[Event]State),
		}
		for _, t := range s.Transitions {
			state.transitions[t.Event] = t.Dst
		}
		mStates[s.State] = state
	}
	var err error
	if _, ok := mStates[initial]; !ok {
		err = ErrInitialStateUndefined
	}
	machine := &machine{
		name:    name,
		states:  mStates,
		initial: initial,
		current: initial,
		err:     err,
	}
	return machine
}

type MachineBuilder struct {
	name     string
	initial  State
	states   map[State]state
	transMap map[State]map[Event]State
}

func NewMachineBuilder(name string) *MachineBuilder {
	return &MachineBuilder{
		name:     name,
		states:   make(map[State]state),
		transMap: make(map[State]map[Event]State),
	}
}

func (b *MachineBuilder) SetInitial(initial State) *MachineBuilder {
	b.initial = initial
	return b
}

func (b *MachineBuilder) AddState(add State) *MachineBuilder {
	b.states[add] = state{transitions: make(map[Event]State)}
	return b
}

func (b *MachineBuilder) AddTransition(src State, event Event, dst State) *MachineBuilder {
	if _, ok := b.transMap[src]; !ok {
		b.transMap[src] = make(map[Event]State)
	}
	b.transMap[src][event] = dst
	return b
}

func (b *MachineBuilder) Build() (Machine, error) {
	if _, ok := b.states[b.initial]; !ok {
		return nil, ErrInitialStateUndefined
	}
	for src, events := range b.transMap {
		if _, ok := b.states[src]; !ok {
			return nil, ErrStateUndefined
		}
		for event, dst := range events {
			if _, ok := b.states[dst]; !ok {
				return nil, ErrStateUndefined
			}
			b.states[src].transitions[event] = dst
		}
	}
	return &machine{
		name:    b.name,
		initial: b.initial,
		current: b.initial,
		states:  b.states,
	}, nil
}
