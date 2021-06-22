package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

/*
	Create member for specific image

	Preconditions

	* The specified images must exist.
	* You can only add a new member to an image which 'visibility' attribute is
		private.
	* You must be the owner of the specified image.

	Synchronous Postconditions

	With correct permissions, you can see the member status of the image as
	pending through API calls.

	More details here:
	http://developer.openstack.org/api-ref-image-v2.html#createImageMember-v2
*/
// CreateOptsBuilder allows extensions to add parameters to the Create request.
type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Member string `json:"member" required:"true"`
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Create(client *golangsdk.ServiceClient, id string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createMemberURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// List members returns list of members for specified image id.
func List(client *golangsdk.ServiceClient, id string) pagination.Pager {
	return pagination.NewPager(client, listMembersURL(client, id), func(r pagination.PageResult) pagination.Page {
		return MemberPage{pagination.SinglePageBase(r)}
	})
}

// Get image member details.
func Get(client *golangsdk.ServiceClient, imageID string, memberID string) (r DetailsResult) {
	_, r.Err = client.Get(getMemberURL(client, imageID, memberID), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete membership for given image. Callee should be image owner.
func Delete(client *golangsdk.ServiceClient, imageID string, memberID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteMemberURL(client, imageID, memberID), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options to an Update request.
type UpdateOpts struct {
	Status  string `json:"status" required:"true"`
	VaultID string `json:"vault_id,omitempty"`
}

// ToUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update function updates member.
func Update(client *golangsdk.ServiceClient, imageID string, memberID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateMemberURL(client, imageID, memberID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
