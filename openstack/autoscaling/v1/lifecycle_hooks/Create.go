package lifecyclehooks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts represents the request parameter struct for lifecycle hook.
type CreateOpts struct {
	// Specifies the lifecycle hook name. The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 32 characters.
	LifecycleHookName string `json:"lifecycle_hook_name" required:"true"`
	// Specifies the lifecycle hook type. Options:
	//   - INSTANCE_TERMINATING: The hook suspends the instance when it is terminated.
	//   - INSTANCE_LAUNCHING: The hook suspends the instance when it is started.
	LifecycleHookType string `json:"lifecycle_hook_type" required:"true"`
	// Specifies the default lifecycle hook callback operation. This operation is performed when the timeout duration expires.
	// Options: ABANDON(Default), CONTINUE
	// ABANDON:
	// If an instance is starting, ABANDON indicates that your customized operations failed, and the instance will be terminated.
	// In such a case, the scaling action fails, and you must create a new instance.
	// If an instance is stopping, ABANDON allows instance termination BUT stops other lifecycle hooks.
	// CONTINUE:
	// If an instance is starting, CONTINUE indicates that your customized operations are successful and the instance can be used.
	// If an instance is stopping, CONTINUE allows instance termination AND the completion of other lifecycle hooks.
	DefaultResult string `json:"default_result,omitempty"`
	// Specifies the lifecycle hook timeout duration, which ranges from 60 to 86400 seconds. The default value is 3600.
	DefaultTimeout int `json:"default_timeout,omitempty"`
	// Specifies a unique topic in SMN. This parameter specifies a notification object for a lifecycle hook.
	// When an instance is suspended by the lifecycle hook, the SMN service sends a notification to the object.
	// This notification contains the basic instance information, your customized notification content, and the token for controlling lifecycle operations.
	NotificationTopicUrn string `json:"notification_topic_urn" required:"true"`
	// Specifies a customized notification, which contains no more than 256 characters. The message cannot contain the following characters: <>&'(){}.
	// After a notification object is configured, the SMN service sends your customized notification to the object.
	NotificationMetadata string `json:"notification_metadata,omitempty"`
}

// This function is used to create a lifecycle hook for an AS group. Up to five lifecycle hooks can be created for one AS group.
// After the creation, when the AS group is resized, the lifecycle hook suspends the involved instance and puts it to a Wait (Adding to AS group) or Wait (Removing from AS group) status.
// This status is retained until the timeout duration ends or you manually perform a callback.
// During the instance waiting duration, you can perform customized operations.
// For example, you can install or configure software on a newly started instance, or download the log file from the instance before the instance terminates.
func Create(client *golangsdk.ServiceClient, opts CreateOpts, asGroupId string) (*LifecycleHook, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	// POST /autoscaling-api/v1/{project_id}/scaling_lifecycle_hook/{scaling_group_id}
	raw, err := client.Post(client.ServiceURL("scaling_lifecycle_hook", asGroupId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LifecycleHook
	err = extract.Into(raw.Body, &res)
	return &res, err
}
