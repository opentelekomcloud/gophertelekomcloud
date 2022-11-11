package apps

type DeleteAppOpts struct {
	// Name of the app to be deleted.
	AppName string `json:"app_name"`
}

// DELETE /v2/{project_id}/apps/{app_name}

type DeleteAppResponse struct {
}
