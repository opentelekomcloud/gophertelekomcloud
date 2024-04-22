package function

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ImportOpts struct {
	FuncName string `json:"func_name" required:"true"`
	FileName string `json:"file_name" required:"true"`
	FileType string `json:"file_type" required:"true"`
	FileCode string `json:"file_code" required:"true"`
	Package  string `json:"package,omitempty"`
}

func Import(client *golangsdk.ServiceClient, opts ImportOpts) (*FuncGraph, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "functions", "import"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FuncGraph
	return &res, extract.Into(raw.Body, &res)
}
