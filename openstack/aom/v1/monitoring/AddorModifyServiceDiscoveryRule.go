package monitoring

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type AddOrModifyServiceDiscoveryRuleOpts struct {
	// Service parameters.
	AppRules []AppRule `json:"appRules,omitempty"`
}

// AppRule represents a single service discovery rule.
type AppRule struct {
	// Creation time.
	// When creating a service discovery rule, leave this parameter blank.
	// When modifying a service discovery rule, enter the returned createTime.
	// No omitempty: the API expects the field present (empty on creation).
	CreateTime string `json:"createTime"`
	// Whether a rule is enabled. Values: true and false.
	// Pointer so that a legitimate `false` value is marshaled and passes the
	// required check (ValidateTags rejects zero values).
	Enable *bool `json:"enable" required:"true"`
	// Rule event name. For service discovery, the fixed value is aom_inventory_rules_event.
	EventName string `json:"eventName" required:"true"`
	// Host ID. Currently, this parameter is not used and can be left blank.
	// No omitempty: the doc example always sends it ("hostid": []).
	Hostid []string `json:"hostid"`
	// Rule ID.
	// When creating a service discovery rule, leave this parameter blank.
	// When modifying a service discovery rule, enter a rule ID.
	// No required tag: the value is legitimately empty on rule creation.
	// No omitempty: the API expects the field present (empty on creation).
	Id string `json:"id"`
	// Rule name, which contains a maximum of 64 characters.
	// It must start with a lowercase letter but cannot end with a hyphen (-).
	// Only digits, lowercase letters, and hyphens are allowed.
	Name string `json:"name" required:"true"`
	// Project ID obtained from IAM. Generally, a project ID contains 32 characters.
	Projectid string `json:"projectid" required:"true"`
	// Rule details.
	Spec AppRulesSpec `json:"spec" required:"true"`
	// Custom description.
	Desc string `json:"desc,omitempty"`
}

// AppRulesSpec represents the details of a service discovery rule.
type AppRulesSpec struct {
	// Service type, which is used only for rule classification and UI display.
	// You can enter any field. For example, enter Java or Python by technology stack,
	// or enter collector or database by function.
	// No required tag: the doc example sends an empty value.
	AppType string `json:"appType,omitempty"`
	// Attribute list. Currently, this parameter is not used and can be left blank.
	// Values: cmdLine and env.
	AttrList []string `json:"attrList,omitempty"`
	// Whether to enable log collection. Values: true and false.
	DetectLog string `json:"detectLog" required:"true"`
	// Discovery rule. If the array contains multiple conditions, only the processes
	// that meet all the conditions will be matched.
	// If the value of checkType is cmdLine, set the value of checkMode to contain.
	// checkContent is in the format of ["xxx"], indicating that the process must contain
	// the xxx parameter.
	// If the value of checkType is env, set the value of checkMode to contain.
	// checkContent is in the format of ["k1","v1"], indicating that the process must contain
	// the environment variable whose name is k1 and value is v1.
	// If the value of checkType is scope, set the value of checkMode to equals.
	// checkContent is in the format of ["hostId1","hostId2"], indicating that the rule takes
	// effect only on specified nodes. If no nodes are specified, the rule applies to all
	// nodes of the project.
	DiscoveryRule []DiscoveryRule `json:"discoveryRule" required:"true"`
	// Whether the rule is the default one. Values: true and false.
	IsDefaultRule string `json:"isDefaultRule" required:"true"`
	// Whether the scenario is a pre-check scenario. No rules will be saved in the pre-check
	// scenario. This scenario is designed only to check whether a rule can detect node
	// processes before it is delivered. Values: true and false.
	IsDetect string `json:"isDetect" required:"true"`
	// Log file suffix. Values: log, trace, and out.
	LogFileFix []string `json:"logFileFix" required:"true"`
	// Log path configuration rule. If cmdLineHash is a fixed string, logs in the specified
	// log path or log file are collected. Otherwise, only the files whose names end with
	// .log or .trace are collected.
	// If the value of nameType is cmdLineHash, args is in the format of ["00001"] and value
	// is in the format of ["/xxx/xx.log"], indicating that the log path is /xxx/xx.log when
	// the startup command is 00001.
	LogPathRule []LogPathRule `json:"logPathRule,omitempty"`
	// Naming rules for discovered services and applications.
	NameRule NameRule `json:"nameRule" required:"true"`
	// Rule priority. Value range: 1 to 9999. Default value: 9999.
	Priority int `json:"priority" required:"true"`
	// Data source.
	DataSource string `json:"dataSource,omitempty"`
	// Whether to support editing. Options: true and false.
	Editable string `json:"editable,omitempty"`
	// Metric configuration.
	AomMetricRelabelConfigs interface{} `json:"aom_metric_relabel_configs,omitempty"`
}

// DiscoveryRule represents a single process match condition.
type DiscoveryRule struct {
	// Matched value.
	CheckContent []string `json:"checkContent" required:"true"`
	// Match condition. Values: contain and equals.
	CheckMode string `json:"checkMode" required:"true"`
	// Match type. Values: cmdLine, env, and scope.
	CheckType string `json:"checkType" required:"true"`
}

// LogPathRule represents a log path configuration rule.
type LogPathRule struct {
	// Command.
	Args []string `json:"args" required:"true"`
	// Value type, which can be cmdLineHash.
	NameType string `json:"nameType" required:"true"`
	// Log path.
	Value []string `json:"value" required:"true"`
}

// NameRule represents the naming rules for discovered services and applications.
type NameRule struct {
	// Service name rule. If there are multiple objects in the array, the character strings
	// extracted from these objects constitute the service name.
	// If the value of nameType is cmdLine, args is in the format of ["start", "end"],
	// indicating that the characters between start and end in the command are extracted.
	// If the value of nameType is env, args is in the format of ["aa"], indicating that
	// the environment variable named aa is extracted.
	// If the value of nameType is str, args is in the format of ["fix"], indicating that
	// the service name is suffixed with fix.
	// If the value of nameType is cmdLineHash, args is in the format of ["0001"] and value
	// is in the format of ["ser"], indicating that the service name is ser when the startup
	// command is 0001.
	AppNameRule []AppNameRule `json:"appNameRule" required:"true"`
	// Application name rule. The value formats are the same as for appNameRule.
	ApplicationNameRule []ApplicationNameRule `json:"applicationNameRule" required:"true"`
}

// AppNameRule represents a service name extraction rule.
type AppNameRule struct {
	// Value type. Values: cmdLineHash, cmdLine, env, and str.
	NameType string `json:"nameType" required:"true"`
	// Input value.
	Args []string `json:"args" required:"true"`
	// Service name, which is mandatory only when the value of nameType is cmdLineHash.
	Value []string `json:"value,omitempty"`
}

// ApplicationNameRule represents an application name extraction rule.
type ApplicationNameRule struct {
	// Value type. Values: cmdLineHash, cmdLine, env, and str.
	NameType string `json:"nameType" required:"true"`
	// Input value.
	Args []string `json:"args" required:"true"`
	// Service name, which is mandatory only when the value of nameType is cmdLineHash.
	Value []string `json:"value,omitempty"`
}

// AddOrModifyServiceDiscoveryRule is used to add or modify one or more service discovery
// rules. A maximum of 100 rules can be added to a project.
// When adding a rule, leave Id and CreateTime blank; when modifying a rule, the rule is
// identified by Id with the returned CreateTime in the request body.
func AddOrModifyServiceDiscoveryRule(client *golangsdk.ServiceClient, opts AddOrModifyServiceDiscoveryRuleOpts) (*ServiceDiscoveryRuleOperationResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/inv/servicediscoveryrules
	raw, err := client.Put(client.ServiceURL("inv", "servicediscoveryrules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ServiceDiscoveryRuleOperationResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// ServiceDiscoveryRuleOperationResp represents the response of the service
// discovery rule add/modify and delete operations.
type ServiceDiscoveryRuleOperationResp struct {
	// Response code.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// Response status code (no longer used).
	ResponseStatus int `json:"responseStatus"`
	// Service discovery rule ID list. This parameter is used during cross-AZ
	// configuration synchronization.
	Id []string `json:"id"`
	// Details about service discovery rules, e.g. {"name": ..., "id": ...} per rule.
	Results []map[string]string `json:"results"`
}
