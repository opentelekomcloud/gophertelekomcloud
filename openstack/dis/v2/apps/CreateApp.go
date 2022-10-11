package apps

type CreateAppRequest struct {
	Body *CreateAppRequestBody `json:"body,omitempty"`
}
type CreateAppRequestBody struct {
	// Unique identifier of the consumer application to be created.
	// The application name contains 1 to 200 characters, including letters, digits, underscores (_), and hyphens (-).
	// Minimum: 1
	// Maximum: 200
	AppName string `json:"app_name"`
}

// POST /v2/{project_id}/apps

type CreateAppResponse struct {
	HttpStatusCode int `json:"-"`
}
