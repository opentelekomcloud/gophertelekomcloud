package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

type Clusters struct {
	suite.Suite
}

func (s *Clusters) SetupTest() {
	th.SetupHTTP()
}

func (s *Clusters) TearDownTest() {
	th.TeardownHTTP()
}

func (s *Clusters) TestCreateRequest() {
	t := s.T()

	th.Mux.HandleFunc("/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, createRequestBody)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, createResponseBody)
	})

	opts := clusters.CreateOpts{
		Instance: &clusters.InstanceSpec{
			Flavor: "css.large.8",
			Volume: &clusters.Volume{
				Type: "COMMON",
				Size: 100,
			},
			Nics: &clusters.Nics{
				VpcID:           vpcID,
				SubnetID:        subnetID,
				SecurityGroupID: sgID,
			},
		},
		Name:        clusterName,
		InstanceNum: 4,
		DiskEncryption: &clusters.DiskEncryption{
			Encrypted: "1",
			CmkID:     cmkID,
		},
		HttpsEnabled:     "false",
		AuthorityEnabled: false,
	}

	created, err := clusters.Create(fake.ServiceClient(), opts)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, clusterName, created.Name)
	th.AssertEquals(t, clusterID, created.ID)
}

func (s *Clusters) TestGetRequest() {
	t := s.T()

	clusterURL := fmt.Sprintf("/clusters/%s", clusterID)
	th.Mux.HandleFunc(clusterURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, getResponseBody)
	})

	cluster, err := clusters.Get(fake.ServiceClient(), clusterID)
	th.AssertNoErr(t, err)
	if cluster == nil {
		t.Fatal("cluster is nil")
	}

	th.AssertEquals(t, clusterID, cluster.ID)
	th.AssertEquals(t, clusterName, cluster.Name)
	th.AssertEquals(t, "200", cluster.Status)
	th.AssertEquals(t, "7.6.2", cluster.Datastore.Version)
	th.AssertEquals(t, "elasticsearch", cluster.Datastore.Type)
}

func (s *Clusters) TestListRequest() {
	t := s.T()

	th.Mux.HandleFunc("/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, listResponseBody)
	})

	clusterSlice, err := clusters.List(fake.ServiceClient())
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(clusterSlice), 3)
	for _, cluster := range clusterSlice {
		t.Logf("%+v", cluster)
	}
}

func (s *Clusters) TestDeleteRequest() {
	t := s.T()

	clusterURL := fmt.Sprintf("/clusters/%s", clusterID)
	th.Mux.HandleFunc(clusterURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.WriteHeader(200)
	})

	err := clusters.Delete(fake.ServiceClient(), clusterID)
	th.AssertNoErr(t, err)
}

func TestClustersMethods(t *testing.T) {
	suite.Run(t, new(Clusters))
}
