package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestMetricData(t *testing.T) {
	client, err := clients.NewCesV1Client()
	th.AssertNoErr(t, err)

}
