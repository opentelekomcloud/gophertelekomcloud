package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rms/advanced"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const Expression = "SELECT id, name FROM resources WHERE provider = 'ecs' AND type = 'cloudservers' AND properties.status = 'SHUTOFF'"
const ExpressionUpdated = "SELECT * FROM resources WHERE provider = 'ecs'"

func TestSchemasList(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	schemas, err := advanced.ListSchemas(client, advanced.ListSchemasOpts{
		DomainId: client.DomainID,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(schemas) > 0)
}

func TestAdvancedQueriesLifecycle(t *testing.T) {
	t.Skip("You are not authorized with rms:resources:list impossible to run within CI")
	client, err := clients.NewRMSClient()
	th.AssertNoErr(t, err)

	queryName := tools.RandomString("rule-", 4)
	addOpts := advanced.CreateQueryOpts{
		DomainId:    client.DomainID,
		Name:        queryName,
		Description: "test-description",
		Expression:  Expression,
	}

	// Test CreateQuery
	query, err := advanced.CreateQuery(client, addOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, query.Name, addOpts.Name)
	th.AssertEquals(t, query.Expression, addOpts.Expression)
	th.AssertEquals(t, query.Description, addOpts.Description)

	t.Cleanup(func() {
		th.AssertNoErr(t, advanced.DeleteQuery(client, client.DomainID, query.Id))
	})

	// Test GetQuery
	retrievedQuery, err := advanced.GetQuery(client, client.DomainID, query.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, retrievedQuery.Name, query.Name)
	th.AssertEquals(t, retrievedQuery.Description, query.Description)
	th.AssertEquals(t, retrievedQuery.Expression, query.Expression)

	// Test UpdateQuery
	updateOpts := advanced.UpdateQueryOpts{
		QueryId:     query.Id,
		DomainId:    client.DomainID,
		Name:        addOpts.Name + "-updated",
		Description: addOpts.Description + "-updated",
		Expression:  ExpressionUpdated,
	}

	updatedQuery, err := advanced.UpdateQuery(client, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updatedQuery.Name, updateOpts.Name)
	th.AssertEquals(t, updatedQuery.Description, updateOpts.Description)
	th.AssertEquals(t, updatedQuery.Expression, updateOpts.Expression)

	// Test RunQuery
	runQuery, err := advanced.RunQuery(client, advanced.RunQueryOpts{
		DomainId:   client.DomainID,
		Expression: ExpressionUpdated,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(runQuery.Results) > 0)

	// Test ListQueries
	listQueryOpts := advanced.ListQueriesOpts{
		DomainId: client.DomainID,
	}

	queries, err := advanced.ListQueries(client, listQueryOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(queries) > 0)
	th.AssertEquals(t, queries[0], *updatedQuery)
}
