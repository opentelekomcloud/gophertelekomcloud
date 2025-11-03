package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/job"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/template"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGeminiTemplateLifecycle(t *testing.T) {
	instanceId := os.Getenv("OS_INSTANCE_ID")
	if instanceId == "" {
		t.Skip("OS_INSTANCE_ID is required for backup test")
	}

	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create gemini template")

	name := tools.RandomString("test-template-", 3)

	createOpts := template.CreateOpts{
		Name:        name,
		Description: "test template",
		Values: map[string]string{
			"max_connections": "10",
			"autocommit":      "OFF",
		},
		DataStore: template.DataStoreOpt{
			Type:    "cassandra",
			Version: "3.11",
		},
	}

	createResp, err := template.Create(client, createOpts)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, createResp.Name, createOpts.Name)
	th.AssertEquals(t, createResp.Description, createOpts.Description)

	t.Cleanup(func() {
		t.Logf("Attempting to delete gemini db template")
		err = template.Delete(client, createResp.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to get gemini template")

	getResp, err := template.Get(client, createResp.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getResp.Name, createResp.Name)

	t.Logf("Attempting to update gemini template")

	updateOpts := template.UpdateOpts{
		ConfigId: createResp.Id,
		Name:     name + "-updated",
		Values: map[string]string{
			"max_connections": "20",
			"autocommit":      "ON",
		},
	}

	err = template.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list gemini templates")

	listResp, err := template.List(client, template.ListOpts{})
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(listResp) > 0, true)

	t.Logf("Attempting to apply gemini template to an instance")

	applyResp, err := template.Apply(client, template.ApplyOpts{
		ConfigId: createResp.Id,
		InstanceIds: []string{
			instanceId,
		},
	})
	th.AssertNoErr(t, err)
	th.AssertNoErr(t, job.WaitForJobCompletion(client, 1200, applyResp.JobId))
}

func TestGeminiParametersLifecycle(t *testing.T) {
	instanceId := os.Getenv("OS_INSTANCE_ID")
	if instanceId == "" {
		t.Skip("OS_INSTANCE_ID is required for backup test")
	}

	client, err := clients.NewGeminiDBClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to update gemini instance parameters")

	_, err = template.UpdateInstanceParameters(client, template.UpdateParametersOpts{
		InstanceId: instanceId,
		Values: map[string]string{
			"request_timeout_in_ms": "10000",
		},
	})

	th.AssertNoErr(t, err)

	t.Logf("Attempting to get gemini instance parameters")

	getResp, err := template.GetInstanceParameters(client, instanceId)
	th.AssertNoErr(t, err)

	var res *template.InstanceParameterResult

	var found bool
	for _, i := range getResp.ConfigurationParameters {
		if i.Name == "request_timeout_in_ms" {
			paramCopy := i
			res = &paramCopy
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Parameter 'request_timeout_in_ms' not found in response")
	}

	th.AssertEquals(t, res.Value, "10000")
	th.AssertEquals(t, res.Name, "request_timeout_in_ms")
}
