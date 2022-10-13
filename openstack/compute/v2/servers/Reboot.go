package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// RebootMethod describes the mechanisms by which a server reboot can be requested.
type RebootMethod string

// These constants determine how a server should be rebooted.
// See the Reboot() function for further details.
const (
	SoftReboot RebootMethod = "SOFT"
	HardReboot RebootMethod = "HARD"
	OSReboot                = SoftReboot
	PowerCycle              = HardReboot
)

// RebootOpts provides options to the reboot request.
type RebootOpts struct {
	// Type is the type of reboot to perform on the server.
	Type RebootMethod `json:"type" required:"true"`
}

/*
Reboot requests that a given server reboot.

Two methods exist for rebooting a server:

HardReboot (aka PowerCycle) starts the server instance by physically cutting
power to the machine, or if a VM, terminating it at the hypervisor level.
It's done. Caput. Full stop.
Then, after a brief while, power is rtored or the VM instance restarted.

SoftReboot (aka OSReboot) simply tells the OS to restart under its own
procedure.
E.g., in Linux, asking it to enter runlevel 6, or executing
"sudo shutdown -r now", or by asking Windows to rtart the machine.
*/
func Reboot(client *golangsdk.ServiceClient, id string, opts RebootOpts) (err error) {
	b, err := build.RequestBody(opts, "reboot")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("servers", id, "action"), b, nil, nil)
	return
}
