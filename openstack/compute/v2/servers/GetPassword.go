package servers

import (
	"crypto/rsa"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetPassword makes a request against the nova API to get the encrypted administrative password.
// ExtractPassword gets the encrypted password.
// If privateKey != nil the password is decrypted with the private key.
// If privateKey == nil the encrypted password is returned and can be decrypted
// with:
//
//	echo '<pwd>' | base64 -D | openssl rsautl -decrypt -inkey <private_key>
func GetPassword(client *golangsdk.ServiceClient, serverId string, privateKey *rsa.PrivateKey) (string, error) {
	raw, err := client.Get(client.ServiceURL("servers", serverId, "os-server-password"), nil, nil)
	if err != nil {
		return "", err
	}

	var res struct {
		Password string `json:"password"`
	}
	err = extract.Into(raw.Body, &res)
	if err == nil && privateKey != nil && res.Password != "" {
		return decryptPassword(res.Password, privateKey)
	}
	return res.Password, err
}
