package nodes

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// List returns collection of nodes.
func List(client *golangsdk.ServiceClient, clusterID string) (*ListNodes, error) {
	url := fmt.Sprintf("https://%s.%s%s", clusterID, client.ResourceBaseURL()[8:], "nodes")

	raw, err := client.Get(url, nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res ListNodes
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListNodes struct {
	Kind       string       `json:"kind"`
	ApiVersion string       `json:"apiVersion"`
	Metadata   MetadataLink `json:"metadata"`
	Nodes      []Node       `json:"items"`
}

type MetadataLink struct {
	SelfLink        string `json:"selfLink"`
	ResourceVersion string `json:"resourceVersion"`
}

type Node struct {
	Metadata MetadataNode `json:"metadata"`
	Spec     Spec         `json:"spec"`
	Status   Status       `json:"status"`
}

type MetadataNode struct {
	Name              string                 `json:"name"`
	SelfLink          string                 `json:"selfLink"`
	ID                string                 `json:"uid"`
	ResourceVersion   string                 `json:"resourceVersion"`
	CreationTimestamp string                 `json:"creationTimestamp"`
	Labels            map[string]interface{} `json:"labels"`
	Annotations       map[string]interface{} `json:"annotations"`
}

type Spec struct {
	ProviderID string  `json:"providerID"`
	Taints     []Taint `json:"taints"`
}

type Taint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}
