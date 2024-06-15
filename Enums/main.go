package main

import "fmt"

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var STATE = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return STATE[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unwknown State: %s", s))
	}
}
