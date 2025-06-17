package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/templates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	MESSAGE1 = "Test message 1"
	MESSAGE2 = "Test message 2"
)

func TestTemplatesWorkflow(t *testing.T) {
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	rName := tools.RandomString("tf-", 8)

	t.Logf("Attempting to create SMN template")
	createOpts := templates.CreateOpts{
		Name:     rName,
		Content:  MESSAGE1,
		Protocol: "email",
	}

	template, err := templates.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created SMN template: %s", template.MessageTemplateID)
	defer func() {
		t.Logf("Attempting to delete SMN template: %s", template.MessageTemplateID)
		err := templates.Delete(client, template.MessageTemplateID)
		th.AssertNoErr(t, err)
		t.Logf("Deleted SMN template: %s", template.MessageTemplateID)
	}()

	t.Logf("Attempting to update SMN template")
	updateOpts := templates.UpdateOpts{
		Content:    MESSAGE2,
		TemplateID: template.MessageTemplateID,
	}

	_, err = templates.Update(client, updateOpts)
	th.AssertNoErr(t, err)

	t.Logf("Updated SMN template: %s", template.MessageTemplateID)

	t.Logf("Attempting to retrieve SMN template")

	getResp, err := templates.Get(client, template.MessageTemplateID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, MESSAGE2, getResp.Content)

	t.Logf("Reetrieved SMN template: %s", getResp.MessageTemplateID)

	t.Logf("Attempting to list SMN templates")

	templatesList, err := templates.List(client, templates.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(templatesList.MessageTemplates) > 0)

	t.Logf("Retrived SMN templates: %d", len(templatesList.MessageTemplates))
}
