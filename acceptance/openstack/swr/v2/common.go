package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/organizations"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/swr/v2/repositories"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

type dependencies struct {
	t      *testing.T
	client *golangsdk.ServiceClient
}

func (d dependencies) createOrganization(name string) {
	th.AssertNoErr(d.t, organizations.Create(d.client, organizations.CreateOpts{Namespace: name}).ExtractErr())
}

func (d dependencies) deleteOrganization(name string) {
	th.AssertNoErr(d.t, organizations.Delete(d.client, name).ExtractErr())
}

func (d dependencies) createRepository(organization, repository string) {
	th.AssertNoErr(d.t, repositories.Create(d.client, organization, repositories.CreateOpts{
		Repository:  repository,
		Category:    "linux",
		Description: "Used repo",
		IsPublic:    false,
	}).ExtractErr())
}

func (d dependencies) deleteRepository(organization, repository string) {
	th.AssertNoErr(d.t, repositories.Delete(d.client, organization, repository).ExtractErr())
}
