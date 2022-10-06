package configurations

import (
	"encoding/base64"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name           string             `json:"scaling_configuration_name" required:"true"`
	InstanceConfig InstanceConfigOpts `json:"instance_config" required:"true"`
}

type InstanceConfigOpts struct {
	ID             string                 `json:"instance_id,omitempty"`
	FlavorRef      string                 `json:"flavorRef,omitempty"`
	ImageRef       string                 `json:"imageRef,omitempty"`
	Disk           []Disk                 `json:"disk,omitempty"`
	SSHKey         string                 `json:"key_name" required:"true"`
	Personality    []Personality          `json:"personality,omitempty"`
	PubicIp        *PublicIp              `json:"public_ip,omitempty"`
	UserData       []byte                 `json:"-"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	SecurityGroups []SecurityGroup        `json:"security_groups,omitempty"`
	MarketType     string                 `json:"market_type,omitempty"`
}

func (opts CreateOpts) toConfigurationCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.InstanceConfig.UserData != nil {
		var userData string
		if _, err := base64.StdEncoding.DecodeString(string(opts.InstanceConfig.UserData)); err != nil {
			userData = base64.StdEncoding.EncodeToString(opts.InstanceConfig.UserData)
		} else {
			userData = string(opts.InstanceConfig.UserData)
		}
		b["instance_config"].(map[string]interface{})["user_data"] = &userData
	}

	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := opts.toConfigurationCreateMap()
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("scaling_configuration"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"scaling_configuration_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
