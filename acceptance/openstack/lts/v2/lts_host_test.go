package v2

import (
	"fmt"
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	hg "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/host-groups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsHostGroupLifecycle(t *testing.T) {
	ak, sk := os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_KEY")
	if ak == "" || sk == "" {
		t.Skip("ACCESS_KEY and SECRET_KEY are required for this test")
	}
	client, err := clients.NewLtsV3Client()
	th.AssertNoErr(t, err)

	userDataScript := fmt.Sprintf(`#!/bin/bash
set +o history
curl https://icagent-eu-de.obs.eu-de.otc.t-systems.com/ICAgent_linux/apm_agent_install.sh > apm_agent_install.sh && REGION=eu-de bash apm_agent_install.sh -ak %s -sk %s -region eu-de -projectid %s -accessip lts-access.eu-de.otc.t-systems.com -obsdomain obs.eu-de.otc.t-systems.com
set -o history
`, ak, sk, client.ProjectID)
	t.Logf("Attempting to Create member for Host Group")
	ecsClient, err := clients.NewComputeV2Client()
	ecs := openstack.CreateServer(t, ecsClient,
		tools.RandomString("lts-hg-", 3),
		"Standard_Debian_11_latest",
		"s2.large.2",
		userDataScript,
	)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete Server: %s", ecs.ID)
		th.AssertNoErr(t, servers.Delete(ecsClient, ecs.ID).ExtractErr())
	})

	err = golangsdk.WaitFor(1000, func() (bool, error) {
		h, err := hg.ListHost(client, hg.ListHostOpts{HostIdList: []string{ecs.ID}})
		if err != nil {
			return false, err
		}

		if len(h.Result) > 0 {
			if h.Result[0].HostStatus == "running" {
				return true, nil
			}
		}

		return false, nil
	})
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-hgroup-", 3)
	createOpts := hg.CreateOpts{
		Name:       name,
		Type:       "linux",
		HostIdList: []string{ecs.ID},
		Tags: []tags.ResourceTag{
			{
				Key: "fizz", Value: "buzz",
			},
		},
	}
	t.Logf("Attempting to Create Host Group")
	created, err := hg.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Host Group")
		_, err = hg.Delete(client, hg.DeleteOpts{
			HostGroupIds: []string{created.ID},
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Update Host Group")
	nameUpdate := tools.RandomString("test-hgroup-up", 3)
	uHg, err := hg.Update(client, hg.UpdateLogGroupOpts{
		ID:         created.ID,
		Name:       nameUpdate,
		HostIdList: &[]string{},
		Tags:       &[]tags.ResourceTag{},
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, nameUpdate, uHg.Name)
	th.AssertEquals(t, 0, len(uHg.HostIdList))
	th.AssertEquals(t, 0, len(uHg.Tags))

	t.Logf("Attempting to List Host Groups")
	hList, err := hg.List(client, hg.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(hList.Result))
}
