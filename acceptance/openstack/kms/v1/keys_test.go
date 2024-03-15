package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/kms/v1/keys"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestKmsKeysLifecycle(t *testing.T) {
	client, err := clients.NewKMSV1Client()
	th.AssertNoErr(t, err)

	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	if kmsID == "" {
		t.Skip("OS_KMS_ID env var is missing but KMSv1 grant test requires")
	}

	createOpts := keys.CreateOpts{
		KeyAlias:       kmsID,
		KeyDescription: "some description",
	}
	createKey, err := keys.Create(client, createOpts).ExtractKeyInfo()
	th.AssertNoErr(t, err)

	defer func() {
		deleteOpts := keys.DeleteOpts{
			KeyID:       createKey.KeyID,
			PendingDays: "7",
		}
		_, err := keys.Delete(client, deleteOpts).Extract()
		th.AssertNoErr(t, err)
	}()

	keyGet, err := keys.Get(client, createKey.KeyID).ExtractKeyInfo()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.KeyAlias, keyGet.KeyAlias)
	th.AssertEquals(t, keyGet.KeyState, "2")

	deleteOpts := keys.DeleteOpts{
		KeyID:       createKey.KeyID,
		PendingDays: "7",
	}
	_, err = keys.Delete(client, deleteOpts).Extract()
	th.AssertNoErr(t, err)

	keyGetDeleted, err := keys.Get(client, createKey.KeyID).ExtractKeyInfo()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, keyGetDeleted.KeyState, "4")

	cancelDeleteOpts := keys.CancelDeleteOpts{
		KeyID: createKey.KeyID,
	}
	_, err = keys.CancelDelete(client, cancelDeleteOpts).Extract()
	th.AssertNoErr(t, err)

	_, err = keys.EnableKey(client, createKey.KeyID).Extract()
	th.AssertNoErr(t, err)

	keyGetEnabled, err := keys.Get(client, createKey.KeyID).ExtractKeyInfo()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, keyGetEnabled.KeyState, "2")

	_, err = keys.DisableKey(client, createKey.KeyID).Extract()
	th.AssertNoErr(t, err)

	keyGetDisabled, err := keys.Get(client, createKey.KeyID).ExtractKeyInfo()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, keyGetDisabled.KeyState, "3")
}

func TestKmsEncryptDataLifecycle(t *testing.T) {
	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	if kmsID == "" {
		t.Skip("OS_KMS_ID env var is missing but KMSv1 grant test requires")
	}

	client, err := clients.NewKMSV1Client()
	th.AssertNoErr(t, err)

	kmsOpts := keys.EncryptDataOpts{
		KeyID:     kmsID,
		PlainText: "hello world",
	}

	kmsEncrypt, err := keys.EncryptData(client, kmsOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, kmsEncrypt)

	kmsDecryptOpt := keys.DecryptDataOpts{
		CipherText: kmsEncrypt.CipherText,
	}

	kmsDecrypt, err := keys.DecryptData(client, kmsDecryptOpt)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, kmsDecrypt)
}
