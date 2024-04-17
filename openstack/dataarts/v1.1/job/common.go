package job

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const (
	HeaderXLanguage   = "X-Language"
	HeaderContentType = "Content-Type"

	ApplicationJson = "application/json"
)

const (
	clustersEndpoint = "clusters"
	jobEndpoint      = "job"
	cdmEndpoint      = "cdm"
)

type JobId struct {
	JobId string `json:"jobId"`
}

func respToJobId(r *http.Response) (*JobId, error) {
	var res *JobId
	err := extract.Into(r.Body, res)
	return res, err
}
