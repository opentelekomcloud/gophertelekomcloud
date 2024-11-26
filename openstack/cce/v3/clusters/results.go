package clusters

import (
	"encoding/json"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type Certificate struct {
	// API type, fixed value Config
	Kind string `json:"kind"`
	// API version, fixed value v1
	ApiVersion string `json:"apiVersion"`
	// Cluster list
	Clusters []CertClusters `json:"clusters"`
	// User list
	Users []CertUsers `json:"users"`
	// Context list
	Contexts []CertContexts `json:"contexts"`
	// The current context
	CurrentContext string `json:"current-context"`
}

type CertClusters struct {
	// Cluster name
	Name string `json:"name"`
	// Cluster information
	Cluster CertCluster `json:"cluster"`
}

type CertCluster struct {
	// Server IP address
	Server string `json:"server"`
	// Certificate data
	CertAuthorityData string `json:"certificate-authority-data"`
}

type CertUsers struct {
	// User name
	Name string `json:"name"`
	// Cluster information
	User CertUser `json:"user"`
}

type CertUser struct {
	// Client certificate
	ClientCertData string `json:"client-certificate-data"`
	// Client key data
	ClientKeyData string `json:"client-key-data"`
}

type CertContexts struct {
	// Context name
	Name string `json:"name"`
	// Context information
	Context CertContext `json:"context"`
}

type CertContext struct {
	// Cluster name
	Cluster string `json:"cluster"`
	// User name
	User string `json:"user"`
}

// UnmarshalJSON helps to unmarshal Status fields into needed values.
// OTC and Huawei have different data types and child fields for `endpoints` field in Cluster Status.
// This function handles the unmarshal for both
func (r *Status) UnmarshalJSON(b []byte) error {
	type tmp Status
	var s struct {
		tmp
		Endpoints []Endpoints `json:"endpoints"`
	}

	err := json.Unmarshal(b, &s)

	if err != nil {
		switch err.(type) {
		case *json.UnmarshalTypeError: // check if type error occurred (handles the different endpoint structure for huawei and otc)
			var s struct {
				tmp
				Endpoints Endpoints `json:"endpoints"`
			}
			err := json.Unmarshal(b, &s)
			if err != nil {
				return err
			}
			*r = Status(s.tmp)
			r.Endpoints = []Endpoints{{Internal: s.Endpoints.Internal,
				External:    s.Endpoints.External,
				ExternalOTC: s.Endpoints.ExternalOTC}}
			return nil
		default:
			return err
		}
	}

	*r = Status(s.tmp)
	r.Endpoints = s.Endpoints

	return err
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

type GetCertResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a cluster.
func (r GetCertResult) Extract() (*Certificate, error) {
	var s Certificate
	err := r.ExtractInto(&s)
	return &s, err
}

// ExtractMap is a function that accepts a result and extracts a kubeconfig.
func (r GetCertResult) ExtractMap() (map[string]interface{}, error) {
	var s map[string]interface{}
	err := r.ExtractInto(&s)
	return s, err
}

// UpdateIpResult represents the result of an update operation. Call its Extract
// method to interpret it as a Cluster.
type UpdateIpResult struct {
	golangsdk.ErrResult
}
