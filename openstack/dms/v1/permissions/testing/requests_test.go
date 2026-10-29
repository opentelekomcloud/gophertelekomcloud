package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v1/permissions"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID  = "0123456789abcdef0123456789abcdef"
	instanceID = "54602a9d-5e22-4239-9123-77e350df4a34"
	topicsPath = "/v1/" + projectID + "/instances/" + instanceID + "/topics"
	grantPath  = topicsPath + "/accesspolicy"
)

func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v1/" + projectID + "/"
	return sc
}

func handleGrant(t *testing.T, status int, expectedBody string) {
	th.Mux.HandleFunc(grantPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, expectedBody)
		w.WriteHeader(status)
	})
}

func TestGrantPermissions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGrant(t, http.StatusNoContent, `{
		"topics": [{
			"name": "topic-test",
			"policies": [
				{"user_name": "root", "access_policy": "all"},
				{"user_name": "reader", "access_policy": "sub"}
			]
		}]
	}`)

	err := permissions.Create(dmsClient(), instanceID, []permissions.CreateOpts{{
		Name: "topic-test",
		Policies: []permissions.CreatePolicy{
			{UserName: "root", AccessPolicy: "all"},
			{UserName: "reader", AccessPolicy: "sub"},
		},
	}})
	th.AssertNoErr(t, err)
}

func TestGrantPermissionsAccepts200(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGrant(t, http.StatusOK, `{"topics": [{"name": "topic-test", "policies": [{"user_name": "root", "access_policy": "pub"}]}]}`)

	err := permissions.Create(dmsClient(), instanceID, []permissions.CreateOpts{{
		Name:     "topic-test",
		Policies: []permissions.CreatePolicy{{UserName: "root", AccessPolicy: "pub"}},
	}})
	th.AssertNoErr(t, err)
}

func TestGrantPermissionsEmptyPolicies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGrant(t, http.StatusNoContent, `{"topics": [{"name": "topic-test", "policies": []}]}`)

	err := permissions.Create(dmsClient(), instanceID, []permissions.CreateOpts{{
		Name:     "topic-test",
		Policies: []permissions.CreatePolicy{},
	}})
	th.AssertNoErr(t, err)
}

func TestGrantPermissionsOmitsEmptyPolicyFields(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGrant(t, http.StatusNoContent, `{"topics": [{"name": "topic-test", "policies": [{}]}]}`)

	err := permissions.Create(dmsClient(), instanceID, []permissions.CreateOpts{{
		Name:     "topic-test",
		Policies: []permissions.CreatePolicy{{}},
	}})
	th.AssertNoErr(t, err)
}

func TestListPermissions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(topicsPath+"/topic-test/accesspolicy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"name": "topic-test",
			"policies": [
				{"owner": false, "user_name": "xxxa", "access_policy": "pub"},
				{"owner": true, "user_name": "root", "access_policy": "all"}
			],
			"topic_type": 1
		}`)
	})

	actual, err := permissions.List(dmsClient(), instanceID, "topic-test")
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &permissions.Permissions{
		Name:      "topic-test",
		TopicType: 1,
		Policies: []permissions.Policy{
			{Owner: false, UserName: "xxxa", AccessPolicy: "pub"},
			{Owner: true, UserName: "root", AccessPolicy: "all"},
		},
	}, actual)
}
