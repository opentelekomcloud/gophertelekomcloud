package testing

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/messages"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const (
	projectID     = "0123456789abcdef0123456789abcdef"
	instanceID    = "54602a9d-5e22-4239-9123-77e350df4a34"
	topic         = "mytest"
	instancePath  = "/v2/" + projectID + "/instances/" + instanceID
	topicPath     = instancePath + "/management/topics/" + topic
	partitionPath = topicPath + "/partitions/0"
)

func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

func failOnRequest(t *testing.T) {
	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})
}

func handleGet(t *testing.T, path, expectedQuery string, status int, body string) {
	th.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertEquals(t, expectedQuery, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = fmt.Fprint(w, body)
	})
}

func assertMissingInput(t *testing.T, err error, argument string) {
	t.Helper()
	var missing golangsdk.ErrMissingInput
	if !errors.As(err, &missing) {
		t.Fatalf("expected ErrMissingInput, got %T: %v", err, err)
	}
	th.AssertEquals(t, argument, missing.Argument)
}

const badTopic = "bad/topic"

func assertInvalidEndpoint(t *testing.T, err error) {
	t.Helper()
	if err == nil || !strings.Contains(err.Error(), "not possible in endpoints") {
		t.Fatalf("expected invalid endpoint error, got %v", err)
	}
}

func assertRequiredQuery(t *testing.T, err error, field string) {
	t.Helper()
	th.AssertEquals(t, fmt.Sprintf("required query parameter [%s] not set", field), fmt.Sprint(err))
}

const errorBody = `{"error_code": "DMS.00400001", "error_msg": "Invalid parameters."}`

const showInstanceMessagesResponse = `
{
  "messages" : [ {
    "topic" : "topic-test",
    "partition" : 0,
    "key" : "k1",
    "value" : "hello world",
    "size" : 21,
    "timestamp" : 1607598463502,
    "huge_message" : true,
    "message_offset" : 4,
    "message_id" : "id-1",
    "app_id" : "app-1",
    "tag" : "tag-1"
  } ],
  "total" : 1,
  "size" : 1
}`

func TestShowInstanceMessagesOffsetQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, instancePath+"/messages",
		"asc=false&end_time=1608609032042&limit=10&offset=0&start_time=1608608432042&topic=topic-test",
		http.StatusOK, showInstanceMessagesResponse)

	asc := false
	res, err := messages.ShowInstanceMessages(dmsClient(), instanceID, messages.ShowInstanceMessagesOpts{
		Topic:     "topic-test",
		Asc:       &asc,
		StartTime: "1608608432042",
		EndTime:   "1608609032042",
		Limit:     "10",
		Offset:    "0",
	})
	th.AssertNoErr(t, err)

	var _ int64 = res.Total
	var _ int64 = res.Size
	var _ int64 = res.Messages[0].MessageOffset
	th.AssertDeepEquals(t, &messages.ShowInstanceMessagesResponse{
		Messages: []messages.MessagesEntity{{
			Topic:         "topic-test",
			Partition:     0,
			Key:           "k1",
			Value:         "hello world",
			Size:          21,
			Timestamp:     1607598463502,
			HugeMessage:   true,
			MessageOffset: 4,
			MessageID:     "id-1",
			AppID:         "app-1",
			Tag:           "tag-1",
		}},
		Total: 1,
		Size:  1,
	}, res)
}

func TestShowInstanceMessagesContentQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, instancePath+"/messages",
		"download=false&keyword=hello&message_offset=0&partition=0&topic=topic-test",
		http.StatusOK, `{"messages": [], "total": 0, "size": 0}`)

	download := false
	res, err := messages.ShowInstanceMessages(dmsClient(), instanceID, messages.ShowInstanceMessagesOpts{
		Topic:         "topic-test",
		Download:      &download,
		MessageOffset: "0",
		Partition:     "0",
		Keyword:       "hello",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(res.Messages))
}

func TestShowInstanceMessagesOmitsUnsetQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, instancePath+"/messages", "topic=topic-test", http.StatusOK, `{"messages": []}`)

	_, err := messages.ShowInstanceMessages(dmsClient(), instanceID, messages.ShowInstanceMessagesOpts{Topic: "topic-test"})
	th.AssertNoErr(t, err)
}

func TestShowInstanceMessagesValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	failOnRequest(t)

	_, err := messages.ShowInstanceMessages(dmsClient(), "", messages.ShowInstanceMessagesOpts{Topic: "topic-test"})
	assertMissingInput(t, err, "instanceID")

	_, err = messages.ShowInstanceMessages(dmsClient(), instanceID, messages.ShowInstanceMessagesOpts{})
	assertRequiredQuery(t, err, "Topic")
}

func TestShowInstanceMessagesBadRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, instancePath+"/messages", "topic=bad", http.StatusBadRequest, errorBody)

	_, err := messages.ShowInstanceMessages(dmsClient(), instanceID, messages.ShowInstanceMessagesOpts{Topic: "bad"})
	var badRequest golangsdk.ErrDefault400
	if !errors.As(err, &badRequest) {
		t.Fatalf("expected ErrDefault400, got %T: %v", err, err)
	}
}

func TestShowInstanceMessagesForbidden(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, instancePath+"/messages", "topic=topic-test", http.StatusForbidden, errorBody)

	_, err := messages.ShowInstanceMessages(dmsClient(), instanceID, messages.ShowInstanceMessagesOpts{Topic: "topic-test"})
	var forbidden golangsdk.ErrDefault403
	if !errors.As(err, &forbidden) {
		t.Fatalf("expected ErrDefault403, got %T: %v", err, err)
	}
}

func TestShowMessages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, topicPath+"/messages",
		"end_time=1568125039164&limit=10&offset=0&partition=-1&start_time=0",
		http.StatusOK, `
{
  "messages" : [ {
    "topic" : "mytest",
    "partition" : 0,
    "message_offset" : 7,
    "size" : 6,
    "timestamp" : 1568125036045
  } ],
  "messages_count" : 1,
  "offsets_count" : 1,
  "offset" : 1
}`)

	offset := 0
	res, err := messages.ShowMessages(dmsClient(), instanceID, topic, messages.ShowMessagesOpts{
		StartTime: "0",
		EndTime:   "1568125039164",
		Limit:     10,
		Offset:    &offset,
		Partition: "-1",
	})
	th.AssertNoErr(t, err)

	var _ int64 = res.Messages[0].MessageOffset
	th.AssertDeepEquals(t, &messages.ShowMessagesResponse{
		Messages: []messages.Messages{{
			Topic:         "mytest",
			Partition:     0,
			MessageOffset: 7,
			Size:          6,
			Timestamp:     1568125036045,
		}},
		MessagesCount: 1,
		OffsetsCount:  1,
		Offset:        1,
	}, res)
}

func TestShowMessagesOmitsUnsetQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, topicPath+"/messages", "", http.StatusOK, `{"messages": []}`)

	_, err := messages.ShowMessages(dmsClient(), instanceID, topic, messages.ShowMessagesOpts{})
	th.AssertNoErr(t, err)
}

func TestShowMessagesValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	failOnRequest(t)

	_, err := messages.ShowMessages(dmsClient(), "", topic, messages.ShowMessagesOpts{})
	assertMissingInput(t, err, "instanceID")

	_, err = messages.ShowMessages(dmsClient(), instanceID, "", messages.ShowMessagesOpts{})
	assertMissingInput(t, err, "topic")

	_, err = messages.ShowMessages(dmsClient(), instanceID, badTopic, messages.ShowMessagesOpts{})
	assertInvalidEndpoint(t, err)
}

func TestShowPartitionMessage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, partitionPath+"/message", "message_offset=0", http.StatusOK, `
{
  "message" : [ {
    "topic" : "mytest",
    "partition" : 0,
    "message_offset" : 7,
    "key" : null,
    "value" : "kasjdf",
    "size" : 6,
    "timestamp" : 1568125036045
  } ]
}`)

	res, err := messages.ShowPartitionMessage(dmsClient(), instanceID, topic, 0, messages.ShowPartitionMessageOpts{
		MessageOffset: "0",
	})
	th.AssertNoErr(t, err)

	var _ int64 = res.Message[0].MessageOffset
	th.AssertDeepEquals(t, &messages.ShowPartitionMessageResponse{
		Message: []messages.ShowPartitionMessageEntity{{
			Key:           "",
			Value:         "kasjdf",
			Topic:         "mytest",
			Partition:     0,
			MessageOffset: 7,
			Size:          6,
			Timestamp:     1568125036045,
		}},
	}, res)
}

func TestShowPartitionMessageValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	failOnRequest(t)

	opts := messages.ShowPartitionMessageOpts{MessageOffset: "7"}

	_, err := messages.ShowPartitionMessage(dmsClient(), "", topic, 0, opts)
	assertMissingInput(t, err, "instanceID")

	_, err = messages.ShowPartitionMessage(dmsClient(), instanceID, "", 0, opts)
	assertMissingInput(t, err, "topic")

	_, err = messages.ShowPartitionMessage(dmsClient(), instanceID, badTopic, 0, opts)
	assertInvalidEndpoint(t, err)

	_, err = messages.ShowPartitionMessage(dmsClient(), instanceID, topic, 0, messages.ShowPartitionMessageOpts{})
	assertRequiredQuery(t, err, "MessageOffset")
}

const partitionOffsetResponse = `
{
  "topic" : "mytest",
  "partition" : 0,
  "offset" : 9,
  "timestamp" : 1568125039164
}`

func TestShowPartitionBeginningMessage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, partitionPath+"/beginning-message", "", http.StatusOK, partitionOffsetResponse)

	res, err := messages.ShowPartitionBeginningMessage(dmsClient(), instanceID, topic, 0)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &messages.ShowPartitionBeginningMessageResponse{
		Topic:     "mytest",
		Partition: 0,
		Offset:    9,
		Timestamp: 1568125039164,
	}, res)
}

func TestShowPartitionEndMessage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, partitionPath+"/end-message", "", http.StatusOK, partitionOffsetResponse)

	res, err := messages.ShowPartitionEndMessage(dmsClient(), instanceID, topic, 0)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &messages.ShowPartitionEndMessageResponse{
		Topic:     "mytest",
		Partition: 0,
		Offset:    9,
		Timestamp: 1568125039164,
	}, res)
}

func TestShowPartitionOffsetValidation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	failOnRequest(t)

	_, err := messages.ShowPartitionBeginningMessage(dmsClient(), "", topic, 0)
	assertMissingInput(t, err, "instanceID")
	_, err = messages.ShowPartitionBeginningMessage(dmsClient(), instanceID, "", 0)
	assertMissingInput(t, err, "topic")

	_, err = messages.ShowPartitionEndMessage(dmsClient(), "", topic, 0)
	assertMissingInput(t, err, "instanceID")
	_, err = messages.ShowPartitionEndMessage(dmsClient(), instanceID, "", 0)
	assertMissingInput(t, err, "topic")
}

func TestShowPartitionOffsetRejectsInvalidTopic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	failOnRequest(t)

	_, err := messages.ShowPartitionBeginningMessage(dmsClient(), instanceID, badTopic, 0)
	assertInvalidEndpoint(t, err)

	_, err = messages.ShowPartitionEndMessage(dmsClient(), instanceID, badTopic, 0)
	assertInvalidEndpoint(t, err)
}

func TestShowPartitionEndMessageNotFound(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	handleGet(t, partitionPath+"/end-message", "", http.StatusNotFound, errorBody)

	_, err := messages.ShowPartitionEndMessage(dmsClient(), instanceID, topic, 0)
	var notFound golangsdk.ErrDefault404
	if !errors.As(err, &notFound) {
		t.Fatalf("expected ErrDefault404, got %T: %v", err, err)
	}
}
