package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const expectedRequest = `
{
  "name": "targetInst",
  "availability_zone": "eu-de-01,eu-de-02",
  "ha": {
    "mode": "ha",
    "replication_mode": "async"
  },
  "flavor_ref": "rds.mysql.s1.large",
  "volume": {
    "type": "ULTRAHIGH",
    "size": 40
  },
  "disk_encryption_id": "2gfdsh-844a-4023-a776-fc5c5fb71fb4",
  "vpc_id": "490a4a08-ef4b-44c5-94be-3051ef9e4fce",
  "subnet_id": "0e2eda62-1d42-4d64-a9d1-4e9aa9cd994f",
  "security_group_id": "2a1f7fc8-3307-42a7-aa6f-42c8b9b8f8c5",
  "backup_strategy": {
    "keep_days": 2,
    "start_time": "19:00-20:00"
  },
  "password": "Demo@12345678",
  "configuration_id": "52e86e87445847a79bf807ceda213165pr01",
  "restore_point": {
    "instance_id": "d8e6ca5a624745bcb546a227aa3ae1cfin01",
    "type": "backup",
    "backup_id": "2f4ddb93-b901-4b08-93d8-1d2e472f30fe"
  }
}
`

const expectedPITRRequest = `
{
  "source": {
    "backup_id": "",
    "instance_id": "d8e6ca5a624745bcb546a227aa3ae1cfin01",
    "restore_time": 1532001446987,
    "type": "timestamp"
  },
  "target": {
    "instance_id": "d8e6ca5a624745bcb546a227aa3ae1cfin01"
  }
}
`

const expectedResponse = `
{
	"instance": {
		"id": "f5ffdd8b1c98434385eb001904209eacin01",
		"name": "demoname",
		"status": "BUILD",
		"datastore": {
			"type": "MySQL",
			"version": "5.6.41"
		},
		"port": "3306",
		"volume": {
			"type": "ULTRAHIGH",
			"size": 40
		},
		"region": "eu-de",
		"backup_strategy": {
			"start_time": "02:00-03:00",
			"keep_days": 7
		},
		"flavor_ref": "rds.mysql.s1.large",
		"availability_zone": "eu-de-01",
		"vpc_id": "19e5d45d-70fd-4a91-87e9-b27e71c9891f",
		"subnet_id": "bd51fb45-2dcb-4296-8783-8623bfe89bb7",
		"security_group_id": "23fd0cd4-15dc-4d65-bdb3-8844cc291be0"
	},
	"job_id": "bf003379-afea-4aa5-aa83-4543542070bc"
}
`

const expectedPITRResponse = `
{
  "job_id": "4c56c0dc-d867-400f-bf3e-d025e4fee686"
}
`

func exampleRestoreOpts() backups.RestoreToNewOpts {
	return backups.RestoreToNewOpts{
		Name: "targetInst",
		Ha: &instances.Ha{
			Mode:            "ha",
			ReplicationMode: "async",
		},
		ConfigurationId: "52e86e87445847a79bf807ceda213165pr01",
		Password:        "Demo@12345678",
		BackupStrategy: &instances.BackupStrategy{
			StartTime: "19:00-20:00",
			KeepDays:  2,
		},
		DiskEncryptionId: "2gfdsh-844a-4023-a776-fc5c5fb71fb4",
		FlavorRef:        "rds.mysql.s1.large",
		Volume: &instances.Volume{
			Type: "ULTRAHIGH",
			Size: 40,
		},
		AvailabilityZone: "eu-de-01,eu-de-02",
		VpcId:            "490a4a08-ef4b-44c5-94be-3051ef9e4fce",
		SubnetId:         "0e2eda62-1d42-4d64-a9d1-4e9aa9cd994f",
		SecurityGroupId:  "2a1f7fc8-3307-42a7-aa6f-42c8b9b8f8c5",
		RestorePoint: backups.RestorePoint{
			InstanceID: "d8e6ca5a624745bcb546a227aa3ae1cfin01",
			Type:       backups.TypeBackup,
			BackupID:   "2f4ddb93-b901-4b08-93d8-1d2e472f30fe",
		},
	}
}

func exampleRestorePITROpts() backups.RestorePITROpts {
	return backups.RestorePITROpts{
		Source: backups.Source{
			InstanceID:  "d8e6ca5a624745bcb546a227aa3ae1cfin01",
			RestoreTime: 1532001446987,
			Type:        "timestamp",
		},
		Target: backups.Target{
			InstanceID: "d8e6ca5a624745bcb546a227aa3ae1cfin01",
		},
	}
}

func TestRestoreRequest(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(func() {
		th.TeardownHTTP()
	})
	th.Mux.HandleFunc("/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, expectedResponse)
	})

	opts := exampleRestoreOpts()
	backup, err := backups.RestoreToNew(client.ServiceClient(), opts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, backup)
}

func TestRestoreRequestPITR(t *testing.T) {
	th.SetupHTTP()
	t.Cleanup(func() {
		th.TeardownHTTP()
	})
	th.Mux.HandleFunc("/instances/recovery", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, expectedPITRResponse)
	})

	opts := exampleRestorePITROpts()
	backup, err := backups.RestorePITR(client.ServiceClient(), opts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, backup)
}
