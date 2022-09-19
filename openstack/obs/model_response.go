package obs

import (
	"encoding/xml"
)

// BaseModel defines base model response from OBS
type BaseModel struct {
	StatusCode      int                 `xml:"-"`
	RequestId       string              `xml:"RequestId" json:"request_id"`
	ResponseHeaders map[string][]string `xml:"-"`
}

// Error defines the error property in DeleteObjectsOutput
type Error struct {
	XMLName   xml.Name `xml:"Error"`
	Key       string   `xml:"Key"`
	VersionId string   `xml:"VersionId"`
	Code      string   `xml:"Code"`
	Message   string   `xml:"Message"`
}

// FetchResponse defines the response fetch policy configuration
type FetchResponse struct {
	Status FetchPolicyStatusType `json:"status"`
	Agency string                `json:"agency"`
}

// SetBucketFetchJobResponse defines the response SetBucketFetchJob configuration
type SetBucketFetchJobResponse struct {
	ID   string `json:"id"`
	Wait int    `json:"Wait"`
}

// GetBucketFetchJobResponse defines the response fetch job configuration
type GetBucketFetchJobResponse struct {
	Err    string      `json:"err"`
	Code   string      `json:"code"`
	Status string      `json:"status"`
	Job    JobResponse `json:"job"`
}

// JobResponse defines the response job configuration
type JobResponse struct {
	Bucket           string `json:"bucket"`
	URL              string `json:"url"`
	Host             string `json:"host"`
	Key              string `json:"key"`
	Md5              string `json:"md5"`
	CallBackURL      string `json:"callbackurl"`
	CallBackBody     string `json:"callbackbody"`
	CallBackBodyType string `json:"callbackbodytype"`
	CallBackHost     string `json:"callbackhost"`
	FileType         string `json:"file_type"`
	IgnoreSameKey    bool   `json:"ignore_same_key"`
}
