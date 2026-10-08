package gpfs

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gpfs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGPFSFileSystemLifecycle(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	if vpcID == "" {
		t.Skip("OS_VPC_ID is required for this test")
	}

	client, err := clients.NewGPFSClient()
	th.AssertNoErr(t, err)

	fsName := strings.ToLower(tools.RandomString("gpfs-sdk-test", 5))

	_, err = client.CreateFS(&gpfs.CreateFSInput{
		FSName:     fsName,
		Redundancy: "3az",
		BucketType: "SFS",
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		_, err = client.DeleteFS(fsName)
		th.AssertNoErr(t, err)
	})

	fileSystems, err := client.ListFS(&gpfs.ListFSInput{
		BucketType: "SFS",
	})
	th.AssertNoErr(t, err)
	th.AssertNotEquals(t, 0, len(fileSystems.Buckets))

	accessRules := []gpfs.AccessRule{
		{
			ID:     "gpfs-sdk-test-rule",
			Action: gpfs.AccessRuleActionFullControl,
			Effect: gpfs.AccessRuleEffectAllow,
			Condition: gpfs.AccessRuleCondition{
				SourceVPC: vpcID,
			},
		},
	}
	_, err = client.CreateFSAccessRules(&gpfs.CreateFSAccessRulesInput{
		FSName:    fsName,
		Statement: accessRules,
	})
	th.AssertNoErr(t, err)
	rulesDeleted := false
	t.Cleanup(func() {
		if !rulesDeleted {
			_, err = client.DeleteFSAccessRules(fsName)
			th.AssertNoErr(t, err)
		}
	})

	err = tools.WaitFor(func() (bool, error) {
		output, err := client.GetFSAccessRules(fsName)
		if err != nil {
			return false, err
		}
		for _, rule := range output.Statement {
			if rule.Condition.SourceVPC == vpcID {
				th.AssertEquals(t, gpfs.AccessRuleActionFullControl, rule.Action)
				th.AssertEquals(t, gpfs.AccessRuleEffectAllow, rule.Effect)
				return true, nil
			}
		}
		return false, nil
	})
	th.AssertNoErr(t, err)

	_, err = client.DeleteFSAccessRules(fsName)
	th.AssertNoErr(t, err)
	rulesDeleted = true
}
