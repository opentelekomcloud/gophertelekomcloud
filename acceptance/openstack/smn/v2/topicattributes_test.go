package v2

import (
	"fmt"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/smn/v2/topicattributes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTopicAttributeWorkflow(t *testing.T) {
	client, err := clients.NewSmnV2Client()
	th.AssertNoErr(t, err)

	topic := createTopic(t, client)
	defer deleteTopic(t, client, topic)

	examplePolicy := fmt.Sprintf(policyTemplate, topic)
	opts := topicattributes.UpdateOpts{
		Value: examplePolicy,
	}
	err = topicattributes.Update(client, topic, attribute, opts).Err
	th.AssertNoErr(t, err)

	listOpts := topicattributes.ListOpts{Name: attribute}
	attributes, err := topicattributes.List(client, topic, listOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, examplePolicy, attributes[attribute])

	err = topicattributes.Delete(client, topic, attribute).Err
	th.AssertNoErr(t, err)
}

const (
	attribute      = "access_policy"
	policyTemplate = `
{
  "Version": "2016-09-07",
  "Id": "__default_policy_ID",
  "Statement": [
    {
      "Sid": "__service_pub_0",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "OBS"
        ]
      },
      "Action": [
        "SMN:Publish",
        "SMN:QueryTopicDetail"
      ],
      "Resource": "%s"
    }
  ]
}
`
)
