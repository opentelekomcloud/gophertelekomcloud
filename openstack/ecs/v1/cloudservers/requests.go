package cloudservers

import (
	"encoding/base64"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type CreateOpts struct {
	// ImageRef ID  the ID of the system image used for creating ECSs.
	ImageRef string `json:"imageRef" required:"true"`

	// FlavorRef ID of the ECS to be created.
	FlavorRef string `json:"flavorRef" required:"true"`

	// Name of the ECS instance.
	Name string `json:"name" required:"true"`

	// UserData to be injected during the ECS creation process.
	UserData []byte `json:"-"`

	// AdminPass sets the root user password. If not set, a randomly-generated
	// password will be created and returned in the response.
	AdminPass string `json:"adminPass,omitempty"`

	// KeyName of the SSH key used for logging in to the ECS.
	KeyName string `json:"key_name,omitempty"`

	// VpcId of the VPC to which the ECS belongs.
	VpcId string `json:"vpcid" required:"true"`

	// Nics information of the ECS.
	Nics []Nic `json:"nics" required:"true"`

	// PublicIp of the ECS.
	PublicIp *PublicIp `json:"publicip,omitempty"`

	// Count of ECSs to be created.
	// If this parameter is not specified, the default value is 1.
	Count int `json:"count,omitempty"`

	// ECS RootVolume configurations.
	RootVolume RootVolume `json:"root_volume" required:"true"`

	// ECS DataVolumes configurations.
	DataVolumes []DataVolume `json:"data_volumes,omitempty"`

	// SecurityGroups of the ECS.
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`

	// AvailabilityZone specifies name of the AZ where the ECS is located.
	AvailabilityZone string `json:"availability_zone" required:"true"`

	// ExtendParam provides the supplementary information about the ECS to be created.
	ExtendParam *ServerExtendParam `json:"extendparam,omitempty"`

	// MetaData specifies the metadata of the ECS to be created.
	MetaData *MetaData `json:"metadata,omitempty"`

	// SchedulerHints schedules ECSs, for example, by configuring an ECS group.
	SchedulerHints *SchedulerHints `json:"os:scheduler_hints,omitempty"`

	// ECS Tags.
	Tags []string `json:"tags,omitempty"`

	ServerTags []ServerTags `json:"server_tags,omitempty"`
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToServerCreateMap() (map[string]interface{}, error)
}

// ToServerCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToServerCreateMap() (map[string]interface{}, error) {
	b, err := build.RequestBodyMap(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.UserData != nil {
		var userData string
		if _, err := base64.StdEncoding.DecodeString(string(opts.UserData)); err != nil {
			userData = base64.StdEncoding.EncodeToString(opts.UserData)
		} else {
			userData = string(opts.UserData)
		}
		b["user_data"] = &userData
	}

	return map[string]interface{}{"server": b}, nil
}

type Nic struct {
	// SubnetId of the ECS.
	SubnetId string `json:"subnet_id" required:"true"`

	// IpAddress of the NIC used by the ECS.
	IpAddress string `json:"ip_address,omitempty"`

	// BindingProfile allows you to customize data.
	// Configure this parameter when creating a HANA ECS.
	BindingProfile BindingProfile `json:"binding:profile,omitempty"`

	// ExtraDhcpOpts indicates extended DHCP options.
	ExtraDhcpOpts []ExtraDhcpOpts `json:"extra_dhcp_opts,omitempty"`
}

type BindingProfile struct {
	// DisableSecurityGroups indicates that a HANA ECS NIC is not added to a security group.
	DisableSecurityGroups *bool `json:"disable_security_groups,omitempty"`
}

type ExtraDhcpOpts struct {
	// Set the parameter value to 26.
	OptName string `json:"opt_name" required:"true"`

	// OptValue specifies the NIC MTU, which ranges from 1280 to 8888.
	OptValue int `json:"opt_value" required:"true"`
}

type PublicIp struct {
	// Id of the existing EIP assigned to the ECS to be created.
	Id string `json:"id,omitempty"`

	// Eip that will be automatically assigned to an ECS.
	Eip *Eip `json:"eip,omitempty"`
}

type Eip struct {
	// Specifies the EIP type
	IpType string `json:"iptype" required:"true"`

	// Specifies the EIP bandwidth.
	BandWidth *BandWidth `json:"bandwidth" required:"true"`
}

type BandWidth struct {
	// Specifies the bandwidth size.
	Size int `json:"size" required:"true"`

	// Specifies the bandwidth sharing type
	ShareType string `json:"sharetype" required:"true"`

	// Specifies the bandwidth billing mode.
	ChargeMode string `json:"chargemode" required:"true"`
}

type RootVolume struct {
	// VolumeType of the ECS system disk.
	VolumeType string `json:"volumetype" required:"true"`

	// System disk Size, in GB.
	Size int `json:"size,omitempty"`

	ExtendParam *VolumeExtendParam `json:"extendparam,omitempty"`

	// Pay attention to this parameter if your ECS is SDI-compliant.
	// If the value of this parameter is true, the created disk is of SCSI type.
	PassThrough *bool `json:"hw:passthrough,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type DataVolume struct {
	// VolumeType of the ECS data disk.
	VolumeType string `json:"volumetype" required:"true"`

	// The data disk Size, in GB.
	Size int `json:"size" required:"true"`

	// MultiAttach is the shared disk information.
	MultiAttach *bool `json:"multiattach,omitempty"`

	// PassThrough indicates whether the data volume uses a SCSI lock.
	PassThrough *bool `json:"hw:passthrough,omitempty"`

	Extendparam *VolumeExtendParam `json:"extendparam,omitempty"`

	// DataImageID If data disks are created using a data disk
	// image, this parameter is mandatory and it does not support metadata.
	DataImageID string `json:"data_image_id,omitempty"`

	// EVS disk Metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type VolumeExtendParam struct {
	SnapshotId string `json:"snapshotId,omitempty"`
}

type ServerExtendParam struct {
	// RegionID is the ID of the region where the ECS resides.
	RegionID string `json:"regionID,omitempty"`

	// SupportAutoRecovery specifies whether automatic recovery is enabled on the ECS.
	SupportAutoRecovery string `json:"support_auto_recovery,omitempty"`
}

type MetaData struct {
	// AdminPass specifies the password of user Administrator for logging in to a Windows ECS.
	AdminPass string `json:"admin_pass,omitempty"`

	// OpSvcUserId specifies the user ID.
	OpSvcUserId string `json:"op_svc_userid,omitempty"`

	// AgencyName specifies the IAM agency name.
	AgencyName string `json:"agency_name,omitempty"`

	// If you have an OS or a software license, you can migrate your services to the cloud
	// platform in BYOL mode to continue using your existing licenses.
	BYOL string `json:"BYOL,omitempty"`
}

type SecurityGroup struct {
	// ID of the security group to which an ECS is to be added
	ID string `json:"id,omitempty"`
}

type SchedulerHints struct {
	// ECS Group ID, which is in UUID format.
	Group string `json:"group,omitempty"`

	// Specifies whether the ECS is created on a Dedicated Host (DeH) or in a shared pool.
	Tenancy string `json:"tenancy,omitempty"`

	// DedicatedHostID specifies a DeH ID.
	DedicatedHostID string `json:"dedicated_host_id,omitempty"`
}

type ServerTags struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value,omitempty"`
}

// Create requests a server to be provisioned to the user in the current tenant.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r JobResult) {
	b, err := opts.ToServerCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200}},
	)
	return
}

// DryRun requests a server to be provisioned to the user in the current tenant.
func DryRun(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r DryRunResult) {
	b, err := opts.ToServerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	b["dry_run"] = true

	_, r.Err = client.Post(createURL(client), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202}},
	)
	return
}

// Get retrieves a particular Server based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return
}

type DeleteOpts struct {
	// Servers to be deleted
	Servers []Server `json:"servers" required:"true"`

	// DeletePublicIP specifies whether to delete the EIP bound to the ECS when deleting the ECS.
	DeletePublicIP bool `json:"delete_publicip,omitempty"`

	// DeleteVolume specifies whether to delete the data disks of the ECS.
	DeleteVolume bool `json:"delete_volume,omitempty"`
}

type Server struct {
	// ID of the ECS to be deleted.
	Id string `json:"id" required:"true"`
}

// ToServerDeleteMap assembles a request body based on the contents of a
// DeleteOpts.
func (opts DeleteOpts) ToServerDeleteMap() (map[string]interface{}, error) {
	return build.RequestBodyMap(opts, "")
}

// Delete requests a server to be deleted to the user in the current tenant.
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (r JobResult) {
	b, err := opts.ToServerDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(deleteURL(client), b, &r.Body, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return
}
