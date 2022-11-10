package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type Cluster struct {
	Datastore        Datastore          `json:"datastore"`
	Instances        []Instance         `json:"instances"`
	Updated          string             `json:"updated"`
	Name             string             `json:"name"`
	Created          string             `json:"created"`
	ID               string             `json:"id"`
	Status           string             `json:"status"`
	Endpoint         string             `json:"endpoint"`
	ActionProgress   map[string]string  `json:"actionProgress"`
	Actions          []string           `json:"actions"`
	FailedReasons    *FailedReasons     `json:"failed_reasons"`
	HttpsEnabled     bool               `json:"httpsEnable"`
	AuthorityEnabled bool               `json:"authorityEnable"`
	DiskEncrypted    bool               `json:"diskEncrypted"`
	CmkID            string             `json:"cmkId"`
	VpcID            string             `json:"vpcId"`
	SubnetID         string             `json:"subnetId"`
	SecurityGroupID  string             `json:"securityGroupId" required:"true"`
	Tags             []tags.ResourceTag `json:"tags"`
}

type Datastore struct {
	// Version - engine version.
	// The default value is 7.6.2.
	Version string `json:"version" required:"true"`
	// Type - Engine type.
	// The default value is `elasticsearch`.
	Type string `json:"type,omitempty"`
}

type Instance struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	SpecCode         string `json:"specCode"`
	AvailabilityZone string `json:"azCode"`
}

type FailedReasons struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
}
