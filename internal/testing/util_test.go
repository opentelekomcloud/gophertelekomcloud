package testing

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/internal"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRemainingKeys(t *testing.T) {
	type User struct {
		UserID    string `json:"user_id"`
		Username  string `json:"username"`
		Location  string `json:"-"`
		CreatedAt string `json:"-"`
		Status    string
		IsAdmin   bool
	}

	userResponse := map[string]any{
		"user_id":      "abcd1234",
		"username":     "jdoe",
		"location":     "Hawaii",
		"created_at":   "2017-06-08T02:49:03.000000",
		"status":       "active",
		"is_admin":     "true",
		"custom_field": "foo",
	}

	expected := map[string]any{
		"created_at":   "2017-06-08T02:49:03.000000",
		"is_admin":     "true",
		"custom_field": "foo",
	}

	actual := internal.RemainingKeys(User{}, userResponse)

	th.AssertDeepEquals(t, expected, actual)
}
