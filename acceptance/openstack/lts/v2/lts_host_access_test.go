package v2

import (
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	ac "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/access-config"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/groups"
	hg "github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/host-groups"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/lts/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestLtsHostAccessLifecycle(t *testing.T) {
	clientV2, err := clients.NewLtsV2Client()
	th.AssertNoErr(t, err)

	clientV20, err := clients.NewLtsV20Client()
	th.AssertNoErr(t, err)

	clientV3, err := clients.NewLtsV3Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test-group-", 3)
	createOpts := groups.CreateOpts{
		LogGroupName: name,
		TTLInDays:    7,
	}
	t.Logf("Attempting to Create Log Group")
	logGroup, err := groups.Create(clientV2, createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Group")
		err = groups.Delete(clientV2, logGroup)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Log Stream")
	sname := tools.RandomString("test-stream-", 3)
	stream, err := streams.Create(clientV2, streams.CreateOpts{
		GroupId:       logGroup,
		LogStreamName: sname,
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Log Stream")
		err = streams.Delete(clientV2, streams.DeleteOpts{
			GroupId:  logGroup,
			StreamId: stream,
		})
		th.AssertNoErr(t, err)
	})

	nameHg := tools.RandomString("test-hgroup-", 3)
	createHgOpts := hg.CreateOpts{
		Name: nameHg,
		Type: "linux",
		Tags: []tags.ResourceTag{
			{
				Key: "fizz", Value: "buzz",
			},
		},
	}
	t.Logf("Attempting to Create Host Group without hosts")
	hgr, err := hg.Create(clientV3, createHgOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Host Group")
		_, err = hg.Delete(clientV3, hg.DeleteOpts{
			HostGroupIds: []string{hgr.ID},
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Create Host Access")
	acName := tools.RandomString("test-ac-", 3)
	opts := ac.CreateOpts{
		Name: acName,
		Type: "AGENT",
		Details: &ac.AccessConfigDetails{
			Paths:      []string{"/var/temp", "/var/log/*"},
			BlackPaths: []string{"/var/temp", "/var/log/*/a.log"},
			Format: &ac.AccessConfigFormat{
				Single: &ac.AccessConfigFormatBody{
					Mode:  "system",
					Value: strconv.FormatInt(time.Now().UnixMilli(), 10),
				},
			},
			WindowsLogInfo: &ac.AccessConfigWindowsLogInfo{
				Categories: []string{"System", "Security", "Setup"},
				TimeOffset: &ac.AccessConfigTimeOffset{
					Offset: 10,
					Unit:   "hour",
				},
				EventLevel: []string{"warning", "error", "critical", "verbose"},
			},
		},
		LogInfo: &ac.LogInfo{
			LogGroupId:  logGroup,
			LogStreamId: stream,
		},
		Tags: []tags.ResourceTag{
			{
				Key: "fizz", Value: "buzz",
			},
		},
		BinaryCollect: pointerto.Bool(true),
		LogSplit:      pointerto.Bool(true),
	}
	acCreate, err := ac.Create(clientV3, opts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to Delete Host Access")
		_, err = ac.Delete(clientV3, ac.DeleteOpts{
			AccessConfigIds: []string{acCreate.ID},
		})
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to Update Host Access")
	updateOpts := ac.UpdateOpts{
		ID: acCreate.ID,
		Details: &ac.AccessConfigDetailsUpdate{
			Paths:      []string{"/wjy/tesxxx"},
			BlackPaths: []string{"/wjy/hei/tesxxx", "/wjy/hei/tesxxxx"},
			Format: &ac.AccessConfigFormat{
				Single: &ac.AccessConfigFormatBody{
					Mode:  "wildcard",
					Value: "1234",
				},
			},
			WindowsLogInfo: &ac.AccessConfigWindowsLogInfoUpdate{
				Categories: []string{"Application", "System"},
				TimeOffset: &ac.AccessConfigTimeOffset{
					Offset: 7,
					Unit:   "day",
				},
				EventLevel: []string{"information", "warning", "error", "critical", "verbose"},
			},
		},
		Tags:          &[]tags.ResourceTag{},
		BinaryCollect: pointerto.Bool(false),
		LogSplit:      pointerto.Bool(false),
	}
	acUpdate, err := ac.Update(clientV3, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(acUpdate.Tags))
	if !reflect.DeepEqual([]string{"/wjy/hei/tesxxx", "/wjy/hei/tesxxxx"}, acUpdate.AccessConfigDetail.BlackPaths) {
		t.Errorf("Expected %v, but got %v", []string{"/wjy/hei/tesxxx", "/wjy/hei/tesxxxx"}, acUpdate.AccessConfigDetail.BlackPaths)
	}
	if !reflect.DeepEqual([]string{"/wjy/tesxxx"}, acUpdate.AccessConfigDetail.Paths) {
		t.Errorf("Expected %v, but got %v", []string{"/wjy/tesxxx"}, acUpdate.AccessConfigDetail.Paths)
	}

	t.Logf("Attempting to Update List Access")
	acList, err := ac.List(clientV3, ac.ListOpts{})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(acList.Result))
	th.AssertEquals(t, acCreate.Name, acList.Result[0].Name)

	agencyProjectId := os.Getenv("OS_LTS_AGENCY_PROJECT_ID")
	agencyDomainName := os.Getenv("OS_LTS_AGENCY_DOMAIN_NAME")
	agencyName := os.Getenv("OS_LTS_AGENCY_NAME")
	agencyStreamName := os.Getenv("OS_LTS_AGENCY_STREAM_NAME")
	agencyStreamId := os.Getenv("OS_LTS_AGENCY_STREAM_ID")
	agencyGroupName := os.Getenv("OS_LTS_AGENCY_GROUP_NAME")
	agencyGroupId := os.Getenv("OS_LTS_AGENCY_GROUP_ID")

	if agencyProjectId != "" || agencyDomainName != "" || agencyName != "" || agencyStreamName != "" ||
		agencyStreamId != "" || agencyGroupName != "" || agencyGroupId != "" {
		t.Logf("Attempting to Create Cross Agency Access")
		crossName := tools.RandomString("rule_", 3)
		access, err := ac.CrossAccess(clientV20, ac.CreateCrossOpts{
			PreviewAgencyList: []ac.PreviewAgencyLogAccess{
				{
					AgencyAccessType: "AGENCYACCESS",
					AgencyLogAccess:  crossName,

					LogStreamName: sname,
					LogStreamId:   stream,
					LogGroupName:  name,
					LogGroupId:    logGroup,

					ProjectId: clientV20.ProjectID,

					LogAgencyStreamName: agencyStreamName,
					LogAgencyStreamId:   agencyStreamId,
					LogAgencyGroupName:  agencyGroupName,
					LogAgencyGroupId:    agencyGroupId,
					AgencyProjectId:     agencyProjectId,
					AgencyDomainName:    agencyDomainName,
					AgencyName:          agencyName,
				},
			},
		})
		th.AssertNoErr(t, err)
		th.AssertEquals(t, crossName, access.AgencyLogAccess)
	}
}
