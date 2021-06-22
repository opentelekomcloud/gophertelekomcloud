package members

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "images"
	membersPath  = "members"
)

func imageMembersURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL(resourcePath, imageID, membersPath)
}

func listMembersURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

func createMemberURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

func imageMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return c.ServiceURL(resourcePath, imageID, membersPath, memberID)
}

func getMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

func updateMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

func deleteMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}
