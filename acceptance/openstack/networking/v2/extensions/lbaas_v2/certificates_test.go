package lbaas_v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/lbaas_v2/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLbaasV2CertificatesList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	listOpts := certificates.ListOpts{}
	allPages, err := certificates.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	lbaasCertificates, err := certificates.ExtractCertificates(allPages)
	th.AssertNoErr(t, err)

	for _, certificate := range lbaasCertificates {
		tools.PrintResource(t, certificate)
	}
}

func TestLbaasV2CertificateLifeCycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create lbaasV2 certificate
	lbaasCertificate, err := createLbaasCertificate(t, client)
	th.AssertNoErr(t, err)
	defer deleteLbaasCertificate(t, client, lbaasCertificate.ID)

	tools.PrintResource(t, lbaasCertificate)

	err = updateLbaasCertificate(t, client, lbaasCertificate.ID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, lbaasCertificate)

	newLbaasCertificate, err := certificates.Get(client, lbaasCertificate.ID).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, newLbaasCertificate)
}
