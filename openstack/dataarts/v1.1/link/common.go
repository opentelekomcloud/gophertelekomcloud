package link

const (
	HeaderXLanguage   = "X-Language"
	HeaderContentType = "Content-Type"

	ApplicationJson = "application/json"
)

const (
	clustersEndpoint = "clusters"
	cdmEndpoint      = "cdm"
	linkEndpoint     = "link"
)

type JobId struct {
	JobId string `json:"jobId"`
}
