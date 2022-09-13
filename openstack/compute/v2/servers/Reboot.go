package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

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
	b, err := opts.ToServerRebootMap()
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("servers", id, "action"), b, nil, nil)
	return
}
