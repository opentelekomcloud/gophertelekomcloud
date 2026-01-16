package templates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListAssociationAlarmsOpts contains the options for querying alarm rules associated with an alarm template.
type ListAssociationAlarmsOpts struct {
	// Specifies the pagination offset. Default: 0, Range: 0-10000
	Offset int `q:"offset,omitempty"`
	// Specifies the number of records on each page. Default: 100, Range: 1-100
	Limit int `q:"limit,omitempty"`
}

// ListAssociationAlarmsResponse contains the response from the ListAssociationAlarms request.
type ListAssociationAlarmsResponse struct {
	// Specifies the list of alarm rules associated with the alarm template.
	Alarms []AssociationAlarm `json:"alarms"`
	// Specifies the total number of alarm rules.
	Count int `json:"count"`
}

// AssociationAlarm represents an alarm rule associated with an alarm template.
type AssociationAlarm struct {
	// Specifies the alarm rule ID.
	AlarmId string `json:"alarm_id"`
	// Specifies the alarm rule name.
	Name string `json:"name"`
	// Specifies the alarm rule description.
	Description string `json:"description"`
}

// ListAssociationAlarms returns a list of alarm rules associated with an alarm template.
func ListAssociationAlarms(client *golangsdk.ServiceClient, templateId string, opts ListAssociationAlarmsOpts) (*ListAssociationAlarmsResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("alarm-templates", templateId, "association-alarms").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/alarm-templates/{template_id}/association-alarms
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ListAssociationAlarmsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
