package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/certificates"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/hosts"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestWafPremiumCertificateWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	opts := certificates.CreateOpts{
		Name:    tools.RandomString("waf-certificate-", 3),
		Content: testCert,
		Key:     testKey,
	}
	certificate, err := certificates.Create(client, opts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium certificate: %s", certificate.ID)
		th.AssertNoErr(t, certificates.Delete(client, certificate.ID))
		t.Logf("Deleted WAF Premium certificate: %s", certificate.ID)
	})

	t.Logf("Attempting to Get WAF Premium certificate: %s", certificate.ID)
	c, err := certificates.Get(client, certificate.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, c.ID, certificate.ID)

	t.Logf("Attempting to Get WAF Premium host with certificate")
	server := hosts.PremiumWafServer{
		FrontProtocol: "HTTPS",
		BackProtocol:  "HTTPS",
		Address:       "10.10.11.11",
		Port:          443,
		Type:          "ipv4",
		VpcId:         vpcID,
	}
	serverOpts := hosts.CreateOpts{
		CertificateId:   c.ID,
		CertificateName: c.Name,
		Hostname:        tools.RandomString("www.waf-demo.com", 3),
		Proxy:           pointerto.Bool(false),
		Server:          []hosts.PremiumWafServer{server},
	}
	h, err := hosts.Create(client, serverOpts)
	th.AssertNoErr(t, err)
	t.Logf("Created WAF host: %s", h.ID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium host: %s", h.ID)
		th.AssertNoErr(t, hosts.Delete(client, h.ID, hosts.DeleteOpts{}))
		t.Logf("Deleted WAF Premium host: %s", h.ID)
	})

	t.Logf("Attempting to Get WAF Premium certificate with host: %s", certificate.ID)
	ch, err := certificates.Get(client, certificate.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ch.ID, certificate.ID)
	th.AssertEquals(t, ch.BoundHosts[0].Hostname, h.Hostname)
	th.AssertEquals(t, ch.BoundHosts[0].ID, h.ID)

	t.Logf("Attempting to List WAF Premium certificate")
	certificatesList, err := certificates.List(client, certificates.ListOpts{})
	th.AssertNoErr(t, err)

	if len(certificatesList) < 1 {
		t.Fatal("empty WAF Premium certificate list")
	}
}
