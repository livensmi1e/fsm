package fsm

import (
	"testing"
)

type testState string

func (s testState) OnEnter() {}
func (s testState) OnExit()  {}
func (s testState) Execute() {}

type testEvent string

func TestFSM_Declarative(t *testing.T) {
	const (
		StateA testState = "A"
		StateB testState = "B"
		EventX testEvent = "X"
	)

	machine := NewMachine("TestFSM", StateA, States{
		{
			State: StateA,
			Transitions: Transitions{
				{EventX, StateB},
			},
		},
		{
			State:       StateB,
			Transitions: Transitions{},
		},
	})

	if machine.CurrentState() != StateA {
		t.Errorf("Expected initial state %v, got %v", StateA, machine.CurrentState())
	}

	if !machine.HasTransition(EventX) {
		t.Errorf("Expected transition %v to exist", EventX)
	}

	if err := machine.Transition(EventX); err != nil {
		t.Errorf("Expected successful transition, got error: %v", err)
	}

	if machine.CurrentState() != StateB {
		t.Errorf("Expected state %v, got %v", StateB, machine.CurrentState())
	}
}

func TestFSM_Builder(t *testing.T) {
	const (
		StateA testState = "A"
		StateB testState = "B"
		EventX testEvent = "X"
	)

	builder := NewMachineBuilder("TestFSM").
		SetInitial(StateA).
		AddState(StateA).
		AddState(StateB).
		AddTransition(StateA, EventX, StateB)

	machine, err := builder.Build()
	if err != nil {
		t.Fatalf("Unexpected error in Build(): %v", err)
	}

	if machine.CurrentState() != StateA {
		t.Errorf("Expected initial state %v, got %v", StateA, machine.CurrentState())
	}

	if !machine.HasTransition(EventX) {
		t.Errorf("Expected transition %v to exist", EventX)
	}

	if err := machine.Transition(EventX); err != nil {
		t.Errorf("Expected successful transition, got error: %v", err)
	}

	if machine.CurrentState() != StateB {
		t.Errorf("Expected state %v, got %v", StateB, machine.CurrentState())
	}
}
