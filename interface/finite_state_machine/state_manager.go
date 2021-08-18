package main

import (
	"errors"
	"fmt"
)

type StateManager struct {
	stateList    map[string]State
	OnChange     func(from, to State)
	currState    State
}

var (
	ErrStateNotFound = errors.New("state not found")
	ErrCanNotTransit = errors.New("state cat not transit")
)

func (sm *StateManager) Add (s State) {
	name := StateName(s)
	if sm.Get(name) != nil {
		fmt.Println("state existed")
	} else {
		sm.stateList[name] = s
	}
}

func (sm *StateManager) Get (name string) State {
	if v, ok := sm.stateList[name]; ok {
		return v
	}
	return nil
}

func (sm *StateManager) CanTransitTo (name string) bool {
	if sm.currState == nil {
		return true
	}
	if sm.currState.Name() == name {
		fmt.Println(ErrCanNotTransit)
		return false
	}
	return sm.currState.CanTransitTo(name)
}

func (sm *StateManager) Transit (name string) error {
	next := sm.Get(name)
	if next == nil {
		return ErrStateNotFound
	}
	pre := sm.currState
	if sm.currState != nil {
		if !sm.CanTransitTo(name) {
			return ErrCanNotTransit
		}
		sm.currState.OnEnd()
	}
	sm.currState = next
	sm.currState.OnBegin()
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.currState)
	}
	return nil
}

func NewStateManager () *StateManager {
	return &StateManager{
		stateList: make(map[string]State),
		OnChange: StateChange,
	}
}

func StateChange (from, to State) {
	fromState := ""
	if from != nil {
		fromState = from.Name()
	}
	fmt.Printf("%s -----> %s\n", fromState, to.Name())
}
