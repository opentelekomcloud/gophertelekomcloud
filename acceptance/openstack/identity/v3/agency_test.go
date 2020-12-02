package v3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/credentials"
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
		Duration: 60,
	}).Extract()
	require.NoError(t, err, "error creating temporary AK/SK")

	assert.NotEmpty(t, cred.AccessKey)
	assert.NotEmpty(t, cred.SecretKey)
	assert.NotEmpty(t, cred.SecurityToken)
	assert.NotEmpty(t, cred.ExpiresAt)
}
