package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/autoscaling"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/quotas"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestInstances(t *testing.T) {
	client, err := clients.NewAutoscalingV1Client()
	th.AssertNoErr(t, err)

	asGroupCreateName := tools.RandomString("as-ins-", 3)
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if networkID == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID env vars are missing but AS Group test requires")
	}

	asCreateName := tools.RandomString("as-ins-", 3)
	keyPairName := clients.EnvOS.GetEnv("KEYPAIR_NAME")
	imageID := clients.EnvOS.GetEnv("IMAGE_ID")
	if keyPairName == "" || imageID == "" {
		t.Skip("OS_KEYPAIR_NAME or OS_IMAGE_ID env vars is missing but AS Configuration test requires")
	}

	configID := autoscaling.CreateASConfig(t, client, asCreateName, imageID, keyPairName)
	t.Cleanup(func() {
		autoscaling.DeleteASConfig(t, client, configID)
	})

	createOpts := groups.CreateOpts{
		Name: asGroupCreateName,
		Networks: []groups.ID{
			{
				ID: networkID,
			},
		},
		SecurityGroup: []groups.ID{
			{
				ID: openstack.DefaultSecurityGroup(t),
			},
		},
		VpcID:                vpcID,
		IsDeletePublicip:     pointerto.Bool(true),
		DesireInstanceNumber: 1,
		MinInstanceNumber:    0,
		MaxInstanceNumber:    5,
		ConfigurationID:      configID,
	}
	t.Logf("Attempting to create AutoScaling Group")
	groupID, err2 := groups.Create(client, createOpts)
	th.AssertNoErr(t, err2)
	t.Logf("Created AutoScaling Group: %s", groupID)

	t.Cleanup(func() {
		autoscaling.DeleteAutoScalingGroup(t, client, groupID)
	})

	err = groups.Enable(client, groupID)
	th.AssertNoErr(t, err)

	err = golangsdk.WaitFor(1000, func() (bool, error) {
		n, err := groups.Get(client, groupID)
		if err != nil {
			return false, err
		}

		if n.ActualInstanceNumber == 1 {
			return true, nil
		}

		return false, nil
	})
	th.AssertNoErr(t, err)

	err = groups.Disable(client, groupID)
	th.AssertNoErr(t, err)

	inst, err := instances.List(client, instances.ListOpts{
		ScalingGroupId: groupID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, inst)

	t.Cleanup(func() {
		err = instances.Delete(client, instances.DeleteOpts{
			InstanceId:     inst.ScalingGroupInstances[0].ID,
			DeleteInstance: "yes",
		})
		th.AssertNoErr(t, err)

		err = golangsdk.WaitFor(1000, func() (bool, error) {
			n, err := groups.Get(client, groupID)
			if err != nil {
				return false, err
			}

			if n.ActualInstanceNumber == 0 {
				return true, nil
			}

			return false, nil
		})
		th.AssertNoErr(t, err)
	})

	quota, err := quotas.ShowPolicyAndInstanceQuota(client, groupID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, quota)
}
