package products

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
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
	url := listURL(c, engineType)
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}
	url += query.String()

	var r ListResp
	_, err = c.Get(url, &r, nil)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
