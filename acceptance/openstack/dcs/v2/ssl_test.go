package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dcs/v2/ssl"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDcsInstanceSSLV2LifeCycle(t *testing.T) {
	client, err := clients.NewDcsV2Client()
	th.AssertNoErr(t, err)

	dcsInstance := createDCSInstance(t, client)

	sslOpts := ssl.SslOpts{
		InstanceId: dcsInstance.InstanceID,
		Enabled:    pointerto.Bool(true),
	}

	t.Logf("Attempting to enable SSL for DCSv2 instance")

	_, err = ssl.Update(client, sslOpts)
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 100, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)

	t.Logf("SSL enabled")

	t.Logf("Attempting to retrieve SSL settings for DCSv2 instance")

	getSsl, err := ssl.Get(client, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, getSsl.Enabled, true)
	th.AssertEquals(t, getSsl.SslValidated, true)

	t.Logf("SSL settings retrieved")

	t.Logf("Attempting to download SSL Cert")
	_, err = ssl.DownloadCert(client, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)
	t.Logf("DCS SSL Cert download successful")

	t.Logf("Attempting to disable SSL for DCSv2 instance")

	sslOpts.Enabled = pointerto.Bool(false)

	_, err = ssl.Update(client, sslOpts)
	th.AssertNoErr(t, err)

	err = waitForInstanceAvailable(client, 100, dcsInstance.InstanceID)
	th.AssertNoErr(t, err)

	t.Logf("SSL disabled for DCSv2 instance")
}
