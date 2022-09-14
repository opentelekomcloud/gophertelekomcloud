package policies

import (
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List returns a Pager which allows you to iterate over a collection of
// backup policies. It accepts a ListOpts struct, which allows you to
// filter the returned collection for greater efficiency.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]BackupPolicy, error) {
	query, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.NewPager(client, client.ServiceURL("policies")+query.String(),
		func(r pagination.PageResult) pagination.Page {
			return BackupPolicyPage{pagination.LinkedPageBase{PageResult: r}}
		}).AllPages()
	if err != nil {
		return nil, err
	}

	policies, err := ExtractBackupPolicies(pages)
	if err != nil {
		return nil, err
	}

	return filterPolicies(policies, opts)
}

func filterPolicies(policies []BackupPolicy, opts ListOpts) ([]BackupPolicy, error) {
	var refinedPolicies []BackupPolicy
	var matched bool
	m := map[string]interface{}{}

	if opts.ID != "" {
		m["ID"] = opts.ID
	}
	if opts.Status != "" {
		m["Status"] = opts.Status
	}
	if len(m) > 0 && len(policies) > 0 {
		for _, policy := range policies {
			matched = true

			for key, value := range m {
				if sVal := getStructPolicyField(&policy, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedPolicies = append(refinedPolicies, policy)
			}
		}
	} else {
		refinedPolicies = policies
	}
	return refinedPolicies, nil
}

func getStructPolicyField(v *BackupPolicy, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

// BackupPolicyPage is the page returned by a pager when traversing over a
// collection of backup policies.
type BackupPolicyPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of backup policies has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r BackupPolicyPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &res, "policies_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

// IsEmpty checks whether a BackupPolicyPage struct is empty.
func (r BackupPolicyPage) IsEmpty() (bool, error) {
	is, err := ExtractBackupPolicies(r)
	return len(is) == 0, err
}

// ExtractBackupPolicies accepts a Page struct, specifically a BackupPolicyPage struct,
// and extracts the elements into a slice of Policy structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractBackupPolicies(r pagination.Page) ([]BackupPolicy, error) {
	var res []BackupPolicy
	err := extract.IntoSlicePtr(r.(BackupPolicyPage).BodyReader(), &res, "policies")
	return res, err
}
