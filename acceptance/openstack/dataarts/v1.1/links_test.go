package v1_1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dataarts/v1.1/link"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const linkName = "testLink"

func TestDataArtsLinksLifecycle(t *testing.T) {
	if os.Getenv("RUN_DATAART_LIFECYCLE") == "" {
		t.Skip("too slow to run in zuul")
	}

	ak := os.Getenv("AWS_ACCESS_KEY")
	sk := os.Getenv("AWS_SECRET_KEY")

	client, err := clients.NewDataArtsV11Client()
	th.AssertNoErr(t, err)

	c := getTestCluster(t, client)

	t.Log("create cluster's links")

	createOpts := link.CreateOpts{Links: []*link.Link{
		{
			Name:          linkName,
			ConnectorName: "obs-connector",
			LinkConfigValues: &link.ConfigValues{
				Configs: []*link.Config{
					{
						Name: "linkConfig",
						Inputs: []*link.Input{
							{
								Name:  "linkConfig.storageType",
								Value: "OBS",
							},
							{
								Name:  "linkConfig.server",
								Value: "obs.eu-de.otc.t-systems.com",
							},
							{
								Name:  "linkConfig.port",
								Value: "443",
							},
							{
								Name: "linkConfig.obsBucketType",

								Value: "OB",
							},
							{
								Name:  "linkConfig.ossEndpoint",
								Value: "oss-cn-hangzhou.aliyuncs.com",
							},
							{
								Name:  "linkConfig.ossAuthType",
								Value: "ACCESS_KEY",
							},
							{
								Name:  "linkConfig.accessKey",
								Value: ak,
								// Value: ak,
							},
							{
								Name:  "linkConfig.securityKey",
								Value: sk,
							},
						},
					},
				},
			},
		},
	}}

	l, err := link.Create(client, c.Id, createOpts, &link.CreateQuery{})
	th.AssertNoErr(t, err)

	t.Log("schedule link cleanup")
	t.Cleanup(func() {
		t.Logf("attempting to delete link: %s", l.Name)
		err := link.Delete(client, c.Id, l.Name)
		th.AssertNoErr(t, err)
		t.Logf("link is deleted: %s", l.Name)
	})

	t.Log("get cluster's links")

	storedLink, err := link.Get(client, c.Id, l.Name)
	tools.PrintResource(t, storedLink)
	th.AssertNoErr(t, err)
}
