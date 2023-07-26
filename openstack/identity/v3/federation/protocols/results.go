package protocols

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Protocol struct {
	ID        string            `json:"id"`
	MappingID string            `json:"mapping_id"`
	Links     map[string]string `json:"links"`
}

type ProtocolPage struct {
	pagination.LinkedPageBase
}

func (p ProtocolPage) IsEmpty() (bool, error) {
	protocols, err := ExtractProtocols(p)
	return len(protocols) == 0, err
}

func ExtractProtocols(p pagination.Page) ([]Protocol, error) {
	var protocols []Protocol

	err := extract.IntoSlicePtr(bytes.NewReader(p.(ProtocolPage).Body), &protocols, "protocols")
	return protocols, err
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Protocol, error) {
	protocol := new(Protocol)
	err := r.ExtractIntoStructPtr(protocol, "protocol")
	if err != nil {
		return nil, err
	}
	return protocol, nil
}

type GetResult struct {
	commonResult
}

type CreateResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}
