package resource

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/backup"
)

// ResourceBackupCapOpts contains the options for querying whether resources can be backed up.
type ResourceBackupCapOpts struct {
	// ID of the resource (server, or EVS disk) to be checked
	// For details about how to obtain the server ID, see the Elastic Cloud Server API Reference.
	// For details about how to obtain the disk ID, see the Elastic Volume Service API Reference.
	ResourceId string `json:"resource_id" required:"true"`
	// Type of the resource to be checked, for example, OS::Nova::Server for an ECS
	ResourceType string `json:"resource_type" required:"true"`
}

// GetResBackupCapabilities will query whether resources can be backed up based on the values in ResourceBackupCapOpts. To extract
// the ResourceCap object from the response, call the ExtractQueryResponse method on the QueryResult.
func GetResBackupCapabilities(client *golangsdk.ServiceClient, opts []ResourceBackupCapOpts) ([]ResourceCapability, error) {
	return doAction(client, opts, "check_protectable", "protectable")
}

func doAction(client *golangsdk.ServiceClient, opts interface{}, parent, label string) ([]ResourceCapability, error) {
	b, err := build.RequestBody(opts, parent)
	if err != nil {
		return nil, err
	}

	// POST https://{endpoint}/v1/{project_id}/providers/{provider_id}/resources/action
	raw, err := client.Post(client.ServiceURL("providers", backup.ProviderID, "resources", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []ResourceCapability
	err = extract.IntoSlicePtr(raw.Body, &res, label)
	return res, err
}

type ResourceCapability struct {
	// Whether backup or restoration is supported
	// true: yes
	// false: no
	Result bool `json:"result"`
	// Resource type
	// Possible values are OS::Nova::Server (ECS) and OS::Ironic::BareMetalServer (BMS).
	ResourceType string `json:"resource_type"`
	// Error code. If an error occurs, a value is returned.
	ErrorCode string `json:"error_code"`
	// Error message, which will be returned if the VM is associated with a backup policy. If an error occurs, a value is returned.
	ErrorMsg string `json:"error_msg"`
	// Resource ID
	ResourceId string `json:"resource_id"`
}
