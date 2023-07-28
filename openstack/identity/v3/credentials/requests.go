package credentials

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type Status string

const parentElement = "credential"

type ListOptsBuilder interface {
	ToCredentialListQuery() (string, error)
}

type ListOpts struct {
	UserID string `json:"user_id,omitempty"`
}

func (opts ListOpts) ToCredentialListQuery() (string, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) (l ListResult) {
	q, err := opts.ToCredentialListQuery()
	if err != nil {
		l.Err = err
		return
	}
	_, l.Err = client.Get(listURL(client)+q, &l.Body, nil)
	return
}

func Get(client *golangsdk.ServiceClient, credentialID string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, credentialID), &r.Body, nil)
	return
}

type CreateOptsBuilder interface {
	ToCredentialCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	UserID      string `json:"user_id"`
	Description string `json:"description"`
}

func (opts CreateOpts) ToCredentialCreateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, parentElement)
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCredentialCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToCredentialUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

func (opts UpdateOpts) ToCredentialUpdateMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, parentElement)
}

func Update(client *golangsdk.ServiceClient, credentialID string, opts UpdateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCredentialUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, credentialID), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, credentialID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, credentialID), nil)
	return
}

type CreateTemporaryOptsBuilder interface {
	ToTempCredentialCreateMap() (map[string]interface{}, error)
}

type CreateTemporaryOpts struct {
	// For obtaining a temporary AK/SK with an agency token for use "assume_role"
	// For user with federated token fill "token" in this field
	Methods []string `json:"methods"`

	// Common token or federated token required for obtaining a temporary AK/SK.
	// You need to choose either the ID in this object or X-Auth-Token in the request header.
	// X-Auth-Token takes priority over the ID in this object.
	Token string `json:"token,omitempty"`
	// Validity period (in seconds) of an AK/SK and security token.
	// The value ranges from 15 minutes to 24 hours.
	// The default value is 15 minutes.
	Duration int `json:"duration-seconds,omitempty"`

	// Name or ID of the domain to which the delegating party belongs
	DomainName string `json:"domain_name,omitempty"`
	DomainID   string `json:"domain_id,omitempty"`

	// Name of the agency created by a delegating party
	AgencyName string `json:"agency_name,omitempty"`
}

// ToTempCredentialCreateMap
func (opts CreateTemporaryOpts) ToTempCredentialCreateMap() (map[string]interface{}, error) {
	// doc: https://docs.otc.t-systems.com/en-us/api/iam/en-us_topic_0097949518.html

	if len(opts.Methods) < 1 {
		return nil, fmt.Errorf("no auth method provided")
	}
	if len(opts.Methods) > 1 {
		return nil, fmt.Errorf("more than one auth method provided")
	}

	method := opts.Methods[0]

	authMap := map[string]interface{}{
		"auth": map[string]interface{}{
			"identity": map[string]interface{}{
				"methods": opts.Methods,
			},
		},
	}

	switch method {
	case "token":
		authMap["token"] = map[string]interface{}{
			"id":               opts.Token,
			"duration-seconds": opts.Duration,
		}
	case "assume_role":
		role := map[string]interface{}{
			"agency_name":      opts.AgencyName,
			"duration-seconds": opts.Duration,
		}
		switch {
		case opts.DomainID != "":
			role["domain_id"] = opts.DomainID
		case opts.DomainName != "":
			role["domain_name"] = opts.DomainName
		default:
			return nil, fmt.Errorf("you need to provide either delegating domain ID or Name")
		}
		authMap["assume_role"] = role
	default:
		return nil, fmt.Errorf("unknown auth method provided: %s", method)
	}

	return authMap, nil
}

func CreateTemporary(client *golangsdk.ServiceClient, opts CreateTemporaryOptsBuilder) (r CreateTemporaryResult) {
	json, err := opts.ToTempCredentialCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createTempURL(client), &json, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	return
}
