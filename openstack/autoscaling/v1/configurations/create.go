package configurations

import (
	"encoding/base64"
	"log"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type CreateOptsBuilder interface {
	ToConfigurationCreateMap() (map[string]interface{}, error)
}

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

func (opts CreateOpts) ToConfigurationCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] ToConfigurationCreateMap b is: %#v", b)
	log.Printf("[DEBUG] ToConfigurationCreateMap opts is: %#v", opts)

	if opts.InstanceConfig.UserData != nil {
		var userData string
		if _, err := base64.StdEncoding.DecodeString(string(opts.InstanceConfig.UserData)); err != nil {
			userData = base64.StdEncoding.EncodeToString(opts.InstanceConfig.UserData)
		} else {
			userData = string(opts.InstanceConfig.UserData)
		}
		b["instance_config"].(map[string]interface{})["user_data"] = &userData
	}
	log.Printf("[DEBUG] ToConfigurationCreateMap b is: %#v", b)
	return b, nil
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToConfigurationCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("scaling_configuration"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type CreateResult struct {
	golangsdk.Result
}

func (r CreateResult) Extract() (string, error) {
	var a struct {
		ID string `json:"scaling_configuration_id"`
	}

	err := r.Result.ExtractInto(&a)
	return a.ID, err
}
