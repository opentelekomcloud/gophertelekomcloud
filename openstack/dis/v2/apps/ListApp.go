package apps

type ListAppOpts struct {
	// Maximum number of apps to list in a single API call. Value range: 1-100 Default value: 10
	// Minimum: 1
	// Maximum: 1000
	// Default: 10
	Limit *int32 `json:"limit,omitempty"`
	// Name of the app to start the list with. The returned app list does not contain this app name.
	StartAppName string `json:"start_app_name,omitempty"`
	// Name of the stream whose apps will be returned.
	StreamName string `json:"stream_name,omitempty"`
}

// GET /v2/{project_id}/apps

type ListAppResponse struct {
	// Specifies whether there are more matching consumer applications to list.
	// true: yes
	// false: no
	HasMoreApp *bool `json:"has_more_app,omitempty"`
	// AppEntry list that meets the current request.
	Apps []DescribeAppResult `json:"apps,omitempty"`
	// Total number of apps that meet criteria.
	TotalNumber *int32 `json:"total_number,omitempty"`
}
type DescribeAppResult struct {
	// Name of the app.
	AppName string `json:"app_name,omitempty"`
	// Unique identifier of the app.
	AppId string `json:"app_id,omitempty"`
	// Time when the app is created, in milliseconds.
	CreateTime *int64 `json:"create_time,omitempty"`
	// List of associated streams.
	CommitCheckPointStreamNames []string `json:"commit_checkpoint_stream_names,omitempty"`
}
