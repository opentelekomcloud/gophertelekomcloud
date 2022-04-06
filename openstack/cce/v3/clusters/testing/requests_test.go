package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/clusters"
	fake "github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v3/common"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGetV3Cluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/daa97872-59d7-11e8-a787-0255ac101f54", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, Output)
	})

	actual, err := clusters.Get(fake.ServiceClient(), "daa97872-59d7-11e8-a787-0255ac101f54").Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)

}

func TestGetV3ClusterOTC(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/daa97872-59d7-11e8-a787-0255ac101f54", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, OutputOTC)
	})

	actual, err := clusters.Get(fake.ServiceClient(), "daa97872-59d7-11e8-a787-0255ac101f54").Extract()
	th.AssertNoErr(t, err)
	expected := ExpectedOTC
	th.AssertDeepEquals(t, expected, actual)

}

func TestListV3Cluster(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = fmt.Fprint(w, ListOutput)
	})

	// count := 0

	actual, err := clusters.List(fake.ServiceClient(), clusters.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract clusters: %v", err)
	}

	expected := ListExpected

	th.AssertDeepEquals(t, expected, actual)
}

func TestListV3ClusterOTC(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = fmt.Fprint(w, ListOutputOTC)
	})

	// count := 0

	actual, err := clusters.List(fake.ServiceClient(), clusters.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract clusters: %v", err)
	}

	expected := ListExpectedOTC

	th.AssertDeepEquals(t, expected, actual)
}

func TestCreateV3Cluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
{
    "kind": "Cluster",
    "apiversion": "v3",
    "metadata": {
        "name": "test-cluster"
           },
    "spec": {
		"category": "cce",
        "type": "VirtualMachine",
        "flavor": "cce.s1.small",
        "version": "v1.7.3-r10",
         "hostNetwork": {
            "vpc": "3305eb40-2707-4940-921c-9f335f84a2ca",
            "subnet": "00e41db7-e56b-4946-bf91-27bb9effd664"
        },
        "containerNetwork": {
            "mode": "overlay_l2"
        },
        "authentication": {
            "mode": "rbac",
			"authenticatingProxy": {}
        }
    }

}
`)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, Output)
	})
	options := clusters.CreateOpts{Kind: "Cluster",
		ApiVersion: "v3",
		Metadata:   clusters.CreateMetaData{Name: "test-cluster"},
		Spec: clusters.Spec{Type: "VirtualMachine",
			Category: "cce",
			Flavor:   "cce.s1.small",
			Version:  "v1.7.3-r10",
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:    "3305eb40-2707-4940-921c-9f335f84a2ca",
				SubnetId: "00e41db7-e56b-4946-bf91-27bb9effd664"},
			ContainerNetwork: clusters.ContainerNetworkSpec{Mode: "overlay_l2"},
			Authentication: clusters.AuthenticationSpec{
				Mode:                "rbac",
				AuthenticatingProxy: make(map[string]string)},
		},
	}
	actual, err := clusters.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)

}

func TestCreateV3TurboCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")

		th.TestJSONRequest(t, r, `
{
    "kind": "Cluster",
    "apiversion": "v3",
    "metadata": {
        "name": "test-turbo-cluster"
           },
    "spec": {
		"category": "turbo",
        "type": "VirtualMachine",
        "flavor": "cce.s2.small",
        "version": "v1.19.10-r0",
         "hostNetwork": {
            "vpc": "3305eb40-2707-4940-921c-9f335f84a2ca",
            "subnet": "00e41db7-e56b-4946-bf91-27bb9effd664"
        },
        "containerNetwork": {
            "mode": "eni"
        },
		"eniNetwork": {
			"eniSubnetId": "417dcc1f-95d7-43e7-8533-ab078d266303",
			"eniSubnetCIDR": "192.168.0.0/24"
		},
        "authentication": {
            "mode": "rbac",
			"authenticatingProxy": {}
        }
    }

}
`)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, Output)
	})
	options := clusters.CreateOpts{Kind: "Cluster",
		ApiVersion: "v3",
		Metadata:   clusters.CreateMetaData{Name: "test-turbo-cluster"},
		Spec: clusters.Spec{Type: "VirtualMachine",
			Category: "turbo",
			Flavor:   "cce.s2.small",
			Version:  "v1.19.10-r0",
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:    "3305eb40-2707-4940-921c-9f335f84a2ca",
				SubnetId: "00e41db7-e56b-4946-bf91-27bb9effd664"},
			ContainerNetwork: clusters.ContainerNetworkSpec{Mode: "eni"},
			EniNetwork: &clusters.EniNetworkSpec{
				SubnetId: "417dcc1f-95d7-43e7-8533-ab078d266303",
				Cidr:     "192.168.0.0/24",
			},
			Authentication: clusters.AuthenticationSpec{
				Mode:                "rbac",
				AuthenticatingProxy: make(map[string]string)},
		},
	}
	actual, err := clusters.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)

}

func TestUpdateV3Cluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/daa97872-59d7-11e8-a787-0255ac101f54", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "spec": {
        "description": "new description"
    }
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = fmt.Fprint(w, Output)
	})
	options := clusters.UpdateOpts{Spec: clusters.UpdateSpec{Description: "new description"}}
	actual, err := clusters.Update(fake.ServiceClient(), "daa97872-59d7-11e8-a787-0255ac101f54", options).Extract()
	th.AssertNoErr(t, err)
	expected := Expected
	th.AssertDeepEquals(t, expected, actual)
}

func TestDeleteV3Cluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/daa97872-59d7-11e8-a787-0255ac101f54", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
	})

	err := clusters.Delete(fake.ServiceClient(), "daa97872-59d7-11e8-a787-0255ac101f54").ExtractErr()
	th.AssertNoErr(t, err)

}

func TestDeleteV3TurboCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/api/v3/projects/c59fd21fd2a94963b822d8985b884673/clusters/daa97872-59d7-11e8-a787-0255ac101f54", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusOK)
	})

	err := clusters.Delete(fake.ServiceClient(), "daa97872-59d7-11e8-a787-0255ac101f54").ExtractErr()
	th.AssertNoErr(t, err)

}
