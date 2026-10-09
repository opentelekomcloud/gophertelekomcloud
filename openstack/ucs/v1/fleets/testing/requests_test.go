package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ucs/v1/fleets"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/clustergroups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, createResponse)
	})

	opts := fleets.CreateOpts{
		Metadata: fleets.CreateMetadata{Name: "group02281605"},
		Spec: &fleets.CreateSpec{
			ClusterIDs:  []string{"514c1a3c-8ec7-11ec-b384-0255ac100189"},
			Description: "test fleet",
		},
	}

	uid, err := fleets.Create(client.ServiceClient(), opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, fleetID, uid)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clustergroups/%s", fleetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, getResponse)
	})

	fleet, err := fleets.Get(client.ServiceClient(), fleetID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, fleetID, fleet.Metadata.UID)
	th.AssertEquals(t, "e2f27cc6-82b5-11ee-84e3-0255ac100032", fleet.Spec.FederationID)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/clustergroups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, listResponse)
	})

	res, err := fleets.List(client.ServiceClient(), fleets.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, res.Total)
	th.AssertEquals(t, fleetID, res.Items[0].Metadata.UID)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clustergroups/%s/description", fleetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{}`)
	})

	opts := fleets.UpdateOpts{Description: "new description"}
	th.AssertNoErr(t, fleets.Update(client.ServiceClient(), fleetID, opts))
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clustergroups/%s", fleetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
	})

	th.AssertNoErr(t, fleets.Delete(client.ServiceClient(), fleetID))
}

func TestAddCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clusters/%s/join", clusterID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, addClusterRequest)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{}`)
	})

	opts := fleets.AddClusterOpts{ClusterGroupID: fleetID}
	th.AssertNoErr(t, fleets.AddCluster(client.ServiceClient(), clusterID, opts))
}

func TestRemoveCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clusters/%s/unjoin", clusterID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{}`)
	})

	th.AssertNoErr(t, fleets.RemoveCluster(client.ServiceClient(), clusterID))
}

func TestEnableFederation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clustergroups/%s/federations", fleetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, `{}`)
	})

	th.AssertNoErr(t, fleets.EnableFederation(client.ServiceClient(), fleetID, fleets.EnableFederationOpts{}))
}

func TestDisableFederation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/clustergroups/%s/federations", fleetID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
	})

	th.AssertNoErr(t, fleets.DisableFederation(client.ServiceClient(), fleetID))
}
