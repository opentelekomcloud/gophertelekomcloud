package configurations

import (
	"encoding/base64"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the AS configuration name.
	// The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	Name string `json:"scaling_configuration_name" required:"true"`
	// Specifies the ECS configuration.
	InstanceConfig InstanceConfigOpts `json:"instance_config" required:"true"`
}

type InstanceConfigOpts struct {
	// Specifies the ECS ID. When using the existing ECS specifications as the template to create AS configurations,
	// specify this parameter. In this case, the flavorRef, imageRef, disk, and security_groups fields do not take effect.
	// If the instance_id field is not specified, flavorRef, imageRef, and disk fields are mandatory.
	ID string `json:"instance_id,omitempty"`
	// Specifies the ECS flavor ID. A maximum of 10 flavors can be selected. Use a comma (,) to separate multiple flavor IDs.
	// You can obtain its value from the API for querying details about flavors and extended flavor information.
	FlavorRef string `json:"flavorRef,omitempty"`
	// Specifies the image ID. Its value is the same as that of image_id for specifying the image selected during ECS creation.
	// Obtain the value using the API for querying images.
	ImageRef string `json:"imageRef,omitempty"`
	// Specifies the disk group information. System disks are mandatory and data disks are optional.
	Disk []Disk `json:"disk,omitempty"`
	// Specifies the name of the SSH key pair used to log in to the ECS.
	SSHKey string `json:"key_name" required:"true"`
	// Specifies information about the injected file. Only text files can be injected.
	// A maximum of five files can be injected at a time and the maximum size of each file is 1 KB.
	Personality []Personality `json:"personality,omitempty"`
	// Specifies the EIP of the ECS. The EIP can be configured in two ways.
	// Do not use an EIP. In this case, this parameter is unavailable.
	// Automatically assign an EIP. You need to specify the information about the new EIP.
	PubicIp *PublicIp `json:"public_ip,omitempty"`
	// Specifies the user data to be injected during the ECS creation process. Text, text files, and gzip files can be injected.
	// Constraints:
	// The content to be injected must be encoded with base64. The maximum size of the content to be injected (before encoding) is 32 KB.
	// Examples:
	// Linux
	// #! /bin/bash
	// echo user_test >> /home/user.txt
	// Windows
	// rem cmd
	// echo 111 > c:\aaa.txt
	UserData []byte `json:"-"`
	// Specifies the ECS metadata.
	Metadata AdminPassMetadata `json:"metadata,omitempty"`
	// Specifies security groups.
	// If the security group is specified both in the AS configuration and AS group, scaled ECS instances
	// will be added to the security group specified in the AS configuration.
	// If the security group is not specified in either of them, scaled ECS instances will be added to the default security group.
	// For your convenience, you are advised to specify the security group in the AS configuration.
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`
	// This parameter is reserved.
	MarketType string `json:"market_type,omitempty"`
}

type AdminPassMetadata struct {
	// Specifies the initial login password of the administrator account for logging in to an ECS using password authentication.
	// The Linux administrator is root, and the Windows administrator is Administrator.
	//
	// Password complexity requirements:
	// Consists of 8 to 26 characters.
	// Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters !@$%^-_=+[{}]:,./?
	// The password cannot contain the username or the username in reversed order.
	// The Windows ECS password cannot contain the username, the username in reversed order, or more than two consecutive characters in the username.
	AdminPass string `json:"admin_pass,omitempty"`
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
