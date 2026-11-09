package testing

import (
	"fmt"
	"net/http"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/products"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

const projectID = "0123456789abcdef0123456789abcdef"

func dmsClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	sc.Endpoint = th.Endpoint() + "v2/" + projectID + "/"
	return sc
}

const listResponse = `
{
  "engine" : "kafka",
  "versions" : [ "1.1.0", "2.3.0" ],
  "products" : [ {
    "type" : "cluster",
    "product_id" : "c6.2u4g.cluster",
    "ecs_flavor_id" : "c6.large.2",
    "billing_code" : "dms.platinum.c6",
    "arch_types" : [ "X86" ],
    "charging_mode" : [ "monthly", "hourly" ],
    "ios" : [ {
      "io_spec" : "dms.physical.storage.high.v2",
      "type" : "evs",
      "available_zones" : [ "az1", "az2" ],
      "unavailable_zones" : [ "az3" ]
    } ],
    "support_features" : [ {
      "name" : "connector_obs",
      "properties" : {
        "max_task" : "10",
        "max_node" : "10",
        "min_task" : "1",
        "min_node" : "2"
      }
    } ],
    "properties" : {
      "max_partition_per_broker" : "250",
      "max_broker" : "30",
      "max_storage_per_node" : "10000",
      "max_consumer_per_broker" : "4000",
      "min_broker" : "3",
      "max_bandwidth_per_broker" : "100",
      "min_storage_per_node" : "200",
      "max_tps_per_broker" : "30000",
      "product_alias" : "kafka.2u4g.cluster"
    }
  } ]
}`

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/kafka/products", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.AssertEquals(t, "product_id=c6.2u4g.cluster", r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, listResponse)
	})

	res, err := products.List(dmsClient(), products.ListOpts{Engine: "kafka", ProductId: "c6.2u4g.cluster"})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "kafka", res.Engine)
	th.AssertDeepEquals(t, []string{"1.1.0", "2.3.0"}, res.Versions)
	th.AssertEquals(t, 1, len(res.Products))

	p := res.Products[0]
	th.AssertEquals(t, "cluster", p.Type)
	th.AssertEquals(t, "c6.2u4g.cluster", p.ProductId)
	th.AssertEquals(t, "c6.large.2", p.ECSFlavorId)
	th.AssertDeepEquals(t, []string{"X86"}, p.ArchTypes)
	th.AssertDeepEquals(t, []string{"monthly", "hourly"}, p.ChargingMode)
	th.AssertDeepEquals(t, []products.EngineIOS{{
		IOSpec:           "dms.physical.storage.high.v2",
		Type:             "evs",
		AvailableZones:   []string{"az1", "az2"},
		UnavailableZones: []string{"az3"},
	}}, p.IOS)
	th.AssertEquals(t, 1, len(p.SupportFeatures))
	th.AssertEquals(t, "connector_obs", p.SupportFeatures[0].Name)
	th.AssertDeepEquals(t, products.EngineSupportFeaturesProperty{
		MaxTask: "10", MinTask: "1", MaxNode: "10", MinNode: "2",
	}, p.SupportFeatures[0].Properties)
	th.AssertEquals(t, "250", p.Properties.MaxPartitionPerBroker)
	th.AssertEquals(t, "30", p.Properties.MaxBroker)
	th.AssertEquals(t, "10000", p.Properties.MaxStoragePerNode)
	th.AssertEquals(t, "4000", p.Properties.MaxConsumerPerBroker)
	th.AssertEquals(t, "3", p.Properties.MinBroker)
	th.AssertEquals(t, "100", p.Properties.MaxBandwidthPerBroker)
	th.AssertEquals(t, "200", p.Properties.MinStoragePerNode)
	th.AssertEquals(t, "30000", p.Properties.MaxTPSPerBroker)
	th.AssertEquals(t, "kafka.2u4g.cluster", p.Properties.ProductAlias)
}

func TestListOmitsEmptyQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/kafka/products", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.AssertEquals(t, "", r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"engine": "kafka", "products": []}`)
	})

	res, err := products.List(dmsClient(), products.ListOpts{Engine: "kafka"})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(res.Products))
}

func TestListRequiresEngine(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
	})

	_, err := products.List(dmsClient(), products.ListOpts{})
	if err == nil {
		t.Fatal("expected error for empty engine")
	}
}
