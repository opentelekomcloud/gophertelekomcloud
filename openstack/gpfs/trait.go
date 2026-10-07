package gpfs

import (
	"encoding/json"
	"errors"
)

type responseModel interface {
	setStatusCode(int)
	setRequestID(string)
	setResponseHeaders(map[string][]string)
}

type serializable interface {
	serialize() (map[string]string, map[string][]string, interface{}, error)
}

type defaultSerializable struct {
	params map[string]string
}

func (input defaultSerializable) serialize() (map[string]string, map[string][]string, interface{}, error) {
	return input.params, nil, nil, nil
}

var emptyInput = &defaultSerializable{}

func newSubResourceInput(name string) *defaultSerializable {
	return &defaultSerializable{params: map[string]string{name: ""}}
}

func (model *BaseModel) setStatusCode(value int)                      { model.StatusCode = value }
func (model *BaseModel) setRequestID(value string)                    { model.RequestId = value }
func (model *BaseModel) setResponseHeaders(value map[string][]string) { model.ResponseHeaders = value }

func (input ListFSInput) serialize() (map[string]string, map[string][]string, interface{}, error) {
	if input.BucketType == "" {
		return nil, nil, nil, errors.New("BucketType is a required parameter")
	}
	return nil, map[string][]string{headerPrefixOBS + headerBucketType: {input.BucketType}}, nil, nil
}

func (input CreateFSInput) serialize() (map[string]string, map[string][]string, interface{}, error) {
	if input.Redundancy == "" {
		return nil, nil, nil, errors.New("Redundancy is a required parameter")
	}
	if input.BucketType == "" {
		return nil, nil, nil, errors.New("BucketType is a required parameter")
	}
	headers := map[string][]string{
		headerPrefixOBS + headerAZRedundancy: {input.Redundancy},
		headerPrefixOBS + headerBucketType:   {input.BucketType},
	}
	if input.Location == "" {
		return nil, headers, nil, nil
	}
	body, err := marshalXML(input.BucketLocation)
	return nil, headers, body, err
}

func (input CreateFSAccessRulesInput) serialize() (map[string]string, map[string][]string, interface{}, error) {
	body, err := json.Marshal(input)
	return map[string]string{subResourceSFSACL: ""}, map[string][]string{headerContentType: {"application/json"}}, body, err
}
