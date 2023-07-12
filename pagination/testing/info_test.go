package testing

import (
	"encoding/json"
	"fmt"
	"github.com/opentelekomcloud/gophertelekomcloud"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

type InfoPageResult struct {
	pagination.PageWithInfo
}

func (p InfoPageResult) IsEmpty() (bool, error) {
	l, err := ExtractStructs(p)
	if err != nil {
		return false, err
	}
	return len(l) == 0, nil
}

type Struct struct {
	ID string `json:"id"`
}

func ExtractStructs(p pagination.Page) ([]Struct, error) {
	var structs []Struct
	err := p.(InfoPageResult).ExtractIntoSlicePtr(&structs, "structs")
	if err != nil {
		return nil, err
	}
	return structs, nil
}

var expectedPolicies = []Struct{
	{
		ID: tools.RandomString("plc-", 20),
	},
	{
		ID: tools.RandomString("plc-", 20),
	},
	{
		ID: tools.RandomString("plc-", 20),
	},
	{
		ID: tools.RandomString("plc-", 20),
	},
}

func createInfoPager(t *testing.T) pagination.Pager {
	th.SetupHTTP()

	th.Mux.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		ms := r.Form["marker"]
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		switch {
		case len(ms) == 0:
			allJson, _ := json.MarshalIndent(expectedPolicies[:2], "", "  ")
			response := fmt.Sprintf(`
{
  "page_info": {
    "next_marker": "%s"
  },
  "structs": %s
}
`, expectedPolicies[2].ID, allJson)
			_, _ = fmt.Fprint(w, response)
		case len(ms) == 1 && ms[0] == expectedPolicies[2].ID:
			allJson, _ := json.MarshalIndent(expectedPolicies[2:], "", "  ")
			response := fmt.Sprintf(`
{
  "page_info": {
  },
  "structs": %s
}
`, allJson)
			_, _ = fmt.Fprint(w, response)
		default:
			t.Errorf("Request with unexpected marker: [%v]", ms)
		}
	})

	client := createClient()
	return pagination.Pager{
		Client:     client,
		InitialURL: th.Server.URL + "/page",
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return InfoPageResult{PageWithInfo: pagination.NewPageWithInfo(r)}
		},
	}
}

func TestInfoPageResult(t *testing.T) {
	pager := createInfoPager(t)
	defer th.TeardownHTTP()

	callCount := 0
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		actual, err := ExtractStructs(page)
		if err != nil {
			return false, err
		}
		t.Logf("Handler invoked with %+v", actual)

		switch callCount {
		case 0:
			th.AssertDeepEquals(t, actual, expectedPolicies[:2])
		case 1:
			th.AssertDeepEquals(t, actual, expectedPolicies[2:])
		}

		callCount += 1
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, callCount)
}

func TestInfoPageAll(t *testing.T) {
	pager := createInfoPager(t)
	defer th.TeardownHTTP()

	page, err := pager.AllPages()
	th.AssertNoErr(t, err)
	actual, err := ExtractStructs(page)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedPolicies, actual)
}
