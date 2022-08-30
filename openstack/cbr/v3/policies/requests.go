package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// PolicyODCreate is policy operation definition
// see https://docs.otc.t-systems.com/en-us/api/cbr/CreatePolicy.html#CreatePolicy__request_PolicyoODCreate
type PolicyODCreate struct {
	DailyBackups          int    `json:"day_backups"`
	WeekBackups           int    `json:"week_backups"`
	YearBackups           int    `json:"year_backups"`
	MonthBackups          int    `json:"month_backups"`
	MaxBackups            int    `json:"max_backups,omitempty"`
	RetentionDurationDays int    `json:"retention_duration_days,omitempty"`
	Timezone              string `json:"timezone,omitempty"`
}

type TriggerProperties struct {
	// Pattern - Scheduling policy of the scheduler. Can't be empty.
	Pattern []string `json:"pattern"`
}

type Trigger struct {
	Properties TriggerProperties `json:"properties"`
}

// OperationType is a Policy type.
// One of `backup` and `replication`.
type OperationType string

type CreateOpts struct {
	// Name specifies the policy name. The value consists of 1 to 64 characters
	// and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name"`
	// OperationDefinition - scheduling configuration
	OperationDefinition *PolicyODCreate `json:"operation_definition"`
	// Enabled - whether to enable the policy, default: true
	Enabled *bool `json:"enabled,omitempty"`
	// OperationType - policy type
	OperationType OperationType `json:"operation_type"`
	// Trigger - time rule for the policy execution
	Trigger *Trigger `json:"trigger"`
}

func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "policy")
}

func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	reqBody, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, err = client.Post(listURL(client), reqBody, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	r.Err = err
	return
}

type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

type ListOpts struct {
	OperationType OperationType `q:"operation_type"`
	VaultID       string        `q:"vault_id"`
}

func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PolicyPage{pagination.SinglePageBase(r)}
	})
}

func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(singleURL(client, id), nil, nil)
	return
}

type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Enabled             *bool           `json:"enabled,omitempty"`
	Name                string          `json:"name,omitempty"`
	OperationDefinition *PolicyODCreate `json:"operation_definition,omitempty"`
	Trigger             *Trigger        `json:"trigger,omitempty"`
}

func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "policy")
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	reqBody, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, err = client.Put(singleURL(client, id), reqBody, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	r.Err = err
	return
}

func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	raw, err := client.Delete(singleURL(client, id), nil)
	return
}
