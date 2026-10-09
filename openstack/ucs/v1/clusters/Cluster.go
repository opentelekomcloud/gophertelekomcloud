package clusters

// Cluster is the response body of GET /v1/clusters/{id} (querying_a_cluster.rst, Table 3).
type Cluster struct {
	Kind       string        `json:"kind"`
	APIVersion string        `json:"apiVersion"`
	Metadata   ObjectMeta    `json:"metadata"`
	Spec       ClusterSpec   `json:"spec"`
	Status     ClusterStatus `json:"status"`
}

// ObjectMeta is cluster metadata (Table 4).
type ObjectMeta struct {
	UID               string               `json:"uid"`
	Name              string               `json:"name"`
	GenerateName      string               `json:"generateName"`
	Namespace         string               `json:"namespace"`
	Labels            map[string]string    `json:"labels"`
	Annotations       map[string]string    `json:"annotations"`
	CreationTimestamp string               `json:"creationTimestamp"`
	UpdateTimestamp   string               `json:"updateTimestamp"`
	ResourceVersion   string               `json:"resourceVersion"`
	Generation        int                  `json:"generation"`
	ManagedFields     []ManagedFieldsEntry `json:"managedFields"`
	OwnerReferences   []OwnerReference     `json:"ownerReferences"`
}

// ManagedFieldsEntry describes fields managed by workflows (Table 5).
type ManagedFieldsEntry struct {
	Manager    string                 `json:"manager"`
	Operation  string                 `json:"operation"`
	APIVersion string                 `json:"apiVersion"`
	Time       string                 `json:"time"`
	FieldsType string                 `json:"fieldsType"`
	FieldsV1   map[string]interface{} `json:"fieldsV1"`
}

// OwnerReference describes ownership of an object (Table 6).
type OwnerReference struct {
	APIVersion         string `json:"apiVersion"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	UID                string `json:"uid"`
	Controller         bool   `json:"controller"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
}

// ClusterSpec is the detailed cluster description (Table 7).
// Note: in the response, provider is a plain string (unlike the map used in the request).
type ClusterSpec struct {
	SyncMode                    string                `json:"syncMode"`
	ClusterGroupID              string                `json:"clusterGroupID"`
	ManageType                  string                `json:"manageType"`
	RuleNamespaces              []RuleNamespace       `json:"ruleNamespaces"`
	APIEndpoint                 string                `json:"apiEndpoint"`
	SecretRef                   *LocalSecretReference `json:"secretRef"`
	InsecureSkipTLSVerification bool                  `json:"insecureSkipTLSVerification"`
	ProxyURL                    string                `json:"proxyURL"`
	Provider                    string                `json:"provider"`
	Type                        string                `json:"type"`
	Category                    string                `json:"category"`
	EnableDistMgt               bool                  `json:"enableDistMgt"`
	Region                      string                `json:"region"`
	Country                     string                `json:"country"`
	City                        string                `json:"city"`
	ProjectID                   string                `json:"projectID"`
	ProjectName                 string                `json:"projectName"`
	Zone                        string                `json:"zone"`
	Taints                      []Taint               `json:"taints"`
	IsDownloadedCert            bool                  `json:"IsDownloadedCert"`
	PolicyID                    string                `json:"policyId"`
	OperatorNamespace           string                `json:"operatorNamespace"`
	ConnectProxyEndpoints       []ConnectEndpoint     `json:"connectProxyEndpoints"`
}

// RuleNamespace associates permission policies with namespaces (Table 8).
type RuleNamespace struct {
	Rules      []RuleInfo `json:"rules"`
	Namespaces []string   `json:"namespaces"`
}

// RuleInfo is a permission policy (Table 9).
type RuleInfo struct {
	RuleID   string `json:"ruleID"`
	RuleName string `json:"ruleName"`
}

// LocalSecretReference references a secret used to access a cluster (Table 10).
type LocalSecretReference struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

// Taint is a cluster taint (Table 11).
type Taint struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Effect    string `json:"effect"`
	TimeAdded string `json:"timeadded"`
}

// ConnectEndpoint is a VPC endpoint service used to reach the proxy (Table 12).
type ConnectEndpoint struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	ID     string `json:"id"`
	Region string `json:"region"`
}

// ClusterStatus is the object status (Table 13).
type ClusterStatus struct {
	KubernetesVersion string            `json:"kubernetesVersion"`
	Conditions        []ConditionStatus `json:"conditions"`
	NodeSummary       *NodeSummary      `json:"nodeSummary"`
	ResourceSummary   *ResourceSummary  `json:"resourceSummary"`
	Endpoints         []Endpoint        `json:"endpoints"`
	Phase             string            `json:"phase"`
	Reason            string            `json:"reason"`
	Message           string            `json:"message"`
	ArrearFreeze      string            `json:"arrearFreeze"`
	PoliceFreeze      string            `json:"policeFreeze"`
	APIEnablements    []APIEnablement   `json:"apiEnablements"`
}

// ConditionStatus is a cluster condition (Table 14).
type ConditionStatus struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	ObservedGeneration int    `json:"observedgeneration"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
}

// NodeSummary holds node statistics (Table 15).
type NodeSummary struct {
	TotalNum int `json:"totalNum"`
	ReadyNum int `json:"readyNum"`
}

// ResourceSummary holds resource statistics (Table 16).
type ResourceSummary struct {
	Allocatable map[string]interface{} `json:"allocatable"`
	Allocating  map[string]interface{} `json:"allocating"`
	Allocated   map[string]interface{} `json:"allocated"`
	Capacity    map[string]interface{} `json:"capacity"`
}

// Endpoint is a cluster endpoint (Table 17).
type Endpoint struct {
	URL    string `json:"url"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

// APIEnablement lists enabled resources (Table 18).
type APIEnablement struct {
	GroupVersion string        `json:"groupVersion"`
	Resources    []APIResource `json:"resources"`
}

// APIResource is a resource type and name (Table 19).
type APIResource struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}
