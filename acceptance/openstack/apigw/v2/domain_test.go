package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/domain"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/apigw/v2/group"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/recordsets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	certificate = `<<EOT
-----BEGIN CERTIFICATE-----
MIIEMDCCAxigAwIBAgISBKmXWIa316BWEURqVnLroXwgMA0GCSqGSIb3DQEBCwUA
MDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD
EwJSMzAeFw0yNDAzMDExMDQ3MTRaFw0yNDA1MzAxMDQ3MTNaMCIxIDAeBgNVBAMT
F3JhbmNoZXItdGVzdC52YXh4Y2hhLmluMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcD
QgAE8ps2bDLATJ8Lyufn4VG8y0zxBK5AJ9p60bE9kTwftr3xKgE0mC25qWt0M1JY
z6opBmChuwGyx/M7L/L4dhp7ZaOCAhkwggIVMA4GA1UdDwEB/wQEAwIHgDAdBgNV
HSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4E
FgQU1Uw51mX15Fovr1jOqEItb7J90lgwHwYDVR0jBBgwFoAUFC6zF7dYVsuuUAlA
5h+vnYsUwsYwVQYIKwYBBQUHAQEESTBHMCEGCCsGAQUFBzABhhVodHRwOi8vcjMu
by5sZW5jci5vcmcwIgYIKwYBBQUHMAKGFmh0dHA6Ly9yMy5pLmxlbmNyLm9yZy8w
IgYDVR0RBBswGYIXcmFuY2hlci10ZXN0LnZheHhjaGEuaW4wEwYDVR0gBAwwCjAI
BgZngQwBAgEwggEEBgorBgEEAdZ5AgQCBIH1BIHyAPAAdQA7U3d1Pi25gE6LMFsG
/kA7Z9hPw/THvQANLXJv4frUFwAAAY351wROAAAEAwBGMEQCICGcAZhHtchF6tLQ
yolMSuU5ZQX9ZQ/Ld1Mqg3t1kHBpAiB5zPoxpc2Nvty+U+lNVx5QI6GqY8+oIOP2
+PNhpJleaAB3AO7N0GTV2xrOxVy3nbTNE6Iyh0Z8vOzew1FIWUZxH7WbAAABjfnX
BE4AAAQDAEgwRgIhALgFmbF4pFUBeEXD8sjaKB80f7ZYQ+WBXS9887VgBZvqAiEA
892Gdv3ZP8uR4ptXUOfXsR7JTCBCsFCAG3rk5e8vt5MwDQYJKoZIhvcNAQELBQAD
ggEBAEk8Gj7laMCp9mx/55lCELBCdSBtCBQKvAJrDKVZHZ35keeOcxWpzVsQrall
WgB5UuIspH/EBFHe7hrwcRXW4IkjXFS0mWz4GqwiQ7EBnk6ITiw4Zyj4ejLPNGKj
v2YSBwrBZqxXVqW981FvTGNhmcp45j2fEA37DwoSWU/wKWHpRqA0akIZOGoWwAfu
kRE9ZuPYykvnnl4qk25KWbIYAO4/65WjMqSkj5imOVA8yNtK3j3qnAcNwQL8VPr/
qH2IF4KjanBUK76qT2wHKkj2GPMP/vcB9Nzg+mY+elax9XUhtpR0ZJVWDzE150Vo
ii8i6whnmTj+S/QRvhOqHmWU0Xk=
-----END CERTIFICATE-----
EOT
`

	privateKey = `<<EOT
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgpR2jFW8FNrJCaSTf
evnYeTMA7AIS/DgMsmoIeD1CWDyhRANCAATymzZsMsBMnwvK5+fhUbzLTPEErkAn
2nrRsT2RPB+2vfEqATSYLbmpa3QzUljPqikGYKG7AbLH8zsv8vh2Gntl
-----END PRIVATE KEY-----
EOT
`
)

func TestDomainLifecycle(t *testing.T) {
	apiGwID := clients.EnvOS.GetEnv("APIGW_ID")

	if apiGwID == "" {
		t.Skip("`apiGwID` need to be defined")
	}

	client, err := clients.NewAPIGWClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to create API Gateway Group")
	grp := CreateGroup(client, t, apiGwID)
	t.Cleanup(func() {
		t.Logf("Attempting to delete API Gateway Group")
		th.AssertNoErr(t, group.Delete(client, apiGwID, grp.ID))
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
		GatewayID: apiGwID,
		GroupID:   grp.ID,
		UrlDomain: rs.Name,
	}
	t.Logf("Attempting to create API Gateway Domain")
	dom, err := domain.Create(client, createOpts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Attempting to delete API Gateway Domain")
		th.AssertNoErr(t, domain.Delete(client, domain.DeleteOpts{
			GatewayID: apiGwID,
			GroupID:   grp.ID,
			DomainID:  dom.ID,
		}))
	})

	updateOpts := domain.UpdateOpts{
		GatewayID:     apiGwID,
		GroupID:       grp.ID,
		DomainID:      dom.ID,
		MinSslVersion: "TLSv1.2",
	}
	t.Logf("Attempting to update API Gateway Domain")
	domUpdated, err := domain.Update(client, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "TLSv1.2", domUpdated.MinSslVersion)

	t.Logf("Attempting to asssign API Gateway Domain Certificate")
	certName := tools.RandomString("acc_domain_certificate_", 5)
	certOpts := domain.CreateCertOpts{
		GatewayID:  apiGwID,
		GroupID:    grp.ID,
		DomainID:   dom.ID,
		Content:    certificate,
		PrivateKey: privateKey,
		Name:       certName,
	}
	domCertificate, err := domain.AssignCertificate(client, certOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete API Gateway Domain Certificate")
		th.AssertNoErr(t, domain.UnbindCertificate(client, domain.CertificateOpts{
			GatewayID:     apiGwID,
			GroupID:       grp.ID,
			DomainID:      dom.ID,
			CertificateID: domCertificate.ID,
		}))
	})

	t.Logf("Attempting to get API Gateway Domain Certificate")
	getCert, err := domain.GetCertificate(client, domain.CertificateOpts{
		GatewayID:     apiGwID,
		GroupID:       grp.ID,
		DomainID:      dom.ID,
		CertificateID: domCertificate.ID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, certName, getCert.Name)
}

func CreateDns(client *golangsdk.ServiceClient, t *testing.T) *recordsets.RecordSet {
	zoneName := "otc-acc.test-gateway.com"
	createOpts := zones.CreateOpts{
		Description: "api gw public zone",
		Name:        zoneName,
		TTL:         300,
	}
	zone, err := zones.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	recordOpts := recordsets.CreateOpts{
		Name:        "domain." + zone.Name,
		Description: "record set api gw",
		Records:     []string{"10.0.0.20"},
		Type:        "A",
	}
	recordSet, err := recordsets.Create(client, zone.ID, recordOpts).Extract()
	th.AssertNoErr(t, err)

	return recordSet
}
