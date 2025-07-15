package v3

import (
	"os"
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
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = policies.Delete(client, policy.ID)
		th.AssertNoErr(t, err)
	})

	allPolicies, err := policies.List(client, policies.ListOpts{})
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

	updated, err := policies.Update(client, policy.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, *updateOpts.Enabled, updated.Enabled)
	th.CheckEquals(t, updateOpts.Name, updated.Name)

	current, err := policies.Get(client, policy.ID)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, *updateOpts.Enabled, current.Enabled)
	th.CheckEquals(t, updateOpts.Name, current.Name)

}

func TestPolicyReplicationLifecycle(t *testing.T) {
	destProjectID := os.Getenv("OS_PROJECT_ID_2")
	destRegionName := os.Getenv("OS_REGION_NAME_2")
	if destProjectID == "" || destRegionName == "" {
		t.Skip("OS_PROJECT_ID_2 and OS_REGION_NAME_2 are mandatory for this test!")
	}
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)
	iTrue := true
	policy, err := policies.Create(client, policies.CreateOpts{
		Name: "test-policy",
		OperationDefinition: &policies.PolicyODCreate{
			DailyBackups:         1,
			WeekBackups:          2,
			YearBackups:          3,
			MonthBackups:         4,
			MaxBackups:           10,
			Timezone:             "UTC+03:00",
			DestinationProjectId: destProjectID,
			DestinationRegion:    destRegionName,
		},
		Trigger: &policies.Trigger{
			Properties: policies.TriggerProperties{
				Pattern: []string{"FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00"},
			},
		},
		Enabled:       &iTrue,
		OperationType: "replication",
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = policies.Delete(client, policy.ID)
		th.AssertNoErr(t, err)
	})

	allPolicies, err := policies.List(client, policies.ListOpts{})
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

	updated, err := policies.Update(client, policy.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, *updateOpts.Enabled, updated.Enabled)
	th.CheckEquals(t, updateOpts.Name, updated.Name)

	current, err := policies.Get(client, policy.ID)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, *updateOpts.Enabled, current.Enabled)
	th.CheckEquals(t, updateOpts.Name, current.Name)

}
