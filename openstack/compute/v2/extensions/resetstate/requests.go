package resetstate

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// ServerState refers to the states usable in ResetState Action
type ServerState string

const (
	// StateActive returns the state of the server as active
	StateActive ServerState = "active"

	// StateError returns the state of the server as error
	StateError ServerState = "error"
)

// ResetState will reset the state of a server
func ResetState(client *golangsdk.ServiceClient, id string, state ServerState) (r ResetResult) {
	stateMap := map[string]interface{}{"state": state}
	_, r.Err = client.Post(actionURL(client, id), map[string]interface{}{"os-resetState": stateMap}, nil, nil)
	return
}
