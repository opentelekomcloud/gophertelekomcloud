package testing

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// LinkedPager sample and test cases.

type LinkedPageResult struct {
	pagination.LinkedPageBase
}

func (r LinkedPageResult) IsEmpty() (bool, error) {
	is, err := ExtractLinkedInts(r)
	return len(is) == 0, err
}

func ExtractLinkedInts(r pagination.Page) ([]int, error) {
	var s struct {
		Ints []int `json:"ints"`
	}

	err := extract.Into(bytes.NewReader((r.(LinkedPageResult)).Body), &s)
	return s.Ints, err
}

func createLinked() pagination.Pager {
	th.SetupHTTP()

	th.Mux.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{ "ints": [1, 2, 3], "links": { "next": "%s/page2" } }`, th.Server.URL)
	})

	th.Mux.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{ "ints": [4, 5, 6], "links": { "next": "%s/page3" } }`, th.Server.URL)
	})

	th.Mux.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{ "ints": [7, 8, 9], "links": { "next": null } }`)
	})

	client := createClient()

	createPage := func(r pagination.PageResult) pagination.Page {
		return LinkedPageResult{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.Pager{
		Client:     client,
		InitialURL: th.Server.URL + "/page1",
		CreatePage: createPage,
	}
}

func TestEnumerateLinked(t *testing.T) {
	pager := createLinked()
	defer th.TeardownHTTP()

	callCount := 0
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		actual, err := ExtractLinkedInts(page)
		if err != nil {
			return false, err
		}

		t.Logf("Handler invoked with %v", actual)

		var expected []int
		switch callCount {
		case 0:
			expected = []int{1, 2, 3}
		case 1:
			expected = []int{4, 5, 6}
		case 2:
			expected = []int{7, 8, 9}
		default:
			t.Fatalf("Unexpected call count: %d", callCount)
			return false, nil
		}

		th.AssertDeepEquals(t, expected, actual)
		callCount++
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 3, callCount)
}

func TestAllPagesLinked(t *testing.T) {
	pager := createLinked()
	defer th.TeardownHTTP()

	page, err := pager.AllPages()
	th.AssertNoErr(t, err)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual, err := ExtractLinkedInts(page)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, expected, actual)
}
