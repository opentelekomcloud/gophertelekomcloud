package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name                string                `json:"func_name" required:"true"`
	Package             string                `json:"package" required:"true"`
	Runtime             string                `json:"runtime" required:"true"`
	Timeout             int                   `json:"timeout" required:"true"`
	Handler             string                `json:"handler" required:"true"`
	DependVersionList   []string              `json:"depend_version_list,omitempty"`
	FuncVpc             *FuncVpc              `json:"func_vpc,omitempty"`
	MemorySize          int                   `json:"memory_size" required:"true"`
	GpuMemory           *int                  `json:"gpu_memory,omitempty"`
	CodeType            string                `json:"code_type" required:"true"`
	CodeURL             string                `json:"code_url,omitempty"`
	CodeFilename        string                `json:"code_filename,omitempty"`
	CustomImage         *CustomImage          `json:"custom_image,omitempty"`
	UserData            string                `json:"user_data,omitempty"`
	EncryptedUserData   string                `json:"encrypted_user_data,omitempty"`
	Xrole               string                `json:"xrole,omitempty"`
	AppXrole            string                `json:"app_xrole,omitempty"`
	Description         string                `json:"description,omitempty"`
	FuncCode            *FuncCode             `json:"func_code,omitempty"`
	MountConfig         *MountConfig          `json:"mount_config,omitempty"`
	InitHandler         string                `json:"initializer_handler,omitempty"`
	InitTimeout         *int                  `json:"initializer_timeout,omitempty"`
	PreStopHandler      string                `json:"pre_stop_handler,omitempty"`
	PreStopTimeout      *int                  `json:"pre_stop_timeout,omitempty"`
	Type                string                `json:"type,omitempty"`
	LogConfig           *FuncLogConfig        `json:"log_config,omitempty"`
	NetworkController   *NetworkControlConfig `json:"network_controller,omitempty"`
	IsStatefulFunction  *bool                 `json:"is_stateful_function,omitempty"`
	EnableDynamicMemory *bool                 `json:"enable_dynamic_memory,omitempty"`
}

type FuncVpc struct {
	DomainId       string   `json:"domain_id,omitempty"`
	ProjectID      string   `json:"namespace,omitempty"`
	VpcName        string   `json:"vpc_name,omitempty"`
	VpcID          string   `json:"vpc_id,omitempty"`
	SubnetName     string   `json:"subnet_name,omitempty"`
	SubnetID       string   `json:"subnet_id,omitempty"`
	CIDR           string   `json:"cidr,omitempty"`
	Gateway        string   `json:"gateway,omitempty"`
	SecurityGroups []string `json:"security_groups"`
}

type CustomImage struct {
	Enabled    *bool  `json:"bool,omitempty"`
	Image      string `json:"image,omitempty"`
	Command    string `json:"command,omitempty"`
	Args       string `json:"args,omitempty"`
	WorkingDir string `json:"working_dir,omitempty"`
	UID        string `json:"uid,omitempty"`
	GID        string `json:"gid,omitempty"`
}

type FuncCode struct {
	File string `json:"file,omitempty"`
	Link string `json:"link,omitempty"`
}

type MountConfig struct {
	MountUser  MountUser   `json:"mount_user" required:"true"`
	FuncMounts []FuncMount `json:"func_mounts" required:"true"`
}

type MountUser struct {
	UserID      string `json:"user_id" required:"true"`
	UserGroupID string `json:"user_group_id" required:"true"`
}

type FuncMount struct {
	MountType      string `json:"mount_type" required:"true"`
	MountResource  string `json:"mount_resource" required:"true"`
	MountSharePath string `json:"mount_share_path,omitempty"`
	LocalMountPath string `json:"local_mount_path" required:"true"`
}

type FuncLogConfig struct {
	GroupName  string `json:"group_name,omitempty"`
	GroupID    string `json:"group_id,omitempty"`
	StreamName string `json:"stream_name,omitempty"`
	StreamID   string `json:"stream_id,omitempty"`
}

type NetworkControlConfig struct {
	DisablePublicNetwork *bool       `json:"disable_public_network,omitempty"`
	TriggerAccessVpcs    []VpcConfig `json:"trigger_access_vpcs,omitempty"`
}

type VpcConfig struct {
	VpcName string `json:"vpc_name,omitempty"`
	VpcID   string `json:"vpc_id,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*FuncGraph, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "functions"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncGraph
	return &res, extract.Into(raw.Body, &res)
}

type FuncGraph struct {
	FuncID                   string               `json:"func_id"`
	FuncURN                  string               `json:"func_urn"`
	FuncName                 string               `json:"func_name"`
	DomainID                 string               `json:"domain_id"`
	ProjectID                string               `json:"namespace"`
	ProjectName              string               `json:"project_name"`
	Package                  string               `json:"package"`
	Runtime                  string               `json:"runtime"`
	Timeout                  int                  `json:"timeout"`
	Handler                  string               `json:"handler"`
	MemorySize               int                  `json:"memory_size"`
	GpuMemory                int                  `json:"gpu_memory"`
	CPU                      int                  `json:"cpu"`
	CodeType                 string               `json:"code_type"`
	CodeURL                  string               `json:"code_url"`
	CodeFilename             string               `json:"code_filename"`
	CodeSize                 int                  `json:"code_size"`
	DomainNames              string               `json:"domain_names"`
	UserData                 string               `json:"user_data"`
	EncryptedUserData        string               `json:"encrypted_user_data"`
	Digest                   string               `json:"digest"`
	Version                  string               `json:"version"`
	ImageName                string               `json:"image_name"`
	Xrole                    string               `json:"xrole"`
	AppXrole                 string               `json:"app_xrole"`
	Description              string               `json:"description"`
	LastModified             string               `json:"last_modified"`
	FuncVpc                  FuncVpc              `json:"func_vpc"`
	MountConfig              MountConfig          `json:"mount_config"`
	ReservedInstanceCount    int                  `json:"reserved_instance_count"`
	DependVersionList        []string             `json:"depend_version_list"`
	StrategyConfig           StrategyConfig       `json:"strategy_config"`
	ExtendConfig             string               `json:"extend_config"`
	Dependencies             []Dependency         `json:"dependencies"`
	InitHandler              string               `json:"initializer_handler"`
	InitTimeout              int                  `json:"initializer_timeout"`
	PreStopHandler           string               `json:"pre_stop_handler"`
	PreStopTimeout           string               `json:"pre_stop_timeout"`
	LongTime                 bool                 `json:"long_time"`
	LogGroupID               string               `json:"log_group_id"`
	LogStreamID              string               `json:"log_stream_id"`
	Type                     string               `json:"type"`
	EnableDynamicMemory      bool                 `json:"enable_dynamic_memory"`
	IsStatefulFunction       bool                 `json:"is_stateful_function"`
	CustomImage              CustomImage          `json:"custom_image"`
	IsBridgeFunction         bool                 `json:"is_bridge_function"`
	ApigRouteEnable          bool                 `json:"apig_route_enable"`
	HeartbeatHandler         string               `json:"heartbeat_handler"`
	EnableClassIsolation     bool                 `json:"enable_class_isolation"`
	GpuType                  string               `json:"gpu_type"`
	AllowEphemeralStorage    bool                 `json:"allow_ephemeral_storage"`
	EphemeralStorage         int                  `json:"ephemeral_storage"`
	NetworkController        NetworkControlConfig `json:"network_controller"`
	ResourceID               string               `json:"resource_id"`
	EnableAuthInHeader       bool                 `json:"enable_auth_in_header"`
	ReservedInstanceIdleMode bool                 `json:"reserved_instance_idle_mode"`
}

type StrategyConfig struct {
	Concurrency   int `json:"concurrency"`
	ConcurrentNum int `json:"concurrent_num"`
}

type Dependency struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	Link         string `json:"link"`
	Runtime      string `json:"runtime"`
	Etag         string `json:"etag"`
	Size         int    `json:"size"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	FileName     string `json:"file_name"`
	Version      int    `json:"version"`
	DepID        string `json:"dep_id"`
	LastModified string `json:"last_modified"`
}
