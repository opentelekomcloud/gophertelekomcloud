package util

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetFuncTemplate(client *golangsdk.ServiceClient, templateId string) (*FuncTemplateResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "templates", templateId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res FuncTemplateResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type FuncTemplateResp struct {
	Id                  string            `json:"id"`
	Type                string            `json:"type"`
	Title               string            `json:"title"`
	TemplateName        string            `json:"template_name"`
	Description         string            `json:"description"`
	Runtime             string            `json:"runtime"`
	Handler             string            `json:"handler"`
	CodeType            string            `json:"code_type"`
	Code                string            `json:"code"`
	Timeout             int               `json:"timeout"`
	MemorySize          int               `json:"memory_size"`
	TriggerMetadataList []TriggerMetadata `json:"trigger_metadata_list"`
	TempDetail          *TempDetail       `json:"temp_detail"`
	UserData            string            `json:"user_data"`
	EncryptedUserData   string            `json:"encrypted_user_data"`
	Dependencies        string            `json:"dependencies"`
	Scene               string            `json:"scene"`
	Service             string            `json:"service"`
}

type TriggerMetadata struct {
	TriggerName string `json:"trigger_name"`
	TriggerType string `json:"trigger_type"`
	EventType   string `json:"event_type"`
	EventData   string `json:"event_data"`
}

type TempDetail struct {
	Input   string `json:"input"`
	Output  string `json:"output"`
	Warning string `json:"warning"`
}
