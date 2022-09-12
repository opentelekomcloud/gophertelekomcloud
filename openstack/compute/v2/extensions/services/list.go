package services

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List makes a request against the API to list services.
func List(client *golangsdk.ServiceClient) ([]Service, error) {
	raw, err := client.Get(client.ServiceURL("os-services"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Service
	err = extract.IntoSlicePtr(raw.Body, &res, "services")
	return res, err
}

// Service represents a Compute service in the OpenStack cloud.
type Service struct {
	// The binary name of the service.
	Binary string `json:"binary"`
	// The reason for disabling a service.
	DisabledReason string `json:"disabled_reason"`
	// The name of the host.
	Host string `json:"host"`
	// The id of the service.
	ID int `json:"id"`
	// The state of the service. One of up or down.
	State string `json:"state"`
	// The status of the service. One of enabled or disabled.
	Status string `json:"status"`
	// The date and time when the resource was updated.
	UpdatedAt time.Time `json:"-"`
	// The availability zone name.
	Zone string `json:"zone"`
}

// UnmarshalJSON to override default
func (r *Service) UnmarshalJSON(b []byte) error {
	type tmp Service
	var s struct {
		tmp
		UpdatedAt golangsdk.JSONRFC3339MilliNoZ `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Service(s.tmp)

	r.UpdatedAt = time.Time(s.UpdatedAt)

	return nil
}
