package zones_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

// Recorded from a signed OTC zone; key material shortened.
const dnssecEnabledJSON = `{
	"zone_name": "example.org.",
	"status": "ENABLE",
	"key_tag": 55451,
	"flag": 257,
	"digest_algorithm": "SHA256",
	"digest_type": 2,
	"digest": "be51f11d7e09131ebe9c0000000000000000000000000000000000000000abcd",
	"signature": "ECDSAP256SHA256",
	"signature_type": 13,
	"ksk_public_key": "/+AJ2gscXUGVea9ofxx5AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
	"ds_record": "example.org. 3600 IN DS 55451 13 2 be51f11d7e09131ebe9c0000000000000000000000000000000000000000abcd",
	"created_at": "2026-08-25T09:00:20.925",
	"updated_at": "2026-08-25T09:00:20.925"
}`

const dnssecDisabledJSON = `{"status": "DISABLE"}`
