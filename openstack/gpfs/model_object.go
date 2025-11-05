package gpfs

import (
	"io"
	"time"
)

// GetObjectMetadataOutput is the result of GetObjectMetadata function
type GetObjectMetadataOutput struct {
	BaseModel
	VersionId               string
	WebsiteRedirectLocation string
	Expiration              string
	Restore                 string
	ObjectType              string
	NextAppendPosition      string
	StorageClass            StorageClassType
	ContentLength           int64
	ContentType             string
	ETag                    string
	AllowOrigin             string
	AllowHeader             string
	AllowMethod             string
	ExposeHeader            string
	MaxAgeSeconds           int
	LastModified            time.Time
	SseHeader               ISseHeader
	Metadata                map[string]string
}

// GetObjectOutput is the result of GetObject function
type GetObjectOutput struct {
	GetObjectMetadataOutput
	DeleteMarker       bool
	CacheControl       string
	ContentDisposition string
	ContentEncoding    string
	ContentLanguage    string
	Expires            string
	Body               io.ReadCloser
}
