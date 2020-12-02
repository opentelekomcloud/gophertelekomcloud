package lbaas_v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/lbaas_v2/certificates"
)

func TestLbaasV2CertificatesList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a NetworkingV2 client: %s", err)
	}

	listOpts := certificates.ListOpts{}
	allPages, err := certificates.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to fetch LbaasV2 pages: %s", err)
	}
	lbaasCertificates, err := certificates.ExtractCertificates(allPages)
	if err != nil {
		t.Fatalf("Unable to extract LbaasV2 pages: %s", err)
	}
	for _, certificate := range lbaasCertificates {
		tools.PrintResource(t, certificate)
	}
}

func TestLbaasV2CertificateLifeCycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a NetworkingV2 client: %s", err)
	}

	// Create lbaasV2 certificate
	lbaasCertificate, err := createLbaasCertificate(t, client)
	if err != nil {
		t.Fatalf("Unable to create LbaasV2 certificate: %s", err)
	}
	defer deleteLbaasCertificate(t, client, lbaasCertificate.ID)

	tools.PrintResource(t, lbaasCertificate)

	err = updateLbaasCertificate(t, client, lbaasCertificate.ID)
	if err != nil {
		t.Fatalf("Unable to update LbaasV2 certificate: %s", err)
	}
	tools.PrintResource(t, lbaasCertificate)

	newLbaasCertificate, err := certificates.Get(client, lbaasCertificate.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get LbaasV2 certificate: %s", err)
	}
	tools.PrintResource(t, newLbaasCertificate)
}

func createLbaasCertificate(t *testing.T, client *golangsdk.ServiceClient) (*certificates.Certificate, error) {
	certificateName := tools.RandomString("create-cert-", 8)
	certificate := "-----BEGIN CERTIFICATE-----\nMIIDpTCCAo2gAwIBAgIJAKdmmOBYnFvoMA0GCSqGSIb3DQEBCwUAMGkxCzAJBgNV\nBAYTAnh4MQswCQYDVQQIDAJ4eDELMAkGA1UEBwwCeHgxCzAJBgNVBAoMAnh4MQsw\nCQYDVQQLDAJ4eDELMAkGA1UEAwwCeHgxGTAXBgkqhkiG9w0BCQEWCnh4QDE2My5j\nb20wHhcNMTcxMjA0MDM0MjQ5WhcNMjAxMjAzMDM0MjQ5WjBpMQswCQYDVQQGEwJ4\neDELMAkGA1UECAwCeHgxCzAJBgNVBAcMAnh4MQswCQYDVQQKDAJ4eDELMAkGA1UE\nCwwCeHgxCzAJBgNVBAMMAnh4MRkwFwYJKoZIhvcNAQkBFgp4eEAxNjMuY29tMIIB\nIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwZ5UJULAjWr7p6FVwGRQRjFN\n2s8tZ/6LC3X82fajpVsYqF1xqEuUDndDXVD09E4u83MS6HO6a3bIVQDp6/klnYld\niE6Vp8HH5BSKaCWKVg8lGWg1UM9wZFnlryi14KgmpIFmcu9nA8yV/6MZAe6RSDmb\n3iyNBmiZ8aZhGw2pI1YwR+15MVqFFGB+7ExkziROi7L8CFCyCezK2/oOOvQsH1dz\nQ8z1JXWdg8/9Zx7Ktvgwu5PQM3cJtSHX6iBPOkMU8Z8TugLlTqQXKZOEgwajwvQ5\nmf2DPkVgM08XAgaLJcLigwD513koAdtJd5v+9irw+5LAuO3JclqwTvwy7u/YwwID\nAQABo1AwTjAdBgNVHQ4EFgQUo5A2tIu+bcUfvGTD7wmEkhXKFjcwHwYDVR0jBBgw\nFoAUo5A2tIu+bcUfvGTD7wmEkhXKFjcwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B\nAQsFAAOCAQEAWJ2rS6Mvlqk3GfEpboezx2J3X7l1z8Sxoqg6ntwB+rezvK3mc9H0\n83qcVeUcoH+0A0lSHyFN4FvRQL6X1hEheHarYwJK4agb231vb5erasuGO463eYEG\nr4SfTuOm7SyiV2xxbaBKrXJtpBp4WLL/s+LF+nklKjaOxkmxUX0sM4CTA7uFJypY\nc8Tdr8lDDNqoUtMD8BrUCJi+7lmMXRcC3Qi3oZJW76ja+kZA5mKVFPd1ATih8TbA\ni34R7EQDtFeiSvBdeKRsPp8c0KT8H1B4lXNkkCQs2WX5p4lm99+ZtLD4glw8x6Ic\ni1YhgnQbn5E0hz55OLu5jvOkKQjPCW+9Aa==\n-----END CERTIFICATE-----"

	createLbaasCertificateOpts := certificates.CreateOpts{
		AdminStateUp: true,
		Name:         certificateName,
		Description:  "some test description",
		Certificate:  certificate,
	}
	lbaasCertificate, err := certificates.Create(client, createLbaasCertificateOpts).Extract()
	if err != nil {
		return nil, err
	}
	t.Logf("Created LbaasV2 certificate: %s", lbaasCertificate.ID)

	return lbaasCertificate, nil
}

func deleteLbaasCertificate(t *testing.T, client *golangsdk.ServiceClient, lbaasCertificateId string) {
	t.Logf("Attempting to delete LbaasV2 certificate: %s", lbaasCertificateId)

	if err := certificates.Delete(client, lbaasCertificateId).Err; err != nil {
		t.Fatalf("Unable to delete LbaasV2 certificate: %s", err)
	}

	t.Logf("LbaasV2 certificate is deleted: %s", lbaasCertificateId)
}

func updateLbaasCertificate(t *testing.T, client *golangsdk.ServiceClient, lbaasCertificateId string) error {
	t.Logf("Attempting to update LbaasV2 certificate")

	certificateNewName := tools.RandomString("update-cert-", 8)

	updateCertificate := "-----BEGIN CERTIFICATE-----\nMIIDpTCCAo2gAwIBAgIJAKdmmOBYnFvoMA0GCSqGSIb3DQEBCwUAMGkxCzAJBgNV\nBAYTAnh4MQswCQYDVQQIDAJ4eDELMAkGA1UEBwwCeHgxCzAJBgNVBAoMAnh4MQsw\nCQYDVQQLDAJ4eDELMAkGA1UEAwwCeHgxGTAXBgkqhkiG9w0BCQEWCnh4QDE2My5j\nb20wHhcNMTcxMjA0MDM0MjQ5WhcNMjAxMjAzMDM0MjQ5WjBpMQswCQYDVQQGEwJ4\neDELMAkGA1UECAwCeHgxCzAJBgNVBAcMAnh4MQswCQYDVQQKDAJ4eDELMAkGA1UE\nCwwCeHgxCzAJBgNVBAMMAnh4MRkwFwYJKoZIhvcNAQkBFgp4eEAxNjMuY29tMIIB\nIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwZ5UJULAjWr7p6FVwGRQRjFN\n2s8tZ/6LC3X82fajpVsYqF1xqEuUDndDXVD09E4u83MS6HO6a3bIVQDp6/klnYld\niE6Vp8HH5BSKaCWKVg8lGWg1UM9wZFnlryi14KgmpIFmcu9nA8yV/6MZAe6RSDmb\n3iyNBmiZ8aZhGw2pI1YwR+15MVqFFGB+7ExkziROi7L8CFCyCezK2/oOOvQsH1dz\nQ8z1JXWdg8/9Zx7Ktvgwu5PQM3cJtSHX6iBPOkMU8Z8TugLlTqQXKZOEgwajwvQ5\nmf2DPkVgM08XAgaLJcLigwD513koAdtJd5v+9irw+5LAuO3JclqwTvwy7u/YwwID\nAQABo1AwTjAdBgNVHQ4EFgQUo5A2tIu+bcUfvGTD7wmEkhXKFjcwHwYDVR0jBBgw\nFoAUo5A2tIu+bcUfvGTD7wmEkhXKFjcwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B\nAQsFAAOCAQEAWJ2rS6Mvlqk3GfEpboezx2J3X7l1z8Sxoqg6ntwB+rezvK3mc9H0\n83qcVeUcoH+0A0lSHyFN4FvRQL6X1hEheHarYwJK4agb231vb5erasuGO463eYEG\nr4SfTuOm7SyiV2xxbaBKrXJtpBp4WLL/s+LF+nklKjaOxkmxUX0sM4CTA7uFJypY\nc8Tdr8lDDNqoUtMD8BrUCJi+7lmMXRcC3Qi3oZJW76ja+kZA5mKVFPd1ATih8TbA\ni34R7EQDtFeiSvBdeKRsPp8c0KT8H1B4lXNkkCQs2WX5p4lm99+ZtLD4glw8x6Ic\ni1YhgnQbn5E0hz55OLu5jvOkKQjPCW+9Aa==\n-----END CERTIFICATE-----"

	updateOpts := certificates.UpdateOpts{
		Name:        certificateNewName,
		Certificate: updateCertificate,
	}

	if err := certificates.Update(client, lbaasCertificateId, updateOpts).Err; err != nil {
		return err
	}
	t.Logf("LbaasV2 certificate successfully updated: %s", lbaasCertificateId)
	return nil
}
