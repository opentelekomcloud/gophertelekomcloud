package obs

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type ListObjsInput struct {
	Prefix        string
	MaxKeys       int
	Delimiter     string
	Origin        string
	RequestHeader string
}

type ListObjectsInput struct {
	ListObjsInput
	Bucket string
	Marker string
}

type ListObjectsOutput struct {
	BaseModel
	XMLName        xml.Name  `xml:"ListBucketResult"`
	Delimiter      string    `xml:"Delimiter"`
	IsTruncated    bool      `xml:"IsTruncated"`
	Marker         string    `xml:"Marker"`
	NextMarker     string    `xml:"NextMarker"`
	MaxKeys        int       `xml:"MaxKeys"`
	Name           string    `xml:"Name"`
	Prefix         string    `xml:"Prefix"`
	Contents       []Content `xml:"Contents"`
	CommonPrefixes []string  `xml:"CommonPrefixes>Prefix"`
	Location       string    `xml:"-"`
}

type ListVersionsInput struct {
	ListObjsInput
	Bucket          string
	KeyMarker       string
	VersionIdMarker string
}

type ListVersionsOutput struct {
	BaseModel
	XMLName             xml.Name       `xml:"ListVersionsResult"`
	Delimiter           string         `xml:"Delimiter"`
	IsTruncated         bool           `xml:"IsTruncated"`
	KeyMarker           string         `xml:"KeyMarker"`
	NextKeyMarker       string         `xml:"NextKeyMarker"`
	VersionIdMarker     string         `xml:"VersionIdMarker"`
	NextVersionIdMarker string         `xml:"NextVersionIdMarker"`
	MaxKeys             int            `xml:"MaxKeys"`
	Name                string         `xml:"Name"`
	Prefix              string         `xml:"Prefix"`
	Versions            []Version      `xml:"Version"`
	DeleteMarkers       []DeleteMarker `xml:"DeleteMarker"`
	CommonPrefixes      []string       `xml:"CommonPrefixes>Prefix"`
	Location            string         `xml:"-"`
}

type ListMultipartUploadsInput struct {
	Bucket         string
	Prefix         string
	MaxUploads     int
	Delimiter      string
	KeyMarker      string
	UploadIdMarker string
}

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

type getBucketAclOutputObs struct {
	BaseModel
	accessControlPolicyObs
}

type DeleteObjectInput struct {
	Bucket    string
	Key       string
	VersionId string
}

type DeleteObjectOutput struct {
	BaseModel
	VersionId    string
	DeleteMarker bool
}

type DeleteObjectsInput struct {
	Bucket  string           `xml:"-"`
	XMLName xml.Name         `xml:"Delete"`
	Quiet   bool             `xml:"Quiet,omitempty"`
	Objects []ObjectToDelete `xml:"Object"`
}

type Error struct {
	XMLName   xml.Name `xml:"Error"`
	Key       string   `xml:"Key"`
	VersionId string   `xml:"VersionId"`
	Code      string   `xml:"Code"`
	Message   string   `xml:"Message"`
}

type DeleteObjectsOutput struct {
	BaseModel
	XMLName  xml.Name  `xml:"DeleteResult"`
	Deleteds []Deleted `xml:"Deleted"`
	Errors   []Error   `xml:"Error"`
}

type SetObjectAclInput struct {
	Bucket    string  `xml:"-"`
	Key       string  `xml:"-"`
	VersionId string  `xml:"-"`
	ACL       AclType `xml:"-"`
	AccessControlPolicy
}

type GetObjectAclInput struct {
	Bucket    string
	Key       string
	VersionId string
}

type GetObjectAclOutput struct {
	BaseModel
	VersionId string
	AccessControlPolicy
}

type RestoreObjectInput struct {
	Bucket    string          `xml:"-"`
	Key       string          `xml:"-"`
	VersionId string          `xml:"-"`
	XMLName   xml.Name        `xml:"RestoreRequest"`
	Days      int             `xml:"Days"`
	Tier      RestoreTierType `xml:"GlacierJobParameters>Tier,omitempty"`
}

type GetObjectMetadataInput struct {
	Bucket        string
	Key           string
	VersionId     string
	Origin        string
	RequestHeader string
	SseHeader     ISseHeader
}

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

type GetObjectInput struct {
	GetObjectMetadataInput
	IfMatch                    string
	IfNoneMatch                string
	IfUnmodifiedSince          time.Time
	IfModifiedSince            time.Time
	RangeStart                 int64
	RangeEnd                   int64
	ImageProcess               string
	ResponseCacheControl       string
	ResponseContentDisposition string
	ResponseContentEncoding    string
	ResponseContentLanguage    string
	ResponseContentType        string
	ResponseExpires            string
}

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

type ObjectOperationInput struct {
	Bucket                  string
	Key                     string
	ACL                     AclType
	GrantReadId             string
	GrantReadAcpId          string
	GrantWriteAcpId         string
	GrantFullControlId      string
	StorageClass            StorageClassType
	WebsiteRedirectLocation string
	Expires                 int64
	SseHeader               ISseHeader
	Metadata                map[string]string
}

type PutObjectBasicInput struct {
	ObjectOperationInput
	ContentType   string
	ContentMD5    string
	ContentLength int64
}

type PutObjectInput struct {
	PutObjectBasicInput
	Body io.Reader
}

type PutFileInput struct {
	PutObjectBasicInput
	SourceFile string
}

type PutObjectOutput struct {
	BaseModel
	VersionId    string
	SseHeader    ISseHeader
	StorageClass StorageClassType
	ETag         string
}

type CopyObjectInput struct {
	ObjectOperationInput
	CopySourceBucket            string
	CopySourceKey               string
	CopySourceVersionId         string
	CopySourceIfMatch           string
	CopySourceIfNoneMatch       string
	CopySourceIfUnmodifiedSince time.Time
	CopySourceIfModifiedSince   time.Time
	SourceSseHeader             ISseHeader
	CacheControl                string
	ContentDisposition          string
	ContentEncoding             string
	ContentLanguage             string
	ContentType                 string
	Expires                     string
	MetadataDirective           MetadataDirectiveType
	SuccessActionRedirect       string
}

type CopyObjectOutput struct {
	BaseModel
	CopySourceVersionId string     `xml:"-"`
	VersionId           string     `xml:"-"`
	SseHeader           ISseHeader `xml:"-"`
	XMLName             xml.Name   `xml:"CopyObjectResult"`
	LastModified        time.Time  `xml:"LastModified"`
	ETag                string     `xml:"ETag"`
}

type AbortMultipartUploadInput struct {
	Bucket   string
	Key      string
	UploadId string
}

type InitiateMultipartUploadInput struct {
	ObjectOperationInput
	ContentType string
}

type InitiateMultipartUploadOutput struct {
	BaseModel
	XMLName   xml.Name `xml:"InitiateMultipartUploadResult"`
	Bucket    string   `xml:"Bucket"`
	Key       string   `xml:"Key"`
	UploadId  string   `xml:"UploadId"`
	SseHeader ISseHeader
}

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

type UploadPartOutput struct {
	BaseModel
	PartNumber int
	ETag       string
	SseHeader  ISseHeader
}

type CompleteMultipartUploadInput struct {
	Bucket   string   `xml:"-"`
	Key      string   `xml:"-"`
	UploadId string   `xml:"-"`
	XMLName  xml.Name `xml:"CompleteMultipartUpload"`
	Parts    []Part   `xml:"Part"`
}

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

type ListPartsInput struct {
	Bucket           string
	Key              string
	UploadId         string
	MaxParts         int
	PartNumberMarker int
}

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

type CopyPartOutput struct {
	BaseModel
	XMLName      xml.Name   `xml:"CopyPartResult"`
	PartNumber   int        `xml:"-"`
	ETag         string     `xml:"ETag"`
	LastModified time.Time  `xml:"LastModified"`
	SseHeader    ISseHeader `xml:"-"`
}

type CreateSignedUrlInput struct {
	Method      HttpMethodType
	Bucket      string
	Key         string
	SubResource SubResourceType
	Expires     int
	Headers     map[string]string
	QueryParams map[string]string
}

type CreateSignedUrlOutput struct {
	SignedUrl                  string
	ActualSignedRequestHeaders http.Header
}

type CreateBrowserBasedSignatureInput struct {
	Bucket     string
	Key        string
	Expires    int
	FormParams map[string]string
}

type CreateBrowserBasedSignatureOutput struct {
	OriginPolicy string
	Policy       string
	Algorithm    string
	Credential   string
	Date         string
	Signature    string
}
