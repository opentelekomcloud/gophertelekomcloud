package addons

import (
	"fmt"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type Addon struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata of an Addon
	Metadata MetaData `json:"metadata" required:"true"`
	// Specifications of an Addon
	Spec Spec `json:"spec" required:"true"`
	// Status of an Addon
	Status Status `json:"status"`
}

// Metadata required to create an addon
type MetaData struct {
	// Addon unique name
	Name string `json:"name"`
	// Addon unique Id
	Id string `json:"uid"`
	// Addon tag, key/value pair format
	Labels map[string]string `json:"labels"`
	// Addon annotation, key/value pair format
	Annotations map[string]string `json:"annotaions"`
	// Time when the add-on instance was updated.
	UpdateTimestamp string `json:"updateTimestamp"`
	// Time when the add-on instance was created.
	CreationTimestamp string `json:"creationTimestamp"`
}

// Specifications to create an addon
type Spec struct {
	// For the addon version.
	Version string `json:"version" required:"true"`
	// Cluster ID.
	ClusterID string `json:"clusterID" required:"true"`
	// Addon Template Name.
	AddonTemplateName string `json:"addonTemplateName" required:"true"`
	// Addon Template Type.
	AddonTemplateType string `json:"addonTemplateType" required:"true"`
	// Addon Template Labels.
	AddonTemplateLabels []string `json:"addonTemplateLabels,omitempty"`
	// Addon Description.
	Description string `json:"description" required:"true"`
	// Addon Parameters
	Values Values `json:"values" required:"true"`
}

type Status struct {
	// The state of the addon
	Status string `json:"status"`
	// Reasons for the addon to become current
	Reason string `json:"reason"`
	// Error Message
	Message string `json:"message"`
	// The target versions of the addon
	TargetVersions []string `json:"targetVersions"`
	// Current add-on version.
	CurrentVersion Version `json:"currentVersion"`
}

type Values struct {
	Basic    map[string]interface{} `json:"basic" required:"true"`
	Advanced map[string]interface{} `json:"custom,omitempty"`
	Flavor   map[string]interface{} `json:"flavor,omitempty"`
}

type Version struct {
	// Add-on version
	Version string `json:"version"`
	// Add-on installation parameters
	Input Input `json:"input"`
	// Whether the add-on version is a stable release
	Stable bool `json:"stable"`
	// Translation information used by the GUI.
	Translate map[string]interface{} `json:"translate"`
	// Cluster versions that support the add-on template
	SupportVersions []SupportVersion `json:"supportVersions"`
	// Creation time of the add-on instance
	CreationTimestamp string `json:"creationTimestamp"`
	// Time when the add-on instance was updated
	UpdateTimestamp string `json:"updateTimestamp"`
}

type Input struct {
	Basic      map[string]interface{} `json:"basic"`
	Parameters map[string]interface{} `json:"parameters"`
}

type SupportVersion struct {
	// Cluster type that supports the add-on template
	ClusterType string `json:"clusterType"`
	// Cluster versions that support the add-on template,
	// the parameter value is a regular expression
	ClusterVersion []string `json:"clusterVersion"`
}

type ClusterIdQueryParam struct {
	ClusterId string `q:"cluster_id"`
}

type AddonTemplateList struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Add-on template list
	Items []AddonTemplate `json:"items" required:"true"`
}

type AddonTemplate struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata of an Addon
	Metadata MetaData `json:"metadata" required:"true"`
	// Specifications of an Addon
	Spec AddonTemplateSpec `json:"spec" required:"true"`
}

type AddonTemplateSpec struct {
	// Template type (helm or static).
	Type string `json:"type" required:"true"`
	// Whether the add-on is installed by default
	Require bool `json:"require" required:"true"`
	// Group to which the template belongs
	Labels []string `json:"labels" required:"true"`
	// URL of the logo image
	LogoURL string `json:"logoURL" required:"true"`
	// URL of the readme file
	ReadmeURL string `json:"readmeURL" required:"true"`
	// Template description
	Description string `json:"description" required:"true"`
	// Template version details
	Versions []Version `json:"versions" required:"true"`
}

func CCEServiceURL(client *golangsdk.ServiceClient, clusterID string, parts ...string) string {
	rbUrl := fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])
	return rbUrl + strings.Join(parts, "/")
}
