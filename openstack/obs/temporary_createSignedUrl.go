package obs

import (
	"errors"
)

// CreateSignedUrl creates signed url with the specified CreateSignedUrlInput, and returns the CreateSignedUrlOutput and error
func (obsClient ObsClient) CreateSignedUrl(input *CreateSignedUrlInput) (output *CreateSignedUrlOutput, err error) {
	if input == nil {
		return nil, errors.New("CreateSignedUrlInput is nil")
	}

	params := make(map[string]string, len(input.QueryParams))
	for key, value := range input.QueryParams {
		params[key] = value
	}

	if input.SubResource != "" {
		params[string(input.SubResource)] = ""
	}

	headers := make(map[string][]string, len(input.Headers))
	for key, value := range input.Headers {
		headers[key] = []string{value}
	}

	if input.Expires <= 0 {
		input.Expires = 300
	}

	requestURL, err := obsClient.doAuthTemporary(string(input.Method), input.Bucket, input.Key, params, headers, int64(input.Expires))
	if err != nil {
		return nil, err
	}

	output = &CreateSignedUrlOutput{
		SignedUrl:                  requestURL,
		ActualSignedRequestHeaders: headers,
	}
	return
}
