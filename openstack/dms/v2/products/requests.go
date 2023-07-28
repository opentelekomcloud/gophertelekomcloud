package products

import (
	"fmt"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"net/url"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get products
func Get(client *golangsdk.ServiceClient, engine string) (*GetResponse, error) {
	if len(engine) == 0 {
		return nil, fmt.Errorf("The parameter \"engine\" cannot be empty, it is required.")
	}
	url := getURL(client)
	url = url + "?engine=" + engine

	var rst golangsdk.Result
	_, err := client.Get(url, &rst.Body, nil)
	if err == nil {
		var r GetResponse
		err = rst.ExtractInto(&r)
		return &r, err
	}
	return nil, err
}

type ListOpts struct {
	ProductId string `q:"product_id"`
}

func List(c *golangsdk.ServiceClient, engineType string, opts ListOpts) (*ListResp, error) {
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	raw, err := c.Get(listURL(c, engineType)+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
