package groups_hcs

import (
	"log"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateGroup is a method of creating group
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToGroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	ur := client.ServiceURL("scaling_group")
	log.Printf("[DEBUG] Create URL is: %#v", ur)
	_, r.Err = client.Post(ur, b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
