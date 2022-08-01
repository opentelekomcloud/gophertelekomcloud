package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/metrics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMetrics(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

	opts := metrics.ListMetricsRequest{
		MetricName: "cpu_util",
		Namespace:  "SYS.ECS",
	}

	page := metrics.ListMetrics(client, opts)

	pages, err := page.AllPages()
	th.AssertNoErr(t, err)

	allMetrics, err := metrics.ExtractAllPagesMetrics(pages)
	th.AssertNoErr(t, err)

	for _, m := range allMetrics.Metrics {
		th.AssertEquals(t, m.Namespace, "SYS.ECS")
		th.AssertEquals(t, m.MetricName, "cpu_util")
	}
}
