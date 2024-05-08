package v1_1

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/connection"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const connectionName = "testConnection"

func TestDataArtsConnectionsLifecycle(t *testing.T) {
	client, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	t.Log("create a connection")

	createOpts := connection.Connection{
		Name:        scriptName,
		Type:        connection.TypeDLI,
		Description: "test description",
	}

	err = connection.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule connection cleanup")
	t.Cleanup(func() {
		t.Logf("attempting to delete connection: %s", connectionName)
		err := connection.Delete(client, connectionName, "")
		th.AssertNoErr(t, err)
		t.Logf("connection is deleted: %s", connectionName)
	})

	t.Log("should wait 5 seconds")
	time.Sleep(5 * time.Second)
	t.Log("get connection")

	storedConnection, err := connection.Get(client, connectionName, "")
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedConnection)

	t.Log("modify connection")

	createOpts.Description = "new test description"

	err = connection.Update(client, createOpts, connection.UpdateOpts{})
	th.AssertNoErr(t, err)

	t.Log("get connection")

	storedConnection, err = connection.Get(client, "testConnection1", "")
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedConnection)

	th.CheckEquals(t, "new test desctiption", storedConnection)
}
