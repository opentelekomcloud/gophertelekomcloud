package v1

import (
	"strconv"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ces/v1/metricdata"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMetricData(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

	ecsClient, err := clients.NewComputeV1Client()
	th.AssertNoErr(t, err)

	ecs := openstack.CreateCloudServer(t, ecsClient, openstack.GetCloudServerCreateOpts(t))
	defer openstack.DeleteCloudServer(t, ecsClient, ecs.ID)

	old := time.Now().UnixMilli()
	newOps := metricdata.MetricDataItem{
		Metric: metricdata.MetricInfo{
			Namespace:  "MINE.APP",
			MetricName: "cpu_util",
			Dimensions: []metricdata.MetricsDimension{
				{
					Name:  "instance_id",
					Value: ecs.ID,
				},
			},
		},
		Ttl:         172800,
		CollectTime: old,
		Value:       0.09,
		Unit:        "%",
		Type:        "float",
	}

	err = metricdata.CreateMetricData(client, metricdata.CreateMetricDataRequest{newOps}).ExtractErr()
	th.AssertNoErr(t, err)

	batchOps := metricdata.BatchListMetricDataRequest{
		Metrics: []metricdata.Metric{
			{
				Namespace:  "MINE.APP",
				MetricName: "cpu_util",
				Dimensions: []metricdata.MetricsDimension{
					{
						Name:  "instance_id",
						Value: ecs.ID,
					},
				},
			},
		},
		From:   old,
		To:     time.Now().UnixMilli(),
		Period: "1",
		Filter: "average",
	}

	batchData, err := metricdata.BatchListMetricData(client, batchOps).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(batchData), 1)

	metricOps := metricdata.ShowMetricDataRequest{
		Namespace:  "SYS.ECS",
		MetricName: "cpu_util",
		Dim:        "instance_id," + ecs.ID,
		Filter:     "average",
		Period:     1,
		From:       strconv.FormatInt(old, 10),
		To:         strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	metricData, err := metricdata.ShowMetricData(client, metricOps).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(metricData.Datapoints), 0)

	eventOps := metricdata.ShowEventDataRequest{
		Namespace: "SYS.ECS",
		Dim:       "instance_id," + ecs.ID,
		Type:      "instance_host_info",
		From:      strconv.FormatInt(old, 10),
		To:        strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	event, err := metricdata.ShowEventData(client, eventOps).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(event.Datapoints), 0)
}
