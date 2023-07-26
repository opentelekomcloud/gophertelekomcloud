package extract

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// randomString duplicates here to avoid cyclic imports
// TODO: this function should be moved to some other package later
func randomString(prefix string, n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	_, _ = rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return prefix + string(bytes)
}

type localCloser struct {
	*bytes.Reader

	closed bool
}

func (lc *localCloser) Close() error {
	lc.closed = true
	return nil
}

func TestInto(t *testing.T) {
	key := "data_key"
	value := randomString("v-", 20)

	expected := map[string]string{key: value}

	t.Run("io.Reader", func(t *testing.T) {
		t.Parallel()

		data := bytes.NewReader([]byte(fmt.Sprintf(`{ "data_key":  "%s"}`, value)))

		actual := make(map[string]string)
		err := Into(data, &actual)

		assert.NoError(t, err) // not exiting after one fail
		assert.EqualValues(t, expected[key], actual[key])
	})

	t.Run("io.ReadCloser", func(t *testing.T) {
		t.Parallel()

		data := bytes.NewReader([]byte(fmt.Sprintf(`{ "data_key":  "%s"}`, value)))
		closer := &localCloser{Reader: data}

		actual := make(map[string]string)
		err := Into(closer, &actual)

		assert.NoError(t, err) // not exiting after one fail
		assert.EqualValues(t, expected[key], actual[key])
		assert.True(t, closer.closed)
	})
}

type TestDataType struct {
	DataKey string `json:"data_key"`
}

type TestDataType2 struct {
	TestDataType

	SecondDataField string `json:"second_data_field"`
}

func readerFromString(src string) io.Reader {
	return bytes.NewReader([]byte(src))
}

func TestIntoStructPtr(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		actual := new(TestDataType)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"data_key": "%s"
		}
		`, value)

		err := IntoStructPtr(readerFromString(data), actual, "")
		require.NoError(t, err)
		require.Equal(t, value, actual.DataKey)
	})

	t.Run("with label", func(t *testing.T) {
		t.Parallel()

		actual := new(TestDataType)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"internal": {
				"data_key": "%s"
			}
		}
		`, value)

		err := IntoStructPtr(readerFromString(data), actual, "internal")
		require.NoError(t, err)
		require.Equal(t, value, actual.DataKey)
	})

	t.Run("with label and embed", func(t *testing.T) {
		t.Parallel()

		actual := new(TestDataType2)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"internal": {
				"data_key": "%s",
				"second_data_field": "%[1]s-2"
			}
		}
		`, value)

		err := IntoStructPtr(readerFromString(data), actual, "internal")
		require.NoError(t, err)
		require.Equal(t, value, actual.DataKey)
		require.Equal(t, value+"-2", actual.SecondDataField)
	})

	t.Run("with label (err)", func(t *testing.T) {
		t.Parallel()

		actual := new(TestDataType)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"internal": {
				"data_key": "%s"
			}
		}
		`, value)

		err := IntoStructPtr(readerFromString(data), actual, "")
		require.NoError(t, err)
		require.Equal(t, "", actual.DataKey)
	})

	t.Run("non pointer", func(t *testing.T) {
		t.Parallel()

		actual := TestDataType{}
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"data_key": "%s"
		}
		`, value)

		err := IntoStructPtr(readerFromString(data), actual, "")
		require.EqualError(t, err, "expected pointer, got struct")
	})

	t.Run("non struct", func(t *testing.T) {
		t.Parallel()

		actual := make(map[string]any)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"data_key": "%s"
		}
		`, value)

		err := IntoStructPtr(readerFromString(data), &actual, "")
		require.EqualError(t, err, "expected pointer to struct, got: *map[string]interface {}")
	})
}

func TestIntoSlicePtr(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		actual := make([]TestDataType, 0)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		[{
			"data_key": "%s"
		}]
		`, value)

		err := IntoSlicePtr(readerFromString(data), &actual, "")
		require.NoError(t, err)
		require.Len(t, actual, 1)
		require.Equal(t, value, actual[0].DataKey)
	})

	t.Run("with label", func(t *testing.T) {
		t.Parallel()

		actual := make([]TestDataType, 0)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"data": [{ "data_key": "%s" }]
		}
		`, value)

		err := IntoSlicePtr(readerFromString(data), &actual, "data")
		require.NoError(t, err)
		require.Len(t, actual, 1)
		require.Equal(t, value, actual[0].DataKey)
	})

	t.Run("with label and embed", func(t *testing.T) {
		t.Parallel()

		actual := make([]TestDataType2, 0)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"internal": [{
				"data_key": "%s",
				"second_data_field": "%[1]s"
			}]
		}
		`, value)

		err := IntoSlicePtr(readerFromString(data), &actual, "internal")
		require.NoError(t, err)
		require.Len(t, actual, 1)
		require.Equal(t, value, actual[0].DataKey)
		require.Equal(t, value, actual[0].SecondDataField)
	})

	t.Run("not pointer", func(t *testing.T) {
		t.Parallel()

		actual := make([]TestDataType, 0)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		[{
			"data_key": "%s"
		}]
		`, value)

		err := IntoSlicePtr(readerFromString(data), actual, "")
		require.EqualError(t, err, "expected pointer, got slice")
	})

	t.Run("not slice", func(t *testing.T) {
		t.Parallel()

		actual := new(TestDataType)
		value := randomString("v-", 20)

		data := fmt.Sprintf(`
		{
			"data_key": "%s"
		}
		`, value)

		err := IntoSlicePtr(readerFromString(data), actual, "")
		require.EqualError(t, err, "expected pointer to slice, got: *extract.TestDataType")
	})
}
