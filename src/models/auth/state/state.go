package state

import "encoding/json"

type State struct {
	Redirect string `json:"redirect"`
}

func ParseState(stateStr string) *State {
	var state State;
	error := json.Unmarshal([]byte(stateStr), &state)
	if error != nil {
		return nil
	}
	return &state
}

func StringifyState(state *State) *[]byte {
	value, error := json.Marshal(state)
	if error != nil {
		return nil
	}
	return &value
}