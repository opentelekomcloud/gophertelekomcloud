package clusters

import (
	"encoding/json"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

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
