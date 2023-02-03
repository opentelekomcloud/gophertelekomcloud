package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// ResizeOpts represents the configuration options used to control a Resize operation.
type ResizeOpts struct {
	// FlavorRef is the ID of the flavor you wish your server to become.
	FlavorRef string `json:"flavorRef" required:"true"`
}

// Resize instructs the provider to change the flavor of the server.
//
// Note that this implies rebuilding it.
//
// Unfortunately, one cannot pass rebuild parameters to the resize function.
// When the resize completes, the server will be in VERIFY_RESIZE state.
// While in this state, you can explore the use of the new server's
// configuration. If you like it, call ConfirmResize() to commit the resize
// permanently. Otherwise, call RevertResize() to restore the old configuration.
func Resize(client *golangsdk.ServiceClient, id string, opts ResizeOpts) (err error) {
	b, err := build.RequestBody(opts, "resize")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("servers", id, "action"), b, nil, nil)
	return
}
