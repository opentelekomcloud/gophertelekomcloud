package addons

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

// MetaData required to create an addon
type MetaData struct {
	// Addon unique name
	Name string `json:"name"`
	// Addon unique Id
	Id string `json:"uid"`
	// Addon tag, key/value pair format
	Labels map[string]string `json:"lables"`
	// Addon annotation, key/value pair format
	Annotations map[string]string `json:"annotaions"`
}

// Spec to create an addon
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
	AddonTemplateLables []string `json:"addonTemplateLables,omitempty"`
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
}

type SupportVersion struct {
	// Cluster type that supports the add-on template
	ClusterType string `json:"clusterType"`
	// Cluster versions that support the add-on template,
	// the parameter value is a regular expression
	ClusterVersion []string `json:"clusterVersion"`
}

type Version struct {
	// Add-on version
	Version string `json:"version"`
	// Add-on installation parameters
	Input map[string]interface{} `json:"input"`
	// Whether the add-on version is a stable release
	Stable bool `json:"stable"`
	// Cluster versions that support the add-on template
	SupportVersions []SupportVersion `json:"supportVersions"`
	// Creation time of the add-on instance
	CreationTimestamp string `json:"creationTimestamp"`
	// Time when the add-on instance was updated
	UpdateTimestamp string `json:"updateTimestamp"`
}

type AddonSpec struct {
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

type AddonTemplate struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata of an Addon
	Metadata MetaData `json:"metadata" required:"true"`
	// Specifications of an Addon
	Spec AddonSpec `json:"spec" required:"true"`
}

type AddonTemplateList struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Add-on template list
	Items []AddonTemplate `json:"items" required:"true"`
}

type InstanceMetadata struct {
	ID                string            `json:"uid"`
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	UpdateTimestamp   string            `json:"updateTimestamp"`
	CreationTimestamp string            `json:"creationTimestamp"`
}

type AddonInstanceSpec struct {
	ClusterID      string                 `json:"clusterID"`
	Version        string                 `json:"version"`
	TemplateName   string                 `json:"addonTemplateName"`
	TemplateType   string                 `json:"addonTemplateType"`
	TemplateLabels []string               `json:"addonTemplateLabels"`
	Descrition     string                 `json:"descrition"`
	Values         map[string]interface{} `json:"values"`
}

type Versions struct {
	Version           string                 `json:"version"`
	Input             map[string]interface{} `json:"input"`
	Stable            bool                   `json:"stable"`
	Translate         map[string]interface{} `json:"translate"`
	UpdateTimestamp   string                 `json:"updateTimestamp"`
	CreationTimestamp string                 `json:"creationTimestamp"`
}

type InstanceStatus struct {
	Status         string   `json:"status"`
	Reason         string   `json:"Reason"`
	Message        string   `json:"message"`
	TargetVersions []string `json:"targetVersions"`
	CurrentVersion Versions `json:"currentVersion"`
}

type AddonInstance struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata of an Addon
	Metadata InstanceMetadata `json:"metadata" required:"true"`
	// Specifications of an Addon
	Spec AddonInstanceSpec `json:"spec" required:"true"`
	// Status of an Addon
	Status InstanceStatus `json:"status"`
}

type AddonInstanceList struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata - Basic information about the add-on. A collection of attributes.
	Metadata string `json:"metadata"`
	// Add-on template list
	Items []AddonInstance `json:"items" required:"true"`
}
