package hypervisors

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get makes a request against the API to get details for specific hypervisor.
func Get(client *golangsdk.ServiceClient, hypervisorID int) (*Hypervisor, error) {
	raw, err := client.Get(client.ServiceURL("os-hypervisors", strconv.Itoa(hypervisorID)), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Hypervisor
	err = extract.IntoStructPtr(raw.Body, &res, "hypervisor")
	return &res, err
}

// Hypervisor represents a hypervisor in the OpenStack cloud.
type Hypervisor struct {
	// A structure that contains cpu information like arch, model, vendor,
	// features and topology.
	CPUInfo CPUInfo `json:"-"`
	// The current_workload is the number of tasks the hypervisor is responsible
	// for. This will be equal or greater than the number of active VMs on the
	// system (it can be greater when VMs are being deleted and the hypervisor is
	// still cleaning up).
	CurrentWorkload int `json:"current_workload"`
	// Status of the hypervisor, either "enabled" or "disabled".
	Status string `json:"status"`
	// State of the hypervisor, either "up" or "down".
	State string `json:"state"`
	// DiskAvailableLeast is the actual free disk on this hypervisor,
	// measured in GB.
	DiskAvailableLeast int `json:"disk_available_least"`
	// HostIP is the hypervisor's IP address.
	HostIP string `json:"host_ip"`
	// FreeDiskGB is the free disk remaining on the hypervisor, measured in GB.
	FreeDiskGB int `json:"-"`
	// FreeRAMMB is the free RAM in the hypervisor, measured in MB.
	FreeRamMB int `json:"free_ram_mb"`
	// HypervisorHostname is the hostname of the hypervisor.
	HypervisorHostname string `json:"hypervisor_hostname"`
	// HypervisorType is the type of hypervisor.
	HypervisorType string `json:"hypervisor_type"`
	// HypervisorVersion is the version of the hypervisor.
	HypervisorVersion int `json:"-"`
	// ID is the unique ID of the hypervisor.
	ID int `json:"id"`
	// LocalGB is the disk space in the hypervisor, measured in GB.
	LocalGB int `json:"-"`
	// LocalGBUsed is the used disk space of the  hypervisor, measured in GB.
	LocalGBUsed int `json:"local_gb_used"`
	// MemoryMB is the total memory of the hypervisor, measured in MB.
	MemoryMB int `json:"memory_mb"`
	// MemoryMBUsed is the used memory of the hypervisor, measured in MB.
	MemoryMBUsed int `json:"memory_mb_used"`
	// RunningVMs is the number of running vms on the hypervisor.
	RunningVMs int `json:"running_vms"`
	// Service is the service this hypervisor represents.
	Service Service `json:"service"`
	// VCPUs is the total number of vcpus on the hypervisor.
	VCPUs int `json:"vcpus"`
	// VCPUsUsed is the number of used vcpus on the hypervisor.
	VCPUsUsed int `json:"vcpus_used"`
}

// Topology represents a CPU Topology.
type Topology struct {
	Sockets int `json:"sockets"`
	Cores   int `json:"cores"`
	Threads int `json:"threads"`
}

// CPUInfo represents CPU information of the hypervisor.
type CPUInfo struct {
	Vendor   string   `json:"vendor"`
	Arch     string   `json:"arch"`
	Model    string   `json:"model"`
	Features []string `json:"features"`
	Topology Topology `json:"topology"`
}

// Service represents a Compute service running on the hypervisor.
type Service struct {
	Host           string `json:"host"`
	ID             int    `json:"id"`
	DisabledReason string `json:"disabled_reason"`
}
