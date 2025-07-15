package secret

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteOpts struct {
	Namespace          string         `json:"-"`
	Name               string         `json:"-"`
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

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*SecretDeleteResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.DeleteWithBody(client.ServiceURL("namespaces", opts.Namespace, "secrets", opts.Name), b, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res SecretDeleteResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SecretDeleteResp struct {
	APIVersion string         `json:"apiVersion"`
	Code       int            `json:"code"`
	Details    *StatusDetails `json:"details"`
	Kind       string         `json:"kind"`
	Message    string         `json:"message"`
	Metadata   *ListMeta      `json:"metadata"`
	Reason     string         `json:"reason"`
	Status     string         `json:"status"`
}

type StatusDetails struct {
	Causes            []StatusCause `json:"causes"`
	Group             string        `json:"group"`
	Kind              string        `json:"kind"`
	Name              string        `json:"name"`
	RetryAfterSeconds int           `json:"retryAfterSeconds"`
	UID               string        `json:"uid"`
}

type StatusCause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type ListMeta struct {
	Continue           string `json:"continue"`
	RemainingItemCount int64  `json:"remainingItemCount"`
	ResourceVersion    string `json:"resourceVersion"`
	SelfLink           string `json:"selfLink"`
}
