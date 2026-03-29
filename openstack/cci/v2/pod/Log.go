package pod

import (
	"io"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// LogOpts contains the options for reading pod logs
type LogOpts struct {
	// Namespace of the pod
	Namespace string `json:"-" required:"true"`

	// Name of the pod
	Name string `json:"-" required:"true"`

	// Container for which to stream logs. Defaults to only container if there is one container in the pod.
	Container string `q:"container,omitempty"`

	// Follow the log stream of the pod
	Follow *bool `q:"follow,omitempty"`

	// InsecureSkipTLSVerifyBackend indicates skipping TLS verification of the backend
	InsecureSkipTLSVerifyBackend *bool `q:"insecureSkipTLSVerifyBackend,omitempty"`

	// LimitBytes limits the number of bytes of log output
	LimitBytes *int `q:"limitBytes,omitempty"`

	// Pretty if 'true', then the output is pretty printed
	Pretty string `q:"pretty,omitempty"`

	// Previous return previous terminated container logs
	Previous *bool `q:"previous,omitempty"`

	// SinceSeconds is a relative time in seconds before the current time from which to show logs
	SinceSeconds *int `q:"sinceSeconds,omitempty"`

	// TailLines is the number of lines from the end of the logs to show
	TailLines *int `q:"tailLines,omitempty"`

	// Timestamps adds an RFC3339 timestamp at the beginning of every line of log output
	Timestamps *bool `q:"timestamps,omitempty"`
}

// Log reads logs of the specified pod
func Log(client *golangsdk.ServiceClient, opts LogOpts) (string, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", opts.Namespace, "pods", opts.Name, "log").
		WithQueryParams(&opts).Build()
	if err != nil {
		return "", err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(raw.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
