package build

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/stretchr/testify/require"
)

func TestValidateTags_nil(t *testing.T) {
	t.Parallel()

	require.NoError(t, ValidateTags(nil))
}

func TestValidateTags_notStruct(t *testing.T) {
	t.Parallel()

	require.NoError(t, ValidateTags("data"))
}

type testStructRequired struct {
	Field string `json:"field" required:"true"`
}

type testStructRequiredPtrField struct {
	Field *string `json:"field" required:"true"`
}

func TestValidateTags_required(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		cases := map[string]any{
			"simple":                     testStructRequired{"data"},
			"pointer_struct":             &testStructRequired{"data"},
			"pointer_value":              testStructRequiredPtrField{pointerto.String("data")},
			"pointer_value_empty_string": testStructRequiredPtrField{pointerto.String("")},
		}

		for name, data := range cases {
			data := data
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				require.NoError(t, ValidateTags(data))
			})
		}
	})

	t.Run("not_ok", func(t *testing.T) {
		t.Parallel()

		cases := map[string]any{
			"simple":                     testStructRequired{},
			"pointer_struct":             &testStructRequired{},
			"pointer_value":              testStructRequiredPtrField{},
			"pointer_value_empty_string": &testStructRequiredPtrField{},
		}

		for name, data := range cases {
			data := data
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				require.EqualError(t, ValidateTags(data), "missing input for argument [Field]")
			})
		}
	})

}

type testStructOr struct {
	Field1 string `json:"field1" or:"Field2"`
	Field2 string `json:"field2" or:"Field1"`
}

type testStructOrPtrField struct {
	Field1 *string `json:"field1" or:"Field2"`
	Field2 *string `json:"field2" or:"Field1"`
}

func TestValidateTags_or(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		// pair-wise testing :D
		cases := map[string]any{
			"simple_one":                 testStructOr{"data", ""},
			"simple_both":                testStructOr{"data1", "data2"},
			"pointer_struct":             &testStructOr{"", "data"},
			"pointer_value":              testStructOrPtrField{nil, pointerto.String("data")},
			"pointer_value_empty_string": &testStructOrPtrField{pointerto.String(""), nil},
		}

		for name, data := range cases {
			data := data
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				require.NoError(t, ValidateTags(data))
			})
		}
	})

	t.Run("not_ok", func(t *testing.T) {
		cases := map[string]any{
			"simple":                     testStructOr{},
			"pointer_struct":             &testStructOr{},
			"pointer_value":              testStructOrPtrField{},
			"pointer_value_empty_string": &testStructOrPtrField{},
		}

		for name, data := range cases {
			data := data
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				err := ValidateTags(data)
				require.NotNil(t, err)

				errText := err.Error()
				require.Contains(t, errText, `multiple errors returned:`)
				require.Contains(t, errText, `at least one of Field1 and Field2 must be provided`)
				require.Contains(t, errText, `at least one of Field2 and Field1 must be provided`)
			})
		}
	})
}

type testStructXor struct {
	Field1 string `json:"field1" xor:"Field2"`
	Field2 string `json:"field2" xor:"Field1"`
}

type testStructXorPtrField struct {
	Field1 *string `json:"field1" xor:"Field2"`
	Field2 *string `json:"field2" xor:"Field1"`
}

func TestValidateTags_xor(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		cases := map[string]any{
			"simple":                     testStructXor{"data", ""},
			"pointer_struct":             &testStructXor{"", "data"},
			"pointer_value":              testStructXorPtrField{nil, pointerto.String("data")},
			"pointer_value_empty_string": &testStructXorPtrField{pointerto.String(""), nil},
		}

		for name, data := range cases {
			data := data
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				require.NoError(t, ValidateTags(data))
			})
		}
	})

	t.Run("neither", func(t *testing.T) {
		cases := map[string]any{
			"simple":                     testStructXor{},
			"pointer_struct":             &testStructXor{},
			"pointer_value":              testStructXorPtrField{},
			"pointer_value_empty_string": &testStructXorPtrField{},
		}

		for name, data := range cases {
			data := data
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				err := ValidateTags(data)
				require.NotNil(t, err)

				errText := err.Error()
				require.Contains(t, errText, `multiple errors returned:`)
				require.Contains(t, errText, `exactly one of Field1 and Field2 must be provided`)
				require.Contains(t, errText, `exactly one of Field2 and Field1 must be provided`)
			})
		}
	})

	t.Run("both", func(t *testing.T) {
		cases := map[string]any{
			"simple":                     testStructXor{"data1", "data2"},
			"pointer_struct":             &testStructXor{"data1", "data2"},
			"pointer_value":              testStructXorPtrField{pointerto.String(""), pointerto.String("")},
			"pointer_value_empty_string": &testStructXorPtrField{pointerto.String(""), pointerto.String("")},
		}

		for name, data := range cases {
			data := data
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				err := ValidateTags(data)
				require.NotNil(t, err)

				errText := err.Error()
				require.Contains(t, errText, `multiple errors returned:`)
				require.Contains(t, errText, `exactly one of Field1 and Field2 must be provided`)
				require.Contains(t, errText, `exactly one of Field2 and Field1 must be provided`)
			})
		}
	})
}
