package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/quotasets"

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

var getExpectedQuotaSet = quotasets.QuotaSet{
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

var getUsageExpectedQuotaSet = quotasets.QuotaUsageSet{
	ID:                 FirstTenantID,
	Volumes:            quotasets.QuotaUsage{InUse: 15, Limit: 16, Reserved: 17},
	Snapshots:          quotasets.QuotaUsage{InUse: 18, Limit: 19, Reserved: 20},
	Gigabytes:          quotasets.QuotaUsage{InUse: 21, Limit: 22, Reserved: 23},
	PerVolumeGigabytes: quotasets.QuotaUsage{InUse: 24, Limit: 25, Reserved: 26},
	Backups:            quotasets.QuotaUsage{InUse: 27, Limit: 28, Reserved: 29},
	BackupGigabytes:    quotasets.QuotaUsage{InUse: 30, Limit: 31, Reserved: 32},
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

var fullUpdateOpts = quotasets.UpdateOpts{
	Volumes:            pointerto.Int(8),
	Snapshots:          pointerto.Int(9),
	Gigabytes:          pointerto.Int(10),
	PerVolumeGigabytes: pointerto.Int(11),
	Backups:            pointerto.Int(12),
	BackupGigabytes:    pointerto.Int(13),
}

var fullUpdateExpectedQuotaSet = quotasets.QuotaSet{
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

var partialUpdateOpts = quotasets.UpdateOpts{
	Volumes:            pointerto.Int(200),
	Snapshots:          pointerto.Int(0),
	Gigabytes:          pointerto.Int(0),
	PerVolumeGigabytes: pointerto.Int(0),
	Backups:            pointerto.Int(0),
	BackupGigabytes:    pointerto.Int(0),
}

var partiualUpdateExpectedQuotaSet = quotasets.QuotaSet{Volumes: 200}

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
