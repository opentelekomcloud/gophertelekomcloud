package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3.0/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const CreateUserWithExternalIdentityRequest = `
{
	"user": {
		"domain_id": "1789d1",
		"enabled": true,
		"name": "jsmith",
		"xuser_id": "external-user-123",
		"xuser_type": "TenantIdp"
	}
}
`

const CreateUserWithExternalIdentityResponse = `
{
	"user": {
		"id": "9fe1d3",
		"domain_id": "1789d1",
		"enabled": true,
		"name": "jsmith",
		"xuser_id": "external-user-123",
		"xuser_type": "TenantIdp"
	}
}
`

const ModifyUserAdminWithExternalIdentityRequest = `
{
	"user": {
		"xuser_id": "external-user-456",
		"xuser_type": "TenantIdp"
	}
}
`

const ModifyUserAdminWithExternalIdentityResponse = `
{
	"user": {
		"id": "9fe1d3",
		"domain_id": "1789d1",
		"enabled": true,
		"name": "jsmith",
		"xuser_id": "external-user-456",
		"xuser_type": "TenantIdp"
	}
}
`

const ModifyUserAdminClearExternalIdentityRequest = `
{
	"user": {
		"xuser_id": "",
		"xuser_type": ""
	}
}
`

const ModifyUserAdminClearExternalIdentityResponse = `
{
	"user": {
		"id": "9fe1d3",
		"domain_id": "1789d1",
		"enabled": true,
		"name": "jsmith",
		"xuser_id": "",
		"xuser_type": ""
	}
}
`

var CreatedUserWithExternalIdentity = users.User{
	ID:        "9fe1d3",
	DomainID:  "1789d1",
	Enabled:   true,
	Name:      "jsmith",
	XuserID:   "external-user-123",
	XuserType: "TenantIdp",
}

var UpdatedUserWithExternalIdentity = users.User{
	ID:        "9fe1d3",
	DomainID:  "1789d1",
	Enabled:   true,
	Name:      "jsmith",
	XuserID:   "external-user-456",
	XuserType: "TenantIdp",
}

var UpdatedUserWithClearedExternalIdentity = users.User{
	ID:        "9fe1d3",
	DomainID:  "1789d1",
	Enabled:   true,
	Name:      "jsmith",
	XuserID:   "",
	XuserType: "",
}

func HandleCreateUserWithExternalIdentitySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-USER/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateUserWithExternalIdentityRequest)

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, CreateUserWithExternalIdentityResponse)
	})
}

func HandleModifyUserAdminWithExternalIdentitySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-USER/users/9fe1d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, ModifyUserAdminWithExternalIdentityRequest)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, ModifyUserAdminWithExternalIdentityResponse)
	})
}

func HandleModifyUserAdminClearExternalIdentitySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/OS-USER/users/9fe1d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, ModifyUserAdminClearExternalIdentityRequest)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, ModifyUserAdminClearExternalIdentityResponse)
	})
}
