package instances

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type JobId struct {
	JobId string `json:"job_id"`
}

func extractJob(err error, raw *http.Response) (*string, error) {
	if err != nil {
		return nil, err
	}

	var res JobId
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
