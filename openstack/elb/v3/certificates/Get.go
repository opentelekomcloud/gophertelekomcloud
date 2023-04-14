package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("certificates", id), &r.Body, nil)
	return
}

type Certificate struct {
	//
	ID string `json:"id"`
	//
	ProjectID string `json:"project_id"`
	//
	Name string `json:"name"`
	//
	Description string `json:"description"`
	//
	Type string `json:"type"`
	//
	Domain string `json:"domain"`
	//
	PrivateKey string `json:"private_key"`
	//
	Certificate string `json:"certificate"`
	//
	AdminStateUp bool `json:"admin_state_up"`
	//
	CreatedAt string `json:"created_at"`
	//
	UpdatedAt string `json:"updated_at"`
	//
	ExpireTime string `json:"expire_time"`
}
