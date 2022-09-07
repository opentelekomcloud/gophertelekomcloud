package secgroups

import (
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func commonList(client *golangsdk.ServiceClient, url string) pagination.Pager {
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SecurityGroupPage{pagination.SinglePageBase(r)}
	})
}

// GroupOpts is the underlying struct responsible for creating or updating
// security groups. It therefore represents the mutable attributes of a
// security group.
type GroupOpts struct {
	// the name of your security group.
	Name string `json:"name" required:"true"`
	// the description of your security group.
	Description string `json:"description" required:"true"`
}

// CreateOpts is the struct responsible for creating a security group.
type CreateOpts GroupOpts

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSecGroupCreateMap() (map[string]interface{}, error)
}

// ToSecGroupCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToSecGroupCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "security_group")
}

// UpdateOpts is the struct responsible for updating an existing security group.
type UpdateOpts GroupOpts

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSecGroupUpdateMap() (map[string]interface{}, error)
}

// ToSecGroupUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToSecGroupUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "security_group")
}

// DeleteWithRetry will try to permanently delete a particular security
// group based on its unique ID and RetryTimeout.
func DeleteWithRetry(client *golangsdk.ServiceClient, id string, timeout int) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		_, err := client.Delete(client.ServiceURL("os-security-groups", id), nil)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault400); ok {
				time.Sleep(10 * time.Second)
				return false, nil
			}
			return false, err
		}
		return true, nil
	})
}

// CreateRuleOpts represents the configuration for adding a new rule to an
// existing security group.
type CreateRuleOpts struct {
	// ID is the ID of the group that this rule will be added to.
	ParentGroupID string `json:"parent_group_id" required:"true"`

	// FromPort is the lower bound of the port range that will be opened.
	// Use -1 to allow all ICMP traffic.
	FromPort int `json:"from_port"`

	// ToPort is the upper bound of the port range that will be opened.
	// Use -1 to allow all ICMP traffic.
	ToPort int `json:"to_port"`

	// IPProtocol the protocol type that will be allowed, e.g. TCP.
	IPProtocol string `json:"ip_protocol" required:"true"`

	// CIDR is the network CIDR to allow traffic from.
	// This is ONLY required if FromGroupID is blank. This represents the IP
	// range that will be the source of network traffic to your security group.
	// Use 0.0.0.0/0 to allow all IP addresses.
	CIDR string `json:"cidr,omitempty" or:"FromGroupID"`

	// FromGroupID represents another security group to allow access.
	// This is ONLY required if CIDR is blank. This value represents the ID of a
	// group that forwards traffic to the parent group. So, instead of accepting
	// network traffic from an entire IP range, you can instead refine the
	// inbound source by an existing security group.
	FromGroupID string `json:"group_id,omitempty" or:"CIDR"`
}

// CreateRuleOptsBuilder allows extensions to add additional parameters to the
// CreateRule request.
type CreateRuleOptsBuilder interface {
	ToRuleCreateMap() (map[string]interface{}, error)
}

// ToRuleCreateMap builds a request body from CreateRuleOpts.
func (opts CreateRuleOpts) ToRuleCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "security_group_rule")
}

func actionMap(prefix, groupName string) map[string]map[string]string {
	return map[string]map[string]string{
		prefix + "SecurityGroup": {"name": groupName},
	}
}
