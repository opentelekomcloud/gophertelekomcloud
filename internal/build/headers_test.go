package build

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/stretchr/testify/require"
)

type HeadersStruct struct {
	Bar    string `h:"x_bar" required:"true"`
	Baz    int    `h:"lorem_ipsum"`
	Boop   *bool  `h:"boop"`
	Foo    []int  `h:"foo"`
	SkipMe string
}

func TestHeaders_ok(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		opts     *HeadersStruct
		expected map[string]string
	}{
		"simple": {
			&HeadersStruct{
				Bar:    "bar",
				Baz:    67,
				Boop:   pointerto.Bool(false),
				SkipMe: "skipping",
			},
			map[string]string{
				"x_bar":       "bar",
				"lorem_ipsum": "67",
				"boop":        "false",
			},
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			data := data

			t.Parallel()

			headers, err := Headers(data.opts)
			require.NoError(t, err)
			require.EqualValues(t, data.expected, headers)
		})
	}
}

func TestHeaders_notOk(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		opts   any
		errMsg string
	}{
		"nil": {
			nil,
			"error building headers: nil options provided",
		},
		"not_struct": {
			map[string]any{},
			"error building headers: options type is not a struct",
		},
		"no_required": {
			&HeadersStruct{},
			"error building headers: required header [Bar] not set",
		},
		"unsupported_type": {
			&HeadersStruct{Bar: "1", Foo: []int{1, 2}},
			"error building headers: value of unsupported type []int",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			data := data

			t.Parallel()

			_, err := Headers(data.opts)
			require.Error(t, err)
			require.EqualError(t, err, data.errMsg)
		})
	}
}
