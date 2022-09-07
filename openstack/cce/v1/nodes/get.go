package nodes

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterID, k8sName string) (*GetNode, error) {
	raw, err := client.Get(fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+
		strings.Join([]string{"nodes", k8sName}, "/"), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res GetNode
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GetNode struct {
	Kind       string       `json:"kind"`
	ApiVersion string       `json:"apiVersion"`
	Metadata   MetadataNode `json:"metadata"`
	Spec       Spec         `json:"spec"`
	Status     Status       `json:"status"`
}

type Status struct {
	Capacity        Capacity       `json:"capacity"`
	Allocatable     Capacity       `json:"allocatable"`
	Conditions      []Condition    `json:"conditions"`
	Addresses       []Address      `json:"addresses"`
	DaemonEndpoints DaemonEndpoint `json:"daemonEndpoints"`
	NodeInfo        NodeInfo       `json:"nodeInfo"`
	Images          []Image        `json:"images"`
}

type Capacity struct {
	CCEEni           string `json:"cce/eni"`
	CCESubEni        string `json:"cce/sub-eni"`
	CPU              string `json:"cpu"`
	EphemeralStorage string `json:"ephemeral-storage"`
	HugePages1Gi     string `json:"hugepages-1Gi"`
	HugePages2Mi     string `json:"hugepages-2Mi"`
	Memory           string `json:"memory"`
	Pods             string `json:"pods"`
}

type Condition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	LastHeartbeatTime  string `json:"lastHeartbeatTime"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}

type Address struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

type DaemonEndpoint struct {
	KubeletEndpoint PortEndpoint `json:"kubeletEndpoint"`
}

type PortEndpoint struct {
	Port int `json:"Port"`
}

type NodeInfo struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OsImage                 string `json:"osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
	EniQuotaSize            string `json:"eniQuotaSize"`
}

type Image struct {
	Names     []string `json:"names"`
	SizeBytes int      `json:"sizeBytes"`
}
