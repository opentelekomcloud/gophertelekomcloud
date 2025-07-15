package v1

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"regexp"
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
	createKey, err := keys.Create(client, createOpts)
	th.AssertNoErr(t, err)

	defer func() {
		deleteOpts := keys.DeleteOpts{
			KeyID:       createKey.KeyID,
			PendingDays: "7",
		}
		_, err := keys.Delete(client, deleteOpts)
		th.AssertNoErr(t, err)
	}()

	keyGet, err := keys.Get(client, createKey.KeyID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.KeyAlias, keyGet.KeyAlias)
	th.AssertEquals(t, keyGet.KeyState, "4")

	_, err = keys.CancelDelete(client, createKey.KeyID)
	th.AssertNoErr(t, err)

	_, err = keys.EnableKey(client, createKey.KeyID)
	th.AssertNoErr(t, err)

	keyGetEnabled, err := keys.Get(client, createKey.KeyID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, keyGetEnabled.KeyState, "2")

	_, err = keys.DisableKey(client, createKey.KeyID)
	th.AssertNoErr(t, err)

	keyGetDisabled, err := keys.Get(client, createKey.KeyID)
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

// Test requries an external kms key in `pending import` state
// and can't be run more than once for the key or an error will occur:
// `The plain key value must be same with imported before when try to re-import a deleted key material`
func TestKmsCMKImportLifecycle(t *testing.T) {
	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	if kmsID == "" {
		t.Skip("OS_KMS_ID env var is missing but KMSv1 grant test requires")
	}

	client, err := clients.NewKMSV1Client()
	th.AssertNoErr(t, err)

	getOpts := keys.GetCMKImportOpts{
		KeyId:             kmsID,
		WrappingAlgorithm: "RSAES_PKCS1_V1_5",
	}

	getResp, err := keys.GetCMKImport(client, getOpts)
	th.AssertNoErr(t, err)

	publicKeyBytes, err := base64.StdEncoding.DecodeString(getResp.PublicKey)
	th.AssertNoErr(t, err)

	pubKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	th.AssertNoErr(t, err)

	rsaPublicKey, ok := pubKey.(*rsa.PublicKey)
	th.AssertEquals(t, ok, true)

	keyMaterial := make([]byte, 32)
	_, err = rand.Read(keyMaterial)
	th.AssertNoErr(t, err)

	encryptedKeyMaterial, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		rsaPublicKey,
		keyMaterial,
	)
	th.AssertNoErr(t, err)

	publicMaterial := base64.StdEncoding.EncodeToString(encryptedKeyMaterial)

	matched, err := regexp.MatchString("^[0-9a-zA-Z+/=]{344,360}$", publicMaterial)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, matched, true)

	kmsOpts := keys.ImportCMKOpts{
		KeyId:                kmsID,
		ImportToken:          getResp.ImportToken,
		EncryptedKeyMaterial: publicMaterial,
	}

	err = keys.ImportCMKMaterial(client, kmsOpts)
	th.AssertNoErr(t, err)

	err = keys.DeleteCMKImport(client, keys.DeleteCMKImportOpts{
		KeyId: kmsID,
	})
	th.AssertNoErr(t, err)
}
