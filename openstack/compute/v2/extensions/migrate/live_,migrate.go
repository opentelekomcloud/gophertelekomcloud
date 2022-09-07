package migrate

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// LiveMigrateOptsBuilder allows extensions to add additional parameters to the
// LiveMigrate request.
type LiveMigrateOptsBuilder interface {
	ToLiveMigrateMap() (map[string]interface{}, error)
}

// LiveMigrateOpts specifies parameters of live migrate action.
type LiveMigrateOpts struct {
	// The host to which to migrate the server.
	// If this parameter is None, the scheduler chooses a host.
	Host *string `json:"host"`

	// Set to True to migrate local disks by using block migration.
	// If the source or destination host uses shared storage and you set
	// this value to True, the live migration fails.
	BlockMigration *bool `json:"block_migration,omitempty"`

	// Set to True to enable over commit when the destination host is checked
	// for available disk space. Set to False to disable over commit. This setting
	// affects only the libvirt virt driver.
	DiskOverCommit *bool `json:"disk_over_commit,omitempty"`
}

// ToLiveMigrateMap constructs a request body from LiveMigrateOpts.
func (opts LiveMigrateOpts) ToLiveMigrateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-migrateLive")
}

// LiveMigrate will initiate a live-migration (without rebooting) of the instance to another host.
func LiveMigrate(client *golangsdk.ServiceClient, id string, opts LiveMigrateOptsBuilder) (r MigrateResult) {
	b, err := opts.ToLiveMigrateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("servers", id, "action"), b, nil, nil)
	return
}