package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/users"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID  = "0123456789abcdef0123456789abcdef"
	instanceID = "54602a9d-5e22-4239-9123-77e350df4a34"
	usersPath  = "/v2/" + projectID + "/instances/" + instanceID + "/users"
)

func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func handleNoContent(t *testing.T, path, method, expectedBody string) {
	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, method)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, expectedBody)
		w.WriteHeader(http.StatusNoContent)
	})
}

func TestCreateUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleNoContent(t, usersPath, "POST", `{"user_name": "test", "user_desc": "test user", "user_passwd": "Cxxx3xxx!"}`)

	err := users.Create(dmsClient(), instanceID, users.CreateOpts{
		UserName:   "test",
		UserDesc:   "test user",
		UserPasswd: "Cxxx3xxx!",
	})
	th.AssertNoErr(t, err)
}

func TestCreateUserOmitsEmptyFields(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleNoContent(t, usersPath, "POST", `{}`)

	th.AssertNoErr(t, users.Create(dmsClient(), instanceID, users.CreateOpts{}))
}

func TestDeleteUsers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleNoContent(t, usersPath, "PUT", `{"action": "delete", "users": ["testuser", "other"]}`)

	err := users.Delete(dmsClient(), instanceID, users.DeleteOpts{
		Action: "delete",
		Users:  []string{"testuser", "other"},
	})
	th.AssertNoErr(t, err)
}

func TestDeleteUsersOmitsEmptyFields(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleNoContent(t, usersPath, "PUT", `{}`)

	th.AssertNoErr(t, users.Delete(dmsClient(), instanceID, users.DeleteOpts{Users: []string{}}))
}

func TestListUsers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(usersPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
			"users": [
				{"user_name": "xxxa", "user_desc": "app user", "role": "guest", "default_app": false, "created_time": 1615431764734},
				{"user_name": "ROOT", "role": "admin", "default_app": true, "created_time": 1617194246328}
			]
		}`)
	})

	actual, err := users.List(dmsClient(), instanceID)
	th.AssertNoErr(t, err)
	var _ int64 = actual[0].CreatedTime
	th.CheckDeepEquals(t, []users.Users{
		{UserName: "xxxa", UserDesc: "app user", Role: "guest", DefaultApp: false, CreatedTime: 1615431764734},
		{UserName: "ROOT", Role: "admin", DefaultApp: true, CreatedTime: 1617194246328},
	}, actual)
}

func TestResetPassword(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleNoContent(t, usersPath+"/test", "PUT", `{"new_password": "Cxxx3xxx!"}`)

	err := users.ResetPassword(dmsClient(), instanceID, "test", users.ResetPasswordOpts{NewPassword: "Cxxx3xxx!"})
	th.AssertNoErr(t, err)
}

func TestResetPasswordOmitsEmptyFields(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleNoContent(t, usersPath+"/test", "PUT", `{}`)

	th.AssertNoErr(t, users.ResetPassword(dmsClient(), instanceID, "test", users.ResetPasswordOpts{}))
}
