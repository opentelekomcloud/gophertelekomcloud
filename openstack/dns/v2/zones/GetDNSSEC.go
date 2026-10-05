package zones

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DNSSEC is the DNSSEC configuration of a public zone. Only Status is set while DNSSEC is disabled.
type DNSSEC struct {
	ZoneName        string `json:"zone_name"`
	Status          string `json:"status"`
	KeyTag          int    `json:"key_tag"`
	Flag            int    `json:"flag"`
	DigestAlgorithm string `json:"digest_algorithm"`
	DigestType      int    `json:"digest_type"`
	Digest          string `json:"digest"`
	Signature       string `json:"signature"`
	SignatureType   int    `json:"signature_type"`
	KSKPublicKey    string `json:"ksk_public_key"`
	DSRecord        string `json:"ds_record"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// GetDNSSEC returns the DNSSEC configuration of a public zone.
func GetDNSSEC(client *golangsdk.ServiceClient, zoneID string) (*DNSSEC, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("zones", zoneID, "dnssec").Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}
	var res DNSSEC
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
