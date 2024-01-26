package obs

import (
	"encoding/xml"
)

// ListBucketsInput is the input parameter of ListBuckets function
type ListBucketsInput struct {
	QueryLocation bool
	BucketType    BucketType
}

// ListBucketsOutput is the result of ListBuckets function
type ListBucketsOutput struct {
	BaseModel
	XMLName xml.Name `xml:"ListAllMyBucketsResult"`
	Owner   Owner    `xml:"Owner"`
	Buckets []Bucket `xml:"Buckets>Bucket"`
}

// CreateBucketInput is the input parameter of CreateBucket function
type CreateBucketInput struct {
	BucketLocation
	Bucket                      string           `xml:"-"`
	ACL                         AclType          `xml:"-"`
	StorageClass                StorageClassType `xml:"-"`
	GrantReadId                 string           `xml:"-"`
	GrantWriteId                string           `xml:"-"`
	GrantReadAcpId              string           `xml:"-"`
	GrantWriteAcpId             string           `xml:"-"`
	GrantFullControlId          string           `xml:"-"`
	GrantReadDeliveredId        string           `xml:"-"`
	GrantFullControlDeliveredId string           `xml:"-"`
	Epid                        string           `xml:"-"`
	IsFSFileInterface           bool             `xml:"-"`
	ObjectLockEnabled           bool             `xml:"-"`
}

// SetBucketStoragePolicyInput is the input parameter of SetBucketStoragePolicy function
type SetBucketStoragePolicyInput struct {
	Bucket string `xml:"-"`
	BucketStoragePolicy
}

type getBucketStoragePolicyOutputS3 struct {
	BaseModel
	BucketStoragePolicy
}

// GetBucketStoragePolicyOutput is the result of GetBucketStoragePolicy function
type GetBucketStoragePolicyOutput struct {
	BaseModel
	StorageClass string
}

type getBucketStoragePolicyOutputObs struct {
	BaseModel
	bucketStoragePolicyObs
}

// SetBucketQuotaInput is the input parameter of SetBucketQuota function
type SetBucketQuotaInput struct {
	Bucket string `xml:"-"`
	BucketQuota
}

// GetBucketQuotaOutput is the result of GetBucketQuota function
type GetBucketQuotaOutput struct {
	BaseModel
	BucketQuota
}

// GetBucketStorageInfoOutput is the result of GetBucketStorageInfo function
type GetBucketStorageInfoOutput struct {
	BaseModel
	XMLName      xml.Name `xml:"GetBucketStorageInfoResult"`
	Size         int64    `xml:"Size"`
	ObjectNumber int      `xml:"ObjectNumber"`
}

type getBucketLocationOutputS3 struct {
	BaseModel
	BucketLocation
}
type getBucketLocationOutputObs struct {
	BaseModel
	bucketLocationObs
}

// GetBucketLocationOutput is the result of GetBucketLocation function
type GetBucketLocationOutput struct {
	BaseModel
	Location string `xml:"-"`
}

// GetBucketAclOutput is the result of GetBucketAcl function
type GetBucketAclOutput struct {
	BaseModel
	AccessControlPolicy
}

// SetBucketAclInput is the input parameter of SetBucketAcl function
type SetBucketAclInput struct {
	Bucket string  `xml:"-"`
	ACL    AclType `xml:"-"`
	AccessControlPolicy
}

// SetBucketPolicyInput is the input parameter of SetBucketPolicy function
type SetBucketPolicyInput struct {
	Bucket string
	Policy string
}

// GetBucketPolicyOutput is the result of GetBucketPolicy function
type GetBucketPolicyOutput struct {
	BaseModel
	Policy string `json:"body"`
}

// SetBucketCorsInput is the input parameter of SetBucketCors function
type SetBucketCorsInput struct {
	Bucket string `xml:"-"`
	BucketCors
}

// GetBucketCorsOutput is the result of GetBucketCors function
type GetBucketCorsOutput struct {
	BaseModel
	BucketCors
}

// SetBucketVersioningInput is the input parameter of SetBucketVersioning function
type SetBucketVersioningInput struct {
	Bucket string `xml:"-"`
	BucketVersioningConfiguration
}

// GetBucketVersioningOutput is the result of GetBucketVersioning function
type GetBucketVersioningOutput struct {
	BaseModel
	BucketVersioningConfiguration
}

// SetBucketWebsiteConfigurationInput is the input parameter of SetBucketWebsiteConfiguration function
type SetBucketWebsiteConfigurationInput struct {
	Bucket string `xml:"-"`
	BucketWebsiteConfiguration
}

// GetBucketWebsiteConfigurationOutput is the result of GetBucketWebsiteConfiguration function
type GetBucketWebsiteConfigurationOutput struct {
	BaseModel
	BucketWebsiteConfiguration
}

// GetBucketMetadataInput is the input parameter of GetBucketMetadata function
type GetBucketMetadataInput struct {
	Bucket        string
	Origin        string
	RequestHeader string
}

// SetObjectMetadataInput is the input parameter of SetObjectMetadata function
type SetObjectMetadataInput struct {
	Bucket                  string
	Key                     string
	VersionId               string
	MetadataDirective       MetadataDirectiveType
	CacheControl            string
	ContentDisposition      string
	ContentEncoding         string
	ContentLanguage         string
	ContentType             string
	Expires                 string
	WebsiteRedirectLocation string
	StorageClass            StorageClassType
	Metadata                map[string]string
}

// SetObjectMetadataOutput is the result of SetObjectMetadata function
type SetObjectMetadataOutput struct {
	BaseModel
	MetadataDirective       MetadataDirectiveType
	CacheControl            string
	ContentDisposition      string
	ContentEncoding         string
	ContentLanguage         string
	ContentType             string
	Expires                 string
	WebsiteRedirectLocation string
	StorageClass            StorageClassType
	Metadata                map[string]string
}

// GetBucketMetadataOutput is the result of GetBucketMetadata function
type GetBucketMetadataOutput struct {
	BaseModel
	StorageClass  StorageClassType
	Location      string
	Version       string
	AllowOrigin   string
	AllowMethod   string
	AllowHeader   string
	MaxAgeSeconds int
	ExposeHeader  string
	Epid          string
	FSStatus      FSStatusType
}

// SetBucketLoggingConfigurationInput is the input parameter of SetBucketLoggingConfiguration function
type SetBucketLoggingConfigurationInput struct {
	Bucket string `xml:"-"`
	BucketLoggingStatus
}

// GetBucketLoggingConfigurationOutput is the result of GetBucketLoggingConfiguration function
type GetBucketLoggingConfigurationOutput struct {
	BaseModel
	BucketLoggingStatus
}

// BucketLifecycleConfiguration defines the bucket lifecycle configuration
type BucketLifecycleConfiguration struct {
	XMLName        xml.Name        `xml:"LifecycleConfiguration"`
	LifecycleRules []LifecycleRule `xml:"Rule"`
}

// SetBucketLifecycleConfigurationInput is the input parameter of SetBucketLifecycleConfiguration function
type SetBucketLifecycleConfigurationInput struct {
	Bucket string `xml:"-"`
	BucketLifecycleConfiguration
}

// GetBucketLifecycleConfigurationOutput is the result of GetBucketLifecycleConfiguration function
type GetBucketLifecycleConfigurationOutput struct {
	BaseModel
	BucketLifecycleConfiguration
}

// SetBucketEncryptionInput is the input parameter of SetBucketEncryption function
type SetBucketEncryptionInput struct {
	Bucket string `xml:"-"`
	BucketEncryptionConfiguration
}

// GetBucketEncryptionOutput is the result of GetBucketEncryption function
type GetBucketEncryptionOutput struct {
	BaseModel
	BucketEncryptionConfiguration
}

// SetBucketTaggingInput is the input parameter of SetBucketTagging function
type SetBucketTaggingInput struct {
	Bucket string `xml:"-"`
	BucketTagging
}

// GetBucketTaggingOutput is the result of GetBucketTagging function
type GetBucketTaggingOutput struct {
	BaseModel
	BucketTagging
}

// SetBucketNotificationInput is the input parameter of SetBucketNotification function
type SetBucketNotificationInput struct {
	Bucket string `xml:"-"`
	BucketNotification
}

type getBucketNotificationOutputS3 struct {
	BaseModel
	bucketNotificationS3
}

// GetBucketNotificationOutput is the result of GetBucketNotification function
type GetBucketNotificationOutput struct {
	BaseModel
	BucketNotification
}

// SetBucketReplicationInput is the input parameter of SetBucketReplication function
type SetBucketReplicationInput struct {
	Bucket string `xml:"-"`
	BucketReplicationConfiguration
}

// GetBucketReplicationOutput is the result of GetBucketReplication function
type GetBucketReplicationOutput struct {
	BaseModel
	BucketReplicationConfiguration
}

// SetWORMPolicyInput is the input parameter of SetWORMPolicy function
type SetWORMPolicyInput struct {
	Bucket string `xml:"-"`
	BucketWormPolicy
}

type GetBucketWORMPolicyOutput struct {
	BaseModel
	BucketWormPolicy
}
