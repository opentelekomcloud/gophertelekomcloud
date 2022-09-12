package secgroups

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// SecurityGroup represents a security group.
type SecurityGroup struct {
	// The unique ID of the group. If Neutron is installed, this ID will be
	// represented as a string UUID; if Neutron is not installed, it will be a
	// numeric ID. For the sake of consistency, we always cast it to a string.
	ID string `json:"-"`
	// The human-readable name of the group, which needs to be unique.
	Name string `json:"name"`
	// The human-readable description of the group.
	Description string `json:"description"`
	// The rules which determine how this security group operates.
	Rules []Rule `json:"rules"`
	// The ID of the tenant to which this security group belongs.
	TenantID string `json:"tenant_id"`
}

func (r *SecurityGroup) UnmarshalJSON(b []byte) error {
	type tmp SecurityGroup
	var s struct {
		tmp
		ID interface{} `json:"id"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = SecurityGroup(s.tmp)

	switch t := s.ID.(type) {
	case float64:
		r.ID = strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		r.ID = t
	}

	return err
}

func extra(err error, raw *http.Response) (*SecurityGroup, error) {
	if err != nil {
		return nil, err
	}

	var res SecurityGroup
	err = extract.IntoStructPtr(raw.Body, &res, "security_group")
	return &res, err
}

// Rule represents a security group rule, a policy which determines how a
// security group operates and what inbound traffic it allows in.
type Rule struct {
	// The unique ID. If Neutron is installed, this ID will be
	// represented as a string UUID; if Neutron is not installed, it will be a
	// numeric ID. For the sake of consistency, we always cast it to a string.
	ID string `json:"-"`
	// The lower bound of the port range which this security group should open up.
	FromPort int `json:"from_port"`
	// The upper bound of the port range which this security group should open up.
	ToPort int `json:"to_port"`
	// The IP protocol (e.g. TCP) which the security group accepts.
	IPProtocol string `json:"ip_protocol"`
	// The CIDR IP range whose traffic can be received.
	IPRange IPRange `json:"ip_range"`
	// The security group ID to which this rule belongs.
	ParentGroupID string `json:"-"`
	// Not documented.
	Group Group
}

func (r *Rule) UnmarshalJSON(b []byte) error {
	type tmp Rule
	var s struct {
		tmp
		ID            interface{} `json:"id"`
		ParentGroupID interface{} `json:"parent_group_id"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Rule(s.tmp)

	switch t := s.ID.(type) {
	case float64:
		r.ID = strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		r.ID = t
	}

	switch t := s.ParentGroupID.(type) {
	case float64:
		r.ParentGroupID = strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		r.ParentGroupID = t
	}

	return err
}

// IPRange represents the IP range whose traffic will be accepted by the
// security group.
type IPRange struct {
	CIDR string
}

// Group represents a group.
type Group struct {
	TenantID string `json:"tenant_id"`
	Name     string
}
