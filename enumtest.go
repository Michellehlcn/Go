// import "fmt"
package main

import "fmt"

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Println("Unknown state"))
	}

	return StateConnected
}
