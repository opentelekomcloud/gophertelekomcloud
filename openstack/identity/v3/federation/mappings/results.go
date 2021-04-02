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
	User   User   `json:"user"`
	Group  Group  `json:"group"`
	Groups string `json:"groups"`
}

type User struct {
	Name string `json:"name"`
}

type Group struct {
	Name   string `json:"name"`
	Domain Domain `json:"domain"`
}

type Domain struct {
	Name string `json:"name"`
	ID   string `json:"id"`
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

// NextPageURL extracts the next/previous/self links from the links section of the result.
func (r MappingPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
			Self     string `json:"self"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Links.Next, err
}

// ExtractMappings returns a slice of Mappings contained in a linked page of results.
func ExtractMappings(r pagination.Page) ([]Mapping, error) {
	var s struct {
		Mappings []Mapping `json:"mappings"`
	}
	err := (r.(MappingPage)).ExtractInto(&s)
	return s.Mappings, err
}

// Extract interprets any group results as a Mapping.
func (r mappingResult) Extract() (*Mapping, error) {
	var s struct {
		Mapping *Mapping `json:"mapping"`
	}
	err := r.ExtractInto(&s)
	return s.Mapping, err
}
