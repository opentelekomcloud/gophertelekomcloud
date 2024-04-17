package alias

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetAlias(client *golangsdk.ServiceClient, funcURN, aliasName string) (*FuncAliases, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "aliases", aliasName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res *FuncAliases
	err = extract.Into(raw.Body, &res)
	return res, err
}
