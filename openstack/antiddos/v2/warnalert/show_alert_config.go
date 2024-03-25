package warnalert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowAlertConfig(client *golangsdk.ServiceClient) (*ShowAlertConfigResponse, error) {
	// GET /v2/{project_id}/warnalert/alertconfig/query
	raw, err := client.Get(client.ServiceURL("warnalert", "alertconfig", "query"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowAlertConfigResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowAlertConfigResponse struct {
	// ID of an alarm group
	TopicUrn string `json:"topic_urn"`
	// Description of an alarm group
	DisplayName string `json:"display_name"`
	// Alarm configuration
	WarnConfig AlertConfigRespWarnConfig `json:"warn_config"`
}

type AlertConfigRespWarnConfig struct {
	// DDoS attacks
	AntiDDoS bool `json:"antiDDoS"`
	// Web shells
	BackDoors bool `json:"back_doors"`
	// Brute force cracking (system logins, FTP, and DB)
	BruceForce bool `json:"bruce_force"`
	// Overly high rights of a database process
	HighPrivilege bool `json:"high_privilege"`
	// Alarms about remote logins
	RemoteLogin bool `json:"remote_login"`
	// Possible values:
	// 0: indicates that alarms are sent once a day.
	// 1: indicates that alarms are sent once every half hour.
	SendFrequency int `json:"send_frequency"`
	// Reserved field
	Waf bool `json:"waf,omitempty"`
	// Weak passwords (system and database)
	WeakPassword bool `json:"weak_password"`
}
