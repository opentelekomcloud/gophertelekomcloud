package alias

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetAlias(client *golangsdk.ServiceClient, funcURN, aliasName string) (*FuncAliasesResp, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "aliases", aliasName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res *FuncAliasesResp
	err = extract.Into(raw.Body, &res)
	return res, err
}
