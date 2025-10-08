package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/quota"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestProjectQuotasList(t *testing.T) {
	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	listOpts := quota.ListProjectQuotasOpts{}
	quotas, err := quota.ListProjectQuotas(client, listOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, quotas)
}

func TestQuotasLifecycle(t *testing.T) {
	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	listOpts := quota.ListQuotasOpts{
		EnterpriseProjectName: pointerto.String("default"),
	}
	quotas, err := quota.ListQuotas(client, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(quotas.QuotaList) > 0)

	originalQuota := quotas.QuotaList[0]
	t.Logf("Original quota: %+v", originalQuota)

	setOpts := quota.SetQuotasOpts{
		QuotaList: []quota.SetQuota{
			{
				EnterpriseProjectId:   originalQuota.EnterpriseProjectId,
				EnterpriseProjectName: originalQuota.EnterpriseProjectName,
				InstanceQuota:         originalQuota.InstanceQuota + 5,
				VcpusQuota:            originalQuota.VcpusQuota + 10,
				RamQuota:              originalQuota.RamQuota + 20,
			},
		},
	}

	setResponse, err := quota.SetQuotas(client, setOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, setResponse)
	th.AssertEquals(t, originalQuota.InstanceQuota+5, setResponse.QuotaList[0].InstanceQuota)

	updateOpts := quota.UpdateQuotasOpts{
		QuotaList: []quota.SetQuota{
			{
				EnterpriseProjectId:   originalQuota.EnterpriseProjectId,
				EnterpriseProjectName: originalQuota.EnterpriseProjectName,
				InstanceQuota:         originalQuota.InstanceQuota + 10,
				VcpusQuota:            originalQuota.VcpusQuota + 20,
				RamQuota:              originalQuota.RamQuota + 40,
			},
		},
	}

	updateResponse, err := quota.UpdateQuotas(client, updateOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, updateResponse)
	th.AssertEquals(t, originalQuota.InstanceQuota+10, updateResponse.QuotaList[0].InstanceQuota)

	restoreOpts := quota.UpdateQuotasOpts{
		QuotaList: []quota.SetQuota{
			{
				EnterpriseProjectId:   originalQuota.EnterpriseProjectId,
				EnterpriseProjectName: originalQuota.EnterpriseProjectName,
				InstanceQuota:         originalQuota.InstanceQuota,
				VcpusQuota:            originalQuota.VcpusQuota,
				RamQuota:              originalQuota.RamQuota,
			},
		},
	}

	restoreResponse, err := quota.UpdateQuotas(client, restoreOpts)
	th.AssertNoErr(t, err)
	t.Logf("Restored quota: %+v", restoreResponse.QuotaList[0])
}
