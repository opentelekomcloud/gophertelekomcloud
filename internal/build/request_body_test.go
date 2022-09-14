package build

import (
	"encoding/json"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Data string `json:"data"`
}

type testStructWPtr struct {
	Data *string `json:"data"`
}

func TestRequestBody_MarshalJSON(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		body     Body
		expected string
	}{
		"int_slice": {
			Body{
				Wrapped: []int{1, 2, 3},
			},
			"[1, 2, 3]",
		},
		"int_slice_wrapped": {
			Body{
				RootTag: "items",
				Wrapped: []int{1, 2, 3},
			},
			`{"items": [1, 2, 3]}`,
		},
		"struct": {
			Body{
				RootTag: "",
				Wrapped: &testStruct{Data: "123"},
			},
			`{"data": "123"}`,
		},
		"struct_wrapped": {
			Body{
				RootTag: "root",
				Wrapped: &testStruct{Data: "123"},
			},
			`{"root": {"data": "123"}}`,
		},
		"struct_w_pointer_field": {
			Body{
				RootTag: "",
				Wrapped: testStructWPtr{Data: pointerto.String("123")},
			},
			`{"data": "123"}`,
		},
		"struct_w_pointer_field_wrapped": {
			Body{
				RootTag: "root",
				Wrapped: testStructWPtr{Data: pointerto.String("123")},
			},
			`{"root": {"data": "123"}}`,
		},
	}

	for name, data := range cases {
		data := data

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := json.Marshal(data.body)
			require.NoError(t, err)

			require.JSONEq(t, data.expected, string(actual))
		})
	}
}

func TestRequestBody_String(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		requestBody Body
		expected    string
	}{
		"simple": {
			Body{Wrapped: "data"},
			`"data"`,
		},
		"with_parent": {
			Body{Wrapped: "data", RootTag: "root"},
			`{
  "root": "data"
}`,
		},
	}

	for name, data := range cases {
		data := data

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, data.expected, data.requestBody.String())
		})
	}
}

func TestRequestBody_String_Err(t *testing.T) {
	t.Parallel()

	body := Body{Wrapped: complex(float64(1), float64(2))}
	require.Equal(t, "!err: json: unsupported type: complex128", body.String())
}

func TestBuildRequestBody(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		source := testStruct{"data"}

		expected := &Body{
			Wrapped: source,
		}

		actual, err := RequestBody(source, "")
		require.NoError(t, err)
		require.EqualValues(t, expected, actual)
	})

	t.Run("validation_fail", func(t *testing.T) {
		source := testStructRequired{}

		_, err := RequestBody(source, "")
		require.EqualError(t, err, "error building request body: missing input for argument [Field]")
	})

	t.Run("nil_body", func(t *testing.T) {
		_, err := RequestBody(nil, "")
		require.EqualError(t, err, "error building request body: nil options provided")
	})
}
