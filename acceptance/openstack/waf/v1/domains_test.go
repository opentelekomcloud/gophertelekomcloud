package v1

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/certificates"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/domains"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/policies"
)

func prepareIp(t *testing.T) *floatingips.FloatingIP {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Errorf("fail to make network v2 client: %s", err)
	}
	ip, err := floatingips.Create(client, floatingips.CreateOpts{
		FloatingNetworkID: "0a2228f2-7f8a-45f1-8e09-9039e1d09975", // this value is hardcoded in tf OTC provider
	}).Extract()
	if err != nil {
		t.Errorf("fail to create floating IP: %s", err)
	}
	return ip
}

func preparePolicy(t *testing.T, client *golangsdk.ServiceClient) *policies.Policy {
	cert, err := policies.Create(client, policies.CreateOpts{Name: "waf_policy_1"}).Extract()
	if err != nil {
		t.Errorf("fail to create WAF policy: %s", err)
	}
	return cert
}

func prepareCertificate(t *testing.T, client *golangsdk.ServiceClient) *certificates.Certificate {
	cert, err := certificates.Create(client, certificates.CreateOpts{
		Name:    "waf_cert_1",
		Content: "-----BEGIN CERTIFICATE-----MIIDIjCCAougAwIBAgIJALV96mEtVF4EMA0GCSqGSIb3DQEBBQUAMGoxCzAJBgNVBAYTAnh4MQswCQYDVQQIEwJ4eDELMAkGA1UEBxMCeHgxCzAJBgNVBAoTAnh4MQswCQYDVQQLEwJ-----END CERTIFICATE-----",
		Key:     "-----BEGIN RSA PRIVATE KEY-----MIICXQIBAAKBgQDFPN9ojPndxSC4E1pqWQVKGHCFlXAAGBOxbGfSzXqzsoyacotueqMqXQbxrPSQFATeVmhZPNVEMdvcAMjYsV/mymtAwVqVA6q/OFdX/b3UHO+b/VqLo3J5SrM-----END RSA PRIVATE KEY-----",
	}).Extract()
	if err != nil {
		t.Errorf("fail to create WAF certificate: %s", err)
	}
	return cert
}

func cleanupIP(t *testing.T, ipID string) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Errorf("fail to make network v2 client: %s", err)
	}
	err = floatingips.Delete(client, ipID).ExtractErr()
	if err != nil {
		t.Errorf("fail to delete floating IP: %s", err)
	}
}

func cleanupPolicy(t *testing.T, client *golangsdk.ServiceClient, policyID string) {
	err := policies.Delete(client, policyID).ExtractErr()
	if err != nil {
		t.Errorf("fail to remove WAF policy: %s", err)
	}
}
func cleanupCertificate(t *testing.T, client *golangsdk.ServiceClient, certID string) {
	err := certificates.Delete(client, certID).ExtractErr()
	if err != nil {
		t.Errorf("fail to remove WAF certificate: %s", err)
	}
}

// TestDomainLifecycle is simple "all-in-one" test for waf domain
func TestDomainLifecycle(t *testing.T) {
	client, err := clients.NewWafV1()
	if err != nil {
		t.Fatalf("Unable to create a RDSv3 client: %s", err)
	}

	ip := prepareIp(t)
	defer cleanupIP(t, ip.ID)

	cert := prepareCertificate(t, client)
	defer cleanupCertificate(t, client, cert.Id)

	policy := preparePolicy(t, client)
	defer cleanupPolicy(t, client, policy.Id)

	iTrue := true
	domain, err := domains.Create(client, domains.CreateOpts{
		HostName:      "a.com",
		CertificateId: cert.Id,
		Server: []domains.ServerOpts{
			{
				ClientProtocol: "HTTPS",
				ServerProtocol: "HTTP",
				Address:        ip.FloatingIP,
				Port:           80,
			},
		},
		Proxy:         &iTrue,
		SipHeaderName: "default",
		SipHeaderList: []string{"X-Forwarded-For"},
	}).Extract()
	if err != nil {
		t.Errorf("failed to create domain: %s", err)
	}
	if err := domains.Delete(client, domain.Id).ExtractErr(); err != nil {
		t.Errorf("failed to delete domain: %s", err)
	}
}
