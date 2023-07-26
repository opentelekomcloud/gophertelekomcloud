package keypairs

import (
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Name is used to refer to this keypair from other services within this region.
	Name string `json:"name"`
}

// List returns a Pager that allows you to iterate over a collection of KeyPairs.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]KeyPair, error) {
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("os-keypairs"),
		CreatePage: func(r pagination.PageResult) pagination.Page {
			return KeyPairPage{pagination.LinkedPageBase{PageResult: r}}
		},
	}.AllPages()

	if err != nil {
		return nil, err
	}

	allkeypairs, err := ExtractKeyPairs(pages)
	if err != nil {
		return nil, err
	}

	return filterKeyPairs(allkeypairs, opts)
}

// filterKeyPairs used to filter keypairs using name
func filterKeyPairs(keypairs []KeyPair, opts ListOpts) ([]KeyPair, error) {
	var refinedKeypairs []KeyPair
	var matched bool
	m := map[string]interface{}{}

	if opts.Name != "" {
		m["Name"] = opts.Name
	}

	if len(m) > 0 && len(keypairs) > 0 {
		for _, keypair := range keypairs {
			matched = true

			for key, value := range m {
				if sVal := getStructKeyPairField(&keypair, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedKeypairs = append(refinedKeypairs, keypair)
			}
		}

	} else {
		refinedKeypairs = keypairs
	}

	return refinedKeypairs, nil
}

func getStructKeyPairField(v *KeyPair, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
