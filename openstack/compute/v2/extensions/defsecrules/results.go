package defsecrules

import (
	"encoding/json"
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/extensions/secgroups"
)

// DefaultRule represents a rule belonging to the "default" security group.
// It is identical to an openstack/compute/v2/extensions/secgroups.Rule.
type DefaultRule secgroups.Rule

func (r *DefaultRule) UnmarshalJSON(b []byte) error {
	var s secgroups.Rule
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = DefaultRule(s)
	return nil
}

func extra(err error, raw *http.Response) (*DefaultRule, error) {
	if err != nil {
		return nil, err
	}

	var res DefaultRule
	err = extract.IntoStructPtr(raw.Body, &res, "security_group_default_rule")
	return &res, err
}
