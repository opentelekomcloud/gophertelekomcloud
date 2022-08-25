package services

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Filter the service list result by binary name of the service.
	Binary string `q:"binary"`
	// Filter the service list result by host name of the service.
	Host string `q:"host"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("os-services")+query.String(), func(r pagination.PageResult) pagination.Page {
		return ServicePage{pagination.SinglePageBase(r)}
	})
}

type ServicePage struct {
	pagination.SinglePageBase
}

func (page ServicePage) IsEmpty() (bool, error) {
	services, err := ExtractServices(page)
	return len(services) == 0, err
}

func ExtractServices(r pagination.Page) ([]Service, error) {
	var res struct {
		Service []Service `json:"services"`
	}
	err := extract.Into(r.(ServicePage).Body, &res)
	return res.Service, err
}

type Service struct {
	// The binary name of the service.
	Binary string `json:"binary"`
	// The reason for disabling a service.
	DisabledReason string `json:"disabled_reason"`
	// The name of the host.
	Host string `json:"host"`
	// The state of the service. One of up or down.
	State string `json:"state"`
	// The status of the service. One of available or unavailable.
	Status string `json:"status"`
	// The date and time stamp when the extension was last updated.
	UpdatedAt time.Time `json:"-"`
	// The availability zone name.
	Zone string `json:"zone"`
	// The following fields are optional
	// The host is frozen or not. Only in cinder-volume service.
	Frozen bool `json:"frozen"`
	// The cluster name. Only in cinder-volume service.
	Cluster string `json:"cluster"`
	// The volume service replication status. Only in cinder-volume service.
	ReplicationStatus string `json:"replication_status"`
	// The ID of active storage backend. Only in cinder-volume service.
	ActiveBackendID string `json:"active_backend_id"`
}

func (r *Service) UnmarshalJSON(b []byte) error {
	type tmp Service
	var res struct {
		tmp
		UpdatedAt golangsdk.JSONRFC3339MilliNoZ `json:"updated_at"`
	}
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	*r = Service(res.tmp)

	r.UpdatedAt = time.Time(res.UpdatedAt)

	return nil
}
