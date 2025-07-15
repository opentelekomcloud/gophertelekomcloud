package v2

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	cs "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/cloud-structuring"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsStructuringLifecycle(t *testing.T) {
	client, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	clientV3, err := clients.NewLtsV3Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}
	t.Logf("Attempting to Create Log Group")
	group, err := groups.Create(client, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Group")
		err = groups.Delete(client, group)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Log Stream")
	sname := tools.RandomString("test-stream-", 3)
	stream, err := streams.Create(client, streams.CreateOpts{
		GroupId:       group,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Stream")
		err = streams.Delete(client, streams.DeleteOpts{
			GroupId:  group,
			StreamId: stream,
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Structuring Rule")
	err = cs.Create(clientV3, cs.CreateOpts{
		LogGroupId:    group,
		LogStreamId:   stream,
		TemplateId:    pointerto.String(""),
		Type:          "built_in",
		Name:          "ELB",
		QuickAnalysis: pointerto.Bool(false),
		DemoFields: []cs.Field{
			{
				Name:       "msec",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "access_log_topic_id",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "time_iso8601",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "log_ver",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "remote_addr",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "remote_port",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "status",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "request_method",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "scheme",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "host",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "router_request_uri",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "server_protocol",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "request_length",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "bytes_sent",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "body_bytes_sent",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "request_time",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "upstream_status",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "upstream_connect_time",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "upstream_header_time",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "upstream_response_time",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "upstream_addr",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "http_user_agent",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "http_referer",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "http_x_forwarded_for",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "lb_name",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "listener_name",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "listener_id",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "pool_name",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "member_name",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "tenant_id",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "eip_address",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "eip_port",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "upstream_addr_priv",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "certificate_id",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "ssl_protocol",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "ssl_cipher",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "sni_domain_name",
				IsAnalysis: pointerto.Bool(false),
			},
			{
				Name:       "tcpinfo_rtt",
				IsAnalysis: pointerto.Bool(false),
			}},
		TagFields: []cs.Field{
			{
				Name:       "hostIP",
				IsAnalysis: pointerto.Bool(true),
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get Structuring Rule")
	get, err := cs.Get(client, group, stream)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ELB", get.Name)
	for _, field := range get.DemoFields {
		if field.Name == "msec" {
			th.AssertEquals(t, false, field.IsAnalysis)
		}
		if field.Name == "access_log_topic_id" {
			th.AssertEquals(t, false, field.IsAnalysis)
		}
	}

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Structuring Rule")
		err = cs.Delete(client, get.ID)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Update Structuring Rule")
	err = cs.Update(clientV3, cs.CreateOpts{
		LogGroupId:    group,
		LogStreamId:   stream,
		TemplateId:    pointerto.String(""),
		Type:          "built_in",
		Name:          "ELB",
		QuickAnalysis: pointerto.Bool(false),
		DemoFields: []cs.Field{
			{
				Name:       "msec",
				IsAnalysis: pointerto.Bool(true),
			},
			{
				Name:       "access_log_topic_id",
				IsAnalysis: pointerto.Bool(true),
			},
		},
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get Updated Structuring Rule")
	getUp, err := cs.Get(client, group, stream)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ELB", getUp.Name)
	for _, field := range getUp.DemoFields {
		if field.Name == "msec" {
			th.AssertEquals(t, true, field.IsAnalysis)
		}
		if field.Name == "access_log_topic_id" {
			th.AssertEquals(t, true, field.IsAnalysis)
		}
	}

	customRuleId := os.Getenv("OS_LTS_CUSTOM_RULE_ID")
	if customRuleId != "" {
		t.Logf("Attempting to Create Custom Structuring Rule")
		err = cs.Create(clientV3, cs.CreateOpts{
			LogGroupId:  group,
			LogStreamId: stream,
			DemoFields: []cs.Field{
				{
					Name:       "key",
					IsAnalysis: pointerto.Bool(true),
				},
			},
			TemplateId:    pointerto.String(customRuleId),
			Type:          "custom",
			Name:          "create-json",
			QuickAnalysis: pointerto.Bool(true),
		})
		th.AssertNoErr(t, err)

		brief, err := cs.ListBrief(clientV3)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, true, len(brief) >= 1)

		full, err := cs.List(clientV3, cs.ListOpts{})
		th.AssertNoErr(t, err)
		th.AssertEquals(t, true, len(full) >= 1)
	}
}
