package backups

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Job struct {
	JobId    string `json:"job_id"`
	BackupId string `json:"backup_id"`
}

func extractJob(err error, raw *http.Response) (*Job, error) {
	if err != nil {
		return nil, err
	}

	var res Job
	err = extract.Into(raw.Body, &res)
	return &res, err
}
