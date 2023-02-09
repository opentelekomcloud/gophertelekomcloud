package ims

import (
	"io"
	"net/http"
	"os"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func getClient(t *testing.T) (*golangsdk.ServiceClient, *golangsdk.ServiceClient) {
	v1, err := clients.NewIMSV1Client()
	th.AssertNoErr(t, err)

	v2, err := clients.NewIMSV2Client()
	th.AssertNoErr(t, err)

	return v1, v2
}

func downloadIMG(t *testing.T) (*os.File, error) {
	img, err := os.CreateTemp("", "ims-rancher.img")
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = img.Close()
		th.AssertNoErr(t, err)
		err = os.Remove(img.Name())
		th.AssertNoErr(t, err)
	})

	resp, err := http.Get("https://releases.rancher.com/os/latest/rancheros-openstack.img")
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		err = resp.Body.Close()
		th.AssertNoErr(t, err)
	})

	_, err = io.Copy(img, resp.Body)
	th.AssertNoErr(t, err)
	return img, err
}
