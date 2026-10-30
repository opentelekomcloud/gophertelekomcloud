package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCertificateList(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	firstID := createCertificate(t, client)
	defer deleteCertificate(t, client, firstID)
	secondID := createCertificate(t, client)
	defer deleteCertificate(t, client, secondID)

	certificateList, err := certificates.List(client, certificates.ListOpts{
		ID: []string{firstID, secondID}, Limit: 1,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(certificateList))
	expected := map[string]bool{firstID: true, secondID: true}
	for _, cert := range certificateList {
		th.AssertEquals(t, true, expected[cert.ID])
		delete(expected, cert.ID)
	}
}

func TestCertificateLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	certificateID := createCertificate(t, client)
	defer deleteCertificate(t, client, certificateID)

	t.Logf("Attempting to update ELBv3 certificate: %s", certificateID)
	certName := tools.RandomString("update-cert-", 3)
	emptyDescription := ""
	updateOpts := certificates.UpdateOpts{
		Name:        certName,
		Description: &emptyDescription,
	}

	updated, err := certificates.Update(client, certificateID, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, certName, updated.Name)
	th.AssertEquals(t, emptyDescription, updated.Description)
	t.Logf("Updated ELBv3 certificate: %s", certificateID)

	newCertificate, err := certificates.Get(client, certificateID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newCertificate.Name)
	th.AssertEquals(t, emptyDescription, newCertificate.Description)
}
