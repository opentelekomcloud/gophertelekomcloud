package pod

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// ConnectGetOpts contains query parameters for the exec GET request
type ConnectGetOpts struct {
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

func ConnectGet(client *golangsdk.ServiceClient, namespace string, name string, opts ConnectGetOpts) error {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", namespace, "pods", name, "exec").
		WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	// GET /apis/cci/v2/namespaces/{namespace}/pods/{name}/exec
	_, err = client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
