package mappings

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// Mapping helps manage mapping rules.
type Mapping struct {
	// ID is the unique ID of the mapping.
	ID string `json:"id"`

	// Resource Links of mappings.
	Links map[string]interface{} `json:"links"`

	// Rules used to map federated users to local users
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Local  []LocalRule  `json:"local"`
	Remote []RemoteRule `json:"remote"`
}

type LocalRule struct {
	User   *UserOpts  `json:"user"`
	Group  *GroupOpts `json:"group"`
	Groups string     `json:"groups"`
}

type RemoteRule struct {
	Type     string   `json:"type"`
	NotAnyOf []string `json:"not_any_of"`
	AnyOneOf []string `json:"any_one_of"`
	Regex    bool     `json:"regex"`
}

type mappingResult struct {
	golangsdk.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Mapping.
type GetResult struct {
	mappingResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Mapping.
type CreateResult struct {
	mappingResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Mapping.
type UpdateResult struct {
	mappingResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// MappingPage is a single page of Mapping results.
type MappingPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Mappings contains any results.
func (r MappingPage) IsEmpty() (bool, error) {
	mappings, err := ExtractMappings(r)
	return len(mappings) == 0, err
}

// ExtractMappings returns a slice of Mappings contained in a linked page of results.
func ExtractMappings(r pagination.Page) ([]Mapping, error) {
	var s []Mapping
	err := (r.(MappingPage)).ExtractIntoSlicePtr(&s, "mappings")
	return s, err
}

// Extract interprets any group results as a Mapping.
func (r mappingResult) Extract() (*Mapping, error) {
	s := new(Mapping)
	err := r.ExtractIntoStructPtr(s, "mapping")
	return s, err
}
