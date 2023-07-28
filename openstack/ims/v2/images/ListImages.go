package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"net/url"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListImagesOpts struct {
	// Specifies the image type. The following types are supported:
	//
	// Public image: The value is gold.
	// Private image: The value is private.
	// Shared image: The value is shared.
	// NOTE:
	// The __imagetype of images you share with other tenants or those other tenants share with you and you have accepted is shared. You can use field owner to distinguish the two types of shared images. You can use member_status to filter out shared images you have accepted.
	ImageType string `q:"__imagetype,omitempty"`
	// Specifies whether the image is available. The value can be true. The value is true for all extension APIs by default. Common users can query only the images for which the value of this parameter is true.
	IsRegistered string `q:"__isregistered,omitempty"`
	// Specifies whether the image is a full-ECS image. The value can be true or false.
	WholeImage *bool `q:"__whole_image,omitempty"`
	// Specifies the ID of the key used to encrypt the image. You can obtain the ID from the IMS console or by calling the Querying Image Details (Native OpenStack API) API.
	SystemCmkId string `q:"__system__cmkid,omitempty"`
	// Specifies the OS architecture, 32 bit or 64 bit.
	OsBit string `q:"__os_bit,omitempty"`
	// Specifies the image OS type. Available values include:
	//
	// Linux
	// Windows
	// Other
	OsType string `q:"__os_type,omitempty"`
	// Specifies the image platform type. The value can be Windows, Ubuntu, RedHat, SUSE, CentOS, Debian, OpenSUSE, Oracle Linux, Fedora, Other, CoreOS, or EulerOS.
	Platform string `q:"__platform,omitempty"`
	// Specifies whether the image supports disk-intensive ECSs. If the image supports disk-intensive ECSs, the value is true. Otherwise, this parameter is not required.
	SupportDiskIntensive string `q:"__support_diskintensive,omitempty"`
	// Specifies whether the image supports high-performance ECSs. If the image supports high-performance ECSs, the value is true. Otherwise, this parameter is not required.
	SupportHighPerformance string `q:"__support_highperformance,omitempty"`
	// Specifies whether the image supports KVM. If yes, the value is true. Otherwise, this parameter is not required.
	SupportKvm string `q:"__support_kvm,omitempty"`
	// Specifies whether the image supports GPU-accelerated ECSs on the KVM platform. See Table 3 for its value. If the image does not support GPU-accelerated ECSs on the KVM platform, this parameter is not required. This attribute cannot co-exist with __support_xen and __support_kvm.
	SupportKvmGpuType string `q:"__support_kvm_gpu_type,omitempty"`
	// Specifies whether the image supports ECSs with the InfiniBand NIC on the KVM platform. If yes, the value is true. Otherwise, this parameter is not required.
	//
	// This attribute cannot co-exist with __support_xen.
	SupportKvmInfiniband string `q:"__support_kvm_infiniband,omitempty"`
	// Specifies whether the image supports large-memory ECSs. If the image supports large-memory ECSs, the value is true. Otherwise, this parameter is not required.
	SupportLargeMemory string `q:"__support_largememory,omitempty"`
	// Specifies whether the image supports Xen. If yes, the value is true. Otherwise, this parameter is not required.
	SupportXen string `q:"__support_xen,omitempty"`
	// Specifies whether the image supports GPU-accelerated ECSs on the Xen platform. See Table 2 for its value. If the image does not support GPU-accelerated ECSs on the Xen platform, this parameter is not required. This attribute cannot co-exist with __support_xen and __support_kvm.
	SupportXenGpuType string `q:"__support_xen_gpu_type,omitempty"`
	// Specifies whether the image supports HANA ECSs on the Xen platform. If yes, the value is true. Otherwise, this parameter is not required.
	//
	// This attribute cannot co-exist with __support_xen and __support_kvm.
	SupportXenHana string `q:"__support_xen_hana,omitempty"`
	// Specifies the container type. The value is bare.
	ContainerFormat string `q:"container_format,omitempty"`
	// Specifies the image format. The value can be vhd, raw, zvhd, or qcow2. The default value is zvhd2.
	DiskFormat string `q:"disk_format,omitempty"`
	// Specifies the enterprise project to which the images to be queried belong.
	// If the value is 0, images of enterprise project default are to be queried.
	// If the value is UUID, images of the enterprise project corresponding to the UUID are to be queried.
	// If the value is all_granted_eps, images of all enterprise projects are to be queried.
	// For more information about enterprise projects and how to obtain enterprise project IDs, see Enterprise Management User Guide.
	EnterpriseProjectId string `q:"enterprise_project_id,omitempty"`
	// Specifies the image ID
	Id string `q:"id,omitempty"`
	// Specifies the number of images to be queried. The value is an integer and is 500 by default.
	Limit int `q:"limit,omitempty"`
	// Specifies the start number from which images are queried. The value is the image ID.
	Marker string `q:"marker,omitempty"`
	// Specifies the member status. The value can be accepted, rejected, or pending. accepted: indicates that the shared image is accepted. rejected indicates that the image shared by others is rejected. pending indicates that the image shared by others needs to be confirmed. To use this parameter, set visibility to shared during the query.
	MemberStatus string `q:"member_status,omitempty"`
	// Specifies the minimum disk space (GB) required for running the image. The value ranges from 1 GB to 1024 GB.
	MinDisk int `q:"min_disk,omitempty"`
	// Specifies the minimum memory size (MB) required for running the image. The parameter value depends on the ECS specifications. Generally, the value is 0.
	MinRam int `q:"min_ram,omitempty"`
	// Specifies the image name. Exact matching is used. For detailed description, see Image Attributes.
	Name string `q:"name,omitempty"`
	// Specifies the tenant to which the image belongs.
	Owner string `q:"owner,omitempty"`
	// Specifies whether the image is protected. The value can be true or false. Set it to true when you query public images. This parameter is optional when you query private images.
	Protected *bool `q:"protected,omitempty"`
	// Specifies whether the query results are sorted in ascending or descending order. Its value can be desc (default) or asc. This parameter is used together with parameter sort_key. The default value is desc.
	SortDir string `q:"sort_dir,omitempty"`
	// Specifies the field for sorting the query results. The value can be an attribute of the image: name, container_format, disk_format, status, id, size, or created_at. The default value is created_at.
	SortKey string `q:"sort_key,omitempty"`
	// Specifies the image status. The value can be one of the following:
	//
	// queued: indicates that the image metadata has already been created, and it is ready for the image file to upload.
	// saving: indicates that the image file is being uploaded to the backend storage.
	// deleted: indicates that the image has been deleted.
	// killed: indicates that an error occurs on the image uploading.
	// active: indicates that the image is available for use.
	Status string `q:"status,omitempty"`
	// Specifies a tag added to an image. Tags can be used as a filter to query images.
	//
	// NOTE:
	// The tagging function has been upgraded. If the tags added before the function upgrade are in the format of "Key.Value", query tags using "Key=Value". For example, an existing tag is a.b. After the tag function upgrade, query the tag using "tag=a=b".
	Tag string `q:"tag,omitempty"`
	// Specifies the environment where the image is used. The value can be FusionCompute, Ironic, DataImage, or IsoImage.
	//
	// For an ECS image (system disk image), the value is FusionCompute.
	// For a data disk image, the value is DataImage.
	// For a BMS image, the value is Ironic.
	// For an ISO image, the value is IsoImage.
	VirtualEnvType string `q:"virtual_env_type,omitempty"`
	// Specifies whether the image is available to other tenants. Available values include:
	//
	// public: public image
	// private: private image
	// shared: shared image
	Visibility string `q:"visibility,omitempty"`
	// Specifies the time when the image was created. Images can be queried by time. The value is in the format of Operator:UTC time.
	//
	// The following operators are supported:
	//
	// gt: greater than
	// gte: greater than or equal to
	// lt: less than
	// lte: less than or equal to
	// eq: equal to
	// neq: not equal to
	// The time format is yyyy-MM-ddThh:mm:ssZ or yyyy-MM-dd hh:mm:ss.
	//
	// For example, to query images created before Oct 28, 2018 10:00:00, set the value of created_at as follows:
	//
	// created_at=lt:2018-10-28T10:00:00Z
	CreatedAt string `q:"created_at,omitempty"`
	// Specifies the time when the image was modified. Images can be queried by time. The value is in the format of Operator:UTC time.
	//
	// The following operators are supported:
	//
	// gt: greater than
	// gte: greater than or equal to
	// lt: less than
	// lte: less than or equal to
	// eq: equal to
	// neq: not equal to
	// The time format is yyyy-MM-ddThh:mm:ssZ or yyyy-MM-dd hh:mm:ss.
	//
	// For example, to query images updated before Oct 28, 2018 10:00:00, set the value of updated_at as follows:
	//
	// updated_at=lt:2018-10-28T10:00:00Z
	UpdatedAt string `q:"updated_at,omitempty"`

	// SizeMin filters on the size_min image property.
	SizeMin int64 `q:"size_min"`
	// SizeMax filters on the size_max image property.
	SizeMax int64 `q:"size_max"`
}

// ListImages This API is used to query images using search criteria and to display the images in a list.
func ListImages(client *golangsdk.ServiceClient, opts ListImagesOpts) ([]ImageInfo, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET /v2/cloudimages
	raw, err := client.Get(client.ServiceURL("cloudimages")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ImageInfo
	err = extract.IntoSlicePtr(raw.Body, &res, "images")
	return res, err
}

type ImageInfo struct {
	// Specifies the backup ID. To create an image using a backup, set the value to the backup ID. Otherwise, this value is left empty.
	BackupId string `json:"__backup_id,omitempty"`
	// Specifies the image source.
	//
	// If the image is a public image, this parameter is left empty.
	DataOrigin string `json:"__data_origin,omitempty"`
	// Specifies the image description. For detailed description, see Image Attributes.
	Description string `json:"__description,omitempty"`
	// Specifies the size (bytes) of the image file.
	ImageSize string `json:"__image_size"`
	// Specifies the image backend storage type. Only UDS is supported currently.
	ImageSourceType string `json:"__image_source_type"`
	// Specifies the image type. The following types are supported:
	//
	// Public image: The value is gold.
	// Private image: The value is private.
	// Shared image: The value is shared.
	Imagetype string `json:"__imagetype"`
	// Specifies whether the image has been registered. The value can be true or false.
	Isregistered string `json:"__isregistered"`
	// Specifies the parent image ID.
	//
	// If the image is a public image or created from an image file, this value is left empty.
	Originalimagename string `json:"__originalimagename,omitempty"`
	// Specifies the OS architecture, 32 bit or 64 bit.
	OsBit string `json:"__os_bit,omitempty"`
	// Specifies the OS type. The value can be Linux, Windows, or Other.
	OsType string `json:"__os_type"`
	// Specifies the OS version.
	OsVersion string `json:"__os_version,omitempty"`
	// Specifies the image platform type. The value can be Windows, Ubuntu, RedHat, SUSE, CentOS, Debian, OpenSUSE, Oracle Linux, Fedora, Other, CoreOS, or EulerOS.
	Platform string `json:"__platform,omitempty"`
	// Specifies whether the image supports disk-intensive ECSs. If the image supports disk-intensive ECSs, the value is true. Otherwise, this parameter is not required.
	SupportDiskintensive string `json:"__support_diskintensive,omitempty"`
	// Specifies whether the image supports high-performance ECSs. If the image supports high-performance ECSs, the value is true. Otherwise, this parameter is not required.
	SupportHighperformance string `json:"__support_highperformance,omitempty"`
	// Specifies whether the image supports KVM. If yes, the value is true. Otherwise, this parameter is not required.
	SupportKvm string `json:"__support_kvm,omitempty"`
	// Specifies whether the image supports GPU-accelerated ECSs on the KVM platform. See Table 3 for its value.
	//
	// If the image does not support GPU-accelerated ECSs on the KVM platform, this parameter is not required. This attribute cannot co-exist with __support_xen and __support_kvm.
	SupportKvmGpuType string `json:"__support_kvm_gpu_type,omitempty"`
	// Specifies whether the image supports ECSs with the InfiniBand NIC on the KVM platform. If yes, the value is true. Otherwise, this parameter is not required.
	//
	// This attribute cannot co-exist with __support_xen.
	SupportKvmInfiniband string `json:"__support_kvm_infiniband,omitempty"`
	// Specifies whether the image supports large-memory ECSs. If the image supports large-memory ECSs, the value is true. Otherwise, this parameter is not required.
	SupportLargememory string `json:"__support_largememory,omitempty"`
	// Specifies whether the image supports Xen. If yes, the value is true. Otherwise, this parameter is not required.
	SupportXen string `json:"__support_xen,omitempty"`
	// Specifies whether the image supports GPU-accelerated ECSs on the Xen platform. See Table 2 for its value. If the image does not support GPU-accelerated ECSs on the Xen platform, this parameter is not required. This attribute cannot co-exist with __support_xen and __support_kvm.
	SupportXenGpuType string `json:"__support_xen_gpu_type,omitempty"`
	// Specifies whether the image supports HANA ECSs on the Xen platform. If yes, the value is true. Otherwise, this parameter is not required.
	//
	// This attribute cannot co-exist with __support_xen and __support_kvm.
	SupportXenHana string `json:"__support_xen_hana,omitempty"`
	// This parameter is unavailable currently.
	Checksum string `json:"checksum,omitempty"`
	// Specifies the container type.
	ContainerFormat string `json:"container_format"`
	// Specifies the time when the image was created. The value is in UTC format.
	CreatedAt time.Time `json:"created_at"`
	// Specifies the image format. The value can be vhd, raw, zvhd, or qcow2. The default value is vhd.
	DiskFormat string `json:"disk_format,omitempty"`
	// Specifies the enterprise project that the image belongs to.
	//
	// If the value is 0 or left blank, the image belongs to the default enterprise project.
	// If the value is a UUID, the image belongs to the enterprise project corresponding to the UUID.
	// For more information about enterprise projects, see Enterprise Management User Guide.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Specifies the URL for uploading and downloading the image file.
	File string `json:"file,omitempty"`
	// Specifies the image ID.
	Id string `json:"id"`
	// Specifies the minimum disk space (GB) required for running the image. The value ranges from 1 GB to 1024 GB.
	MinDisk int `json:"min_disk"`
	// Specifies the minimum memory size (MB) required for running the image. The parameter value depends on the ECS specifications. Generally, the value is 0.
	MinRam int `json:"min_ram"`
	// Specifies the image name. For detailed description, see Image Attributes.
	Name string `json:"name"`
	// Specifies the tenant to which the image belongs.
	Owner string `json:"owner"`
	// Specifies whether the image is protected. A protected image cannot be deleted. The value can be true or false.
	Protected *bool `json:"protected"`
	// Specifies the image schema.
	Schema string `json:"schema,omitempty"`
	// Specifies the image URL.
	Self string `json:"self"`
	// This parameter is unavailable currently.
	Size int `json:"size,omitempty"`
	// Specifies the image status. The value can be one of the following:
	//
	// queued: indicates that the image metadata has already been created, and it is ready for the image file to upload.
	// saving: indicates that the image file is being uploaded to the backend storage.
	// deleted: indicates that the image has been deleted.
	// killed: indicates that an error occurs on the image uploading.
	// active: indicates that the image is available for use.
	Status string `json:"status"`
	// Specifies tags of the image, through which you can manage private images in your own way. You can use the image tag API to add different tags to each image and filter images by tag.
	Tags []string `json:"tags"`
	// Specifies the time when the image was updated. The value is in UTC format.
	UpdatedAt time.Time `json:"updated_at"`
	// Specifies the environment where the image is used. The value can be FusionCompute, Ironic, DataImage, or IsoImage.
	//
	// For an ECS image, the value is FusionCompute.
	// For a data disk image, the value is DataImage.
	// For a BMS image, the value is Ironic.
	// For an ISO image, the value is IsoImage.
	VirtualEnvType string `json:"virtual_env_type"`
	// This parameter is unavailable currently.
	VirtualSize int `json:"virtual_size,omitempty"`
	// Specifies whether the image is available to other tenants. Available values include:
	//
	// private: private image
	// public: public image
	// shared: shared image
	Visibility string `json:"visibility"`
	// Specifies whether the image supports password/private key injection using Cloud-Init.
	//
	// If the value is set to true, password/private key injection using Cloud-Init is not supported.
	//
	// NOTE:
	// This parameter is valid only for ECS system disk images.
	SupportFcInject string `json:"__support_fc_inject,omitempty"`
	// Specifies the ECS boot mode. Available values include:
	//
	// bios indicates the BIOS boot mode.
	// uefi indicates the UEFI boot mode.
	HwFirmwareType string `json:"hw_firmware_type,omitempty"`
	// Specifies the maximum memory (MB) of the image. You can set this parameter based on the ECS specifications. Generally, you do not need to set this parameter.
	MaxRam string `json:"max_ram,omitempty"`
	// Specifies the ID of the key used to encrypt the image.
	SystemCmkid string `json:"__system__cmkid,omitempty"`
	// Specifies additional attributes of the image. The value is a list (in JSON format) of advanced features supported by the image.
	OsFeatureList string `json:"__os_feature_list,omitempty"`
	// Specifies whether the image supports NIC multi-queue. The value can be true or false.
	HwVifMultiqueueEnabled string `json:"hw_vif_multiqueue_enabled,omitempty"`
	// Specifies whether the image supports lazy loading. The value can be true, false, True, or False.
	Lazyloading string `json:"__lazyloading,omitempty"`
	// Specifies that the image is created from an external image file. Value: file
	RootOrigin string `json:"__root_origin,omitempty"`
	// Specifies the ECS system disk slot number corresponding to the image.
	//
	// Example value: 0
	SequenceNum string `json:"__sequence_num,omitempty"`
	// Specifies the time when the image status became active.
	ActiveAt string `json:"active_at"`
	// Specifies whether the image uses AMD's x86 architecture. The value can be true or false.
	SupportAmd string `json:"__support_amd,omitempty"`
	// Specifies the location where the image is stored.
	ImageLocation string `json:"__image_location"`
	// Specifies the charging identifier for the image.
	AccountCode string `json:"__account_code"`

	SupportKvmNvmeSpdk    string `json:"__support_kvm_nvme_spdk"`
	ImageLoginUser        string `json:"__image_login_user"`
	SupportKvmHi1822Hiovs string `json:"__support_kvm_hi1822_hiovs"`
}
