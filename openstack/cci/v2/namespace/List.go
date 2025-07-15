package namespace

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the corresponding query parameters.
type ListOpts struct {
	// AllowWatchBookmarks requests watch events with type "BOOKMARK"
	AllowWatchBookmarks *bool `q:"allowWatchBookmarks,omitempty"`

	// The continue option should be set when retrieving more results from the server
	Continue string `q:"continue,omitempty"`

	// A selector to restrict the list of returned objects by their fields
	FieldSelector string `q:"fieldSelector,omitempty"`

	// A selector to restrict the list of returned objects by their labels
	LabelSelector string `q:"labelSelector,omitempty"`

	// Limit is a maximum number of responses to return for a list call
	Limit *int `q:"limit,omitempty"`

	// ResourceVersion sets a constraint on what resource versions a request may be served from
	ResourceVersion string `q:"resourceVersion,omitempty"`

	// ResourceVersionMatch determines how resourceVersion is applied to list calls
	ResourceVersionMatch string `q:"resourceVersionMatch,omitempty"`

	// SendInitialEvents when set together with Watch option
	SendInitialEvents *bool `q:"sendInitialEvents,omitempty"`

	// Timeout for the list/watch call
	TimeoutSeconds *int `q:"timeoutSeconds,omitempty"`

	// Watch for changes to the described resources
	Watch *bool `q:"watch,omitempty"`

	// If 'true', then the output is pretty printed
	Pretty *bool `q:"pretty,omitempty"`
}

// List returns collection of namespaces
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Namespace, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return NamespacePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}

	return ExtractNamespaces(pages)
}

// NamespacePage is the page returned by a pager when traversing over a
// collection of namespaces.
type NamespacePage struct {
	pagination.NewSinglePageBase
}

// ExtractNamespaces accepts a Page struct, specifically a NamespacePage struct,
// and extracts the elements into a slice of Namespace structs.
func ExtractNamespaces(r pagination.NewPage) ([]Namespace, error) {
	var s struct {
		Items []Namespace `json:"items"`
	}
	err := extract.Into(bytes.NewReader((r.(NamespacePage)).Body), &s)
	return s.Items, err
}

// ListMeta represents metadata that is required for paginated list responses
type ListMeta struct {
	// Continue may be set if the user set a limit on the number of items returned
	Continue string `json:"continue,omitempty"`

	// RemainingItemCount is the number of subsequent items in the list which are not included
	RemainingItemCount *int64 `json:"remainingItemCount,omitempty"`

	// ResourceVersion is an opaque value that allows clients to determine when objects have changed
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// SelfLink is a URL representing this object
	SelfLink string `json:"selfLink,omitempty"`
}
