package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/csbs/v1/policies"
)

func TestPoliciesList(t *testing.T) {
	client, err := clients.NewCsbsV1Client()
	if err != nil {
		t.Fatalf("Unable to create a CSBSv1 client: %s", err)
	}

	backupPolicies, err := policies.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to get list of backup policies: %s", err)
	}

	policiesExtracted, err := policies.ExtractBackupPolicies(backupPolicies)
	if err != nil {
		t.Fatalf("Unable to extract policies: %s", err)
	}

	for _, policy := range policiesExtracted {
		tools.PrintResource(t, policy)
	}
}
