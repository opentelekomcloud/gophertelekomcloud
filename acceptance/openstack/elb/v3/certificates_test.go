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

	listOpts := certificates.ListOpts{}
	certificatePages, err := certificates.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	certificateList, err := certificates.ExtractCertificates(certificatePages)
	th.AssertNoErr(t, err)

	for _, cert := range certificateList {
		tools.PrintResource(t, cert)
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
		Description: emptyDescription,
	}

	_, err = certificates.Update(client, certificateID, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Updated ELBv3 certificate: %s", certificateID)

	newCertificate, err := certificates.Get(client, certificateID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateOpts.Name, newCertificate.Name)
	th.AssertEquals(t, emptyDescription, newCertificate.Description)
}
