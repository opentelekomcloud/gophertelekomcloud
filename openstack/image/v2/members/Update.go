package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToImageMemberUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options to an Update request.
type UpdateOpts struct {
	Status  string `json:"status" required:"true"`
	VaultID string `json:"vault_id,omitempty"`
}

// ToImageMemberUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToImageMemberUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update function updates member.
func Update(client *golangsdk.ServiceClient, imageID string, memberID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToImageMemberUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(client.ServiceURL("images", imageID, "members", memberID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
