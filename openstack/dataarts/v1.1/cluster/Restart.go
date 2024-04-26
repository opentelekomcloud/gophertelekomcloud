package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

const (
	RestartImmediately = "IMMEDIATELY"
	RestartForcely     = "FORCELY"
	RestartSoftly      = "SOFTLY"
)

type RestartOpts struct {
	// Restart Cluster restart. For details about how to define the cluster to restart, see RestartStruct.
	Restart RestartStruct `json:"restart" required:"true"`
}

type RestartStruct struct {
	// RestartDelayTime Restart delay, in seconds.
	RestartDelayTime int `json:"restartDelayTime,omitempty"`
	// Restart mode
	// IMMEDIATELY: immediate restart
	// FORCELY: forcible restart
	// SOFTLY: common restart
	// The default value is IMMEDIATELY. Forcibly restarting the service process will interrupt the service process and restart the VMs in the cluster.
	RestartMode string `json:"restartMode,omitempty"`
	// RestartLevel Restart level
	// SERVICE: service restart
	// VM: VM restart
	// The default value is SERVICE.
	RestartLevel string `json:"restartLevel,omitempty"`
	// Type of the cluster to be restarted. The value can be set to cdm only.
	Type string `json:"type,omitempty"`
	// Instance Reserved field. When restartLevel is set to SERVICE, this parameter is mandatory and an empty string should be entered.
	Instance string `json:"instance,omitempty"`
	// Reserved field. When restartLevel is set to SERVICE, this parameter is mandatory and an empty string should be entered.
	Group string `json:"group,omitempty"`
}

// Restart is used to restart a cluster.
// Send request POST /v1.1/{project_id}/clusters/{cluster_id}/action
func Restart(client *golangsdk.ServiceClient, clusterId string, opts RestartOpts) (*JobId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(
		client.ServiceURL(clustersEndpoint, clusterId, actionEndpoint),
		b,
		nil,
		&golangsdk.RequestOpts{
			MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
		},
	)
	if err != nil {
		return nil, err
	}

	return respToJobId(resp)
}
