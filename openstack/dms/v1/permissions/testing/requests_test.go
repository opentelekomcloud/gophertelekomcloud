package testing

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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

func failOnRequest(t *testing.T) {
	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})
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

func TestGrantPermissionsValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	failOnRequest(t)

	policy := permissions.CreatePolicy{UserName: "root", AccessPolicy: "all"}
	valid := []permissions.CreateOpts{{Name: "topic-test", Policies: []permissions.CreatePolicy{policy}}}

	var missing golangsdk.ErrMissingInput
	err := permissions.Create(dmsClient(), "", valid)
	if !errors.As(err, &missing) || missing.Argument != "instanceId" {
		t.Fatalf("expected missing instanceId, got %v", err)
	}
	for _, opts := range [][]permissions.CreateOpts{nil, {}} {
		err = permissions.Create(dmsClient(), instanceID, opts)
		if !errors.As(err, &missing) || missing.Argument != "topics" {
			t.Fatalf("expected missing topics, got %v", err)
		}
	}

	cases := []struct {
		opts     []permissions.CreateOpts
		contains []string
	}{
		{[]permissions.CreateOpts{{Policies: []permissions.CreatePolicy{policy}}}, []string{"topics[0]", "Name"}},
		{[]permissions.CreateOpts{valid[0], {Name: "other"}}, []string{"topics[1]", "Policies"}},
	}
	for _, c := range cases {
		err = permissions.Create(dmsClient(), instanceID, c.opts)
		if err == nil {
			t.Fatalf("expected error containing %v, got nil", c.contains)
		}
		for _, s := range c.contains {
			if !strings.Contains(err.Error(), s) {
				t.Errorf("expected error containing %q, got %v", s, err)
			}
		}
	}
}

func TestGrantPermissionsForbidden(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(grantPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprint(w, `{"error_code": "DMS.00403001", "error_msg": "Authentication failed."}`)
	})

	err := permissions.Create(dmsClient(), instanceID, []permissions.CreateOpts{{
		Name:     "topic-test",
		Policies: []permissions.CreatePolicy{{UserName: "root", AccessPolicy: "all"}},
	}})
	var forbidden golangsdk.ErrDefault403
	if !errors.As(err, &forbidden) {
		t.Fatalf("expected ErrDefault403, got %T: %v", err, err)
	}
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

func TestListPermissionsValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	failOnRequest(t)

	var missing golangsdk.ErrMissingInput
	_, err := permissions.List(dmsClient(), "", "topic-test")
	if !errors.As(err, &missing) || missing.Argument != "instanceId" {
		t.Fatalf("expected missing instanceId, got %v", err)
	}
	_, err = permissions.List(dmsClient(), instanceID, "")
	if !errors.As(err, &missing) || missing.Argument != "topicName" {
		t.Fatalf("expected missing topicName, got %v", err)
	}
}
