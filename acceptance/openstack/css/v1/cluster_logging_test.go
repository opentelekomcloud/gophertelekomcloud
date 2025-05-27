package v1

import (
	"log"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/logs"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCSSLoggingFullLifecycle(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	agency := clients.EnvOS.GetEnv("AGENCY_NAME")
	if agency == "" {
		t.Skipf("OS_AGENCY_NAME is required for this test")
	}
	bucketName := clients.EnvOS.GetEnv("BUCKET_NAME")
	if bucketName == "" {
		t.Skipf("OS_BUCKET_NAME is required for this test")
	}
	period := clients.EnvOS.GetEnv("LOG_STARTING_PERIOD")
	if period == "" {
		t.Skip("`OS_LOG_STARTING_PERIOD` must be defined")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	got, err := logs.GetConfiguration(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("CSS log configuration:")

	tools.PrintResource(t, got)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	if got.LogSwitch {
		log.Println("The logging has been already enabled.")

	} else {

		basicOpts := logs.EnableLogsOpts{
			Agency:   agency,
			Bucket:   bucketName,
			BasePath: "css/log",
		}

		err = logs.EnableLogs(client, clusterID, basicOpts)
		th.AssertNoErr(t, err)

		log.Println("Cluster logging enabled.")

		th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
	}

	if got.AutoEnable {
		log.Println("Cluster automatic backup for CSS logging has been already enabled.")
	} else {

		opts := logs.EnableAutomaticBackupOpts{
			Period: period,
		}

		err = logs.EnableAutomaticBackups(client, clusterID, opts)
		th.AssertNoErr(t, err)
		log.Println("Cluster automatic backup for CSS logging enabled.")

		th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
	}

	err = logs.DisableAutomaticBackups(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("Cluster automatic backup for CSS logging disabled.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	err = logs.DisableLogs(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("Cluster logging disabled.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestGetCSSLoggingConfiguration(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	got, err := logs.GetConfiguration(client, clusterID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, got)
}

func TestUpdateCSSLoggingConfigurations(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	agency := clients.EnvOS.GetEnv("AGENCY_NAME")
	if agency == "" {
		t.Skipf("OS_AGENCY_NAME is required for this test")
	}
	bucketName := clients.EnvOS.GetEnv("BUCKET_NAME")
	if bucketName == "" {
		t.Skipf("OS_BUCKET_NAME is required for this test")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	updatedOpts := logs.UpdateLogConfigurationOpts{
		Agency:   agency,
		Bucket:   bucketName,
		BasePath: "css/log",
	}

	err = logs.UpdateLogs(client, clusterID, updatedOpts)
	th.AssertNoErr(t, err)
}
