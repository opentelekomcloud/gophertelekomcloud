package alarm

import "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteKeywordRule(client *golangsdk.ServiceClient, ruleId string) (err error) {
	// DELETE /v2/{project_id}/lts/alarms/keywords-alarm-rule/{keywords_alarm_rule_id}
	_, err = client.Delete(client.ServiceURL("lts", "alarms", "keywords-alarm-rule", ruleId), &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	return
}
