package testing

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	tokens3 "github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/tokens"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

var catalog3 = tokens3.ServiceCatalog{
	Entries: []tokens3.CatalogEntry{
		{
			Type: "same",
			Name: "same",
			Endpoints: []tokens3.Endpoint{
				{
					ID:        "1",
					Region:    "same",
					Interface: "public",
					URL:       "https://public.correct.com/",
				},
				{
					ID:        "2",
					Region:    "same",
					Interface: "admin",
					URL:       "https://admin.correct.com/",
				},
				{
					ID:        "3",
					Region:    "same",
					Interface: "internal",
					URL:       "https://internal.correct.com/",
				},
				{
					ID:        "4",
					Region:    "different",
					Interface: "public",
					URL:       "https://badregion.com/",
				},
			},
		},
		{
			Type: "same",
			Name: "different",
			Endpoints: []tokens3.Endpoint{
				{
					ID:        "5",
					Region:    "same",
					Interface: "public",
					URL:       "https://badname.com/",
				},
				{
					ID:        "6",
					Region:    "different",
					Interface: "public",
					URL:       "https://badname.com/+badregion",
				},
			},
		},
		{
			Type: "different",
			Name: "different",
			Endpoints: []tokens3.Endpoint{
				{
					ID:        "7",
					Region:    "same",
					Interface: "public",
					URL:       "https://badtype.com/+badname",
				},
				{
					ID:        "8",
					Region:    "different",
					Interface: "public",
					URL:       "https://badtype.com/+badregion+badname",
				},
			},
		},
	},
}

func TestV3EndpointExact(t *testing.T) {
	expectedURLs := map[golangsdk.Availability]string{
		golangsdk.AvailabilityPublic:   "https://public.correct.com/",
		golangsdk.AvailabilityAdmin:    "https://admin.correct.com/",
		golangsdk.AvailabilityInternal: "https://internal.correct.com/",
	}

	for availability, expected := range expectedURLs {
		actual, err := openstack.V3EndpointURL(&catalog3, golangsdk.EndpointOpts{
			Type:         "same",
			Name:         "same",
			Region:       "same",
			Availability: availability,
		})
		th.AssertNoErr(t, err)
		th.CheckEquals(t, expected, actual)
	}
}

func TestV3EndpointNone(t *testing.T) {
	_, actual := openstack.V3EndpointURL(&catalog3, golangsdk.EndpointOpts{
		Type:         "nope",
		Availability: golangsdk.AvailabilityPublic,
	})
	expected := &golangsdk.ErrEndpointNotFound{}
	th.CheckEquals(t, expected.Error(), actual.Error())
}

func TestV3EndpointMultiple(t *testing.T) {
	_, err := openstack.V3EndpointURL(&catalog3, golangsdk.EndpointOpts{
		Type:         "same",
		Region:       "same",
		Availability: golangsdk.AvailabilityPublic,
	})
	if !strings.HasPrefix(err.Error(), "Discovered 2 matching endpoints:") {
		t.Errorf("Received unexpected error: %v", err)
	}
}

func TestV3EndpointBadAvailability(t *testing.T) {
	_, err := openstack.V3EndpointURL(&catalog3, golangsdk.EndpointOpts{
		Type:         "same",
		Name:         "same",
		Region:       "same",
		Availability: "wat",
	})
	th.CheckEquals(t, "Unexpected availability in endpoint query: wat", err.Error())
}
