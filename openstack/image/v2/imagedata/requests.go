package imagedata

import (
	"io"
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Stage performs PUT call on the existing image object in the Imageservice with
// the provided file.
// Existing image object must be in the "queued" status.
func Stage(client *golangsdk.ServiceClient, id string, data io.Reader) (r StageResult) {
	_, r.Err = client.Put(client.ServiceURL("images", id, "stage"), data, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/octet-stream"},
		OkCodes:     []int{204},
	})
	return
}

// Download retrieves an image.
func Download(client *golangsdk.ServiceClient, id string) (r DownloadResult) {
	var resp *http.Response
	resp, r.Err = client.Get(client.ServiceURL("images", id, "file"), nil, nil)
	if resp != nil {
		r.Body = nil
		r.reader = resp.Body
		r.Header = resp.Header
	}
	return
}
