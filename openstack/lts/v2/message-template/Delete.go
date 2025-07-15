package message_template

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	DomainId string `json:"_" required:"true"`
	// Array of names of templates to be deleted.
	TemplateNames []string `json:"template_names" required:"true"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// DELETE /v2/{project_id}/{domain_id}/lts/events/notification/templates
	_, err = client.DeleteWithBody(client.ServiceURL(opts.DomainId, "lts", "events", "notification", "templates"), b, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{204},
	})
	return
}
