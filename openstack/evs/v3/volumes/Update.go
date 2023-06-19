package volumes

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contain options for updating an existing Volume. This object is passed
// to the volumes.Update function. For more information about the parameters, see
// the Volume object.
type UpdateOpts struct {
	VolumeId string
	// Specifies the disk name. The value can contain a maximum of 255 bytes.
	Name string `json:"name,omitempty"`
	// Specifies the disk description. The value can contain a maximum of 255 bytes.
	Description string `json:"description,omitempty"`
	// Specifies the disk metadata.
	// The length of the key or value in the metadata cannot exceed 255 bytes.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Specifies also the disk name. You can specify either parameter name or display_name.
	// If both parameters are specified, the name value is used. The value can contain a maximum of 255 bytes.
	DisplayName string `json:"display_name,omitempty"`
	// Specifies also the disk description. You can specify either parameter description or display_description.
	// If both parameters are specified, the description value is used. The value can contain a maximum of 255 bytes.
	DisplayDescription string `json:"display_description,omitempty"`
}

// Update will update the Volume with provided information. To extract the updated
// Volume from the response, call the Extract method on the UpdateResult.
func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Volume, error) {
	b, err := build.RequestBody(opts, "volume")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/volumes/{volume_id}
	raw, err := client.Put(client.ServiceURL("volumes", opts.VolumeId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}

// Attachment represents a Volume Attachment record
type Attachment struct {
	// Specifies the time when the disk was attached.
	// Time format: UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	AttachedAt time.Time `json:"-"`
	// Specifies the ID of the attachment information.
	AttachmentID string `json:"attachment_id"`
	// Specifies the device name.
	Device string `json:"device"`
	// Specifies the name of the physical host accommodating the server to which the disk is attached.
	HostName string `json:"host_name"`
	// Specifies the ID of the attached resource.
	ID string `json:"id"`
	// Specifies the ID of the server to which the disk is attached.
	ServerID string `json:"server_id"`
	// Specifies the disk ID.
	VolumeID string `json:"volume_id"`
}

// UnmarshalJSON is our unmarshalling helper
func (r *Attachment) UnmarshalJSON(b []byte) error {
	type tmp Attachment
	var s struct {
		tmp
		AttachedAt golangsdk.JSONRFC3339MilliNoZ `json:"attached_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Attachment(s.tmp)

	r.AttachedAt = time.Time(s.AttachedAt)

	return err
}
