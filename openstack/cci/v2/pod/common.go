package pod

// Pod represents a pod in Kubernetes
type Pod struct {
	APIVersion string     `json:"apiVersion,omitempty"`
	Kind       string     `json:"kind,omitempty"`
	Metadata   ObjectMeta `json:"metadata,omitempty"`
	Spec       PodSpec    `json:"spec,omitempty"`
	Status     PodStatus  `json:"status,omitempty"`
}

// PodSpec represents the specification of a pod
type PodSpec struct {
	ActiveDeadlineSeconds         *int64                     `json:"activeDeadlineSeconds,omitempty"`
	Affinity                      *Affinity                  `json:"affinity,omitempty"`
	Containers                    []Container                `json:"containers"`
	DNSConfig                     *PodDNSConfig              `json:"dnsConfig,omitempty"`
	DNSPolicy                     string                     `json:"dnsPolicy,omitempty"`
	EphemeralContainers           []EphemeralContainer       `json:"ephemeralContainers,omitempty"`
	HostAliases                   []HostAlias                `json:"hostAliases,omitempty"`
	Hostname                      string                     `json:"hostname,omitempty"`
	ImagePullSecrets              []LocalObjectReference     `json:"imagePullSecrets,omitempty"`
	InitContainers                []Container                `json:"initContainers,omitempty"`
	NodeName                      string                     `json:"nodeName,omitempty"`
	Overhead                      map[string]string          `json:"overhead,omitempty"`
	ReadinessGates                []PodReadinessGate         `json:"readinessGates,omitempty"`
	RestartPolicy                 string                     `json:"restartPolicy,omitempty"`
	SchedulerName                 string                     `json:"schedulerName,omitempty"`
	SecurityContext               *PodSecurityContext        `json:"securityContext,omitempty"`
	SetHostnameAsFQDN             *bool                      `json:"setHostnameAsFQDN,omitempty"`
	ShareProcessNamespace         *bool                      `json:"shareProcessNamespace,omitempty"`
	TerminationGracePeriodSeconds *int64                     `json:"terminationGracePeriodSeconds,omitempty"`
	Tolerations                   []Toleration               `json:"tolerations,omitempty"`
	TopologySpreadConstraints     []TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
	Volumes                       []Volume                   `json:"volumes,omitempty"`
}

// Toleration represents a toleration that is applied to a pod to schedule it onto nodes with matching taints
type Toleration struct {
	// Effect indicates the taint effect to match. Empty means match all taint effects.
	Effect string `json:"effect,omitempty"`

	// Key is the taint key that the toleration applies to. Empty means match all taint keys.
	Key string `json:"key,omitempty"`

	// Operator represents a key's relationship to the value. Valid operators are Exists and Equal.
	Operator string `json:"operator,omitempty"`

	// TolerationSeconds represents the period of time the toleration tolerates the taint.
	TolerationSeconds *int64 `json:"tolerationSeconds,omitempty"`

	// Value is the taint value the toleration matches to.
	Value string `json:"value,omitempty"`
}

// Affinity represents pod affinity and anti-affinity
type Affinity struct {
	NodeAffinity    *NodeAffinity    `json:"nodeAffinity,omitempty"`
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

// NodeAffinity represents node affinity scheduling rules
type NodeAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution *NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NodeSelector represents a node selector
type NodeSelector struct {
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// NodeSelectorTerm contains a list of node selector requirements
type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty"`
}

// NodeSelectorRequirement represents a node selector requirement
type NodeSelectorRequirement struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

// PodAntiAffinity represents pod anti-affinity scheduling rules
type PodAntiAffinity struct {
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// WeightedPodAffinityTerm represents a weighted pod affinity term
type WeightedPodAffinityTerm struct {
	PodAffinityTerm PodAffinityTerm `json:"podAffinityTerm"`
	Weight          int32           `json:"weight"`
}

// PodAffinityTerm represents a pod affinity term
type PodAffinityTerm struct {
	LabelSelector *LabelSelector `json:"labelSelector,omitempty"`
	Namespaces    []string       `json:"namespaces,omitempty"`
	TopologyKey   string         `json:"topologyKey"`
}

// Container represents a container in a pod
type Container struct {
	Args                     []string             `json:"args,omitempty"`
	Command                  []string             `json:"command,omitempty"`
	Env                      []EnvVar             `json:"env,omitempty"`
	EnvFrom                  []EnvFromSource      `json:"envFrom,omitempty"`
	Image                    string               `json:"image,omitempty"`
	ImagePullPolicy          string               `json:"imagePullPolicy,omitempty"`
	Lifecycle                *Lifecycle           `json:"lifecycle,omitempty"`
	LivenessProbe            *Probe               `json:"livenessProbe,omitempty"`
	Name                     string               `json:"name"`
	Ports                    []ContainerPort      `json:"ports,omitempty"`
	ReadinessProbe           *Probe               `json:"readinessProbe,omitempty"`
	Resources                ResourceRequirements `json:"resources,omitempty"`
	SecurityContext          *SecurityContext     `json:"securityContext,omitempty"`
	StartupProbe             *Probe               `json:"startupProbe,omitempty"`
	Stdin                    bool                 `json:"stdin,omitempty"`
	StdinOnce                bool                 `json:"stdinOnce,omitempty"`
	TerminationMessagePath   string               `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy string               `json:"terminationMessagePolicy,omitempty"`
	TTY                      bool                 `json:"tty,omitempty"`
	VolumeMounts             []VolumeMount        `json:"volumeMounts,omitempty"`
	WorkingDir               string               `json:"workingDir,omitempty"`
}

// ContainerPort represents a port in a container running in a pod
type ContainerPort struct {
	ContainerPort int32  `json:"containerPort"`
	Name          string `json:"name,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

// EnvVar represents an environment variable
type EnvVar struct {
	Name      string        `json:"name"`
	Value     string        `json:"value,omitempty"`
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty"`
}

// EnvVarSource represents a source for an environment variable's value
type EnvVarSource struct {
	ConfigMapKeyRef  *ConfigMapKeySelector  `json:"configMapKeyRef,omitempty"`
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	SecretKeyRef     *SecretKeySelector     `json:"secretKeyRef,omitempty"`
}

// ConfigMapKeySelector selects a key from a ConfigMap
type ConfigMapKeySelector struct {
	Key      string `json:"key"`
	Name     string `json:"name,omitempty"`
	Optional *bool  `json:"optional,omitempty"`
}

// ObjectFieldSelector selects a field from the pod
type ObjectFieldSelector struct {
	APIVersion string `json:"apiVersion,omitempty"`
	FieldPath  string `json:"fieldPath"`
}

// ResourceFieldSelector selects a resource of the container
type ResourceFieldSelector struct {
	ContainerName string `json:"containerName,omitempty"`
	Divisor       string `json:"divisor,omitempty"`
	Resource      string `json:"resource"`
}

// SecretKeySelector selects a key from a Secret
type SecretKeySelector struct {
	Key      string `json:"key"`
	Name     string `json:"name,omitempty"`
	Optional *bool  `json:"optional,omitempty"`
}

// EnvFromSource represents a source to populate environment variables
type EnvFromSource struct {
	ConfigMapRef *ConfigMapEnvSource `json:"configMapRef,omitempty"`
	Prefix       string              `json:"prefix,omitempty"`
	SecretRef    *SecretEnvSource    `json:"secretRef,omitempty"`
}

// ConfigMapEnvSource selects a ConfigMap to populate environment variables
type ConfigMapEnvSource struct {
	Name     string `json:"name,omitempty"`
	Optional *bool  `json:"optional,omitempty"`
}

// SecretEnvSource selects a Secret to populate environment variables
type SecretEnvSource struct {
	Name     string `json:"name,omitempty"`
	Optional *bool  `json:"optional,omitempty"`
}

// Lifecycle represents container lifecycle hooks
type Lifecycle struct {
	PostStart *LifecycleHandler `json:"postStart,omitempty"`
	PreStop   *LifecycleHandler `json:"preStop,omitempty"`
}

// LifecycleHandler represents a lifecycle handler
type LifecycleHandler struct {
	Exec    *ExecAction    `json:"exec,omitempty"`
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty"`
}

// ExecAction represents an exec lifecycle action
type ExecAction struct {
	Command []string `json:"command,omitempty"`
}

// HTTPGetAction represents an HTTP GET lifecycle action
type HTTPGetAction struct {
	Host        string       `json:"host,omitempty"`
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`
	Path        string       `json:"path,omitempty"`
	Port        string       `json:"port"`
	Scheme      string       `json:"scheme,omitempty"`
}

// HTTPHeader represents an HTTP header
type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Probe represents a container probe
type Probe struct {
	Exec                          *ExecAction      `json:"exec,omitempty"`
	FailureThreshold              int32            `json:"failureThreshold,omitempty"`
	HTTPGet                       *HTTPGetAction   `json:"httpGet,omitempty"`
	InitialDelaySeconds           int32            `json:"initialDelaySeconds,omitempty"`
	PeriodSeconds                 int32            `json:"periodSeconds,omitempty"`
	SuccessThreshold              int32            `json:"successThreshold,omitempty"`
	TCPSocket                     *TCPSocketAction `json:"tcpSocket,omitempty"`
	TerminationGracePeriodSeconds *int64           `json:"terminationGracePeriodSeconds,omitempty"`
	TimeoutSeconds                int32            `json:"timeoutSeconds,omitempty"`
}

// TCPSocketAction represents a TCP socket action
type TCPSocketAction struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port"`
}

// ResourceRequirements represents compute resource requirements
type ResourceRequirements struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests,omitempty"`
}

// SecurityContext represents security options for a container
type SecurityContext struct {
	Capabilities           *Capabilities `json:"capabilities,omitempty"`
	ProcMount              string        `json:"procMount,omitempty"`
	ReadOnlyRootFilesystem *bool         `json:"readOnlyRootFilesystem,omitempty"`
	RunAsGroup             *int64        `json:"runAsGroup,omitempty"`
	RunAsNonRoot           *bool         `json:"runAsNonRoot,omitempty"`
	RunAsUser              *int64        `json:"runAsUser,omitempty"`
}

// Capabilities represents container capabilities
type Capabilities struct {
	Add  []string `json:"add,omitempty"`
	Drop []string `json:"drop,omitempty"`
}

// VolumeMount represents the mounting of a volume in a container
type VolumeMount struct {
	ExtendPathMode string `json:"extendPathMode,omitempty"`
	MountPath      string `json:"mountPath"`
	Name           string `json:"name"`
	ReadOnly       bool   `json:"readOnly,omitempty"`
	SubPath        string `json:"subPath,omitempty"`
	SubPathExpr    string `json:"subPathExpr,omitempty"`
}

// Volume represents a pod volume
type Volume struct {
	ConfigMap             *ConfigMapVolumeSource             `json:"configMap,omitempty"`
	DownwardAPI           *DownwardAPIVolumeSource           `json:"downwardAPI,omitempty"`
	EmptyDir              *EmptyDirVolumeSource              `json:"emptyDir,omitempty"`
	Ephemeral             *EphemeralVolumeSource             `json:"ephemeral,omitempty"`
	Name                  string                             `json:"name"`
	NFS                   *NFSVolumeSource                   `json:"nfs,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
	Projected             *ProjectedVolumeSource             `json:"projected,omitempty"`
	Secret                *SecretVolumeSource                `json:"secret,omitempty"`
}

// ObjectMeta represents metadata about the object
type ObjectMeta struct {
	Annotations                map[string]string    `json:"annotations,omitempty"`
	ClusterName                string               `json:"clusterName,omitempty"`
	CreationTimestamp          string               `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64               `json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp          string               `json:"deletionTimestamp,omitempty"`
	Enable                     *bool                `json:"enable,omitempty"`
	Finalizers                 []string             `json:"finalizers,omitempty"`
	GenerateName               string               `json:"generateName,omitempty"`
	Generation                 *int64               `json:"generation,omitempty"`
	Labels                     map[string]string    `json:"labels,omitempty"`
	ManagedFields              []ManagedFieldsEntry `json:"managedFields,omitempty"`
	Name                       string               `json:"name,omitempty"`
	Namespace                  string               `json:"namespace,omitempty"`
	OwnerReferences            []OwnerReference     `json:"ownerReferences,omitempty"`
	ResourceVersion            string               `json:"resourceVersion,omitempty"`
	SelfLink                   string               `json:"selfLink,omitempty"`
	UID                        string               `json:"uid,omitempty"`
}

// ManagedFieldsEntry contains information about the manager that created or modified a resource
type ManagedFieldsEntry struct {
	APIVersion  string      `json:"apiVersion,omitempty"`
	FieldsType  string      `json:"fieldsType,omitempty"`
	FieldsV1    interface{} `json:"fieldsV1,omitempty"`
	Manager     string      `json:"manager,omitempty"`
	Operation   string      `json:"operation,omitempty"`
	Subresource string      `json:"subresource,omitempty"`
	Time        string      `json:"time,omitempty"`
}

// OwnerReference represents a reference to an object's owner
type OwnerReference struct {
	APIVersion         string `json:"apiVersion"`
	BlockOwnerDeletion *bool  `json:"blockOwnerDeletion,omitempty"`
	Controller         *bool  `json:"controller,omitempty"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	UID                string `json:"uid"`
}

// PodStatus represents the current status of a pod
type PodStatus struct {
	Conditions                 []PodCondition    `json:"conditions,omitempty"`
	ContainerStatuses          []ContainerStatus `json:"containerStatuses,omitempty"`
	EphemeralContainerStatuses []ContainerStatus `json:"ephemeralContainerStatuses,omitempty"`
	HostIP                     string            `json:"hostIP,omitempty"`
	InitContainerStatuses      []ContainerStatus `json:"initContainerStatuses,omitempty"`
	Message                    string            `json:"message,omitempty"`
	NominatedNodeName          string            `json:"nominatedNodeName,omitempty"`
	Phase                      string            `json:"phase,omitempty"`
	PodIP                      string            `json:"podIP,omitempty"`
	PodIPs                     []PodIP           `json:"podIPs,omitempty"`
	QOSClass                   string            `json:"qosClass,omitempty"`
	Reason                     string            `json:"reason,omitempty"`
	StartTime                  string            `json:"startTime,omitempty"`
}

// PodCondition represents a pod condition
type PodCondition struct {
	LastProbeTime      string `json:"lastProbeTime,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	Message            string `json:"message,omitempty"`
	Reason             string `json:"reason,omitempty"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}

// ContainerStatus represents the status of a container
type ContainerStatus struct {
	ContainerID  string          `json:"containerID,omitempty"`
	Image        string          `json:"image"`
	ImageID      string          `json:"imageID"`
	LastState    *ContainerState `json:"lastState,omitempty"`
	Name         string          `json:"name"`
	Ready        bool            `json:"ready"`
	RestartCount int32           `json:"restartCount"`
	Started      *bool           `json:"started,omitempty"`
	State        *ContainerState `json:"state,omitempty"`
}

// ContainerState represents the state of a container
type ContainerState struct {
	Running    *ContainerStateRunning    `json:"running,omitempty"`
	Terminated *ContainerStateTerminated `json:"terminated,omitempty"`
	Waiting    *ContainerStateWaiting    `json:"waiting,omitempty"`
}

// ContainerStateRunning represents a running container state
type ContainerStateRunning struct {
	StartedAt string `json:"startedAt,omitempty"`
}

// ContainerStateTerminated represents a terminated container state
type ContainerStateTerminated struct {
	ContainerID string `json:"containerID,omitempty"`
	ExitCode    int32  `json:"exitCode"`
	FinishedAt  string `json:"finishedAt,omitempty"`
	Message     string `json:"message,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Signal      int32  `json:"signal,omitempty"`
	StartedAt   string `json:"startedAt,omitempty"`
}

// ContainerStateWaiting represents a waiting container state
type ContainerStateWaiting struct {
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// PodIP represents a pod IP address
type PodIP struct {
	IP string `json:"ip,omitempty"`
}

// PodDNSConfig represents DNS configurations
type PodDNSConfig struct {
	Nameservers []string             `json:"nameservers,omitempty"`
	Options     []PodDNSConfigOption `json:"options,omitempty"`
	Searches    []string             `json:"searches,omitempty"`
}

// PodDNSConfigOption represents a DNS configuration option
type PodDNSConfigOption struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// EphemeralContainer represents an ephemeral container
type EphemeralContainer struct {
	Args                     []string         `json:"args,omitempty"`
	Command                  []string         `json:"command,omitempty"`
	Env                      []EnvVar         `json:"env,omitempty"`
	EnvFrom                  []EnvFromSource  `json:"envFrom,omitempty"`
	Image                    string           `json:"image,omitempty"`
	Name                     string           `json:"name"`
	SecurityContext          *SecurityContext `json:"securityContext,omitempty"`
	Stdin                    bool             `json:"stdin,omitempty"`
	StdinOnce                bool             `json:"stdinOnce,omitempty"`
	TargetContainerName      string           `json:"targetContainerName,omitempty"`
	TerminationMessagePath   string           `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy string           `json:"terminationMessagePolicy,omitempty"`
	TTY                      bool             `json:"tty,omitempty"`
	VolumeMounts             []VolumeMount    `json:"volumeMounts,omitempty"`
	WorkingDir               string           `json:"workingDir,omitempty"`
}

// HostAlias represents a mapping between IP and hostnames
type HostAlias struct {
	Hostnames []string `json:"hostnames,omitempty"`
	IP        string   `json:"ip,omitempty"`
}

// LocalObjectReference contains enough information to locate the referenced object
type LocalObjectReference struct {
	Name string `json:"name,omitempty"`
}

// PodReadinessGate represents a readiness gate
type PodReadinessGate struct {
	ConditionType string `json:"conditionType"`
}

// PodSecurityContext holds pod-level security attributes
type PodSecurityContext struct {
	FSGroup             *int64   `json:"fsGroup,omitempty"`
	FSGroupChangePolicy string   `json:"fsGroupChangePolicy,omitempty"`
	RunAsGroup          *int64   `json:"runAsGroup,omitempty"`
	RunAsNonRoot        *bool    `json:"runAsNonRoot,omitempty"`
	RunAsUser           *int64   `json:"runAsUser,omitempty"`
	SupplementalGroups  []int64  `json:"supplementalGroups,omitempty"`
	Sysctls             []Sysctl `json:"sysctls,omitempty"`
}

// Sysctl represents a kernel parameter to be set
type Sysctl struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TopologySpreadConstraint represents how to spread pods
type TopologySpreadConstraint struct {
	LabelSelector      *LabelSelector `json:"labelSelector,omitempty"`
	MatchLabelKeys     []string       `json:"matchLabelKeys,omitempty"`
	MaxSkew            int32          `json:"maxSkew"`
	MinDomains         *int32         `json:"minDomains,omitempty"`
	NodeAffinityPolicy string         `json:"nodeAffinityPolicy,omitempty"`
	NodeTaintsPolicy   string         `json:"nodeTaintsPolicy,omitempty"`
	TopologyKey        string         `json:"topologyKey"`
	WhenUnsatisfiable  string         `json:"whenUnsatisfiable"`
}

// LabelSelector represents label selectors
type LabelSelector struct {
	MatchExpressions []LabelSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchLabels      map[string]string          `json:"matchLabels,omitempty"`
}

// LabelSelectorRequirement represents a requirement for label selection
type LabelSelectorRequirement struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

// ConfigMapVolumeSource adapts a ConfigMap into a volume
type ConfigMapVolumeSource struct {
	DefaultMode *int32      `json:"defaultMode,omitempty"`
	Items       []KeyToPath `json:"items,omitempty"`
	Name        string      `json:"name,omitempty"`
	Optional    *bool       `json:"optional,omitempty"`
}

// DownwardAPIVolumeSource represents a downward API volume source
type DownwardAPIVolumeSource struct {
	DefaultMode *int32                  `json:"defaultMode,omitempty"`
	Items       []DownwardAPIVolumeFile `json:"items,omitempty"`
}

// DownwardAPIVolumeFile represents information to create the file containing the pod field
type DownwardAPIVolumeFile struct {
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty"`
	Mode             *int32                 `json:"mode,omitempty"`
	Path             string                 `json:"path"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}

// EmptyDirVolumeSource represents an empty directory volume
type EmptyDirVolumeSource struct {
	Medium    string `json:"medium,omitempty"`
	SizeLimit string `json:"sizeLimit,omitempty"`
}

// EphemeralVolumeSource represents an ephemeral volume source
type EphemeralVolumeSource struct {
	VolumeClaimTemplate *PersistentVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty"`
}

// PersistentVolumeClaimTemplate represents a template for PVC creation
type PersistentVolumeClaimTemplate struct {
	Metadata ObjectMeta                `json:"metadata,omitempty"`
	Spec     PersistentVolumeClaimSpec `json:"spec"`
}

// PersistentVolumeClaimSpec represents the specification of a PVC
type PersistentVolumeClaimSpec struct {
	AccessModes      []string                   `json:"accessModes,omitempty"`
	DataSource       *TypedLocalObjectReference `json:"dataSource,omitempty"`
	DataSourceRef    *TypedLocalObjectReference `json:"dataSourceRef,omitempty"`
	Resources        ResourceRequirements       `json:"resources,omitempty"`
	Selector         *LabelSelector             `json:"selector,omitempty"`
	StorageClassName string                     `json:"storageClassName,omitempty"`
	VolumeMode       string                     `json:"volumeMode,omitempty"`
	VolumeName       string                     `json:"volumeName,omitempty"`
}

// TypedLocalObjectReference contains enough information to locate the referenced object
type TypedLocalObjectReference struct {
	APIGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
}

// NFSVolumeSource represents an NFS mount
type NFSVolumeSource struct {
	Path     string `json:"path"`
	ReadOnly bool   `json:"readOnly,omitempty"`
	Server   string `json:"server"`
}

// PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim
type PersistentVolumeClaimVolumeSource struct {
	ClaimName string `json:"claimName"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
}

// ProjectedVolumeSource represents a projected volume source
type ProjectedVolumeSource struct {
	DefaultMode *int32             `json:"defaultMode,omitempty"`
	Sources     []VolumeProjection `json:"sources,omitempty"`
}

// VolumeProjection represents a projected volume source
type VolumeProjection struct {
	ConfigMap   *ConfigMapProjection   `json:"configMap,omitempty"`
	DownwardAPI *DownwardAPIProjection `json:"downwardAPI,omitempty"`
	Secret      *SecretProjection      `json:"secret,omitempty"`
}

// ConfigMapProjection adapts a ConfigMap into a projected volume
type ConfigMapProjection struct {
	Items    []KeyToPath `json:"items,omitempty"`
	Name     string      `json:"name,omitempty"`
	Optional *bool       `json:"optional,omitempty"`
}

// DownwardAPIProjection represents downward API info for projecting
type DownwardAPIProjection struct {
	Items []DownwardAPIVolumeFile `json:"items,omitempty"`
}

// SecretProjection adapts a Secret into a projected volume
type SecretProjection struct {
	Items    []KeyToPath `json:"items,omitempty"`
	Name     string      `json:"name,omitempty"`
	Optional *bool       `json:"optional,omitempty"`
}

// SecretVolumeSource adapts a Secret into a volume
type SecretVolumeSource struct {
	DefaultMode *int32      `json:"defaultMode,omitempty"`
	Items       []KeyToPath `json:"items,omitempty"`
	Optional    *bool       `json:"optional,omitempty"`
	SecretName  string      `json:"secretName,omitempty"`
}

// KeyToPath maps a string key to a path within a volume
type KeyToPath struct {
	Key  string `json:"key"`
	Mode *int32 `json:"mode,omitempty"`
	Path string `json:"path"`
}
