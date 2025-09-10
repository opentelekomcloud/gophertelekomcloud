package servicemesh

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// API version. The value is fixed at v1 and cannot be changed.
	APIVersion string `json:"apiVersion" required:"true"`
	// API type. The value is fixed at Mesh or mesh and cannot be changed.
	Kind string `json:"kind" required:"true"`
	// Basic information about the service mesh. Metadata is a collection of attributes.
	Metadata MeshMetadata `json:"metadata" required:"true"`
	// Detailed description of the service mesh. ASM creates or updates the service mesh by spec.
	Spec MeshSpec `json:"spec" required:"true"`
}

type MeshMetadata struct {
	// Service mesh name.
	// Enter 4 to 64 characters. The name must start with a lowercase letter and not end with a hyphen (-).
	// Only lowercase letters, digits, and hyphens (-) are allowed.
	Name string `json:"name" required:"true"`
}

type MeshSpec struct {
	// Service mesh type.
	// InCluster: service mesh with an in-cluster control plane.
	// The value is InCluster for the service mesh of the Basic edition.
	Type string `json:"type" required:"true"`
	// Service mesh version.
	Version string `json:"version" required:"true"`
	// Extensions of the service mesh.
	ExtendParams *MeshExtendParams `json:"extendParams" required:"true"`
	// Whether the service mesh supports IPv6.
	IPv6Enable bool `json:"ipv6Enable,omitempty"`
	// Service mesh configuration.
	Config *MeshConfig `json:"config,omitempty"`
}

type MeshExtendParams struct {
	// Cluster information in the service mesh.
	Clusters []MeshCluster `json:"clusters" required:"true"`
}

type MeshCluster struct {
	// Cluster ID, which is unique and can be used to query the cluster to be added.
	ClusterID string `json:"clusterID" required:"true"`
	// Sidecar injection configuration.
	Injection *InjectionConfig `json:"injection,omitempty"`
	// Installation configuration of service mesh components.
	Installation *InstallationConfig `json:"installation" required:"true"`
}

type InjectionConfig struct {
	// Namespaces where sidecars to be injected.
	Namespaces *Selector `json:"namespaces,omitempty"`
}

type InstallationConfig struct {
	// Nodes where service mesh components are installed.
	Nodes *Selector `json:"nodes" required:"true"`
}

type Selector struct {
	// Field selector.
	FieldSelector *FieldSelector `json:"fieldSelector" required:"true"`
}

type FieldSelector struct {
	// Key
	Key string `json:"key" required:"true"`
	// Operator. The value can only be In.
	Operator string `json:"operator" required:"true"`
	// Values
	Values []string `json:"values" required:"true"`
}

type MeshConfig struct {
	// Data plane configuration of the service mesh.
	ProxyConfig *ProxyConfig `json:"proxyConfig,omitempty"`
	// Observability configuration of the service mesh.
	TelemetryConfig *TelemetryConfig `json:"telemetryConfig,omitempty"`
}

type ProxyConfig struct {
	// IP address ranges that will be included for outbound traffic redirection. Use commas (,) to separate the IP address ranges.
	IncludeIPRanges string `json:"includeIPRanges,omitempty"`
	// IP address ranges that will be excluded for outbound traffic redirection. Use commas (,) to separate the IP address ranges.
	ExcludeIPRanges string `json:"excludeIPRanges,omitempty"`
	// Ports that will be excluded for outbound traffic redirection. Use commas (,) to separate the ports.
	ExcludeOutboundPorts string `json:"excludeOutboundPorts,omitempty"`
	// Ports that will be excluded for inbound traffic redirection. Use commas (,) to separate the ports.
	ExcludeInboundPorts string `json:"excludeInboundPorts,omitempty"`
	// Ports that will be included for outbound traffic redirection. Use commas (,) to separate the ports.
	IncludeOutboundPorts string `json:"includeOutboundPorts,omitempty"`
	// Ports that will be included for inbound traffic redirection. Use commas (,) to separate the ports.
	IncludeInboundPorts string `json:"includeInboundPorts,omitempty"`
}

type TelemetryConfig struct {
	// Tracing configuration, which is used to report traces in the service mesh.
	Tracing *Tracing `json:"tracing,omitempty"`
}

type Tracing struct {
	// Tracing sampling rate.
	RandomSamplingPercentage float64 `json:"randomSamplingPercentage,omitempty"`
	// Name of the default provider that tracing reports data to, which must match the name field in extensionProviders or use the preset provider apm-otel of ASM.
	// If apm-otel is used, ensure that APM 2.0 is supported in the current region and the service mesh version is later than 1.18.
	DefaultProviders []string `json:"defaultProviders,omitempty"`
	// User-defined provider. Currently, Zipkin is supported.
	// If you configure the Zipkin provider, ensure that the service mesh version is 1.15 or later.
	ExtensionProviders []TracingExtensionProvider `json:"extensionProviders,omitempty"`
}

type TracingExtensionProvider struct {
	// Provider name.
	Name string `json:"name,omitempty"`
	// Self-defined configuration of Zipkin.
	Zipkin ZipkinTracingProvider `json:"zipkin,omitempty"`
}

type ZipkinTracingProvider struct {
	// Service address of Zipkin.
	Service string `json:"service,omitempty"`
	// Service port of Zipkin.
	Port int `json:"port,omitempty"`
}

// This function is used to create a service mesh.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*ServiceMesh, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/meshes
	raw, err := client.Post(client.ServiceURL("meshes"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ServiceMesh
	return &res, extract.Into(raw.Body, &res)
}
