package flavors

import "github.com/opentelekomcloud/gophertelekomcloud"

// toFlavorExtraSpecUpdateMap assembles a body for an Update request based on the contents of a ExtraSpecOpts.
func toFlavorExtraSpecUpdateMap(opts map[string]string) (map[string]string, string, error) {
	if len(opts) != 1 {
		err := golangsdk.ErrInvalidInput{}
		err.Argument = "flavors.ExtraSpecOpts"
		err.Info = "Must have 1 and only one key-value pair"
		return nil, "", err
	}

	var key string
	for k := range opts {
		key = k
	}

	return opts, key, nil
}

// UpdateExtraSpec will update the value of the specified flavor's extraAcc spec for the key in opts.
func UpdateExtraSpec(client *golangsdk.ServiceClient, flavorID string, opts map[string]string) (map[string]string, error) {
	b, key, err := toFlavorExtraSpecUpdateMap(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("flavors", flavorID, "os-extra_specs", key), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extraSpe(err, raw)
}
