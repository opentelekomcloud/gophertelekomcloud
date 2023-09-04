package v3

import (
	"fmt"
	"testing"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/credentials"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemporaryCredential(t *testing.T) {
	cloud, err := openstack.NewEnv("OS_").Cloud()
	if err != nil {
		t.Fatal(err)
	}
	if cloud.AuthType != "agency" || cloud.AuthInfo.AgencyName == "" || cloud.AuthInfo.DelegatedProject == "" {
		t.Skip("Agency authorization is not set up")
	}

	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	cred, err := credentials.CreateTemporary(client, credentials.CreateTemporaryOpts{
		Methods:  []string{"token"},
		Token:    client.Token(),
		Duration: 1800,
	}).Extract()
	require.NoError(t, err, "error creating temporary AK/SK")

	currentTime := time.Now()
	returnedTime, err := time.Parse(time.RFC3339Nano, cred.ExpiresAt)
	th.AssertNoErr(t, err)

	expiration := returnedTime.Sub(currentTime).Minutes()
	assert.Condition(t, func() bool {
		if expiration <= 30.0 {
			fmt.Println(expiration)
			return false
		}
		return true
	}, "returned time should be 30 min or more")

	assert.NotEmpty(t, cred.AccessKey)
	assert.NotEmpty(t, cred.SecretKey)
	assert.NotEmpty(t, cred.SecurityToken)
	assert.NotEmpty(t, cred.ExpiresAt)
}

func TestTemporaryCredentialAgency(t *testing.T) {
	cloud, err := openstack.NewEnv("OS_").Cloud()
	if err != nil {
		t.Fatal(err)
	}
	if cloud.AuthType != "agency" || cloud.AuthInfo.AgencyName == "" || cloud.AuthInfo.DelegatedProject == "" {
		t.Skip("Agency authorization is not set up")
	}

	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	cred, err := credentials.CreateTemporary(client, credentials.CreateTemporaryOpts{
		Methods:    []string{"assume_role"},
		Token:      client.Token(),
		Duration:   1800,
		AgencyName: cloud.AuthInfo.AgencyName,
		DomainID:   client.DomainID,
	}).Extract()
	require.NoError(t, err, "error creating temporary AK/SK")

	currentTime := time.Now()
	returnedTime, err := time.Parse(time.RFC3339Nano, cred.ExpiresAt)
	th.AssertNoErr(t, err)

	expiration := returnedTime.Sub(currentTime).Minutes()
	assert.Condition(t, func() bool {
		if expiration <= 30.0 {
			fmt.Println(expiration)
			return false
		}
		return true
	}, "returned time should be 30 min or more")

	assert.NotEmpty(t, cred.AccessKey)
	assert.NotEmpty(t, cred.SecretKey)
	assert.NotEmpty(t, cred.SecurityToken)
	assert.NotEmpty(t, cred.ExpiresAt)
}
