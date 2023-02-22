package images

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Update implements image updated request.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToImageUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Patch(client.ServiceURL("images", id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/openstack-images-v2.1-json-patch"},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	// returns value implementing json.Marshaler which when marshaled matches
	// the patch schema:
	// http://specs.openstack.org/openstack/glance-specs/specs/api/v2/http-patch-image-api-v2.html
	ToImageUpdateMap() ([]interface{}, error)
}

// UpdateOpts implements UpdateOpts
type UpdateOpts []Patch

// ToImageUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToImageUpdateMap() ([]interface{}, error) {
	m := make([]interface{}, len(opts))
	for i, patch := range opts {
		patchJSON := patch.ToImagePatchMap()
		m[i] = patchJSON
	}
	return m, nil
}

// Patch represents a single update to an existing image. Multiple updates
// to an image can be submitted at the same time.
type Patch interface {
	ToImagePatchMap() map[string]interface{}
}

// UpdateVisibility represents an updated visibility property request.
type UpdateVisibility struct {
	Visibility ImageVisibility
}

// ToImagePatchMap assembles a request body based on UpdateVisibility.
func (r UpdateVisibility) ToImagePatchMap() map[string]interface{} {
	return map[string]interface{}{
		"op":    "replace",
		"path":  "/visibility",
		"value": r.Visibility,
	}
}

// ReplaceImageName represents an updated image_name property request.
type ReplaceImageName struct {
	NewName string
}

// ToImagePatchMap assembles a request body based on ReplaceImageName.
func (r ReplaceImageName) ToImagePatchMap() map[string]interface{} {
	return map[string]interface{}{
		"op":    "replace",
		"path":  "/name",
		"value": r.NewName,
	}
}

// ReplaceImageChecksum represents an updated checksum property request.
type ReplaceImageChecksum struct {
	Checksum string
}

// ReplaceImageChecksum assembles a request body based on ReplaceImageChecksum.
func (r ReplaceImageChecksum) ToImagePatchMap() map[string]interface{} {
	return map[string]interface{}{
		"op":    "replace",
		"path":  "/checksum",
		"value": r.Checksum,
	}
}

// ReplaceImageTags represents an updated tags property request.
type ReplaceImageTags struct {
	NewTags []string
}

// ToImagePatchMap assembles a request body based on ReplaceImageTags.
func (r ReplaceImageTags) ToImagePatchMap() map[string]interface{} {
	return map[string]interface{}{
		"op":    "replace",
		"path":  "/tags",
		"value": r.NewTags,
	}
}

// UpdateOp represents a valid update operation.
type UpdateOp string

const (
	AddOp     UpdateOp = "add"
	ReplaceOp UpdateOp = "replace"
	RemoveOp  UpdateOp = "remove"
)

// UpdateImageProperty represents an update property request.
type UpdateImageProperty struct {
	Op    UpdateOp
	Name  string
	Value string
}

// ToImagePatchMap assembles a request body based on UpdateImageProperty.
func (r UpdateImageProperty) ToImagePatchMap() map[string]interface{} {
	updateMap := map[string]interface{}{
		"op":   r.Op,
		"path": fmt.Sprintf("/%s", r.Name),
	}

	if r.Value != "" {
		updateMap["value"] = r.Value
	}

	return updateMap
}
