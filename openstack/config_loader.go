package openstack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	defaultEnvVar = "envvars"
)

var (
	yamlSuffixes = []string{".yaml", ".yml"}
	jsonSuffixes = []string{".json"}

	configFiles = fileList("clouds")
	secureFiles = fileList("secure")
	vendorFiles = fileList("clouds-public")
)

func configSearchPath() []string {
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	userConfigDir, _ := filepath.Abs("~/.config/openstack")
	unixConfigDir, _ := filepath.Abs("/etc/openstack")
	return []string{
		cwd,
		userConfigDir,
		unixConfigDir,
	}
}

func fileList(name string) []string {
	paths := configSearchPath()
	var suffixes []string
	suffixes = append(suffixes, yamlSuffixes...)
	suffixes = append(suffixes, jsonSuffixes...)
	size := len(suffixes) * len(paths)
	files := make([]string, size, size)
	for _, path := range paths {
		for _, suffix := range suffixes {
			files = append(files, filepath.Join(path, name+suffix))
		}
	}
	return files
}

/*
AuthOptionsFromEnv fills out an identity.AuthOptions structure with the
settings found on the various OpenStack OS_* environment variables.

The following variables provide sources of truth: OS_AUTH_URL, OS_USERNAME,
OS_PASSWORD, OS_TENANT_ID, and OS_TENANT_NAME.

Of these, OS_USERNAME, OS_PASSWORD, and OS_AUTH_URL must have settings,
or an error will result.  OS_TENANT_ID, OS_TENANT_NAME, OS_PROJECT_ID, and
OS_PROJECT_NAME are optional.

OS_TENANT_ID and OS_TENANT_NAME are mutually exclusive to OS_PROJECT_ID and
OS_PROJECT_NAME. If OS_PROJECT_ID and OS_PROJECT_NAME are set, they will
still be referred as "tenant" in Gophercloud.

To use this function, first set the OS_* environment variables (for example,
by sourcing an `openrc` file), then:

	opts, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(opts)
*/
func AuthOptionsFromEnv(envs ...*env) (golangsdk.AuthOptions, error) {
	e := NewEnv("")
	if len(envs) > 0 {
		e = envs[0]
	}

	authURL := e.GetEnv("AUTH_URL")
	token := e.GetEnv("TOKEN", "TOKEN_ID")
	username := e.GetEnv("USERNAME")
	userID := e.GetEnv("USERID", "USER_ID")
	password := e.GetEnv("PASSWORD")
	tenantID := e.GetEnv("PROJECT_ID", "TENANT_ID", "OS_AGENCY_DOMAIN_ID")
	tenantName := e.GetEnv("PROJECT_NAME", "TENANT_NAME", "OS_AGENCY_DOMAIN_NAME")
	domainID := e.GetEnv("DOMAIN_ID")
	domainName := e.GetEnv("DOMAIN_NAME")

	access := noEnv.GetEnv("ACCESS_KEY", "AWS_ACCESS_KEY_ID")
	secret := noEnv.GetEnv("SECRET_KEY", "AWS_SECRET_ACCESS_KEY")

	agencyName := e.GetEnv("AGENCY", "AGENCY_NAME")

	ao := golangsdk.AuthOptions{
		IdentityEndpoint: authURL,
		Username:         username,
		UserID:           userID,
		Password:         password,
		TokenID:          token,
		Scope: &golangsdk.AuthScope{
			ProjectID:   tenantID,
			ProjectName: tenantName,
			DomainID:    domainID,
			DomainName:  domainName,
		},
		AgencyName: agencyName,
		AKSKAuthOptions: &golangsdk.AKSKAuthOptions{
			Access: access,
			Secret: secret,
		},
	}
	return ao, nil
}

// This is helper for env-prefixed loading
type Env interface {
	CloudFromEnv() *Cloud
	GetEnv(keys ...string) string
	GetPrefix() string
	LoadOpenstackConfig() (*Config, error)
}

type env struct {
	Prefix string
}

func NewEnv(prefix string) Env {
	if !strings.HasSuffix(prefix, "_") {
		prefix += "_"
	}
	return &env{Prefix: prefix}
}

var noEnv = NewEnv("")

func (e *env) GetPrefix() string {
	return e.Prefix
}

func (e *env) CloudFromEnv() *Cloud {
	authOpts, _ := AuthOptionsFromEnv(e)
	verify := true
	if v := e.GetEnv("INSECURE"); v != "" {
		verify = v != "1" && v != "true"
	}
	cloud := &Cloud{
		Cloud:   e.GetEnv("CLOUD"),
		Profile: e.GetEnv("PROFILE"),
		AuthInfo: AuthInfo{
			AuthURL:           authOpts.IdentityEndpoint,
			Token:             authOpts.TokenID,
			Username:          authOpts.Username,
			UserID:            authOpts.UserID,
			Password:          authOpts.Password,
			ProjectName:       authOpts.TenantName,
			ProjectID:         authOpts.TenantID,
			UserDomainName:    e.GetEnv("USER_DOMAIN_NAME"),
			UserDomainID:      e.GetEnv("USER_DOMAIN_ID"),
			ProjectDomainName: e.GetEnv("PROJECT_DOMAIN_NAME"),
			ProjectDomainID:   e.GetEnv("PROJECT_DOMAIN_ID"),
			DomainName:        authOpts.DomainName,
			DomainID:          authOpts.DomainID,
			DefaultDomain:     e.GetEnv("DEFAULT_DOMAIN"),
			AccessKey:         authOpts.AKSKAuthOptions.Access,
			SecretKey:         authOpts.AKSKAuthOptions.Secret,
		},
		AuthType:           AuthType(e.GetEnv("AUTH_TYPE")),
		RegionName:         e.GetEnv("REGION_NAME", "REGION_ID"),
		EndpointType:       e.GetEnv("ENDPOINT_TYPE"),
		Interface:          e.GetEnv("INTERFACE"),
		IdentityAPIVersion: e.GetEnv("IDENTITY_API_VERSION"),
		VolumeAPIVersion:   e.GetEnv("VOLUME_API_VERSION"),
		Verify:             &verify,
		CACertFile:         e.GetEnv("CA_CERT", "CA_CERT_FILE"),
		ClientCertFile:     e.GetEnv("CLIENT_CERT", "CLIENT_CERT_FILE"),
		ClientKeyFile:      e.GetEnv("CLIENT_KEY", "CLIENT_KEY_FILE"),
	}
	return cloud
}

// GetEnv returns first non-empty value of given environment variables
func (e *env) GetEnv(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(e.Prefix + key); value != "" {
			return value
		}
	}
	return ""
}

// VendorConfig represents a collection of PublicCloud entries in clouds-public.yaml file.
// The format of the clouds-public.yml is documented at
// https://docs.openstack.org/python-openstackclient/latest/configuration/
type VendorConfig struct {
	Clouds map[string]Cloud `yaml:"public-clouds" json:"public-clouds"`
}

// Config represents a collection of Cloud entries in a clouds.yaml file.
// The format of clouds.yaml is documented at
// https://docs.openstack.org/os-client-config/latest/user/configuration.html.
type Config struct {
	DefaultCloud string           `yaml:"-" json:"-"`
	Clouds       map[string]Cloud `yaml:"clouds" json:"clouds"`
}

// AuthType represents a valid method of authentication.
type AuthType string

// AuthInfo represents the auth section of a cloud entry or
// auth options entered explicitly in ClientOpts.
type AuthInfo struct {
	// AuthURL is the keystone/identity endpoint URL.
	AuthURL string `yaml:"auth_url,omitempty" json:"auth_url,omitempty"`

	// Token is a pre-generated authentication token.
	Token string `yaml:"token,omitempty" json:"token,omitempty"`

	// Username is the username of the user.
	Username string `yaml:"username,omitempty" json:"username,omitempty"`

	// UserID is the unique ID of a user.
	UserID string `yaml:"user_id,omitempty" json:"user_id,omitempty"`

	// Password is the password of the user.
	Password string `yaml:"password,omitempty" json:"password,omitempty"`

	// ProjectName is the common/human-readable name of a project.
	// Users can be scoped to a project.
	// ProjectName on its own is not enough to ensure a unique scope. It must
	// also be combined with either a ProjectDomainName or ProjectDomainID.
	// ProjectName cannot be combined with ProjectID in a scope.
	ProjectName string `yaml:"project_name,omitempty" json:"project_name,omitempty"`

	// ProjectID is the unique ID of a project.
	// It can be used to scope a user to a specific project.
	ProjectID string `yaml:"project_id,omitempty" json:"project_id,omitempty"`

	// UserDomainName is the name of the domain where a user resides.
	// It is used to identify the source domain of a user.
	UserDomainName string `yaml:"user_domain_name,omitempty" json:"user_domain_name,omitempty"`

	// UserDomainID is the unique ID of the domain where a user resides.
	// It is used to identify the source domain of a user.
	UserDomainID string `yaml:"user_domain_id,omitempty" json:"user_domain_id,omitempty"`

	// ProjectDomainName is the name of the domain where a project resides.
	// It is used to identify the source domain of a project.
	// ProjectDomainName can be used in addition to a ProjectName when scoping
	// a user to a specific project.
	ProjectDomainName string `yaml:"project_domain_name,omitempty" json:"project_domain_name,omitempty"`

	// ProjectDomainID is the name of the domain where a project resides.
	// It is used to identify the source domain of a project.
	// ProjectDomainID can be used in addition to a ProjectName when scoping
	// a user to a specific project.
	ProjectDomainID string `yaml:"project_domain_id,omitempty" json:"project_domain_id,omitempty"`

	// DomainName is the name of a domain which can be used to identify the
	// source domain of either a user or a project.
	// If UserDomainName and ProjectDomainName are not specified, then DomainName
	// is used as a default choice.
	// It can also be used be used to specify a domain-only scope.
	DomainName string `yaml:"domain_name,omitempty" json:"domain_name,omitempty"`

	// DomainID is the unique ID of a domain which can be used to identify the
	// source domain of either a user or a project.
	// If UserDomainID and ProjectDomainID are not specified, then DomainID is
	// used as a default choice.
	// It can also be used be used to specify a domain-only scope.
	DomainID string `yaml:"domain_id,omitempty" json:"domain_id,omitempty"`

	// DefaultDomain is the domain ID to fall back on if no other domain has
	// been specified and a domain is required for scope.
	DefaultDomain string `yaml:"default_domain,omitempty" json:"default_domain,omitempty"`

	// AK/SK auth means
	AccessKey string `yaml:"ak,omitempty" json:"access_key,omitempty"`
	SecretKey string `yaml:"sk,omitempty" json:"secret_key,omitempty"`
}

// Cloud represents an entry in a clouds.yaml/public-clouds.yaml/secure.yaml file.
type Cloud struct {
	Cloud      string        `yaml:"cloud,omitempty" json:"cloud,omitempty"`
	Profile    string        `yaml:"profile,omitempty" json:"profile,omitempty"`
	AuthInfo   AuthInfo      `yaml:"auth,omitempty" json:"auth,omitempty"`
	AuthType   AuthType      `yaml:"auth_type,omitempty" json:"auth_type,omitempty"`
	RegionName string        `yaml:"region_name,omitempty" json:"region_name,omitempty"`
	Regions    []interface{} `yaml:"regions,omitempty" json:"regions,omitempty"`

	// EndpointType and Interface both specify whether to use the public, internal,
	// or admin interface of a service. They should be considered synonymous, but
	// EndpointType will take precedence when both are specified.
	EndpointType string `yaml:"endpoint_type,omitempty" json:"endpoint_type,omitempty"`
	Interface    string `yaml:"interface,omitempty" json:"interface,omitempty"`

	// API Version overrides.
	IdentityAPIVersion string `yaml:"identity_api_version,omitempty" json:"identity_api_version,omitempty"`
	VolumeAPIVersion   string `yaml:"volume_api_version,omitempty" json:"volume_api_version,omitempty"`

	// Verify whether or not SSL API requests should be verified.
	Verify *bool `yaml:"verify,omitempty" json:"verify,omitempty"`

	// CACertFile a path to a CA Cert bundle that can be used as part of
	// verifying SSL API requests.
	CACertFile string `yaml:"cacert,omitempty" json:"cacert,omitempty"`

	// ClientCertFile a path to a client certificate to use as part of the SSL
	// transaction.
	ClientCertFile string `yaml:"cert,omitempty" json:"cert,omitempty"`

	// ClientKeyFile a path to a client key to use as part of the SSL
	// transaction.
	ClientKeyFile string `yaml:"key,omitempty" json:"key,omitempty"`
}

func loadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func loadCloudFile(path string) (*Config, error) {
	data, err := loadFile(path)
	if err != nil {
		return nil, err
	}
	clouds := new(Config)
	if err := json.Unmarshal(data, clouds); err != nil {
		return nil, err
	}
	return clouds, err
}

func loadVendorFile(path string) (*VendorConfig, error) {
	data, err := loadFile(path)
	if err != nil {
		return nil, err
	}
	clouds := new(VendorConfig)
	if err := json.Unmarshal(data, clouds); err != nil {
		return nil, err
	}
	return clouds, err
}

func mergeCloudConfigs(config, override *Config) (*Config, error) {
	resultClouds := new(Config)
	for k, cfg := range config.Clouds {
		if over, ok := override.Clouds[k]; ok {
			cld, err := mergeClouds(cfg, over)
			if err != nil {
				return nil, err
			}
			resultClouds.Clouds[k] = *cld
		}
	}
	return resultClouds, nil
}

func selectExisting(files []string) string {
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			return file
		}
	}
	return ""
}

// mergeClouds merges two Config recursively (the AuthInfo also gets merged).
// In case both Config define a value, the value in the 'override' cloud takes precedence
func mergeClouds(cloud, override interface{}) (*Cloud, error) {
	overrideJson, err := json.Marshal(override)
	if err != nil {
		return nil, err
	}
	cloudJson, err := json.Marshal(cloud)
	if err != nil {
		return nil, err
	}
	var overrideInterface interface{}
	err = json.Unmarshal(overrideJson, &overrideInterface)
	if err != nil {
		return nil, err
	}
	var cloudInterface interface{}
	err = json.Unmarshal(cloudJson, &cloudInterface)
	if err != nil {
		return nil, err
	}
	var mergedCloud Cloud
	mergedInterface := mergeInterfaces(overrideInterface, cloudInterface)
	mergedJson, err := json.Marshal(mergedInterface)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mergedJson, &mergedCloud); err != nil {
		return nil, err
	}
	return &mergedCloud, nil
}

// merges two interfaces. In cases where a value is defined for both 'overridingInterface' and
// 'inferiorInterface' the value in 'overridingInterface' will take precedence.
func mergeInterfaces(overridingInterface, inferiorInterface interface{}) interface{} {
	switch overriding := overridingInterface.(type) {
	case map[string]interface{}:
		interfaceMap, ok := inferiorInterface.(map[string]interface{})
		if !ok {
			return overriding
		}
		for k, v := range interfaceMap {
			if overridingValue, ok := overriding[k]; ok {
				overriding[k] = mergeInterfaces(overridingValue, v)
			} else {
				overriding[k] = v
			}
		}
	case []interface{}:
		list, ok := inferiorInterface.([]interface{})
		if !ok {
			return overriding
		}
		for i := range list {
			overriding = append(overriding, list[i])
		}
		return overriding
	case nil:
		// mergeClouds(nil, map[string]interface{...}) -> map[string]interface{...}
		v, ok := inferiorInterface.(map[string]interface{})
		if ok {
			return v
		}
	}
	// We don't want to override with empty values
	if reflect.DeepEqual(overridingInterface, nil) || reflect.DeepEqual(reflect.Zero(reflect.TypeOf(overridingInterface)).Interface(), overridingInterface) {
		return inferiorInterface
	} else {
		return overridingInterface
	}
}

func prepend(item string, slice []string) []string {
	newSize := len(slice) + 1
	result := make([]string, newSize, newSize)
	result[0] = item
	for i, v := range slice {
		result[i+1] = v
	}
	return result
}

func mergeWithSecure(cloudConfig *Config, securePath string) *Config {
	s, err := loadCloudFile(securePath)
	if err != nil {
		log.Printf("Failed to load %s as secure config", securePath)
		return cloudConfig
	}
	cc, err := mergeCloudConfigs(cloudConfig, s)
	if err != nil {
		log.Printf("Failed to merge %s into cloud config", securePath)
		return cloudConfig
	}
	return cc
}

func mergeWithVendors(cloudConfig *Config, vendorPath string) *Config {
	v, err := loadVendorFile(vendorPath)
	if err != nil {
		log.Printf("Failed to load %s as vendor config", vendorPath)
		return cloudConfig
	}
	cc, err := mergeCloudConfigs(cloudConfig, &Config{Clouds: v.Clouds})
	if err != nil {
		log.Printf("Failed to merge %s into vendor config", vendorPath)
		return cloudConfig
	}
	return cc
}

// LoadCloudConfig utilize all existing cloud configurations to create cloud configuration:
// env variables, clouds.yaml, secure.yaml, clouds-public.yaml
func (e *env) LoadOpenstackConfig() (*Config, error) {
	var configs, secure, vendors []string
	copy(configs, configFiles)
	copy(secure, secureFiles)
	copy(vendors, vendorFiles)

	// find config files
	if c := e.GetEnv("CLIENT_CONFIG_FILE"); c != "" {
		configs = prepend(c, configs)
	}
	configPath := selectExisting(configFiles)
	if s := e.GetEnv("CLIENT_SECURE_FILE"); s != "" {
		secure = prepend(s, secure)
	}
	securePath := selectExisting(secureFiles)
	if v := e.GetEnv("CLIENT_VENDOR_FILE"); v != "" {
		vendors = prepend(v, vendors)
	}
	vendorPath := selectExisting(vendors)

	cloudConfig := new(Config)

	// load clouds.yaml
	if configPath != "" {
		c, err := loadCloudFile(configPath)
		if err != nil {
			log.Printf("Failed to load %s as cloud config", securePath)
		}
		if c != nil {
			cloudConfig = c
		}
	}

	// merge with secure.yaml
	if securePath != "" {
		cloudConfig = mergeWithSecure(cloudConfig, securePath)
	}

	// append cloud from envvars
	envVarKey := e.GetEnv("CLOUD_NAME")
	if envVarKey == "" {
		envVarKey = defaultEnvVar
	}
	if _, ok := cloudConfig.Clouds[envVarKey]; ok {
		return nil, fmt.Errorf("%sCLOUD_NAME=`%s` duplicates cloud defined in file", e.Prefix, envVarKey)
	}
	cloudConfig.Clouds[envVarKey] = *NewEnv(envVarKey).CloudFromEnv()

	cloudName := e.GetEnv("CLOUD")
	if cloudName == "" && len(cloudConfig.Clouds) == 1 {
		for k, _ := range cloudConfig.Clouds {
			cloudName = k
		}
	}
	cloudConfig.DefaultCloud = cloudName

	// merge with clouds-public.yaml
	if vendorPath != "" {
		cloudConfig = mergeWithVendors(cloudConfig, vendorPath)
	}
	return cloudConfig, nil
}
