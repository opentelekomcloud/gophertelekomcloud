package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, templateId string) (err error) {
	// DELETE /v2/{project_id}/notifications/message_template/{message_template_id}
	_, err = client.Delete(client.ServiceURL("message_template", templateId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}
