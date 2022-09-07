package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreatedCluster struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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

type CreateResult struct {
	golangsdk.Result
}

func (r CreateResult) Extract() (*CreatedCluster, error) {
	var s struct {
		Cluster *CreatedCluster `json:"cluster"`
	}
	err := r.ExtractInto(&s)
	return s.Cluster, err
}

type ClusterPage struct {
	pagination.SinglePageBase
}

func ExtractClusters(r pagination.Page) ([]Cluster, error) {
	var clusters []Cluster
	err := (r.(ClusterPage)).ExtractIntoSlicePtr(&clusters, "clusters")
	return clusters, err
}

func (p ClusterPage) GetBody() []byte {
	return p.Body
}

func (p ClusterPage) IsEmpty() (bool, error) {
	clusterSlice, err := ExtractClusters(p)
	return len(clusterSlice) == 0, err
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*Cluster, error) {
	cluster := new(Cluster)
	err := r.ExtractInto(cluster)
	return cluster, err
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type CertificateResult struct {
	golangsdk.Result
}

func (r CertificateResult) Extract() (string, error) {
	var cert struct {
		CertBase64 string `json:"certBase64"`
	}

	err := r.ExtractInto(&cert)
	return cert.CertBase64, err
}

type ExtendResult struct {
	golangsdk.Result
}

type ExtendedInstance struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	ShardID string `json:"shard_id"`
}

type ExtendedCluster struct {
	ID        string             `json:"id"`
	Instances []ExtendedInstance `json:"instances"`
}

func (r ExtendResult) Extract() (*ExtendedCluster, error) {
	c := new(ExtendedCluster)
	err := r.ExtractInto(c)
	return c, err
}
