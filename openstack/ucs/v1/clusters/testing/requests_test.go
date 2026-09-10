package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ucs/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, createResponse)
	})

	opts := clusters.CreateOpts{
		Kind:       "Cluster",
		APIVersion: "v1",
		Metadata:   clusters.CreateMetadata{Name: "cce-cluster"},
		Spec: clusters.CreateSpec{
			Category:   "self",
			Type:       "turbo",
			Provider:   map[string]string{"CCE": "cce"},
			Country:    "DE",
			ManageType: "discrete",
		},
	}

	uid, err := clusters.Create(client.ServiceClient(), opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, clusterID, uid)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clusters/%s", clusterID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, getResponse)
	})

	cluster, err := clusters.Get(client.ServiceClient(), clusterID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, clusterID, cluster.Metadata.UID)
	th.AssertEquals(t, "cce", cluster.Spec.Provider)
	th.AssertEquals(t, "Available", cluster.Status.Phase)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, listResponse)
	})

	res, err := clusters.List(client.ServiceClient(), clusters.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, res.Total)
	th.AssertEquals(t, 1, len(res.Items))
	th.AssertEquals(t, clusterID, res.Items[0].Metadata.UID)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clusters/%s", clusterID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{}`)
	})

	opts := clusters.UpdateOpts{
		Kind:       "Cluster",
		APIVersion: "v1",
		Spec:       &clusters.UpdateSpec{Country: "AL", City: "AL"},
	}
	th.AssertNoErr(t, clusters.Update(client.ServiceClient(), clusterID, opts))
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clusters/%s", clusterID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
	})

	th.AssertNoErr(t, clusters.Delete(client.ServiceClient(), clusterID))
}

func TestActivate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clusters/%s/activation", clusterID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{}`)
	})

	th.AssertNoErr(t, clusters.Activate(client.ServiceClient(), clusterID))
}
