package v1

import (
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v1/schemas"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDDMSchemasLifecycle(t *testing.T) {
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	subnetID := clients.EnvOS.GetEnv("NETWORK_ID")
	secGroupId := clients.EnvOS.GetEnv("SECURITY_GROUP")
	if subnetID == "" || vpcID == "" || secGroupId == "" {
		t.Skip("OS_NETWORK_ID or OS_VPC_ID or OS_SECURITY_GROUP env vars are missing but are required for DDM instances test")
	}

	// CREATE DDM CLIENT
	client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)
	// CREATE RDS CLIENT
	rdsclient, err := clients.NewRdsV3()
	th.AssertNoErr(t, err)

	cc, err := clients.CloudAndClient()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := CreateDDMInstance(t, client)
	// CREATE RDS INSTANCE
	// RDS INSTANCE MUST BE MYSQL 5.7, 8.0 WITH LowerCaseTableNames SET TO 1
	rdsInstance := CreateRDS(t, rdsclient, cc.RegionName)

	// CLEANUP
	t.Cleanup(func() {
		DeleteRDS(t, rdsclient, rdsInstance.Id)
		DeleteDDMInstance(t, client, ddmInstance.Id)
	})

	// CREATE SCHEMA
	schemaName := "ddm_schema_acc_test_1"
	usedRds := schemas.DatabaseInstancesParam{
		Id:            rdsInstance.Id,
		AdminUser:     "root",
		AdminPassword: "acc-test-password1!",
	}
	createDatabaseDetail := schemas.CreateDatabaseDetail{
		Name:        schemaName,
		ShardMode:   "cluster",
		ShardNumber: 8,
		ShardUnit:   8,
		UsedRds:     []schemas.DatabaseInstancesParam{usedRds},
	}
	createSchemaOpts := schemas.CreateSchemaOpts{
		Databases: []schemas.CreateDatabaseDetail{createDatabaseDetail},
	}
	t.Logf("Creating DDM Schema: %s", schemaName)
	_, err = schemas.CreateSchema(client, ddmInstance.Id, createSchemaOpts)
	th.AssertNoErr(t, err)
	err = golangsdk.WaitFor(1200, func() (bool, error) {
		schemaDetails, err := schemas.QuerySchemaDetails(client, ddmInstance.Id, schemaName)
		if err != nil {
			return false, err
		}
		time.Sleep(5 * time.Second)
		t.Logf("Schema status: %s", schemaDetails.Database.Status)
		if schemaDetails.Database.Status == "RUNNING" {
			th.AssertEquals(t, schemaDetails.Database.Name, schemaName)
			return true, nil
		}
		return false, nil
	})
	th.AssertNoErr(t, err)
	t.Logf("Created DDM Schema: %s", schemaName)

	// QUERY  SCHEMAS
	t.Logf("Listing DDM Schemas: %s", schemaName)
	listSchemaOpts := schemas.QuerySchemasOpts{}
	ddmSchemaList, err := schemas.QuerySchemas(client, ddmInstance.Id, listSchemaOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ddmSchemaList[0].Name, schemaName)
	t.Logf("DDM Schemas:\n %#v", ddmSchemaList)

	// DELETE SCHEMA
	t.Logf("Deleting DDM Schema: %s", schemaName)
	err = golangsdk.WaitFor(1000, func() (bool, error) {
		_, err := schemas.DeleteSchema(client, ddmInstance.Id, schemaName, true)
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	th.AssertNoErr(t, err)
	t.Logf("Deleted DDM Schema: %s", schemaName)
}

func TestDDMQueryAvailableDbInstances(t *testing.T) {
	// CREATE CLIENT
	client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)

	ddmInstance := CreateDDMInstance(t, client)
	t.Cleanup(func() {
		DeleteDDMInstance(t, client, ddmInstance.Id)
	})

	queryAvailableDbOpts := schemas.QueryAvailableDbOpts{}
	_, err = schemas.QueryAvailableDb(client, ddmInstance.Id, queryAvailableDbOpts)
	th.AssertNoErr(t, err)
}
