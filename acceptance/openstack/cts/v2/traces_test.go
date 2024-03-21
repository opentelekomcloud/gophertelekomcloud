package v2

import (
	"strconv"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cts/v2/traces"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTraces(t *testing.T) {
	cv2, err := clients.NewCTSV2Client()
	th.AssertNoErr(t, err)

	now := time.Now().UTC()
	toMilliseconds := now.UnixNano() / int64(time.Millisecond)

	oneDay := now.Add(-time.Hour * 24 * 1)
	fromMilliseconds := oneDay.UnixNano() / int64(time.Millisecond)

	listOpts := traces.ListTracesOpts{
		To:    strconv.FormatInt(toMilliseconds, 10),
		From:  strconv.FormatInt(fromMilliseconds, 10),
		Limit: "20",
	}

	var listResp traces.ListTracesResponse

	for i := 0; i < 2; i++ {
		list, err := traces.List(cv2, "system", listOpts)
		th.AssertNoErr(t, err)
		tools.PrintResource(t, list)

		listOpts.Next = list.MetaData.Marker
		listResp = *list
	}

	t.Logf("Number of Tracker Traces in latest API call : %d", len(listResp.Traces))
	th.AssertEquals(t, true, len(listResp.Traces) > 0)
}
