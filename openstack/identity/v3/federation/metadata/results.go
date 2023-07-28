package metadata

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type Metadata struct {
	ID           string `json:"id"`
	ProviderID   string `json:"idp_id"`
	EntityID     string `json:"entity_id"`
	ProtocolID   string `json:"protocol_id"`
	DomainID     string `json:"domain_id"`
	XAccountType string `json:"xaccount_type"`
	UpdateTime   string `json:"update_time"`
	Data         string `json:"data"`
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*Metadata, error) {
	metadata := new(Metadata)
	err := r.ExtractIntoStructPtr(metadata, "")
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

type ImportResult struct {
	golangsdk.Result
}
