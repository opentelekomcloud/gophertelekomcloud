package golangsdk

import (
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	type testCase struct {
		Name        string
		isFail      bool
		Endpoints   []string
		QueryParams interface{}
	}

	type network struct {
		Name      string `q:"display_name"`
		Status    string `q:"status"`
		NetworkID string `q:"network_id"`
	}

	type empty struct{}

	type withoutTags struct {
		Name string
	}

	cases := []testCase{
		{
			Name:        "positive case without query parameters",
			Endpoints:   []string{"servers"},
			QueryParams: nil,
		},
		{
			Name:      "positive case with query parameters",
			Endpoints: []string{"servers", "network"},
			QueryParams: &network{
				Name:      "network_name",
				Status:    "active",
				NetworkID: "7bf9e4a9-fab1-4415-a6f7-ff7284ee1870",
			},
		},
		{
			Name:        "positive case with empty struct, without query parameters",
			Endpoints:   []string{"servers", "network"},
			QueryParams: &empty{},
		},
		{
			Name:        "positive case with struct without query tags",
			Endpoints:   []string{"servers", "network"},
			QueryParams: &withoutTags{Name: "anyname"},
		},
		{
			Name:      "negative case with slash '/' in the endpoint name",
			isFail:    true,
			Endpoints: []string{"servers/"},
		},
		{
			Name:        "negative case with slash '/' in the second endpoint name",
			isFail:      true,
			Endpoints:   []string{"servers", "/asdf/"},
			QueryParams: nil,
		},
		{
			Name:        "negative case with characters ' /!?$#=&+_\"' ' in the second endpoint name",
			isFail:      true,
			Endpoints:   []string{"servers", "/!?$#=&+_\"'"},
			QueryParams: nil,
		},
		{
			Name:        "negative case then query params is not a struct",
			isFail:      true,
			Endpoints:   []string{"servers", "network"},
			QueryParams: "anystring",
		},
	}

	for _, c := range cases {
		t.Log("starting test case:", c.Name)
		u, err := NewURLBuilder().WithEndpoints(c.Endpoints...).WithQueryParams(c.QueryParams).Build()

		if c.isFail {
			assert.Error(t, err)
			continue
		}

		assert.NoError(t, err)

		uObj, err := url.Parse(u.String())
		assert.NoError(t, err)

		assert.Equal(t, strings.Join(c.Endpoints, "/"), uObj.Path)

		if c.QueryParams == nil {
			continue
		}

		if c.Name == "positive case with struct without query tags" ||
			c.Name == "positive case with empty struct, without query parameters" {
			q := uObj.Query()
			assert.Equal(t, 0, len(q))
			continue
		}

		assert.Equal(t, true, uObj.Query().Has("display_name"))
		assert.Equal(t, "network_name", uObj.Query().Get("display_name"))

		assert.Equal(t, true, uObj.Query().Has("status"))
		assert.Equal(t, "active", uObj.Query().Get("status"))

		assert.Equal(t, true, uObj.Query().Has("network_id"))
		assert.Equal(t, "7bf9e4a9-fab1-4415-a6f7-ff7284ee1870", uObj.Query().Get("network_id"))
	}

}
