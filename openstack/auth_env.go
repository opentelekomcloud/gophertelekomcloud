package openstack

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

/*
AuthOptionsFromEnv fills out an identity.AuthOptions structure with the
settings found on the various OpenStack OS_* environment variables.

The following variables provide sources of truth: OS_AUTH_URL, OS_USERNAME,
OS_PASSWORD, OS_TENANT_ID, and OS_TENANT_NAME.

Of these, OS_USERNAME, OS_PASSWORD, and OS_AUTH_URL must have settings,
or an error will result.  OS_TENANT_ID, OS_TENANT_NAME, OS_PROJECT_ID, and
OS_PROJECT_NAME are optional.

OS_TENANT_ID and OS_TENANT_NAME are mutually exclusive to OS_PROJECT_ID and
OS_PROJECT_NAME. If OS_PROJECT_ID and OS_PROJECT_NAME are set, they will
still be referred as "tenant" in Gophercloud.

To use this function, first set the OS_* environment variables (for example,
by sourcing an `openrc` file), then:

	opts, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.OldAuthenticatedClient(opts)
*/
func AuthOptionsFromEnv(envs ...*env) (golangsdk.AuthOptions, error) {
	e := NewEnv(defaultPrefix)
	if len(envs) > 0 {
		e = envs[0]
	}

	authURL := e.GetEnv("AUTH_URL")
	token := e.GetEnv("TOKEN", "TOKEN_ID")
	username := e.GetEnv("USERNAME")
	userID := e.GetEnv("USERID", "USER_ID")
	password := e.GetEnv("PASSWORD")
	projectID := e.GetEnv("PROJECT_ID", "TENANT_ID")
	projectName := e.GetEnv("PROJECT_NAME", "TENANT_NAME")
	domainID := e.GetEnv("DOMAIN_ID")
	domainName := e.GetEnv("DOMAIN_NAME")

	access := noEnv.GetEnv("AWS_ACCESS_KEY_ID")
	if access == "" {
		access = e.GetEnv("ACCESS_KEY", "ACCESS_KEY_ID")
	}
	secret := noEnv.GetEnv("AWS_ACCESS_KEY_SECRET")
	if secret == "" {
		secret = e.GetEnv("SECRET_KEY", "ACCESS_KEY_SECRET")
	}

	ao := golangsdk.AuthOptions{
		IdentityEndpoint: authURL,
		Username:         username,
		UserID:           userID,
		Password:         password,
		DomainID:         domainID,
		DomainName:       domainName,
		TenantID:         projectID,
		TenantName:       projectName,
		TokenID:          token,
		AccessKey:        access,
		SecretKey:        secret,
		AgencyName:       e.GetEnv("AGENCY_NAME"),
		AgencyDomainName: e.GetEnv("AGENCY_DOMAIN_NAME"),
		DelegatedProject: e.GetEnv("DELEGATED_PROJECT"),
	}
	return ao, nil
}
