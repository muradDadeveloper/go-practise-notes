package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetruing
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetruing:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	sn2 := transition(ns)
	fmt.Println(sn2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected

	case StateConnected, StateRetruing:
		return StateIdle

	case StateError:
		return StateError

	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
