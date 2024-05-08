package v1_1

import (
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1/script"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const scriptName = "testScript"

func TestDataArtsScriptsLifecycle(t *testing.T) {
	client, err := clients.NewDataArtsV1Client()
	th.AssertNoErr(t, err)

	t.Log("create a script")

	createOpts := script.Script{
		Name:           scriptName,
		Type:           "Shell",
		Content:        "echo 123456",
		ConnectionName: "anyConnection",
	}

	// st, err := script.Get(client, scriptName, "")
	// tools.PrintResource(t, st)

	err = script.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Log("schedule script cleanup")
	t.Cleanup(func() {
		t.Logf("attempting to delete script: %s", scriptName)
		err := script.Delete(client, scriptName, nil)
		th.AssertNoErr(t, err)
		t.Logf("script is deleted: %s", scriptName)
	})

	t.Log("should wait 5 seconds")
	time.Sleep(5 * time.Second)
	t.Log("get script")

	// TODO: remove this block before a merge request.
	storedScripts, err := script.List(client, script.ListOpts{}, "")
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedScripts)

	storedScript, err := script.Get(client, scriptName, "")
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedScript)

	t.Log("modify script")

	createOpts.Content = "echo 123456789"

	err = script.Update(client, scriptName, createOpts)
	th.AssertNoErr(t, err)

	t.Log("get script")

	storedScript, err = script.Get(client, scriptName, "")
	th.AssertNoErr(t, err)
	tools.PrintResource(t, storedScript)
	th.CheckEquals(t, "echo 123456789", storedScript.Content)
}
