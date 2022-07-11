package keys

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	// "github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateOpts struct {
	// Alias of a CMK
	KeyAlias string `json:"key_alias" required:"true"`
	// CMK description
	KeyDescription string `json:"key_description,omitempty"`
	// Region where a CMK resides
	Realm string `json:"realm,omitempty"`
	// Purpose of a CMK (The default value is Encrypt_Decrypt)
	KeyUsage string `json:"key_usage,omitempty"`
}

type DeleteOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// Number of days after which a CMK is scheduled to be deleted
	// (The value ranges from 7 to 1096.)
	PendingDays string `json:"pending_days" required:"true"`
}

type CancelDeleteOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
}

type UpdateAliasOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	KeyAlias string `json:"key_alias" required:"true"`
}

type UpdateDesOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	KeyDescription string `json:"key_description" required:"true"`
}

type DataEncryptOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	EncryptionContext string `json:"encryption_context,omitempty"`
	// 36-byte serial number of a request message
	DatakeyLength string `json:"datakey_length,omitempty"`
}

type EncryptDEKOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// CMK description
	EncryptionContext string `json:"encryption_context,omitempty"`
	// 36-byte serial number of a request message
	DataKeyPlainLength string `json:"datakey_plain_length,omitempty"`
	// Both the plaintext (64 bytes) of a DEK and the SHA-256 hash value (32 bytes)
	// of the plaintext are expressed as a hexadecimal character string.
	PlainText string `json:"plain_text" required:"true"`
}

// ListOpts holds options for listing Volumes. It is passed to the volumes.List
// function.
type ListOpts struct {
	// State of a CMK
	KeyState string `json:"key_state,omitempty"`
	Limit    string `json:"limit,omitempty"`
	Marker   string `json:"marker,omitempty"`
}

type RotationOpts struct {
	// ID of a CMK
	KeyID string `json:"key_id" required:"true"`
	// Rotation interval of a CMK
	Interval int `json:"rotation_interval"`
	// 36-byte serial number of a request message
	Sequence string `json:"sequence,omitempty"`
}

// ToKeyCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToKeyCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ToKeyDeleteMap assembles a request body based on the contents of a
// DeleteOpts.
func (opts DeleteOpts) ToKeyDeleteMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ToKeyCancelDeleteMap assembles a request body based on the contents of a
// CancelDeleteOpts.
func (opts CancelDeleteOpts) ToKeyCancelDeleteMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ToKeyUpdateAliasMap assembles a request body based on the contents of a
// UpdateAliasOpts.
func (opts UpdateAliasOpts) ToKeyUpdateAliasMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ToKeyUpdateDesMap assembles a request body based on the contents of a
// UpdateDesOpts.
func (opts UpdateDesOpts) ToKeyUpdateDesMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func (opts DataEncryptOpts) ToDataEncryptMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func (opts EncryptDEKOpts) ToEncryptDEKMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func (opts ListOpts) ToKeyListMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// ToKeyRotationMap assembles a request body based on the contents of a
// RotationOpts.
func (opts RotationOpts) ToKeyRotationMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

type CreateOptsBuilder interface {
	ToKeyCreateMap() (map[string]interface{}, error)
}

type DeleteOptsBuilder interface {
	ToKeyDeleteMap() (map[string]interface{}, error)
}

type CancelDeleteOptsBuilder interface {
	ToKeyCancelDeleteMap() (map[string]interface{}, error)
}

type UpdateAliasOptsBuilder interface {
	ToKeyUpdateAliasMap() (map[string]interface{}, error)
}

type UpdateDesOptsBuilder interface {
	ToKeyUpdateDesMap() (map[string]interface{}, error)
}

type DataEncryptOptsBuilder interface {
	ToDataEncryptMap() (map[string]interface{}, error)
}

type EncryptDEKOptsBuilder interface {
	ToEncryptDEKMap() (map[string]interface{}, error)
}

type ListOptsBuilder interface {
	ToKeyListMap() (map[string]interface{}, error)
}

type RotationOptsBuilder interface {
	ToKeyRotationMap() (map[string]interface{}, error)
}

// Create will create a new key based on the values in CreateOpts. To ExtractKeyInfo
// the key object from the response, call the ExtractKeyInfo method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToKeyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get retrieves the key with the provided ID. To extract the key object
// from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	b := map[string]interface{}{"key_id": id}
	_, r.Err = client.Post(getURL(client), &b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will delete the existing key with the provided ID.
func Delete(client *golangsdk.ServiceClient, opts DeleteOptsBuilder) (r DeleteResult) {
	b, err := opts.ToKeyDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(deleteURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: &r.Body,
	})
	return
}

func UpdateAlias(client *golangsdk.ServiceClient, opts UpdateAliasOptsBuilder) (r UpdateAliasResult) {
	b, err := opts.ToKeyUpdateAliasMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(updateAliasURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func UpdateDes(client *golangsdk.ServiceClient, opts UpdateDesOptsBuilder) (r UpdateDesResult) {
	b, err := opts.ToKeyUpdateDesMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(updateDesURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func DataEncryptGet(client *golangsdk.ServiceClient, opts DataEncryptOptsBuilder) (r DataEncryptResult) {
	b, err := opts.ToDataEncryptMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(dataEncryptURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func DataEncryptGetWithoutPlaintext(client *golangsdk.ServiceClient, opts DataEncryptOptsBuilder) (r DataEncryptResult) {
	b, err := opts.ToDataEncryptMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(dataEncryptURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func EncryptDEKGet(client *golangsdk.ServiceClient, opts EncryptDEKOptsBuilder) (r EncryptDEKResult) {
	b, err := opts.ToEncryptDEKMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(encryptDEKURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func EnableKey(client *golangsdk.ServiceClient, id string) (r ExtractUpdateKeyStateResult) {
	b := map[string]interface{}{"key_id": id}
	_, r.Err = client.Post(enableKeyURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func DisableKey(client *golangsdk.ServiceClient, id string) (r ExtractUpdateKeyStateResult) {
	b := map[string]interface{}{"key_id": id}
	_, r.Err = client.Post(disableKeyURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) (r ListResult) {
	b, err := opts.ToKeyListMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(listURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func ListAllKeys(client *golangsdk.ServiceClient, opts ListOptsBuilder) (r ListResult) {
	b, err := opts.ToKeyListMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(listURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func EnableKeyRotation(client *golangsdk.ServiceClient, opts RotationOptsBuilder) (r golangsdk.ErrResult) {
	b, err := opts.ToKeyRotationMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(enableKeyRotationURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func DisableKeyRotation(client *golangsdk.ServiceClient, opts RotationOptsBuilder) (r golangsdk.ErrResult) {
	b, err := opts.ToKeyRotationMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(disableKeyRotationURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func GetKeyRotationStatus(client *golangsdk.ServiceClient, opts RotationOptsBuilder) (r GetRotationResult) {
	b, err := opts.ToKeyRotationMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(getKeyRotationStatusURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

func UpdateKeyRotationInterval(client *golangsdk.ServiceClient, opts RotationOptsBuilder) (r golangsdk.ErrResult) {
	b, err := opts.ToKeyRotationMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(updateKeyRotationIntervalURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// CancelDelete will cancel the scheduled deletion for a CMK only when the CMK's status is Scheduled deletion with the provided ID.
func CancelDelete(client *golangsdk.ServiceClient, opts CancelDeleteOptsBuilder) (r DeleteResult) {
	b, err := opts.ToKeyCancelDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(cancelDeleteURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: &r.Body,
	})
	return
}
