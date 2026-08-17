package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/favorites"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	expectedCreateRequest = `
{
  "log_group_id": "d91fff37-9d10-47f1-85de-c2840724908f",
  "log_group_name": "lts-group-sgq1",
  "log_stream_id": "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
  "log_stream_name": "lts-topic-xxxxtest",
  "eps_id": "0",
  "favorite_resource_id": "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
  "favorite_resource_type": "log_stream",
  "is_global": true
}`
	createResponse = `
{
  "create_time": 1669018970929,
  "eps_id": "0",
  "favorite_resource_id": "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
  "is_global": true,
  "favorite_resource_type": "LOG_STREAM",
  "log_group_id": "d91fff37-9d10-47f1-85de-c2840724908f",
  "log_group_name": "lts-group-sgq1",
  "log_stream_id": "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
  "log_stream_name": "lts-topic-xxxxtest",
  "project_id": "2a473356cca5487f8373be891bffc1cf"
}`
)

var createOpts = favorites.CreateOpts{
	EnterpriseProjectID: "0",
	ResourceID:          "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
	ResourceType:        "log_stream",
	LogGroupID:          "d91fff37-9d10-47f1-85de-c2840724908f",
	LogGroupName:        "lts-group-sgq1",
	LogStreamID:         "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
	LogStreamName:       "lts-topic-xxxxtest",
	IsGlobal:            true,
}

var expectedFavorite = &favorites.Favorite{
	CreateTime:          1669018970929,
	EnterpriseProjectID: "0",
	ResourceID:          "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
	ResourceType:        "LOG_STREAM",
	LogGroupID:          "d91fff37-9d10-47f1-85de-c2840724908f",
	LogGroupName:        "lts-group-sgq1",
	LogStreamID:         "f2fb0a2d-d4cd-4bc9-ac12-93c6d255883c",
	LogStreamName:       "lts-topic-xxxxtest",
	ProjectID:           "2a473356cca5487f8373be891bffc1cf",
	IsGlobal:            true,
}

func handleCreate(t *testing.T, status int, response string) {
	th.Mux.HandleFunc("/lts/favorite", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPost)
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json;charset=utf8")
		th.TestJSONRequest(t, r, expectedCreateRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = fmt.Fprint(w, response)
	})
}
