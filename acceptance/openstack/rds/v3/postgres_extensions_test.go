package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	extensions "github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/postgres-extensions"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRdsPostgresExtensions(t *testing.T) {
	rdsId := os.Getenv("OS_RDS_ID")
	if rdsId == "" {
		t.Skip("OS_RDS_ID env var required for the test is missing")
	}

	client, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	createOpts := extensions.PostgresExtensionOpts{
		DatabaseName:  "postgres",
		ExtensionName: "pg_stat_statements",
	}
	err = extensions.Create(client, rdsId, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = extensions.Delete(client, rdsId, createOpts)
	})

	getExtensionsRaw, err := extensions.List(client, rdsId, extensions.ListOpts{
		DatabaseName: "postgres",
	})
	th.AssertNoErr(t, err)
	for _, extension := range getExtensionsRaw.Extensions {
		if extension.Name == "pg_stat_statements" {
			th.AssertEquals(t, extension.Created, true)
		}
	}

	err = extensions.Update(client, rdsId, createOpts)
	getExtensionsRaw, err = extensions.List(client, rdsId, extensions.ListOpts{
		DatabaseName: "postgres",
	})
	th.AssertNoErr(t, err)
	for _, extension := range getExtensionsRaw.Extensions {
		if extension.Name == "pg_stat_statements" {
			th.AssertEquals(t, extension.Created, true)
		}
	}
}
