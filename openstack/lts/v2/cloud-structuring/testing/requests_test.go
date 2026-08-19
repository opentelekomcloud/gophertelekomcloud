package testing

import (
	"testing"

	cloud_structuring "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/cloud-structuring"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListCustom(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleListCustom(t)

	templates, err := cloud_structuring.List(fake.ServiceClient(), cloud_structuring.ListOpts{
		ID: "template-id",
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(templates))
	template := templates[0]
	th.AssertEquals(t, "2a473356cca5487f8373be89xxxxxxxx", template.ProjectId)
	th.AssertEquals(t, "custom-json", template.Name)
	th.AssertEquals(t, "json", template.Type)
	th.AssertEquals(t, "sample log", template.DemoLog)
	th.AssertEquals(t, "sample-label", template.DemoLabel)
	th.AssertEquals(t, int64(1641258099551), template.CreatedAt)
	th.AssertEquals(t, "43a8cc7b-b632-4c36-a65d-8150e98219f1", template.ID)

	th.AssertEquals(t, 1, len(template.DemoFields))
	demoField := template.DemoFields[0]
	th.AssertEquals(t, "message", demoField.Name)
	th.AssertEquals(t, "value", demoField.Content)
	th.AssertEquals(t, "string", demoField.Type)
	th.AssertEquals(t, false, demoField.IsAnalysis)
	th.AssertEquals(t, 0, demoField.Index)
	th.AssertEquals(t, "root", demoField.Relation)
	th.AssertEquals(t, "alias", demoField.UserDefinedName)

	th.AssertEquals(t, 1, len(template.TagFields))
	tagField := template.TagFields[0]
	th.AssertEquals(t, "host", tagField.Name)
	th.AssertEquals(t, "host-1", tagField.Content)
	th.AssertEquals(t, "string", tagField.Type)
	th.AssertEquals(t, true, tagField.IsAnalysis)
	th.AssertEquals(t, 0, tagField.Index)

	th.AssertEquals(t, "json", template.Rule.Type)
	th.AssertEquals(t, `{"layers":1}`, template.Rule.Param)
}

func TestListCustomError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleListCustomError(t)

	_, err := cloud_structuring.List(fake.ServiceClient(), cloud_structuring.ListOpts{})

	if err == nil {
		t.Fatal("expected server error")
	}
}
