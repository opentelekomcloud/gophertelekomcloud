package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateImageOpts struct {
	ImageId string `json:"-" required:"true"`
	// Specifies the operation. The value can be add, replace, or remove.
	Op string `json:"op"`
	// Specifies the name of the attribute to be modified. / needs to be added in front of it.
	//
	// You can modify the following attributes:
	//
	// name: specifies the image name.
	// __description: specifies the image description.
	// __support_xen: Xen is supported.
	// __support_largememory: Ultra-large memory is supported.
	// __support_diskintensive: Intensive storage is supported.
	// __support_highperformance: High-performance computing (HPC) is supported.
	// __support_xen_gpu_type: GPU-accelerated ECSs that use Xen for virtualization are supported.
	// __support_kvm_gpu_type: GPU-accelerated ECSs that use KVM for virtualization are supported.
	// __support_xen_hana: HANA ECSs that use Xen for virtualization are supported.
	// __is_config_init: specifies whether initialization configuration is complete.
	// enterprise_project_id: specifies the enterprise project ID.
	// min_ram: specifies the minimum memory.
	// hw_vif_multiqueue_enabled: The NIC multi-queue feature is supported.
	// hw_firmware_type: specifies the boot mode. The value can be bios or uefi.
	// You can add or delete extension attributes.
	Path string `json:"path"`
	// Specifies the new value of the attribute. For detailed description, see Image Attributes.
	Value string `json:"value"`
}

// UpdateImage This API is used to modify image attributes and update image information.
// Only information of images in active status can be changed.
func UpdateImage(client *golangsdk.ServiceClient, opts UpdateImageOpts) (*ImageInfo, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PATCH /v2/cloudimages/{image_id}
	raw, err := client.Patch(client.ServiceURL("cloudimages", opts.ImageId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ImageInfo
	err = extract.Into(raw.Body, &res)
	return &res, err
}
