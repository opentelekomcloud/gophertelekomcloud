package v1

import (
	"log"
	"strconv"
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

	indexPrefix := clients.EnvOS.GetEnv("CSS_INDEX_PREFIX")
	if indexPrefix == "" {
		t.Skip("`OS_CSS_INDEX_PREFIX` must be defined")
	}
	keepDays := clients.EnvOS.GetEnv("CSS_KEEP_DAYS")
	if keepDays == "" {
		t.Skip("`OS_CSS_KEEP_DAYS` must be defined")
	}

	i, err := strconv.Atoi(keepDays)
	th.AssertNoErr(t, err)

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	got, err := logs.GetBaseLogConfiguration(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("CSS base log configuration:")

	tools.PrintResource(t, got)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	m, ok := got.(*logs.BaseLogConfiguration)
	if !ok {
		log.Fatal("An interface was expected!")
	}

	if m.LogSwitch {
		log.Println("The logging has been already enabled.")

	} else {

		basicOpts := logs.LogsOpts{
			Agency:   agency,
			Bucket:   bucketName,
			BasePath: "css/log",
		}

		err = logs.EnableBaseLogs(client, clusterID, &basicOpts)
		th.AssertNoErr(t, err)

		log.Println("Cluster logging enabled.")

		th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
	}

	if m.AutoEnable {
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

	realTimeOpts := logs.LogsOpts{
		IndexPrefix:     indexPrefix,
		KeepDays:        i,
		TargetClusterId: clusterID,
	}

	err = logs.EnableRealTimeLogs(client, clusterID, &realTimeOpts)
	th.AssertNoErr(t, err)

	log.Println("Cluster log ingestion enabled.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	got, err = logs.GetRealTimeLogConfiguration(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("CSS real time log configuration:")

	tools.PrintResource(t, got)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	err = logs.DisableAutomaticBackups(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("Cluster automatic backup for CSS logging disabled.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	err = logs.DisableBaseLogs(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("Cluster logging disabled.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	err = logs.DisableRealTimeLogs(client, clusterID)
	th.AssertNoErr(t, err)

	log.Println("Cluster log ingestion disabled.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
}

func TestEnableLogSwitch(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

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

	logSwitchOpts := logs.LogsOpts{
		Agency:   agency,
		Bucket:   bucketName,
		BasePath: "css/log",
	}

	automaticLogBackupOpts := logs.EnableAutomaticBackupOpts{
		Period: "00:00 GMT+08:00",
	}

	err = logs.EnableBaseLogs(client, clusterID, &logSwitchOpts)
	th.AssertNoErr(t, err)

	err = logs.EnableAutomaticBackups(client, clusterID, automaticLogBackupOpts)
	th.AssertNoErr(t, err)

	err = logs.DisableAutomaticBackups(client, clusterID)
	th.AssertNoErr(t, err)

	err = logs.DisableBaseLogs(client, clusterID)
	th.AssertNoErr(t, err)

}
func TestEnableLogIngestion(t *testing.T) {
	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	indexPrefix := clients.EnvOS.GetEnv("CSS_INDEX_PREFIX")
	if indexPrefix == "" {
		t.Skip("`OS_CSS_INDEX_PREFIX` must be defined")
	}
	keepDays := clients.EnvOS.GetEnv("CSS_KEEP_DAYS")
	if keepDays == "" {
		t.Skip("`OS_CSS_KEEP_DAYS` must be defined")
	}

	i, err := strconv.Atoi(keepDays)
	th.AssertNoErr(t, err)

	logSwitchOpts := logs.LogsOpts{
		IndexPrefix:     indexPrefix,
		KeepDays:        i,
		TargetClusterId: clusterID,
	}

	err = logs.EnableRealTimeLogs(client, clusterID, &logSwitchOpts)
	th.AssertNoErr(t, err)

	err = logs.DisableRealTimeLogs(client, clusterID)
	th.AssertNoErr(t, err)

}

func TestGetCSSLoggingConfiguration(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	action := clients.EnvOS.GetEnv("CSS_LOG_CONFIGURATION")

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	got, err := logs.GetConfiguration(client, clusterID, &action)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, got)
}

func TestUpdateCSSLoggingBaseConfigurations(t *testing.T) {
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
		BasePath: "css/log/2",
	}

	err = logs.UpdateBaseLogs(client, clusterID, updatedOpts)
	th.AssertNoErr(t, err)
}

func TestUpdateCSSLoggingRealTimeConfigurations(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	indexPrefix := clients.EnvOS.GetEnv("CSS_INDEX_PREFIX")
	if indexPrefix == "" {
		t.Skip("`OS_CSS_INDEX_PREFIX` must be defined")
	}
	keepDays := clients.EnvOS.GetEnv("CSS_KEEP_DAYS")
	if keepDays == "" {
		t.Skip("`OS_CSS_KEEP_DAYS` must be defined")
	}

	i, err := strconv.Atoi(keepDays)
	th.AssertNoErr(t, err)

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	updatedOpts := logs.UpdateLogConfigurationOpts{
		IndexPrefix:     indexPrefix,
		KeepDays:        i,
		TargetClusterId: clusterID,
	}

	err = logs.UpdateRealTimeLogs(client, clusterID, updatedOpts)
	th.AssertNoErr(t, err)
}
