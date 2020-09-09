package testing

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestServiceURL(t *testing.T) {
	c := &golangsdk.ServiceClient{Endpoint: "http://123.45.67.8/"}
	expected := "http://123.45.67.8/more/parts/here"
	actual := c.ServiceURL("more", "parts", "here")
	th.CheckEquals(t, expected, actual)
}
