package build

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dns/v2/zones"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/objectstorage/v1/accounts"
	"github.com/stretchr/testify/require"
)

// Backward compatibility test.
// Proving we can replace deprecated methods with new ones.
//
// Compare results of old methods with results of currently used methods.
//
// Those tests can be used as an example of replacing deprecated methods.

func TestCompareHeaderOpts(t *testing.T) {
	headerOpts := &accounts.UpdateOpts{
		ContentType:       "application/lolson",
		DetectContentType: true,
		TempURLKey:        "key1",
		TempURLKey2:       "key2",
	}

	expected, err := headerOpts.ToAccountUpdateMap()
	require.NoError(t, err)

	actual, err := Headers(headerOpts)
	require.NoError(t, err)

	require.EqualValues(t, expected, actual)
}

func TestCompareRequestBodyOpts(t *testing.T) {
	bodyOpts := &zones.CreateOpts{
		Email:       "mail@me.plz",
		Description: "email here",
		Name:        "this is a name",
		TTL:         1600,
		ZoneType:    "TYPE1",
		Router: &zones.RouterOpts{
			RouterID:     "1a23d5a8-1027-4b82-96a2-780b945fa294",
			RouterRegion: "nord",
		},
		Tags: []tags.ResourceTag{
			{
				Key:   "Foo",
				Value: "Bar",
			},
		},
	}

	expected, err := bodyOpts.ToZoneCreateMap()
	require.NoError(t, err)

	expectedJSON, err := extract.JsonMarshal(expected)
	require.NoError(t, err)

	actual, err := RequestBody(bodyOpts, "")
	require.NoError(t, err)

	actualJSON, err := extract.JsonMarshal(actual)
	require.NoError(t, err)

	require.JSONEq(t, string(expectedJSON), string(actualJSON))
}
