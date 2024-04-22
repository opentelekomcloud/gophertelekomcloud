package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateFuncMetadataOpts struct {
	FuncUrn              string                `json:"-"`
	Name                 string                `json:"func_name" required:"true"`
	Runtime              string                `json:"runtime" required:"true"`
	Timeout              int                   `json:"timeout" required:"true"`
	Handler              string                `json:"handler" required:"true"`
	MemorySize           int                   `json:"memory_size" required:"true"`
	GpuMemory            *int                  `json:"gpu_memory,omitempty"`
	UserData             string                `json:"user_data,omitempty"`
	EncryptedUserData    string                `json:"encrypted_user_data,omitempty"`
	Xrole                string                `json:"xrole,omitempty"`
	AppXrole             string                `json:"app_xrole,omitempty"`
	Description          string                `json:"description,omitempty"`
	FuncVpc              *FuncVpc              `json:"func_vpc,omitempty"`
	MountConfig          *MountConfig          `json:"mount_config,omitempty"`
	StrategyConfig       *StrategyConfig       `json:"strategy_config,omitempty"`
	CustomImage          *CustomImage          `json:"custom_image,omitempty"`
	Package              string                `json:"package"`
	ExtendConfig         string                `json:"extend_config,omitempty"`
	InitHandler          string                `json:"initializer_handler,omitempty"`
	InitTimeout          *int                  `json:"initializer_timeout,omitempty"`
	PreStopHandler       string                `json:"pre_stop_handler,omitempty"`
	PreStopTimeout       *int                  `json:"pre_stop_timeout,omitempty"`
	EphemeralStorage     *int                  `json:"ephemeral_storage,omitempty"`
	LogConfig            *FuncLogConfig        `json:"log_config,omitempty"`
	NetworkController    *NetworkControlConfig `json:"network_controller,omitempty"`
	IsStatefulFunction   *bool                 `json:"is_stateful_function,omitempty"`
	EnableDynamicMemory  *bool                 `json:"enable_dynamic_memory,omitempty"`
	EnableAuthInHeader   *bool                 `json:"enable_auth_in_header,omitempty"`
	DomainNames          string                `json:"domain_names,omitempty"`
	RestoreHookHandler   string                `json:"restore_hook_handler,omitempty"`
	RestoreHookTimeout   *int                  `json:"restore_hook_timeout,omitempty"`
	HeartbeatHandler     string                `json:"heartbeat_handler,omitempty"`
	EnableClassIsolation *bool                 `json:"enable_class_isolation,omitempty"`
	GpuType              string                `json:"gpu_type,omitempty"`
}

func UpdateFuncMetadata(client *golangsdk.ServiceClient, opts UpdateFuncMetadataOpts) (*FuncGraph, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "functions", opts.FuncUrn, "config"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncGraph
	return &res, extract.Into(raw.Body, &res)
}
