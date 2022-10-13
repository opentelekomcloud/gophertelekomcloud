package build

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/stretchr/testify/require"
)

type QueryStruct struct {
	Bar     string            `q:"x_bar" required:"true"`
	Baz     int               `q:"lorem_ipsum"`
	Foo     []int             `q:"foo"`
	FooStr  []string          `q:"foostr"`
	Map1    map[string]string `q:"map"`
	Map2    map[string]int    `q:"map_invalid"` // not supported
	Req     *bool             `q:"req"`
	NotHere string
}

func TestQueryString_ok(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		opts     *QueryStruct
		expected string
	}{
		"simple": {
			&QueryStruct{
				Bar:     "AAA",
				Baz:     200,
				Req:     pointerto.Bool(false),
				NotHere: "no",
			},
			"?lorem_ipsum=200&req=false&x_bar=AAA",
		},
		"with_int_slice": {
			&QueryStruct{
				Bar: "AAA",
				Foo: []int{1, 2, 3},
			},
			"?foo=1&foo=2&foo=3&x_bar=AAA",
		},
		"with_str_slice": {
			&QueryStruct{
				Bar:    "AAA",
				FooStr: []string{"a", "b"},
			},
			"?foostr=a&foostr=b&x_bar=AAA",
		},
		"with_str_map": {
			&QueryStruct{
				Bar:  "AAA",
				Map1: map[string]string{"foo": "bar"},
			},
			"?map=%7B%27foo%27%3A%27bar%27%7D&x_bar=AAA",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			data := data

			t.Parallel()

			query, err := QueryString(data.opts)

			require.NoError(t, err)
			require.EqualValues(t, data.expected, query.String())
		})
	}
}

func TestQueryString_notOk(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		opts   interface{}
		errMsg string
	}{
		"nil opts": {
			nil,
			"error building query string: nil options provided",
		},
		"missing_required": {
			&QueryStruct{},
			"error building query string: required query parameter [Bar] not set",
		},
		"not_struct": {
			map[string]interface{}{},
			"error building query string: options type is not a struct",
		},
		"with_non_str_map": {
			&QueryStruct{Bar: "1", Map2: map[string]int{"foo": 1}},
			"error building query string: expected map[string]string, got map[string]int",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			data := data

			t.Parallel()

			_, err := QueryString(data.opts)

			require.Error(t, err)
			require.EqualError(t, err, data.errMsg)
		})
	}
}
