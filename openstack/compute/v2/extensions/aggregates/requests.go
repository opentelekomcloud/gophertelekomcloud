package aggregates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type CreateOpts struct {
	// The name of the host aggregate.
	Name string `json:"name" required:"true"`

	// The availability zone of the host aggregate.
	// You should use a custom availability zone rather than
	// the default returned by the os-availability-zone API.
	// The availability zone must not include ‘:’ in its name.
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

func (opts CreateOpts) ToAggregatesCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "aggregate")
}

type UpdateOpts struct {
	// The name of the host aggregate.
	Name string `json:"name,omitempty"`

	// The availability zone of the host aggregate.
	// You should use a custom availability zone rather than
	// the default returned by the os-availability-zone API.
	// The availability zone must not include ‘:’ in its name.
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

func (opts UpdateOpts) ToAggregatesUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "aggregate")
}

type RemoveHostOpts struct {
	// The name of the host.
	Host string `json:"host" required:"true"`
}

func (opts RemoveHostOpts) ToAggregatesRemoveHostMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "remove_host")
}

type SetMetadataOpts struct {
	Metadata map[string]interface{} `json:"metadata" required:"true"`
}

func (opts SetMetadataOpts) ToSetMetadataMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "set_metadata")
}
