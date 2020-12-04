package v3

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/catalog"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

func TestGetCatalog(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}
	var endpoints []string
	err = catalog.List(client).EachPage(func(page pagination.Page) (bool, error) {
		entries, err := catalog.ExtractServiceCatalog(page)
		if err != nil {
			return false, err
		}
		for _, ent := range entries {
			for _, ep := range ent.Endpoints {
				endpoints = append(endpoints, ep.URL)
			}
		}
		return true, nil
	})
	require.NoError(t, err)
	require.True(t, len(endpoints) > 0)
}
