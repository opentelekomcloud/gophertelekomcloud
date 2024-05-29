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
	SecurityGroupID  string             `json:"securityGroupId"`
	Tags             []tags.ResourceTag `json:"tags"`
	PublicKibana     *PublicKibana      `json:"publicKibanaResp"`
	PublicNetwork    *PublicNetwork     `json:"elbWhiteList"`
	PublicIp         string             `json:"publicIp"`
	VpcEpIp          string             `json:"vpcepIp"`
	BandwidthSize    int                `json:"bandwidthSize"`
	BackupAvailable  bool               `json:"backupAvailable"`
}

type Datastore struct {
	// Version - engine version.
	// The default value is 7.6.2.
	Version string `json:"version" required:"true"`
	// Type - Engine type.
	// The value is `elasticsearch` or 'opensearch'.
	Type string `json:"type" required:"true"`
}

type Instance struct {
	Type             string             `json:"type"`
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	Status           string             `json:"status"`
	SpecCode         string             `json:"specCode"`
	AvailabilityZone string             `json:"azCode"`
	Volume           *ShowClusterVolume `json:"volume"`
}

type PublicKibana struct {
	// Bandwidth range. Unit: Mbit/s
	Bandwidth int `json:"eipSize"`
	// Kibana public network access information.
	ElbWhiteList *PublicNetwork `json:"elbWhiteListResp"`
	// Specifies the IP address for accessing Kibana.
	PublicIp string `json:"publicKibanaIp"`
}

type PublicNetwork struct {
	// Whether the public network access control is enabled.
	// true: Public network access control is enabled.
	// false: Public network access control is disabled.
	Enabled bool `json:"enableWhiteList"`
	// Whitelist of public network for accessing Kibana.
	Whitelist string `json:"whiteList"`
}

type ShowClusterVolume struct {
	// Instance disk type
	Type string `json:"type"`
	// Instance disk size
	Size int `json:"size"`
}

type FailedReasons struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
}
