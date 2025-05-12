package network

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteOpts struct {
	Namespace string     `json:"-" required:"true"`
	Name      string     `json:"-" required:"true"`
	Body      DeleteBody `json:"-" required:"true"`
}

type DeleteBody struct {
	APIVersion         string         `json:"apiVersion,omitempty"`
	Kind               string         `json:"kind,omitempty"`
	GracePeriodSeconds *int64         `json:"gracePeriodSeconds,omitempty"`
	PropagationPolicy  string         `json:"propagationPolicy,omitempty"`
	DryRun             []string       `json:"dryRun,omitempty"`
	OrphanDependents   *bool          `json:"orphanDependents,omitempty"`
	Preconditions      *Preconditions `json:"preconditions,omitempty"`
}

type Preconditions struct {
	ResourceVersion string `json:"resourceVersion,omitempty"`
	UID             string `json:"uid,omitempty"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*DeleteResult, error) {
	url := client.ServiceURL("namespaces", opts.Namespace, "networks", opts.Name)

	b, err := build.RequestBody(opts.Body, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.DeleteWithBody(url, b, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DeleteResult struct {
	APIVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Status     NetworkStatus `json:"status"`
	Message    string        `json:"message"`
	Code       int           `json:"code"`
	Reason     string        `json:"reason"`
	Details    StatusDetails `json:"details"`
	Metadata   Metadata      `json:"metadata"`
}

type StatusDetails struct {
	Name              string        `json:"name"`
	Group             string        `json:"group"`
	Kind              string        `json:"kind"`
	Causes            []StatusCause `json:"causes"`
	RetryAfterSeconds int           `json:"retryAfterSeconds"`
	UID               string        `json:"uid"`
}

type StatusCause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type Metadata struct {
	Continue           string `json:"continue"`
	RemainingItemCount *int64 `json:"remainingItemCount"`
	ResourceVersion    string `json:"resourceVersion"`
	SelfLink           string `json:"selfLink"`
}
