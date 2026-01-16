package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v2/templates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAlarmTemplatesCRUD(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create alarm template")
	createOpts := templates.CreateOpts{
		TemplateName:        "test-template-acc",
		TemplateDescription: "Test alarm template for acceptance tests",
		Policies: []templates.Policy{
			{
				Namespace:          "SYS.ECS",
				DimensionName:      "instance_id",
				MetricName:         "cpu_util",
				Period:             300,
				Filter:             "average",
				ComparisonOperator: ">",
				Value:              80,
				Unit:               "%",
				Count:              3,
				AlarmLevel:         2,
				SuppressDuration:   300,
			},
		},
	}

	templateId, err := templates.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created alarm template: %s", templateId)

	t.Cleanup(func() {
		t.Log("Attempting to delete alarm template")
		_, err := templates.Delete(client, templates.DeleteOpts{
			TemplateIds:          []string{templateId},
			DeleteAssociateAlarm: false,
		})
		th.AssertNoErr(t, err)
	})

	t.Log("Attempting to get alarm template details")
	template, err := templates.Get(client, templateId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, template.TemplateName, "test-template-acc")
	th.AssertEquals(t, template.TemplateDescription, "Test alarm template for acceptance tests")
	th.AssertEquals(t, len(template.Policies), 1)
	th.AssertEquals(t, template.Policies[0].MetricName, "cpu_util")

	t.Log("Attempting to list alarm templates")
	listResp, err := templates.List(client, templates.ListOpts{
		TemplateName: "test-template-acc",
		TemplateType: "custom",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, listResp.Count >= 1, true)

	t.Log("Attempting to update alarm template")
	err = templates.Update(client, templateId, templates.UpdateOpts{
		TemplateName:        "test-template-acc-updated",
		TemplateDescription: "Updated description",
		Policies: []templates.Policy{
			{
				Namespace:          "SYS.ECS",
				DimensionName:      "instance_id",
				MetricName:         "cpu_util",
				Period:             300,
				Filter:             "average",
				ComparisonOperator: ">=",
				Value:              90,
				Unit:               "%",
				Count:              5,
				AlarmLevel:         1,
				SuppressDuration:   600,
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Log("Attempting to verify template was updated")
	template, err = templates.Get(client, templateId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, template.TemplateName, "test-template-acc-updated")
	th.AssertEquals(t, template.TemplateDescription, "Updated description")
	th.AssertEquals(t, template.Policies[0].ComparisonOperator, ">=")
	th.AssertEquals(t, template.Policies[0].Value, float64(90))
	th.AssertEquals(t, template.Policies[0].AlarmLevel, 1)
}

func TestAlarmTemplatesList(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to list system alarm templates")
	listResp, err := templates.List(client, templates.ListOpts{
		TemplateType: "system",
		Limit:        10,
	})
	th.AssertNoErr(t, err)

	t.Logf("Found %d system alarm templates", listResp.Count)
	for _, tmpl := range listResp.AlarmTemplates {
		t.Logf("Template ID: %s, Name: %s, Type: %s",
			tmpl.TemplateId, tmpl.TemplateName, tmpl.TemplateType)
	}
}

func TestAlarmTemplateAssociationAlarms(t *testing.T) {
	client, err := clients.NewCesV2Client()
	th.AssertNoErr(t, err)

	t.Log("Attempting to create alarm template")
	createOpts := templates.CreateOpts{
		TemplateName:        "test-template-assoc-acc",
		TemplateDescription: "Test template for association alarms",
		Policies: []templates.Policy{
			{
				Namespace:          "SYS.ECS",
				DimensionName:      "instance_id",
				MetricName:         "cpu_util",
				Period:             300,
				Filter:             "average",
				ComparisonOperator: ">",
				Value:              80,
				Unit:               "%",
				Count:              3,
				AlarmLevel:         2,
				SuppressDuration:   300,
			},
		},
	}

	templateId, err := templates.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Log("Attempting to delete alarm template")
		_, err := templates.Delete(client, templates.DeleteOpts{
			TemplateIds:          []string{templateId},
			DeleteAssociateAlarm: true,
		})
		th.AssertNoErr(t, err)
	})

	t.Log("Attempting to list association alarms")
	listResp, err := templates.ListAssociationAlarms(client, templateId, templates.ListAssociationAlarmsOpts{
		Limit: 10,
	})
	th.AssertNoErr(t, err)

	t.Logf("Found %d associated alarm rules", listResp.Count)
	for _, alarm := range listResp.Alarms {
		t.Logf("Alarm ID: %s, Name: %s", alarm.AlarmId, alarm.Name)
	}
}
