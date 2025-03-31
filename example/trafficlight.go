package main

import (
	"fmt"
	"time"

	"github.com/livensmi1e/fsm"
)

type StateGreen struct{}

func (s *StateGreen) OnEnter() { fmt.Println("🚦 Green: Go!") }
func (s *StateGreen) OnExit()  { fmt.Println("🚦 Green: Time's up!") }
func (s *StateGreen) Execute() { fmt.Println("🚦 Green: Running...") }

type StateAmber struct{}

func (s *StateAmber) OnEnter() { fmt.Println("🟡 Amber: Caution!") }
func (s *StateAmber) OnExit()  { fmt.Println("🟡 Amber: Switching...") }
func (s *StateAmber) Execute() { fmt.Println("🟡 Amber: Running...") }

type StateRed struct{}

func (s *StateRed) OnEnter() { fmt.Println("🔴 Red: Stop!") }
func (s *StateRed) OnExit()  { fmt.Println("🔴 Red: Switching...") }
func (s *StateRed) Execute() { fmt.Println("🔴 Red: Running...") }

type TimerExpire struct{}

func main() {
	green := &StateGreen{}
	red := &StateRed{}
	amber := &StateAmber{}

	timeexpire := TimerExpire{}

	fsm, err := fsm.NewMachineBuilder("traffic light").
		SetInitial(green).
		AddState(green).
		AddState(red).
		AddState(amber).
		AddTransition(green, timeexpire, amber).
		AddTransition(amber, timeexpire, red).
		AddTransition(red, timeexpire, green).
		Build()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := 0; i < 6; i++ {
		fsm.CurrentState().Execute()
		time.Sleep(1 * time.Second)
		if err := fsm.Transition(timeexpire); err != nil {
			fmt.Println("Transition Error:", err)
		}
	}
}
