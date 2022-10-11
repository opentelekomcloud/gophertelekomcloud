package apps

type DescribeAppRequest struct {
	// Name of the app to be queried.
	AppName string `json:"app_name"`
}

// GET /v2/{project_id}/apps/{app_name}

type DescribeAppResponse struct {
	// Name of the app.
	AppName *string `json:"app_name,omitempty"`
	// Unique identifier of the app.
	AppId *string `json:"app_id,omitempty"`
	// Time when the app is created, in milliseconds.
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}
