package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPolicyLifecycle(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	iTrue := true
	policy, err := policies.Create(client, policies.CreateOpts{
		Name: "test-policy",
		OperationDefinition: &policies.PolicyODCreate{
			DailyBackups: 1,
			WeekBackups:  2,
			YearBackups:  3,
			MonthBackups: 4,
			MaxBackups:   10,
			Timezone:     "UTC+03:00",
		},
		Trigger: &policies.Trigger{
			Properties: policies.TriggerProperties{
				Pattern: []string{"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00"},
			},
		},
		Enabled:       &iTrue,
		OperationType: "backup",
	}).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		err = policies.Delete(client, policy.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	allPages, err := policies.List(client, policies.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	allPolicies, err := policies.ExtractPolicies(allPages)
	th.AssertNoErr(t, err)
	found := false
	for _, p := range allPolicies {
		if p.ID == policy.ID {
			found = true
			break
		}
	}
	th.CheckEquals(t, true, found)

	iFalse := false
	updateOpts := policies.UpdateOpts{
		Enabled: &iFalse,
		Name:    "name2",
	}

	updated, err := policies.Update(client, policy.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, *updateOpts.Enabled, updated.Enabled)
	th.CheckEquals(t, updateOpts.Name, updated.Name)

	current, err := policies.Get(client, policy.ID).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, *updateOpts.Enabled, current.Enabled)
	th.CheckEquals(t, updateOpts.Name, current.Name)

}
