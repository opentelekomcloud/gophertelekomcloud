package ssl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*SslInfo, error) {
	raw, err := client.Get(client.ServiceURL("instances", id, "ssl"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res SslInfo
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SslInfo struct {
	Enabled      bool   `json:"enabled"`
	Ip           string `json:"ip"`
	Port         string `json:"port"`
	DomainName   string `json:"domain_name"`
	SslExpiredAt string `json:"ssl_expired_at"`
	SslValidated bool   `json:"ssl_validated"`
}
