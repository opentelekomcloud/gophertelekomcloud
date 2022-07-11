package keys

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	resourcePath = "kms"
)

func getURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "describe-key")
}

func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "create-key")
}

func deleteURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "schedule-key-deletion")
}

func updateAliasURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "update-key-alias")
}

func updateDesURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "update-key-description")
}

func dataEncryptURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "create-datakey")
}

func encryptDEKURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "encrypt-datakey")
}

func enableKeyURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "enable-key")
}

func enableKeyRotationURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "enable-key-rotation")
}

func disableKeyRotationURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "disable-key-rotation")
}

func getKeyRotationStatusURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "get-key-rotation-status")
}

func updateKeyRotationIntervalURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "update-key-rotation-interval")
}

func disableKeyURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "disable-key")
}

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "list-keys")
}

func cancelDeleteURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath, "cancel-key-deletion")
}
