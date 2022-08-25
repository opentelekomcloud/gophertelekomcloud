package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeTypeCreateMap() (map[string]interface{}, error)
}

// CreateOpts are options for creating a volume type.
type CreateOpts struct {
	// See VolumeType.
	ExtraSpecs map[string]interface{} `json:"extra_specs,omitempty"`
	// See VolumeType.
	Name string `json:"name,omitempty"`
}

// ToVolumeTypeCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToVolumeTypeCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "volume_type")
}

// Create will create a new volume. To extract the created volume type object,
// call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeTypeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("types"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

// Delete will delete the volume type with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(client.ServiceURL("types", id), nil)
	return
}

// Get will retrieve the volume type with the provided ID. To extract the volume
// type from the result, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("types", id), nil, nil)
	return
}

// List returns all volume types.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, client.ServiceURL("types"), func(r pagination.PageResult) pagination.Page {
		return VolumeTypePage{pagination.SinglePageBase(r)}
	})
}
