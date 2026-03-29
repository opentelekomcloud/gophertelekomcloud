package secret

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through the API
type ListOpts struct {
	// AllowWatchBookmarks requests watch events with type "BOOKMARK"
	AllowWatchBookmarks *bool `q:"allowWatchBookmarks,omitempty"`

	// Continue option should be set when retrieving more results from the server
	Continue string `q:"continue,omitempty"`

	// FieldSelector is a selector to restrict the list of returned objects by their fields
	FieldSelector string `q:"fieldSelector,omitempty"`

	// LabelSelector is a selector to restrict the list of returned objects by their labels
	LabelSelector string `q:"labelSelector,omitempty"`

	// Limit is a maximum number of responses to return for a list call
	Limit *int `q:"limit,omitempty"`

	// ResourceVersion sets a constraint on what resource versions a request may be served from
	ResourceVersion string `q:"resourceVersion,omitempty"`

	// ResourceVersionMatch determines how resourceVersion is applied to list calls
	ResourceVersionMatch string `q:"resourceVersionMatch,omitempty"`

	// SendInitialEvents when set together with Watch option
	SendInitialEvents *bool `q:"sendInitialEvents,omitempty"`

	// TimeoutSeconds is timeout for the list/watch call
	TimeoutSeconds *int `q:"timeoutSeconds,omitempty"`

	// Watch for changes to the described resources
	Watch *bool `q:"watch,omitempty"`

	// Pretty if 'true', then the output is pretty printed
	Pretty string `q:"pretty,omitempty"`
}

// List returns collection of secrets in a namespace
func List(client *golangsdk.ServiceClient, namespace string, opts ListOpts) ([]Secret, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", namespace, "secrets").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SecretPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}

	return ExtractSecrets(pages)
}

// SecretPage is the page returned by a pager when traversing over a collection of secrets
type SecretPage struct {
	pagination.NewSinglePageBase
}

// ExtractSecrets accepts a Page struct and extracts the elements into a slice of Secret structs
func ExtractSecrets(r pagination.NewPage) ([]Secret, error) {
	var s struct {
		Items []Secret `json:"items"`
	}
	err := extract.Into(bytes.NewReader((r.(SecretPage)).Body), &s)
	return s.Items, err
}
