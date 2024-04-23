package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateFuncCodeOpts struct {
	FuncUrn           string    `json:"-"`
	CodeType          string    `json:"code_type" required:"true"`
	CodeURL           string    `json:"code_url,omitempty"`
	CodeFilename      string    `json:"code_filename,omitempty"`
	FuncCode          *FuncCode `json:"func_code,omitempty"`
	DependVersionList []string  `json:"depend_version_list,omitempty"`
}

func UpdateFuncCode(client *golangsdk.ServiceClient, opts UpdateFuncCodeOpts) (*FuncGraphCode, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("fgs", "functions", opts.FuncUrn, "code"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncGraphCode
	return &res, extract.Into(raw.Body, &res)
}
