package cluster

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const (
	HeaderXLanguage = "X-Language"
	RequestedLang   = "en"
)

type JobId struct {
	JobId string `json:"jobId"`
}

func respToJobId(r *http.Response) (*JobId, error) {
	var res *JobId
	err := extract.Into(r.Body, res)
	return res, err
}

func respToJobIdSlice(r *http.Response) (*JobId, error) {
	var res *JobId
	err := extract.Into(r.Body, res)
	return res, err
}
