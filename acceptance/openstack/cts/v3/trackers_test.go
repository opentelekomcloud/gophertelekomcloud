package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v3/keyevent"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestKeyEventLifecycle(t *testing.T) {
	client, err := clients.NewCTSV3Client()
	th.AssertNoErr(t, err)

	event, err := keyevent.Create(client, keyevent.CreateNotificationOpts{
		NotificationName: "keyevent_test_notification",
		OperationType:    "customized",
		Operations: []keyevent.Operations{
			{
				ServiceType:  "OBS",
				ResourceType: "bucket",
				TraceNames:   []string{"createBucket"},
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = keyevent.Delete(client, []string{event.NotificationId})
		th.AssertNoErr(t, err)
	})

	list, err := keyevent.List(client, keyevent.ListNotificationsOpts{
		NotificationType: "smn",
		NotificationName: event.NotificationName,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, list)

	update, err := keyevent.Update(client, keyevent.UpdateNotificationOpts{
		NotificationName: "keyevent_test_update",
		Status:           "disabled",
		OperationType:    "customized",
		NotificationId:   event.NotificationId,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, update)
}
