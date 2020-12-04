package v3

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/catalog"
)

func TestGetCatalog(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}
	allPages, err := catalog.List(client).AllPages()
	require.NoError(t, err)
	allServices, err := catalog.ExtractServiceCatalog(allPages)
	require.True(t, len(allServices) > 0)
}
