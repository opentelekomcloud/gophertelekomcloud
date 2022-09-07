package testing

import (
	"fmt"
	"net/http"
	"testing"

	quotasets2 "github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/quotasets"

	"github.com/opentelekomcloud/gophertelekomcloud"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const FirstTenantID = "555544443333222211110000ffffeeee"

var getExpectedJSONBody = `
{
	"quota_set" : {
		"volumes" : 8,
		"snapshots" : 9,
		"gigabytes" : 10,
		"per_volume_gigabytes" : 11,
		"backups" : 12,
		"backup_gigabytes" : 13
	}
}`

var getExpectedQuotaSet = quotasets2.QuotaSet{
	Volumes:            8,
	Snapshots:          9,
	Gigabytes:          10,
	PerVolumeGigabytes: 11,
	Backups:            12,
	BackupGigabytes:    13,
}

var getUsageExpectedJSONBody = `
{
  "quota_set": {
    "id": "555544443333222211110000ffffeeee",
    "volumes": {
      "in_use": 15,
      "limit": 16,
      "reserved": 17
    },
    "snapshots": {
      "in_use": 18,
      "limit": 19,
      "reserved": 20
    },
    "gigabytes": {
      "in_use": 21,
      "limit": 22,
      "reserved": 23
    },
    "per_volume_gigabytes": {
      "in_use": 24,
      "limit": 25,
      "reserved": 26
    },
    "backups": {
      "in_use": 27,
      "limit": 28,
      "reserved": 29
    },
    "backup_gigabytes": {
      "in_use": 30,
      "limit": 31,
      "reserved": 32
    }
  }
}
`

var getUsageExpectedQuotaSet = quotasets2.QuotaUsageSet{
	ID:                 FirstTenantID,
	Volumes:            quotasets2.QuotaUsage{InUse: 15, Limit: 16, Reserved: 17},
	Snapshots:          quotasets2.QuotaUsage{InUse: 18, Limit: 19, Reserved: 20},
	Gigabytes:          quotasets2.QuotaUsage{InUse: 21, Limit: 22, Reserved: 23},
	PerVolumeGigabytes: quotasets2.QuotaUsage{InUse: 24, Limit: 25, Reserved: 26},
	Backups:            quotasets2.QuotaUsage{InUse: 27, Limit: 28, Reserved: 29},
	BackupGigabytes:    quotasets2.QuotaUsage{InUse: 30, Limit: 31, Reserved: 32},
}

var fullUpdateExpectedJSONBody = `
{
	"quota_set": {
		"volumes": 8,
		"snapshots": 9,
		"gigabytes": 10,
		"per_volume_gigabytes": 11,
		"backups": 12,
		"backup_gigabytes": 13
	}
}`

var fullUpdateOpts = quotasets2.UpdateOpts{
	Volumes:            golangsdk.IntToPointer(8),
	Snapshots:          golangsdk.IntToPointer(9),
	Gigabytes:          golangsdk.IntToPointer(10),
	PerVolumeGigabytes: golangsdk.IntToPointer(11),
	Backups:            golangsdk.IntToPointer(12),
	BackupGigabytes:    golangsdk.IntToPointer(13),
}

var fullUpdateExpectedQuotaSet = quotasets2.QuotaSet{
	Volumes:            8,
	Snapshots:          9,
	Gigabytes:          10,
	PerVolumeGigabytes: 11,
	Backups:            12,
	BackupGigabytes:    13,
}

var partialUpdateExpectedJSONBody = `
{
	"quota_set": {
		"volumes": 200,
		"snapshots": 0,
		"gigabytes": 0,
		"per_volume_gigabytes": 0,
		"backups": 0,
		"backup_gigabytes": 0
	}
}`

var partialUpdateOpts = quotasets2.UpdateOpts{
	Volumes:            golangsdk.IntToPointer(200),
	Snapshots:          golangsdk.IntToPointer(0),
	Gigabytes:          golangsdk.IntToPointer(0),
	PerVolumeGigabytes: golangsdk.IntToPointer(0),
	Backups:            golangsdk.IntToPointer(0),
	BackupGigabytes:    golangsdk.IntToPointer(0),
}

var partiualUpdateExpectedQuotaSet = quotasets2.QuotaSet{Volumes: 200}

// HandleSuccessfulRequest configures the test server to respond to an HTTP request.
func HandleSuccessfulRequest(t *testing.T, httpMethod, uriPath, jsonOutput string, uriQueryParams map[string]string) {

	th.Mux.HandleFunc(uriPath, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, httpMethod)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")

		if uriQueryParams != nil {
			th.TestFormValues(t, r, uriQueryParams)
		}

		_, _ = fmt.Fprint(w, jsonOutput)
	})
}

// HandleDeleteSuccessfully tests quotaset deletion.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-quota-sets/"+FirstTenantID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
	})
}
