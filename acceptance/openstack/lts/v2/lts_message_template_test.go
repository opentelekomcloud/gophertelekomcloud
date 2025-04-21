package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	message_template "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/message-template"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsMessageTemplateLifecycle(t *testing.T) {
	clientV2, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-alarm-template-", 3)
	t.Logf("Attempting to Create LTS Message Template")
	messageTemplate, err := message_template.Create(clientV2, message_template.CreateOpts{
		DomainId:    clientV2.DomainID,
		Name:        name,
		Description: "test",
		Source:      "LTS",
		Language:    "en-us",
		Templates: []message_template.Templates{
			{
				Type:    "sms",
				Content: "Severity: ${event_severity};\nOccurred: ${starts_at};\nResource ID: ${resources};\nStatistical type: by keyword;\nExpression: $event.annotations.condition_expression;\nCurrent value: $event.annotations.current_value;\nStatistical period: $event.annotations.frequency;",
			},
			{
				Type:    "email",
				Content: "Severity: ${event_severity};\nOccurred: ${starts_at};\nAlarm source: $event.metadata.resource_provider;\nResource type: $event.metadata.resource_type;\nResource ID: ${resources};\nStatistical type: by keyword;\nExpression: $event.annotations.condition_expression;\nCurrent value: $event.annotations.current_value;\nStatistical period: $event.annotations.frequency;\nQuery time: $event.annotations.results[0].time;\nQuery log: $event.annotations.results[0].raw_results;",
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete LTS Message Template")
		err = message_template.Delete(clientV2, message_template.DeleteOpts{
			DomainId:      clientV2.DomainID,
			TemplateNames: []string{messageTemplate.Name},
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to List LTS Message Template")
	listTemplates, err := message_template.List(clientV2, clientV2.DomainID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(listTemplates) > 1)

	t.Logf("Attempting to Update LTS Message Template")
	messageTemplateUp, err := message_template.Update(clientV2, message_template.CreateOpts{
		DomainId:    clientV2.DomainID,
		Name:        name,
		Description: "test",
		Source:      "LTS",
		Language:    "en-us",
		Templates: []message_template.Templates{
			{
				Type:    "sms",
				Content: "Severity: ${event_severity};\nOccurred: ${starts_at};",
			},
			{
				Type:    "email",
				Content: "Severity: ${event_severity};\nOccurred: ${starts_at};",
			},
		},
	})
	th.AssertNoErr(t, err)
	t.Logf("Attempting to Get LTS Message Template")
	getTemplate, err := message_template.Get(clientV2, clientV2.DomainID, messageTemplateUp.Name)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "Severity: ${event_severity};\nOccurred: ${starts_at};", getTemplate.Templates[0].Content)

	// t.Logf("Attempting to Get Email preview of Message Template")
	// previewTemplate := "Severity: ${event_severity};\nOccurred: ${starts_at};\nAlarm source: $event.metadata.resource_provider;\nResource type: $event.metadata.resource_type;\nResource ID: ${resources};\nStatistical type: by keyword;\nExpression: $event.annotations.condition_expression;\nCurrent value: $event.annotations.current_value;\nStatistical period: $event.annotations.frequency;\nQuery time: $event.annotations.results[0].time;\nQuery log: $event.annotations.results[0].raw_results;"
	// previewSubject := "${region_name}[${event_severity}_${event_type}_${clear_type}] generated an alarm at ${starts_at}"
	// preview, err := message_template.Preview(clientV2, message_template.PreviewOpts{
	// 	DomainId: clientV2.DomainID,
	// 	Template: previewTemplate,
	// 	Language: "en-us",
	// 	Source:   "LTS",
	// 	Subject:  previewSubject,
	// })
	// th.AssertNoErr(t, err)
	// th.AssertEquals(t, "LTS", preview.Template)
	// th.AssertEquals(t, previewSubject, preview.Subject)
	// th.AssertEquals(t, previewTemplate, preview.Template)
}
