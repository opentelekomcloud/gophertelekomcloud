package load_balancer

import (
	"log"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatingListenerOpts struct {
	// These parameters passed to the load_balancers.UpdatingLoadBalancingListeners function.
	// Listener object.
	Listener EsListenerRequest `json:"listener" required:"true"`
	// Type: searchTool indicates that the load balancer is modified for Elasticsearch/OpenSearch.
	// viewTool indicates that the load balancer is modified for Kibana/OpenSearch Dashboards.
	// The default value is searchTool.
	Type string `json:"type,omitempty"`
}

type EsListenerRequest struct {
	// These parameters are part of an EsListenerRequest object.
	// ID of the server certificate used by the listener.
	DefaultTlsContainerRef string `json:"default_tls_container_ref" required:"true"`
	// ID of the CA certificate used by the listener.
	// This parameter is mandatory when bidirectional authentication is to be updated.
	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref,omitempty"`
}

// UpdatingLoadBalancingListeners will update load balancing listeners for a CSS cluster based on UpdatingListenerOpts.
func UpdatingLoadBalancingListeners(client *golangsdk.ServiceClient, clusterID string, listenerID string, opts UpdatingListenerOpts) (*EsListenerResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("clusters", clusterID, "es-listeners", listenerID)

	log.Println(b)

	raw, err := client.Put(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		return nil, err
	}

	var res EsListenerResponse
	err = extract.IntoStructPtr(raw.Body, &res, "listener")
	return &res, err
}

// func Update(client *golangsdk.ServiceClient, aclID string, opts CreateOpts) (*AclResp, error) {
// 	b, err := build.RequestBody(opts, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "acls", aclID), b, nil, &golangsdk.RequestOpts{
// 		OkCodes: []int{200},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var res AclResp

// 	err = extract.Into(raw.Body, &res)
// 	return &res, err
// }
