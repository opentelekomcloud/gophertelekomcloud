package nodes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

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

type GetNode struct {
	Kind       string       `json:"kind"`
	ApiVersion string       `json:"apiVersion"`
	Metadata   MetadataNode `json:"metadata"`
	Spec       Spec         `json:"spec"`
	Status     Status       `json:"status"`
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

type ListResult struct {
	golangsdk.Result
}

func (r ListResult) Extract() (*ListNodes, error) {
	s := new(ListNodes)
	err := r.ExtractIntoStructPtr(s, "")
	return s, err
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*GetNode, error) {
	s := new(GetNode)
	err := r.ExtractIntoStructPtr(s, "")
	return s, err
}

type UpdateResult struct {
	golangsdk.Result
}

func (r UpdateResult) Extract() (*GetNode, error) {
	s := new(GetNode)
	err := r.ExtractIntoStructPtr(s, "")
	return s, err
}
