package notifications_test

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/notifications"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	requests := 0
	th.Mux.HandleFunc("/scaling_notification/group-id", func(w http.ResponseWriter, r *http.Request) {
		requests++
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertEquals(t, "", r.URL.RawQuery)
		_, _ = w.Write([]byte(`{"topics":[` + notificationJSON + `,{
			"topic_urn":"urn:smn:eu-de:project-id:second-topic",
			"topic_scene":["SCALING_UP","SCALING_UP_FAIL"],"topic_name":"second-topic"
		}]}`))
	})
	actual, err := notifications.List(client.ServiceClient(), "group-id")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, requests)
	th.AssertDeepEquals(t, []notifications.Notification{
		{
			TopicURN: topicURN,
			TopicScene: []string{
				"SCALING_UP", "SCALING_UP_FAIL", "SCALING_DOWN", "SCALING_DOWN_FAIL", "SCALING_GROUP_ABNORMAL",
			},
			TopicName: "test-topic",
		},
		{
			TopicURN:   "urn:smn:eu-de:project-id:second-topic",
			TopicScene: []string{"SCALING_UP", "SCALING_UP_FAIL"}, TopicName: "second-topic",
		},
	}, actual)
}

func TestListEmpty(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/scaling_notification/group-id", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"topics":[]}`))
	})
	actual, err := notifications.List(client.ServiceClient(), "group-id")
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, []notifications.Notification{}, actual)
}

func TestListErrors(t *testing.T) {
	for _, status := range []int{http.StatusNotFound, http.StatusForbidden, http.StatusOK} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			th.SetupHTTP()
			defer th.TeardownHTTP()
			th.Mux.HandleFunc("/scaling_notification/group-id", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"topics":`))
			})
			actual, err := notifications.List(client.ServiceClient(), "group-id")
			if err == nil || actual != nil {
				t.Fatalf("expected error and nil notifications, got %#v, %v", actual, err)
			}
		})
	}
}

func TestListInvalidGroupID(t *testing.T) {
	actual, err := notifications.List(client.ServiceClient(), "../scaling_group/other-group")
	if err == nil || actual != nil {
		t.Fatalf("expected invalid group ID error, got %#v, %v", actual, err)
	}
}
