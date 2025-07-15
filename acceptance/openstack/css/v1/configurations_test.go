package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	pc "github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/parameter-configuration"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestConfigurationsWorkflow(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`CSS_CLUSTER_ID` must be defined")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to MODIFY Parameter Configuration for cluster: %s", clusterID)
	opts := pc.ModifyOpts{
		Edit: map[string]interface{}{
			"modify": map[string]interface{}{
				"elasticsearch.yml": map[string]interface{}{
					"http.cors.allow-credentials":  true,
					"http.cors.allow-headers":      "X-Requested-With, Content-Type",
					"thread_pool.force_merge.size": "5",
				},
			},
		},
	}
	_, err = pc.Modify(client, opts, clusterID)
	th.AssertNoErr(t, err)

	th.AssertNoErr(t, waitForState(client, clusterID, 30*60))

	t.Cleanup(func() {
		t.Logf("Attempting to DELETE Parameter Configuration for cluster: %s", clusterID)
		resetOpts := pc.ModifyOpts{
			Edit: map[string]interface{}{
				"reset": map[string]interface{}{
					"elasticsearch.yml": map[string]interface{}{
						"http.cors.allow-credentials":  "",
						"http.cors.allow-origin":       "",
						"http.cors.max-age":            "",
						"http.cors.allow-headers":      "",
						"http.cors.enabled":            "",
						"http.cors.allow-methods":      "",
						"reindex.remote.whitelist":     "",
						"indices.queries.cache.size":   "",
						"thread_pool.force_merge.size": "",
					},
				},
			},
		}
		_, err = pc.Modify(client, resetOpts, clusterID)
		th.AssertNoErr(t, err)

		th.AssertNoErr(t, waitForState(client, clusterID, 30*60))
	})

	cfgs, err := pc.List(client, clusterID)
	th.AssertNoErr(t, err)
	for key, config := range cfgs.Templates {
		if key == "http.cors.allow-headers" {
			th.AssertEquals(t, "X-Requested-With, Content-Type", config.Value)
		}
		if key == "http.cors.allow-credentials" {
			th.AssertEquals(t, "true", config.Value)
		}
		if key == "thread_pool.force_merge.size" {
			th.AssertEquals(t, "5", config.Value)
		}
	}
}

func waitForState(client *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		cfgList, err := pc.ListTask(client, id)
		if err != nil {
			return false, err
		}
		for _, task := range cfgList {
			if task.Status == "running" {
				return false, nil
			}
		}
		return true, nil
	})
}
