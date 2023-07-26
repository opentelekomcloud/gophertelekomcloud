package ruletypes

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// The result of listing the qos rule types
type RuleType struct {
	Type string `json:"type"`
}

type ListRuleTypesPage struct {
	pagination.SinglePageBase
}

func (r ListRuleTypesPage) IsEmpty() (bool, error) {
	v, err := ExtractRuleTypes(r)
	return len(v) == 0, err
}

func ExtractRuleTypes(r pagination.Page) ([]RuleType, error) {
	var s struct {
		RuleTypes []RuleType `json:"rule_types"`
	}

	err := extract.Into(bytes.NewReader((r.(ListRuleTypesPage)).Body), &s)
	return s.RuleTypes, err
}
