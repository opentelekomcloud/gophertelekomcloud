package servers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// MetadatumOpts is a map of length one that contains a key-value pair.
type MetadatumOpts map[string]string

// toMetadatumCreateMap assembles a body for a Create request based on the contents of a MetadataumOpts.
func toMetadatumCreateMap(opts MetadatumOpts) (map[string]interface{}, string, error) {
	if len(opts) != 1 {
		err := golangsdk.ErrInvalidInput{}
		err.Argument = "servers.MetadatumOpts"
		err.Info = "Must have 1 and only 1 key-value pair"
		return nil, "", err
	}
	metadatum := map[string]interface{}{"meta": opts}
	var key string
	for k := range metadatum["meta"].(MetadatumOpts) {
		key = k
	}
	return metadatum, key, nil
}

// CreateMetadatum will create or update the key-value pair with the given key for the given server ID.
func CreateMetadatum(client *golangsdk.ServiceClient, id string, opts MetadatumOpts) (map[string]string, error) {
	b, key, err := toMetadatumCreateMap(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("servers", id, "metadata", key), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraTum(err, raw)
}
