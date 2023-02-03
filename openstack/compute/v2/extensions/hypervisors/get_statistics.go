package hypervisors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetStatistics makes a request against the API to get hypervisors statistics.
func GetStatistics(client *golangsdk.ServiceClient) (*Statistics, error) {
	raw, err := client.Get(client.ServiceURL("os-hypervisors", "statistics"), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Statistics
	err = extract.IntoStructPtr(raw.Body, &res, "hypervisor_statistics")
	return &res, err
}

// Statistics represents a summary statistics for all enabled
// hypervisors over all compute nodes in the OpenStack cloud.
type Statistics struct {
	// The number of hypervisors.
	Count int `json:"count"`
	// The current_workload is the number of tasks the hypervisor is responsible for
	CurrentWorkload int `json:"current_workload"`
	// The actual free disk on this hypervisor(in GB).
	DiskAvailableLeast int `json:"disk_available_least"`
	// The free disk remaining on this hypervisor(in GB).
	FreeDiskGB int `json:"free_disk_gb"`
	// The free RAM in this hypervisor(in MB).
	FreeRamMB int `json:"free_ram_mb"`
	// The disk in this hypervisor(in GB).
	LocalGB int `json:"local_gb"`
	// The disk used in this hypervisor(in GB).
	LocalGBUsed int `json:"local_gb_used"`
	// The memory of this hypervisor(in MB).
	MemoryMB int `json:"memory_mb"`
	// The memory used in this hypervisor(in MB).
	MemoryMBUsed int `json:"memory_mb_used"`
	// The total number of running vms on all hypervisors.
	RunningVMs int `json:"running_vms"`
	// The number of vcpu in this hypervisor.
	VCPUs int `json:"vcpus"`
	// The number of vcpu used in this hypervisor.
	VCPUsUsed int `json:"vcpus_used"`
}
