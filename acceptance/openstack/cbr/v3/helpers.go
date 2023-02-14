package v3

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/backups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/checkpoint"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateCheckpoint(t *testing.T, client *golangsdk.ServiceClient, createOpts checkpoint.CreateOpts) *checkpoint.Checkpoint {
	backup, err := checkpoint.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		errBack := backups.Delete(client, backup.ID)
		th.AssertNoErr(t, errBack)
		th.AssertNoErr(t, waitForBackupDelete(client, 600, backup.ID))
	})

	err = golangsdk.WaitFor(600, func() (bool, error) {
		checkp, err := checkpoint.Get(client, backup.ID)
		if err != nil {
			return false, err
		}
		if checkp.Status == "available" {
			return true, nil
		}
		if checkp.Status == "error" {
			return false, fmt.Errorf("error creating a checkpoint")
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	return backup
}

func RestoreBackup(t *testing.T, client *golangsdk.ServiceClient, id string, opts backups.RestoreBackupOpts) error {
	errRest := backups.RestoreBackup(client, id, opts)
	th.AssertNoErr(t, errRest)

	err := golangsdk.WaitFor(600, func() (bool, error) {
		back, err := backups.Get(client, id)
		if err != nil {
			return false, err
		}
		if back.Status == "available" {
			return true, nil
		}
		if back.Status == "error" {
			return false, fmt.Errorf("error restoring a backup")
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	return nil
}

func waitForBackupDelete(client *golangsdk.ServiceClient, secs int, id string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := backups.Get(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		return false, nil
	})
}
