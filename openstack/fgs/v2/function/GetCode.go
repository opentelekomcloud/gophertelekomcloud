package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetCode(client *golangsdk.ServiceClient, funcURN string) (*FuncGraphCode, error) {
	raw, err := client.Get(client.ServiceURL("fgs", "functions", funcURN, "code"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res FuncGraphCode
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type FuncGraphCode struct {
	FuncURN           string         `json:"func_urn"`
	FuncName          string         `json:"func_name"`
	DomainID          string         `json:"domain_id"`
	Runtime           string         `json:"runtime"`
	CodeType          string         `json:"code_type"`
	CodeURL           string         `json:"code_url"`
	CodeFilename      string         `json:"code_filename"`
	CodeSize          int            `json:"code_size"`
	Digest            string         `json:"digest"`
	LastModified      string         `json:"last_modified"`
	FuncCode          FuncCode       `json:"func_code"`
	DependVersionList []string       `json:"depend_version_list"`
	StrategyConfig    StrategyConfig `json:"strategy_config"`
	Dependencies      []Dependency   `json:"dependencies"`
}
