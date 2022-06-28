package v3

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/backups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestBackupLifecycle(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)
	id := "a18fee07-fa83-4a13-b0a6-41698c322907"
	// Itrue := true
	page, err := backups.Get(client, id).Extract()
	fmt.Println(page)
	th.AssertNoErr(t, err)

	allPages, err := backups.List(client, backups.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	allBackups, err := backups.ExtractBackups(allPages)
	fmt.Printf("%+v", allBackups)

}
