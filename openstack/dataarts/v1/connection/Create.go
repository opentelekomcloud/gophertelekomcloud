package connection

import (
	"encoding/json"
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

const connectionsEndpoint = "connections"

const (
	TypeDWS        = "DWS"
	TypeDLI        = "DLI"
	TypeSparkSQL   = "SparkSQL"
	TypeHive       = "HIVE"
	TypeRDS        = "RDS"
	TypeCloudTable = "CloudTable"
	TypeHOST       = "HOST"
)

const (
	MethodAgent  = "agent"
	MethodDirect = "direct"
)

const HeaderWorkspace = "workspace"

type Connection struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// Connection name. The name contains a maximum of 100 characters, including only letters, numbers, hyphens (-), and underscores (_).
	// The connection name must be unique.
	Name string `json:"name" required:"true"`
	// Connection type. Should be one of: DWS, DLI, SparkSQL, HIVE, RDS, CloudTable, HOST.
	Type string `json:"type" required:"true"`
	// Config connection configuration. The configuration item varies with the connection type.
	// You do not need to set the config parameter for DLI connections.
	// For other types of connections, see the description of connection configuration items.
	Config interface{} `json:"config,omitempty"`
	// Description of the connection. The description contains a maximum of 255 characters.
	Description string `json:"description,omitempty"`
}

func (c *Connection) UnmarshalJSON(data []byte) error {
	var obj map[string]json.RawMessage

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	n, ok := obj["name"]
	if !ok {
		return fmt.Errorf("missed field in json, field: name")
	}

	if err := json.Unmarshal(n, &c.Name); err != nil {
		return err
	}

	t, ok := obj["type"]
	if !ok {
		return fmt.Errorf("missed field in json, field: type")
	}

	if err := json.Unmarshal(t, &c.Type); err != nil {
		return err
	}

	if d, ok := obj["description"]; ok {
		if err := json.Unmarshal(d, &c.Description); err != nil {
			return err
		}
	}

	switch c.Type {
	case TypeDWS:
		d := &DWSConfig{}
		if err := json.Unmarshal(obj["config"], d); err != nil {
			return err
		}
		c.Config = *d
	case TypeDLI:
		// There aren't any configurations for DLI type.
		return nil
	case TypeSparkSQL:
		s := &SparkConfig{}
		if err := json.Unmarshal(obj["config"], s); err != nil {
			return err
		}
		c.Config = *s
	case TypeHive:
		h := &HiveConfig{}
		if err := json.Unmarshal(obj["config"], h); err != nil {
			return err
		}
		c.Config = *h
	case TypeRDS:
		r := &RDSConfig{}
		if err := json.Unmarshal(obj["config"], r); err != nil {
			return err
		}
		c.Config = *r
	case TypeCloudTable:
		ct := &CloudTableConfig{}
		if err := json.Unmarshal(obj["config"], ct); err != nil {
			return err
		}
		c.Config = *ct
	case TypeHOST:
		h := &HOSTConfig{}
		if err := json.Unmarshal(obj["config"], h); err != nil {
			return err
		}
		c.Config = *h
	default:
		return fmt.Errorf("connection type is not supported")
	}

	return nil
}

type DWSConfig struct {
	// Name of a DWS cluster.
	ClusterName string `json:"clusterName,omitempty"`
	// IP address for accessing the DWS cluster.
	IP string `json:"ip,omitempty"`
	// Port for accessing the DWS cluster.
	Port string `json:"port,omitempty"`
	// Username of the database. This username is the username entered during the creation of the DWS cluster.
	Username string `json:"userName" required:"true"`
	// Password for accessing the database. This password is the password entered during the creation of the DWS cluster.
	Password string `json:"password" required:"true"`
	// Specifies whether to enable the SSL connection.
	SSLEnable bool `json:"sslEnable" required:"true"`
	// Name of a KMS key.
	KMSKey string `json:"kmsKey" required:"true"`
	// Name of a CDM cluster.
	AgentName string `json:"agentName" required:"true"`
}

type SparkConfig struct {
	// Name of an MRS cluster.
	ClusterName string `json:"clusterName" required:"true"`
	// Method to connect.
	//   agent: connected through an agent.
	//   direct: connected directly.
	ConnectionMethod string `json:"connectionMethod" required:"true"`
	// Username of the MRS cluster. This parameter is mandatory when connectionMethod is set to agent.
	Username string `json:"userName,omitempty"`
	// Password for accessing the MRS cluster. This parameter is mandatory when connectionMethod is set to agent.
	Password string `json:"password,omitempty"`
	// Name of a CDM cluster. This parameter is mandatory when connectionMethod is set to agent.
	AgentName string `json:"agentName,omitempty"`
	// Name of a KMS key. This parameter is mandatory when connectionMethod is set to agent.
	KMSKey string `json:"kmsKey,omitempty"`
}

type HiveConfig struct {
	// Name of an MRS cluster.
	ClusterName string `json:"clusterName" required:"true"`
	// Method to connect.
	//   agent: connected through an agent.
	//   direct: connected directly.
	ConnectionMethod string `json:"connectionMethod" required:"true"`
	// Username of the MRS cluster. This parameter is mandatory when connectionMethod is set to agent.
	Username string `json:"userName,omitempty"`
	// Password for accessing the MRS cluster. This parameter is mandatory when connectionMethod is set to agent.
	Password string `json:"password,omitempty"`
	// Name of a CDM cluster. This parameter is mandatory when connectionMethod is set to agent.
	AgentName string `json:"agentName,omitempty"`
	// Name of a KMS key. This parameter is mandatory when connectionMethod is set to agent.
	KMSKey string `json:"kmsKey,omitempty"`
}

type RDSConfig struct {
	// Address for accessing RDS.
	IP string `json:"ip" required:"true"`
	// Port for accessing RDS.
	Port string `json:"port" required:"true"`
	// Username of the database. This username is the username entered during the creation of the cluster.
	Username string `json:"userName" required:"true"`
	// Password for accessing the database. This password is the password entered during the creation of the cluster.
	Password string `json:"password" required:"true"`
	// Name of a CDM cluster.
	AgentName string `json:"agentName" required:"true"`
	// Name of a KMS key.
	KMSKey string `json:"kmsKey" required:"true"`
	// Name of the driver.
	DriverName string `json:"driverName" required:"true"`
	// Path of the driver on OBS.
	DriverPath string `json:"driverPath" required:"true"`
}

type CloudTableConfig struct {
	// Name of a CloudTable cluster.
	ClusterName string `json:"clusterName" required:"true"`
}

type HOSTConfig struct {
	// IP address of the host.
	IP string `json:"ip" required:"true"`
	// SSH port number of the host.
	Port string `json:"port" required:"true"`
	// Username for logging in to the host.
	Username string `json:"userName" required:"true"`
	// Password for logging in to the host.
	Password string `json:"password" required:"true"`
	// Name of a CDM cluster.
	AgentName string `json:"agentName" required:"true"`
	// Name of a KMS key.
	KMSKey string `json:"kmsKey" required:"true"`
}

// Create is used to create a connection. The supported connection types include DWS, DLI, Spark SQL, RDS, CloudTable, and Hive.
// Send request  /v1/{project_id}/connections
func Create(client *golangsdk.ServiceClient, opts Connection) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	reqOpts := &golangsdk.RequestOpts{
		OkCodes: []int{204},
	}

	if opts.Workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: opts.Workspace}
	}

	_, err = client.Post(client.ServiceURL(connectionsEndpoint), b, nil, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
