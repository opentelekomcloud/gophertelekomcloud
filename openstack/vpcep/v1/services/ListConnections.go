package services

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListConnectionsOpts struct {
	// Specifies the unique ID of the VPC endpoint.
	ID string `q:"id"`
	// Specifies the packet ID of the VPC endpoint.
	MarkerId string `q:"marker_id"`
	// Specifies the connection status of the VPC endpoint.
	// pendingAcceptance: The VPC endpoint is to be accepted.
	// accepted: The VPC endpoint has been accepted.
	// rejected: The VPC endpoint has been rejected.
	// failed: The VPC endpoint service failed to be created.
	Status Status `q:"status"`
	// Specifies the sorting field of the VPC endpoint list. The field can be:
	// created_at: VPC endpoints are sorted by creation time.
	// updated_at: VPC endpoints are sorted by update time.
	SortKey string `q:"sort_key"`
	// Specifies the sorting method of the VPC endpoint list. The method can be:
	// desc: VPC endpoints are sorted in descending order.
	// asc: VPC endpoints are sorted in ascending order.
	SortDir string `q:"sort_dir"`
	// Specifies the maximum number of VPC endpoint services displayed on each page.
	Limit *int `q:"limit"`
	// Specifies the offset.
	Offset *int `q:"offset"`
}

func ListConnections(client *golangsdk.ServiceClient, ServiceID string, opts ListConnectionsOpts) ([]Connection, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpc-endpoint-services", ServiceID, "connections").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ServiceConnectionsPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractConnections(pages)
}

type ServiceConnectionsPage struct {
	pagination.NewSinglePageBase
}

func ExtractConnections(r pagination.NewPage) ([]Connection, error) {
	var s struct {
		Connections []Connection `json:"connections"`
	}
	err := extract.Into(bytes.NewReader((r.(ServiceConnectionsPage)).Body), &s)
	return s.Connections, err
}

type Connection struct {
	// Specifies the unique ID of the VPC endpoint.
	ID string `json:"id"`
	// Specifies the packet ID of the VPC endpoint.
	MarkerId int `json:"marker_id"`
	// Specifies the creation time of the VPC endpoint.
	CreatedAt string `json:"created_at"`
	// Specifies the update time of the VPC endpoint.
	UpdatedAt string `json:"updated_at"`
	// Specifies the user's domain ID.
	DomainId string `json:"domain_id"`
	// Specifies the connection status of the VPC endpoint.
	// pendingAcceptance: The VPC endpoint is to be accepted.
	// creating: The VPC endpoint is being created.
	// accepted: The VPC endpoint has been accepted.
	// rejected: The VPC endpoint has been rejected.
	// failed: The VPC endpoint service failed to be created.
	// deleting: The VPC endpoint is being deleted.
	Status string `json:"status"`
	// Specifies the error message.
	Error []ErrorParameters `json:"error"`
}
