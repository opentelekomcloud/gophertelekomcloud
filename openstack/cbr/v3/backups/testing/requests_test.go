package testing

import (
	"encoding/json"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/backups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestBackupGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleBackup(t)

	actual, err := backups.Get(fake.ServiceClient(), backupID).Extract()
	th.AssertNoErr(t, err)

	expectedRespose := new(struct {
		Backup *backups.Backup `json:"backup"`
	})
	th.AssertNoErr(t, json.Unmarshal([]byte(getBackupResponse), expectedRespose))

	th.AssertDeepEquals(t, expectedRespose.Backup, actual)
}
