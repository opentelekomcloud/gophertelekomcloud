package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	c "github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/certifcates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCertificateDownload(t *testing.T) {
	t.Skip("No need to run this test in CI not secure")
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)
	_, err = c.Get(client)
	th.AssertNoErr(t, err)
}
