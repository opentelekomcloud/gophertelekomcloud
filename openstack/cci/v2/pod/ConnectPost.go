package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// ConnectPostOpts contains query parameters for the exec POST request
type ConnectPostOpts struct {
	// Command to execute
	Command string `q:"command,omitempty"`

	// Container in which to execute the command
	Container string `q:"container,omitempty"`

	// Redirect the standard error stream of the pod for this call
	Stderr *bool `q:"stderr,omitempty"`

	// Redirect the standard input stream of the pod for this call
	Stdin *bool `q:"stdin,omitempty"`

	// Redirect the standard output stream of the pod for this call
	Stdout *bool `q:"stdout,omitempty"`

	// TTY if true indicates that a tty will be allocated for the exec call
	TTY *bool `q:"tty,omitempty"`
}

func ConnectPost(client *golangsdk.ServiceClient, namespace string, name string, opts ConnectPostOpts) error {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", namespace, "pods", name, "exec").
		WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	// POST /apis/cci/v2/namespaces/{namespace}/pods/{name}/exec
	_, err = client.Post(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
