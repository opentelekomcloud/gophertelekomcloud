package v2

import (
	"os"
	"strings"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/util"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestFunctionGraphUtils(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	createResp, _ := createFunctionGraph(t, client)

	funcUrn := strings.TrimSuffix(createResp.FuncURN, ":latest")

	defer func(client *golangsdk.ServiceClient, id string) {
		err = function.Delete(client, id)
		th.AssertNoErr(t, err)
	}(client, funcUrn)

	timestamp := time.Now().Add(-5 * time.Minute).Format(time.RFC3339)

	listPeriod, err := util.ListStatsPeriod(client, funcUrn, timestamp)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listPeriod)

	// API doesn't work
	// th.AssertNoErr(t, util.EnableFuncLts(client))

	// getLts, err := util.GetFuncLts(client, funcUrn)
	// th.AssertNoErr(t, err)
}

func TestFunctionGraphStatList(t *testing.T) {
	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	listStats, err := util.ListStats(client, util.ListStatsOpts{
		Filter: "monthly_report",
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listStats)
}

func TestFunctionGraphTemplate(t *testing.T) {
	t.Skip("API not published")
	templateID := os.Getenv("TEMPLATE_ID")

	client, err := clients.NewFuncGraphClient()
	th.AssertNoErr(t, err)

	templateResp, err := util.GetFuncTemplate(client, templateID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, templateResp)
}
