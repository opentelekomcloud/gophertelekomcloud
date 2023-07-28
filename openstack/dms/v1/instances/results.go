package instances

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// InstanceCreate response
type InstanceCreate struct {
	InstanceID string `json:"instance_id"`
}

// CreateResult is a struct that contains all the return parameters of creation
type CreateResult struct {
	golangsdk.Result
}

// Extract from CreateResult
func (r CreateResult) Extract() (*InstanceCreate, error) {
	s := new(InstanceCreate)
	err := r.ExtractIntoStructPtr(s, "")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteResult is a struct which contains the result of deletion
type DeleteResult struct {
	Err error
}

type ListDmsResponse struct {
	Instances  []Instance `json:"instances"`
	TotalCount int        `json:"instance_num"`
}

// Instance response
type Instance struct {
	Name              string   `json:"name"`
	Engine            string   `json:"engine"`
	EngineVersion     string   `json:"engine_version"`
	Specification     string   `json:"specification"`
	StorageSpace      int      `json:"storage_space"`
	PartitionNum      string   `json:"partition_num"`
	UsedStorageSpace  int      `json:"used_storage_space"`
	ConnectAddress    string   `json:"connect_address"`
	Port              int      `json:"port"`
	Status            string   `json:"status"`
	Description       string   `json:"description"`
	InstanceID        string   `json:"instance_id"`
	ResourceSpecCode  string   `json:"resource_spec_code"`
	Type              string   `json:"type"`
	ChargingMode      int      `json:"charging_mode"`
	VpcID             string   `json:"vpc_id"`
	VpcName           string   `json:"vpc_name"`
	CreatedAt         string   `json:"created_at"`
	ErrorCode         string   `json:"error_code"`
	ProductID         string   `json:"product_id"`
	SecurityGroupID   string   `json:"security_group_id"`
	SecurityGroupName string   `json:"security_group_name"`
	SubnetID          string   `json:"subnet_id"`
	SubnetName        string   `json:"subnet_name"`
	SubnetCIDR        string   `json:"subnet_cidr"`
	AvailableZones    []string `json:"available_zones"`
	UserID            string   `json:"user_id"`
	UserName          string   `json:"user_name"`
	AccessUser        string   `json:"access_user"`
	TotalStorageSpace int      `json:"total_storage_space"`
	StorageResourceID string   `json:"storage_resource_id"`
	StorageSpecCode   string   `json:"storage_spec_code"`
	RetentionPolicy   string   `json:"retention_policy"`
	KafkaPublicStatus string   `json:"kafka_public_status"`
	PublicBandwidth   int      `json:"public_bandwidth"`
	SslEnable         bool     `json:"ssl_enable"`
	ServiceType       string   `json:"service_type"`
	StorageType       string   `json:"storage_type"`
	OrderID           string   `json:"order_id"`
	MaintainBegin     string   `json:"maintain_begin"`
	MaintainEnd       string   `json:"maintain_end"`
}

// UpdateResult is a struct from which can get the result of update method
type UpdateResult struct {
	Err error
}

// GetResult contains the body of getting detailed
type GetResult struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult) Extract() (*Instance, error) {
	s := new(Instance)
	err := r.ExtractIntoStructPtr(s, "")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type DmsPage struct {
	pagination.SinglePageBase
}

func (r DmsPage) IsEmpty() (bool, error) {
	data, err := ExtractDmsInstances(r)
	if err != nil {
		return false, err
	}
	return len(data.Instances) == 0, err
}

// ExtractDmsInstances is a function that takes a ListResult and returns the services' information.
func ExtractDmsInstances(r pagination.Page) (*ListDmsResponse, error) {
	var s ListDmsResponse

	err := extract.IntoStructPtr(bytes.NewReader((r.(DmsPage)).Body), &s, "")
	if err != nil {
		return nil, err
	}
	return &s, nil
}
