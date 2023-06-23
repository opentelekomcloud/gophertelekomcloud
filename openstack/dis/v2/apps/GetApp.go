package apps

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetApp(client *golangsdk.ServiceClient, appName string) (*GetAppResponse, error) {
	// GET /v2/{project_id}/apps/{app_name}
	raw, err := client.Get(client.ServiceURL("apps", appName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetAppResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetAppResponse struct {
	// Name of the app.
	AppName string `json:"app_name,omitempty"`
	// Unique identifier of the app.
	AppId string `json:"app_id,omitempty"`
	// Time when the app is created, in milliseconds.
	CreateTime *int64 `json:"create_time,omitempty"`
	// List of associated streams.
	CommitCheckPointStreamNames []string `json:"commit_checkpoint_stream_names,omitempty"`
}
