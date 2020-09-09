package lockunlock

import "github.com/opentelekomcloud/gophertelekomcloud"

func actionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

// Lock is the operation responsible for locking a Compute server.
func Lock(client *golangsdk.ServiceClient, id string) (r LockResult) {
	_, r.Err = client.Post(actionURL(client, id), map[string]interface{}{"lock": nil}, nil, nil)
	return
}

// Unlock is the operation responsible for unlocking a Compute server.
func Unlock(client *golangsdk.ServiceClient, id string) (r UnlockResult) {
	_, r.Err = client.Post(actionURL(client, id), map[string]interface{}{"unlock": nil}, nil, nil)
	return
}
