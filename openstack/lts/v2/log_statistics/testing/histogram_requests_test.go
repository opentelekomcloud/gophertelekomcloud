package testing

import (
	"net/http"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/log_statistics"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func TestListLogHistogram(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleHistogram(t, http.StatusOK, histogramResponse)

	actual, err := log_statistics.ListLogHistogram(fake.ServiceClient(), histogramOpts())
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &log_statistics.LogHistogramResponse{
		Count: 1,
		Histogram: []log_statistics.LogHistogram{
			{Num: 1, StartTime: 1637821594579, EndTime: 1637821595000},
			{Num: 0, StartTime: 1637821654000, EndTime: 1637821654579},
		},
		IsQueryComplete: true,
	}, actual)
}

func TestListLogHistogramExtractsZeroValues(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleHistogram(t, http.StatusOK, `{}`)

	actual, err := log_statistics.ListLogHistogram(fake.ServiceClient(), histogramOpts())
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &log_statistics.LogHistogramResponse{}, actual)
}

func TestListLogHistogramRejectsMissingInput(t *testing.T) {
	_, err := log_statistics.ListLogHistogram(fake.ServiceClient(), log_statistics.ListLogHistogramOpts{})
	if err == nil {
		t.Fatal("expected missing required input to return an error")
	}
}

func TestListLogHistogramReturnsBadRequestError(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	handleHistogram(
		t,
		http.StatusBadRequest,
		`{"error_code":"LTS.0601","error_msg":"must be less than or equal to 86400000"}`,
	)

	_, err := log_statistics.ListLogHistogram(fake.ServiceClient(), histogramOpts())
	if err == nil {
		t.Fatal("expected bad request error")
	}
}
