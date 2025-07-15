package job

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
