package autoscaling

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/autoscaling/v1/groups"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func CreateAutoScalingGroup(t *testing.T, client *golangsdk.ServiceClient, networkID, vpcID, asName string) string {
	defaultSGID := openstack.DefaultSecurityGroup(t)
	deletePublicIP := true

	createOpts := groups.CreateOpts{
		Name: asName,
		Networks: []groups.NetworkOpts{
			{
				ID: networkID,
			},
		},
		SecurityGroup: []groups.SecurityGroupOpts{
			{
				ID: defaultSGID,
			},
		},
		VpcID:            vpcID,
		IsDeletePublicip: &deletePublicIP,
	}
	t.Logf("Attempting to create AutoScaling Group")
	groupID, err := groups.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created AutoScaling Group: %s", groupID)

	return groupID
}

func DeleteAutoScalingGroup(t *testing.T, client *golangsdk.ServiceClient, groupID string) {
	t.Logf("Attempting to delete AutoScaling Group")
	err := groups.Delete(client, groupID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Deleted AutoScaling Group: %s", groupID)
}
