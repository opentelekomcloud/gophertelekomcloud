package lifecyclehooks

// CreateLifecycleHookResponse represents the response parameter struct for lifecycle hook.
type LifecycleHook struct {
	// Specifies the lifecycle hook name.
	LifecycleHookName string `json:"lifecycle_hook_name"`
	// Specifies the lifecycle hook type. Values: INSTANCE_TERMINATING, INSTANCE_LAUNCHING
	LifecycleHookType string `json:"lifecycle_hook_type"`
	// Specifies the default lifecycle hook callback operation. Values: ABANDON, CONTINUE
	DefaultResult string `json:"default_result"`
	// Specifies the lifecycle hook timeout duration in seconds.
	DefaultTimeout int `json:"default_timeout"`
	// Specifies a unique topic in SMN for notification.
	NotificationTopicUrn string `json:"notification_topic_urn"`
	// Specifies the topic name in SMN.
	NotificationTopicName string `json:"notification_topic_name"`
	// Specifies the notification message.
	NotificationMetadata string `json:"notification_metadata"`
	// Specifies the UTC-compliant time when the lifecycle hook is created.
	CreateTime string `json:"create_time"`
}
