package network

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	Namespace       string         `json:"-"`
	DryRun          string         `json:"-" q:"dryRun,omitempty"`
	FieldManager    string         `json:"-" q:"fieldManager,omitempty"`
	FieldValidation string         `json:"-" q:"fieldValidation,omitempty"`
	Pretty          string         `json:"-" q:"pretty,omitempty"`
	APIVersion      string         `json:"apiVersion,omitempty"`
	Kind            string         `json:"kind,omitempty"`
	Metadata        *ObjectMeta    `json:"metadata,omitempty"`
	Spec            *NetworkSpec   `json:"spec,omitempty"`
	Status          *NetworkStatus `json:"status,omitempty"`
}

type ObjectMeta struct {
	Annotations                map[string]string    `json:"annotations,omitempty"`
	ClusterName                string               `json:"clusterName,omitempty"`
	CreationTimestamp          string               `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64               `json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp          string               `json:"deletionTimestamp,omitempty"`
	Enable                     *bool                `json:"enable,omitempty"`
	Finalizers                 []string             `json:"finalizers,omitempty"`
	GenerateName               string               `json:"generateName,omitempty"`
	Generation                 *int64               `json:"generation,omitempty"`
	Labels                     map[string]string    `json:"labels,omitempty"`
	ManagedFields              []ManagedFieldsEntry `json:"managedFields,omitempty"`
	Name                       string               `json:"name,omitempty"`
	Namespace                  string               `json:"namespace,omitempty"`
	OwnerReferences            []OwnerReference     `json:"ownerReferences,omitempty"`
	ResourceVersion            string               `json:"resourceVersion,omitempty"`
	SelfLink                   string               `json:"selfLink,omitempty"`
	UID                        string               `json:"uid,omitempty"`
}

type ManagedFieldsEntry struct {
	APIVersion  string      `json:"apiVersion,omitempty"`
	FieldsType  string      `json:"fieldsType,omitempty"`
	FieldsV1    interface{} `json:"fieldsV1,omitempty"`
	Manager     string      `json:"manager,omitempty"`
	Operation   string      `json:"operation,omitempty"`
	Subresource string      `json:"subresource,omitempty"`
	Time        string      `json:"time,omitempty"`
}

type OwnerReference struct {
	APIVersion         string `json:"apiVersion"`
	BlockOwnerDeletion *bool  `json:"blockOwnerDeletion,omitempty"`
	Controller         *bool  `json:"controller,omitempty"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	UID                string `json:"uid"`
}

type NetworkSpec struct {
	IPFamilies     []string     `json:"ipFamilies,omitempty"`
	NetworkType    string       `json:"networkType,omitempty"`
	SecurityGroups []string     `json:"securityGroups,omitempty"`
	Subnets        []SubnetConf `json:"subnets,omitempty"`
}

type SubnetConf struct {
	SubnetID string `json:"subnetID,omitempty"`
}

type NetworkStatus struct {
	Conditions  []Condition  `json:"conditions,omitempty"`
	Status      string       `json:"status,omitempty"`
	SubnetAttrs []SubnetAttr `json:"subnetAttrs,omitempty"`
}

type Condition struct {
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	Message            string `json:"message,omitempty"`
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	Reason             string `json:"reason,omitempty"`
	Status             string `json:"status,omitempty"`
	Type               string `json:"type,omitempty"`
}

type SubnetAttr struct {
	NetworkID  string `json:"networkID,omitempty"`
	SubnetV4ID string `json:"subnetV4ID,omitempty"`
	SubnetV6ID string `json:"subnetV6ID,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Network, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("namespaces", opts.Namespace, "networks").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var r Network
	_, err = client.Post(client.ServiceURL(url.String()), b, &r, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type Network struct {
	APIVersion string        `json:"apiVersion,omitempty"`
	Kind       string        `json:"kind,omitempty"`
	Metadata   ObjectMeta    `json:"metadata,omitempty"`
	Spec       NetworkSpec   `json:"spec,omitempty"`
	Status     NetworkStatus `json:"status,omitempty"`
}
