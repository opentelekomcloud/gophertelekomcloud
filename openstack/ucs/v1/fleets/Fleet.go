package fleets

// ClusterGroup is the response body of GET /v1/clustergroups/{id}.
type ClusterGroup struct {
	Kind       string             `json:"kind"`
	APIVersion string             `json:"apiVersion"`
	Metadata   ObjectMeta         `json:"metadata"`
	Spec       ClusterGroupSpec   `json:"spec"`
	Status     ClusterGroupStatus `json:"status"`
}

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

type ManagedFieldsEntry struct {
	Manager    string                 `json:"manager"`
	Operation  string                 `json:"operation"`
	APIVersion string                 `json:"apiVersion"`
	Time       string                 `json:"time"`
	FieldsType string                 `json:"fieldsType"`
	FieldsV1   map[string]interface{} `json:"fieldsV1"`
}

type OwnerReference struct {
	APIVersion         string `json:"apiVersion"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	UID                string `json:"uid"`
	Controller         bool   `json:"controller"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
}

type ClusterGroupSpec struct {
	RuleNamespaces                []string          `json:"ruleNamespaces"`
	FederationID                  string            `json:"federationId"`
	Description                   string            `json:"description"`
	DNSSuffix                     []string          `json:"dnsSuffix"`
	FederationExpirationTimestamp string            `json:"federationExpirationTimestamp"`
	PolicyID                      string            `json:"policyId"`
	FederationVersion             string            `json:"federationVersion"`
	ConnectGatewayEndpoints       []ConnectEndpoint `json:"connectGatewayEndpoints"`
}

type ConnectEndpoint struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	ID     string `json:"id"`
	Region string `json:"region"`
}

type ClusterGroupStatus struct {
	Conditions []ClusterGroupCondition `json:"conditions"`
}

type ClusterGroupCondition struct {
	Type               string `json:"type"`
	Status             string `json:"status"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
	LastTransitionTime string `json:"lastTransitionTime"`
}
