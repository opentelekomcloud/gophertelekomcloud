package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/certificates"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/domains"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func prepareIp(t *testing.T) *floatingips.FloatingIP {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)
	ip, err := floatingips.Create(client, floatingips.CreateOpts{
		FloatingNetworkID: "0a2228f2-7f8a-45f1-8e09-9039e1d09975", // this value is hardcoded in tf OTC provider
	}).Extract()
	th.AssertNoErr(t, err)
	return ip
}

func preparePolicy(t *testing.T, client *golangsdk.ServiceClient) *policies.Policy {
	randomName := tools.RandomString("waf_policy_", 3)
	cert, err := policies.Create(client, policies.CreateOpts{Name: randomName}).Extract()
	th.AssertNoErr(t, err)
	return cert
}

func prepareCertificate(t *testing.T, client *golangsdk.ServiceClient) *certificates.Certificate {
	randomName := tools.RandomString("waf_cert_", 3)
	cert, err := certificates.Create(client, certificates.CreateOpts{
		Name:    randomName,
		Content: testCert,
		Key:     testKey,
	}).Extract()
	th.AssertNoErr(t, err)
	return cert
}

func cleanupIP(t *testing.T, ipID string) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)
	err = floatingips.Delete(client, ipID).ExtractErr()
	th.AssertNoErr(t, err)
}

func cleanupPolicy(t *testing.T, client *golangsdk.ServiceClient, policyID string) {
	err := policies.Delete(client, policyID).ExtractErr()
	th.AssertNoErr(t, err)
}

func cleanupCertificate(t *testing.T, client *golangsdk.ServiceClient, certID string) {
	err := certificates.Delete(client, certID).ExtractErr()
	th.AssertNoErr(t, err)
}

// TestDomainLifecycle is simple "all-in-one" test for waf domain
func TestDomainLifecycle(t *testing.T) {
	client, err := clients.NewWafV1Client()
	th.AssertNoErr(t, err)

	ip := prepareIp(t)
	defer cleanupIP(t, ip.ID)

	cert := prepareCertificate(t, client)
	defer cleanupCertificate(t, client, cert.Id)

	policy := preparePolicy(t, client)
	defer cleanupPolicy(t, client, policy.Id)

	iTrue := true
	createOpts := domains.CreateOpts{
		HostName:      "a.com",
		CertificateId: cert.Id,
		Server: []domains.ServerOpts{
			{
				ClientProtocol: "HTTPS",
				ServerProtocol: "HTTPS",
				Address:        ip.FloatingIP,
				Port:           443,
			},
		},
		Proxy:         &iTrue,
		SipHeaderName: "default",
		SipHeaderList: []string{"X-Forwarded-For"},
	}

	domain, err := domains.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.HostName, domain.HostName)
	th.AssertEquals(t, cert.Id, domain.CertificateId)
	th.AssertEquals(t, len(createOpts.Server), len(domain.Server))

	updateOpts := domains.UpdateOpts{
		TLS:    "TLS v1.1",
		Cipher: "cipher_1",
	}
	domain, err = domains.Update(client, domain.Id, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Cipher, domain.Cipher)

	err = domains.Delete(client, domain.Id).ExtractErr()
	th.AssertNoErr(t, err)
}
