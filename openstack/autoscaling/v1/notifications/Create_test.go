package notifications_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/notifications"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const topicURN = "urn:smn:eu-de:project-id:test-topic"

const notificationJSON = `{
	"topic_urn":"urn:smn:eu-de:project-id:test-topic",
	"topic_scene":["SCALING_UP","SCALING_UP_FAIL","SCALING_DOWN","SCALING_DOWN_FAIL","SCALING_GROUP_ABNORMAL"],
	"topic_name":"test-topic"
}`

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/scaling_notification/group-id", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodPut)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"topic_urn":"urn:smn:eu-de:project-id:test-topic","topic_scene":["SCALING_UP","SCALING_UP_FAIL","SCALING_DOWN","SCALING_DOWN_FAIL","SCALING_GROUP_ABNORMAL"]}`)
		_, _ = w.Write([]byte(notificationJSON))
	})
	scenes := []string{"SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL"}
	actual, err := notifications.Create(client.ServiceClient(), notifications.CreateOpts{
		TopicURN: topicURN, TopicScene: scenes,
	}, "group-id")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, notifications.Notification{
		TopicURN: topicURN, TopicScene: scenes, TopicName: "test-topic",
	}, *actual)
}

func TestCreateRequiredFields(t *testing.T) {
	for _, opts := range []notifications.CreateOpts{
		{},
		{TopicURN: topicURN},
		{TopicScene: []string{"SCALING_UP"}},
	} {
		actual, err := notifications.Create(client.ServiceClient(), opts, "group-id")
		if err == nil || actual != nil {
			t.Fatalf("expected required field error, got %#v, %v", actual, err)
		}
	}
}

func TestCreateErrors(t *testing.T) {
	for _, status := range []int{http.StatusBadRequest, http.StatusForbidden, http.StatusConflict, http.StatusOK} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/scaling_notification/group-id", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"topic_urn":`))
			})
			actual, err := notifications.Create(client.ServiceClient(), notifications.CreateOpts{
				TopicURN: topicURN, TopicScene: []string{"SCALING_UP"},
			}, "group-id")
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil notification, got %#v, %v", actual, err)
			}
		})
	}
}

func TestCreateInvalidGroupID(t *testing.T) {
	actual, err := notifications.Create(client.ServiceClient(), notifications.CreateOpts{
		TopicURN: topicURN, TopicScene: []string{"SCALING_UP"},
	}, "../scaling_group/other-group")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid group ID error, got %#v, %v", actual, err)
	}
}
