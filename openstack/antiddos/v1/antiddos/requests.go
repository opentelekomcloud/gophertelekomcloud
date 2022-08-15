package antiddos

import (
	"reflect"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type CreateOpts struct {
	// Whether to enable L7 defense
	EnableL7 bool `json:"enable_L7,"`

	// Position ID of traffic. The value ranges from 1 to 9.
	TrafficPosId int `json:"traffic_pos_id,"`

	// Position ID of number of HTTP requests. The value ranges from 1 to 15.
	HttpRequestPosId int `json:"http_request_pos_id,"`

	// Position ID of access limit during cleaning. The value ranges from 1 to 8.
	CleaningAccessPosId int `json:"cleaning_access_pos_id,"`

	// Application type ID. Possible values: 0 1
	AppTypeId int `json:"app_type_id,"`
}

type CreateOptsBuilder interface {
	ToCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToCreateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *golangsdk.ServiceClient, floatingIpId string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client, floatingIpId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(client *golangsdk.ServiceClient, floatingIpId string) (r DeleteResult) {
	url := DeleteURL(client, floatingIpId)
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		JSONResponse: &r.Body,
		OkCodes:      []int{200},
	})
	return
}

func GetStatus(client *golangsdk.ServiceClient, floatingIpId string) (r GetStatusResult) {
	url := GetStatusURL(client, floatingIpId)
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
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

type UpdateOpts struct {
	// Whether to enable L7 defense
	EnableL7 bool `json:"enable_L7,"`

	// Position ID of traffic. The value ranges from 1 to 9.
	TrafficPosId int `json:"traffic_pos_id,"`

	// Position ID of number of HTTP requests. The value ranges from 1 to 15.
	HttpRequestPosId int `json:"http_request_pos_id,"`

	// Position ID of access limit during cleaning. The value ranges from 1 to 8.
	CleaningAccessPosId int `json:"cleaning_access_pos_id,"`

	// Application type ID. Possible values: 0 1
	AppTypeId int `json:"app_type_id,"`
}

type UpdateOptsBuilder interface {
	ToUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToUpdateMap() (map[string]interface{}, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *golangsdk.ServiceClient, floatingIpId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(UpdateURL(client, floatingIpId), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
