package v2

import (
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/cert"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/domain"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCertificateLifecycle(t *testing.T) {
	gatewayID := os.Getenv("GATEWAY_ID")
	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create API Gateway Group")
	grp := CreateGroup(client, t, gatewayID)
	t.Cleanup(func() {
		t.Logf("Attempting to delete API Gateway Group")
		th.AssertNoErr(t, group.Delete(client, gatewayID, grp.ID))
	})

	t.Logf("Attempting to create Public DNS zone with A record")
	clientNetwork, err := clients.NewDNSV2Client()
	th.AssertNoErr(t, err)
	rs := CreateDns(clientNetwork, t)
	t.Cleanup(func() {
		t.Logf("Attempting to delete Public DNS zone")
		_, err := zones.Delete(clientNetwork, rs.ZoneID).Extract()
		th.AssertNoErr(t, err)
	})

	createOpts := domain.CreateOpts{
		GatewayID: gatewayID,
		GroupID:   grp.ID,
		UrlDomain: rs.Name,
	}
	t.Logf("Attempting to create API Gateway Domain")
	dom, err := domain.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to delete API Gateway Domain")
		th.AssertNoErr(t, domain.Delete(client, domain.DeleteOpts{
			GatewayID: gatewayID,
			GroupID:   grp.ID,
			DomainID:  dom.ID,
		}))
	})

	createResp := CreateTestCertificate(client, t, gatewayID, dom.UrlDomain)

	t.Cleanup(func() {
		t.Logf("Attempting to delete certificate: %s", createResp.ID)
		th.AssertNoErr(t, cert.Delete(client, createResp.ID))
	})

	newCert, newPk, err := openstack.GenerateTestCertKeyPair(dom.UrlDomain)

	updateOpts := cert.UpdateOpts{
		Name:        createResp.Name + "_updated",
		CertContent: newCert,
		PrivateKey:  newPk,
		Type:        "instance",
		InstanceID:  gatewayID,
	}

	t.Logf("Attempting to update certificate: %s", createResp.ID)
	updateResp, err := cert.Update(client, createResp.ID, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, updateResp.Name)

	t.Logf("Attempting to get certificate: %s", createResp.ID)
	getResp, err := cert.Get(client, createResp.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, getResp)

	bindOpts := cert.BindOpts{
		InstanceID: gatewayID,
		GroupID:    grp.ID,
		DomainID:   dom.ID,
		CertificateIDs: []string{
			createResp.ID,
		},
	}

	t.Logf("Attempting to bind domain with certificates: %s", createResp.ID)
	err = cert.Bind(client, bindOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to unbind domain certificate: %s", createResp.ID)
	err = cert.Unbind(client, bindOpts)
	th.AssertNoErr(t, err)

	bindCertOpts := cert.AttachDomainOpts{
		CertificateID: createResp.ID,
		Domains: []cert.AttachDomainInfo{
			{
				Domain: dom.UrlDomain,
			},
		},
	}

	t.Logf("Attempting to bind cert to domain: %s", createResp.ID)
	err = cert.BindCertToDomain(client, bindCertOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to unbind cert from domain: %s", createResp.ID)
	err = cert.UnbindCertFromDomain(client, bindCertOpts)
	th.AssertNoErr(t, err)
}

func TestCertificateList(t *testing.T) {
	client, err := clients.NewAPIGWClient()
	gatewayID := os.Getenv("GATEWAY_ID")
	if gatewayID == "" {
		t.Skip("`GATEWAY_ID` needs to be defined")
	}
	th.AssertNoErr(t, err)
	t.Log("Attempting to list certificates")
	allPages, err := cert.List(client, cert.ListOpts{
		InstanceId: gatewayID,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, allPages)
}

func CreateTestCertificate(client *golangsdk.ServiceClient, t *testing.T, gatewayID, domainName string) *cert.CertificateResp {
	certificate, privateKey, err := openstack.GenerateTestCertKeyPair(domainName)

	opts := cert.CreateOpts{
		Name:        tools.RandomString("cert_", 5),
		CertContent: certificate,
		PrivateKey:  privateKey,
		Type:        "instance",
		InstanceID:  gatewayID,
	}

	t.Logf("Attempting to create certificate: %s", opts.Name)
	createResp, err := cert.Create(client, opts)
	th.AssertNoErr(t, err)
	return createResp
}
