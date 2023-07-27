//go:build acceptance
// +build acceptance

package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/domains"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/roles"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRolesList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	listOpts := roles.ListOpts{
		DomainID: "default",
	}

	allPages, err := roles.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allRoles, err := roles.ExtractRoles(allPages)
	th.AssertNoErr(t, err)

	for _, role := range allRoles {
		tools.PrintResource(t, role)
	}
}

func TestRolesGet(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	role, err := FindRole(t, client)
	th.AssertNoErr(t, err)

	p, err := roles.Get(client, role.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, p)
}

func TestRoleCRUD(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	createOpts := roles.CreateOpts{
		Name:     "testrole",
		DomainID: "default",
		Extra: map[string]interface{}{
			"description": "test role description",
		},
	}

	// Create Role in the default domain
	role, err := CreateRole(t, client, &createOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		DeleteRole(t, client, role.ID)
	})

	tools.PrintResource(t, role)
	tools.PrintResource(t, role.Extra)

	updateOpts := roles.UpdateOpts{
		Extra: map[string]interface{}{
			"description": "updated test role description",
		},
	}

	newRole, err := roles.Update(client, role.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newRole)
	tools.PrintResource(t, newRole.Extra)
}

func TestRoleAssignToUserOnProject(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	project, err := CreateProject(t, client, nil)
	if err != nil {
		t.Fatal("Unable to create a project")
	}
	t.Cleanup(func() {
		DeleteProject(t, client, project.ID)
	})

	role, err := FindRole(t, client)
	th.AssertNoErr(t, err)

	user, err := CreateUser(t, client, nil)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		DeleteUser(t, client, user.ID)
	})

	t.Logf("Attempting to assign a role %s to a user %s on a project %s", role.Name, user.Name, project.Name)
	err = roles.Assign(client, role.ID, roles.AssignOpts{
		UserID:    user.ID,
		ProjectID: project.ID,
	}).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("Successfully assigned a role %s to a user %s on a project %s", role.Name, user.Name, project.Name)
	t.Cleanup(func() {
		UnassignRole(t, client, role.ID, &roles.UnassignOpts{
			UserID:    user.ID,
			ProjectID: project.ID,
		})
	})

	allPages, err := roles.ListAssignments(client, roles.ListAssignmentsOpts{
		ScopeProjectID: project.ID,
		UserID:         user.ID,
	}).AllPages()
	th.AssertNoErr(t, err)

	allRoleAssignments, err := roles.ExtractRoleAssignments(allPages)
	th.AssertNoErr(t, err)

	t.Logf("Role assignments of user %s on project %s:", user.Name, project.Name)
	for _, roleAssignment := range allRoleAssignments {
		tools.PrintResource(t, roleAssignment)
	}
}

func TestRoleAssignToUserOnDomain(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	domain, err := CreateDomain(t, client, &domains.CreateOpts{
		Enabled: pointerto.Bool(false),
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		DeleteDomain(t, client, domain.ID)
	})

	role, err := FindRole(t, client)
	th.AssertNoErr(t, err)

	user, err := CreateUser(t, client, nil)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		DeleteUser(t, client, user.ID)
	})

	t.Logf("Attempting to assign a role %s to a user %s on a domain %s", role.Name, user.Name, domain.Name)
	err = roles.Assign(client, role.ID, roles.AssignOpts{
		UserID:   user.ID,
		DomainID: domain.ID,
	}).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("Successfully assigned a role %s to a user %s on a domain %s", role.Name, user.Name, domain.Name)
	t.Cleanup(func() {
		UnassignRole(t, client, role.ID, &roles.UnassignOpts{
			UserID:   user.ID,
			DomainID: domain.ID,
		})
	})

	allPages, err := roles.ListAssignments(client, roles.ListAssignmentsOpts{
		ScopeDomainID: domain.ID,
		UserID:        user.ID,
	}).AllPages()
	th.AssertNoErr(t, err)

	allRoleAssignments, err := roles.ExtractRoleAssignments(allPages)
	th.AssertNoErr(t, err)

	t.Logf("Role assignments of user %s on domain %s:", user.Name, domain.Name)
	for _, roleAssignment := range allRoleAssignments {
		tools.PrintResource(t, roleAssignment)
	}
}

func TestRoleAssignToGroupOnDomain(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	domain, err := CreateDomain(t, client, &domains.CreateOpts{
		Enabled: pointerto.Bool(false),
	})
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		DeleteDomain(t, client, domain.ID)
	})

	role, err := FindRole(t, client)
	th.AssertNoErr(t, err)

	group, err := CreateGroup(t, client, nil)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		DeleteGroup(t, client, group.ID)
	})

	t.Logf("Attempting to assign a role %s to a group %s on a domain %s", role.Name, group.Name, domain.Name)
	err = roles.Assign(client, role.ID, roles.AssignOpts{
		GroupID:  group.ID,
		DomainID: domain.ID,
	}).ExtractErr()
	th.AssertNoErr(t, err)

	t.Logf("Successfully assigned a role %s to a group %s on a domain %s", role.Name, group.Name, domain.Name)
	t.Cleanup(func() {
		UnassignRole(t, client, role.ID, &roles.UnassignOpts{
			GroupID:  group.ID,
			DomainID: domain.ID,
		})
	})

	allPages, err := roles.ListAssignments(client, roles.ListAssignmentsOpts{
		GroupID:       group.ID,
		ScopeDomainID: domain.ID,
	}).AllPages()
	th.AssertNoErr(t, err)

	allRoleAssignments, err := roles.ExtractRoleAssignments(allPages)
	th.AssertNoErr(t, err)

	t.Logf("Role assignments of group %s on domain %s:", group.Name, domain.Name)
	for _, roleAssignment := range allRoleAssignments {
		tools.PrintResource(t, roleAssignment)
	}
}

func TestRoleAssignToGroupOnProject(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	project, err := CreateProject(t, client, nil)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		DeleteProject(t, client, project.ID)
	})

	role, err := FindRole(t, client)
	th.AssertNoErr(t, err)

	group, err := CreateGroup(t, client, nil)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		DeleteGroup(t, client, group.ID)
	})

	t.Logf("Attempting to assign a role %s to a group %s on a project %s", role.Name, group.Name, project.Name)
	err = roles.Assign(client, role.ID, roles.AssignOpts{
		GroupID:   group.ID,
		ProjectID: project.ID,
	}).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Successfully assigned a role %s to a group %s on a project %s", role.Name, group.Name, project.Name)

	t.Cleanup(func() {
		UnassignRole(t, client, role.ID, &roles.UnassignOpts{
			GroupID:   group.ID,
			ProjectID: project.ID,
		})
	})

	allPages, err := roles.ListAssignments(client, roles.ListAssignmentsOpts{
		GroupID:        group.ID,
		ScopeProjectID: project.ID,
	}).AllPages()
	th.AssertNoErr(t, err)

	allRoleAssignments, err := roles.ExtractRoleAssignments(allPages)
	th.AssertNoErr(t, err)

	t.Logf("Role assignments of group %s on project %s:", group.Name, project.Name)
	for _, roleAssignment := range allRoleAssignments {
		tools.PrintResource(t, roleAssignment)
	}
}
