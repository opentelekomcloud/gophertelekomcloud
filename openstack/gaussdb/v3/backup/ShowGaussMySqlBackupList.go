package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowGaussMySqlBackupList(client *golangsdk.ServiceClient, opts ) (*, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/mysql/v3/{project_id}/backups
	raw, err := client.Get(client.ServiceURL("backups")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res
	err = extract.Into(raw.Body, &res)
	return &res, err
}
