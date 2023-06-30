package obs

import (
	"encoding/xml"
	"io"
	"time"
)

// ListMultipartUploadsInput is the input parameter of ListMultipartUploads function
type ListMultipartUploadsInput struct {
	Bucket         string
	Prefix         string
	MaxUploads     int
	Delimiter      string
	KeyMarker      string
	UploadIdMarker string
}

// ListMultipartUploadsOutput is the result of ListMultipartUploads function
type ListMultipartUploadsOutput struct {
	BaseModel
	XMLName            xml.Name `xml:"ListMultipartUploadsResult"`
	Bucket             string   `xml:"Bucket"`
	KeyMarker          string   `xml:"KeyMarker"`
	NextKeyMarker      string   `xml:"NextKeyMarker"`
	UploadIdMarker     string   `xml:"UploadIdMarker"`
	NextUploadIdMarker string   `xml:"NextUploadIdMarker"`
	Delimiter          string   `xml:"Delimiter"`
	IsTruncated        bool     `xml:"IsTruncated"`
	MaxUploads         int      `xml:"MaxUploads"`
	Prefix             string   `xml:"Prefix"`
	Uploads            []Upload `xml:"Upload"`
	CommonPrefixes     []string `xml:"CommonPrefixes>Prefix"`
}

// AbortMultipartUploadInput is the input parameter of AbortMultipartUpload function
type AbortMultipartUploadInput struct {
	Bucket   string
	Key      string
	UploadId string
}

// InitiateMultipartUploadInput is the input parameter of InitiateMultipartUpload function
type InitiateMultipartUploadInput struct {
	ObjectOperationInput
	ContentType string
}

// InitiateMultipartUploadOutput is the result of InitiateMultipartUpload function
type InitiateMultipartUploadOutput struct {
	BaseModel
	XMLName   xml.Name `xml:"InitiateMultipartUploadResult"`
	Bucket    string   `xml:"Bucket"`
	Key       string   `xml:"Key"`
	UploadId  string   `xml:"UploadId"`
	SseHeader ISseHeader
}

// UploadPartInput is the input parameter of UploadPart function
type UploadPartInput struct {
	Bucket     string
	Key        string
	PartNumber int
	UploadId   string
	ContentMD5 string
	SseHeader  ISseHeader
	Body       io.Reader
	SourceFile string
	Offset     int64
	PartSize   int64
}

// UploadPartOutput is the result of UploadPart function
type UploadPartOutput struct {
	BaseModel
	PartNumber int
	ETag       string
	SseHeader  ISseHeader
}

// CompleteMultipartUploadInput is the input parameter of CompleteMultipartUpload function
type CompleteMultipartUploadInput struct {
	Bucket   string   `xml:"-"`
	Key      string   `xml:"-"`
	UploadId string   `xml:"-"`
	XMLName  xml.Name `xml:"CompleteMultipartUpload"`
	Parts    []Part   `xml:"Part"`
}

// CompleteMultipartUploadOutput is the result of CompleteMultipartUpload function
type CompleteMultipartUploadOutput struct {
	BaseModel
	VersionId string     `xml:"-"`
	SseHeader ISseHeader `xml:"-"`
	XMLName   xml.Name   `xml:"CompleteMultipartUploadResult"`
	Location  string     `xml:"Location"`
	Bucket    string     `xml:"Bucket"`
	Key       string     `xml:"Key"`
	ETag      string     `xml:"ETag"`
}

// ListPartsInput is the input parameter of ListParts function
type ListPartsInput struct {
	Bucket           string
	Key              string
	UploadId         string
	MaxParts         int
	PartNumberMarker int
}

// ListPartsOutput is the result of ListParts function
type ListPartsOutput struct {
	BaseModel
	XMLName              xml.Name         `xml:"ListPartsResult"`
	Bucket               string           `xml:"Bucket"`
	Key                  string           `xml:"Key"`
	UploadId             string           `xml:"UploadId"`
	PartNumberMarker     int              `xml:"PartNumberMarker"`
	NextPartNumberMarker int              `xml:"NextPartNumberMarker"`
	MaxParts             int              `xml:"MaxParts"`
	IsTruncated          bool             `xml:"IsTruncated"`
	StorageClass         StorageClassType `xml:"StorageClass"`
	Initiator            Initiator        `xml:"Initiator"`
	Owner                Owner            `xml:"Owner"`
	Parts                []Part           `xml:"Part"`
}

// CopyPartInput is the input parameter of CopyPart function
type CopyPartInput struct {
	Bucket               string
	Key                  string
	UploadId             string
	PartNumber           int
	CopySourceBucket     string
	CopySourceKey        string
	CopySourceVersionId  string
	CopySourceRangeStart int64
	CopySourceRangeEnd   int64
	SseHeader            ISseHeader
	SourceSseHeader      ISseHeader
}

// CopyPartOutput is the result of CopyPart function
type CopyPartOutput struct {
	BaseModel
	XMLName      xml.Name   `xml:"CopyPartResult"`
	PartNumber   int        `xml:"-"`
	ETag         string     `xml:"ETag"`
	LastModified time.Time  `xml:"LastModified"`
	SseHeader    ISseHeader `xml:"-"`
}
