package antiddos

import (
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, floatingIpId string) (r DeleteResult) {
	url := DeleteURL(client, floatingIpId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200},
	})
	return
}

type GetTaskOpts struct {
	// Task ID (nonnegative integer) character string
	TaskId string `q:"task_id"`
}

type GetTaskOptsBuilder interface {
	ToGetTaskQuery() (string, error)
}

func (opts GetTaskOpts) ToGetTaskQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func GetTask(client *golangsdk.ServiceClient, opts GetTaskOptsBuilder) (r GetTaskResult) {
	url := GetTaskURL(client)
	if opts != nil {
		query, err := opts.ToGetTaskQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func getStructField(v *DDosStatus, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
