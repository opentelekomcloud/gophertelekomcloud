package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOpsBuilder is used for creating instance parameters.
// any struct providing the parameters should implement this interface
type CreateOpsBuilder interface {
	ToInstanceCreateMap() (map[string]interface{}, error)
}

// CreateOps is a struct that contains all the parameters.

// InstanceBackupPolicy for dcs

// PeriodicalBackupPlan for dcs

type ListDcsInstanceOpts struct {
	Id            string `q:"id"`
	Name          string `q:"name"`
	Type          string `q:"type"`
	DataStoreType string `q:"datastore_type"`
	VpcId         string `q:"vpc_id"`
	SubnetId      string `q:"subnet_id"`
	Offset        int    `q:"offset"`
	Limit         int    `q:"limit"`
}

type ListDcsBuilder interface {
	ToDcsListDetailQuery() (string, error)
}

func (opts ListDcsInstanceOpts) ToDcsListDetailQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

// ToInstanceCreateMap is used for type convert
func (ops CreateOps) ToInstanceCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(ops, "")
}

// UpdateOptsBuilder is an interface which can build the map paramter of update function
type UpdateOptsBuilder interface {
	ToInstanceUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	// DCS instance name.
	// An instance name is a string of 4–64 characters
	// that contain letters, digits, underscores (_), and hyphens (-).
	// An instance name must start with letters.
	Name string `json:"name,omitempty"`
	// Brief description of the DCS instance.
	// A brief description supports up to 1024 characters.
	Description *string `json:"description,omitempty"`
	// Backup policy.
	// This parameter is available for master/standby DCS instances.
	InstanceBackupPolicy *InstanceBackupPolicy `json:"instance_backup_policy,omitempty"`
	// Time at which the maintenance time window starts.
	// Format: HH:mm:ss
	MaintainBegin string `json:"maintain_begin,omitempty"`
	// Time at which the maintenance time window ends.
	// Format: HH:mm:ss
	MaintainEnd string `json:"maintain_end,omitempty"`
	// Security group ID.
	SecurityGroupID string `json:"security_group_id,omitempty"`
}

// ToInstanceUpdateMap is used for type convert
func (opts UpdateOpts) ToInstanceUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// UpdatePasswordOptsBuilder is an interface which can build the map paramter of update password function
type UpdatePasswordOptsBuilder interface {
	ToPasswordUpdateMap() (map[string]interface{}, error)
}

// UpdatePasswordOpts is a struct which represents the parameters of update function
type UpdatePasswordOpts struct {
	// Old password. It may be empty.
	OldPassword string `json:"old_password" required:"true"`
	// New password.
	// Password complexity requirements:
	// A string of 6–32 characters.
	// Must be different from the old password.
	// Contains at least two types of the following characters:
	// Uppercase letters
	// Lowercase letters
	// Digits
	// Special characters `~!@#$%^&*()-_=+\|[{}]:'",<.>/?
	NewPassword string `json:"new_password" required:"true"`
}

// ToPasswordUpdateMap is used for type convert
func (opts UpdatePasswordOpts) ToPasswordUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ExtendOptsBuilder is an interface which can build the map paramter of extend function
type ExtendOptsBuilder interface {
	ToExtendMap() (map[string]interface{}, error)
}

// ExtendOpts is a struct which represents the parameters of extend function
type ExtendOpts struct {
	// New specifications (memory space) of the DCS instance.
	// The new specification value to which the DCS instance
	// will be scaled up must be greater than the current specification value.
	// Unit: GB.
	NewCapacity int `json:"new_capacity" required:"true"`
	// New order ID.
	OrderID string `json:"order_id,omitempty"`
}

// ToExtendMap is used for type convert
func (opts ExtendOpts) ToExtendMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}
