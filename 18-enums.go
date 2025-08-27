package main

import "fmt"

//Enums have a fixed number of possible values with distinct names

//ServerState enum has int type assigned to it
type ServerState int

//Define ServerState's values as constants. iota generates successive constant values automatically
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

//Print out values of ServerState or convert to strings 
var stateName = map[ServerState]string {
	StateIdle: "Idle",
	StateConnected: "Connected",
	StateError: "Error",
	StateRetrying: "Retrying...",
	//In case of many values, automate using stringer tool and go:generate
} 

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
newStateTransition := transition(StateIdle)
fmt.Println(newStateTransition)

newSecondStateTransition := transition(newStateTransition)
fmt.Println(newSecondStateTransition)
}


//Use transition to emulate state transition for a server. It will take an existing state and return a new state

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected

	case StateConnected, StateRetrying:
		//Checks to determine the next server state
		return StateIdle
	
	case StateError:
			return StateError
	
	default:
		panic(fmt.Errorf("Unknown State: %s", s))		
	}
}