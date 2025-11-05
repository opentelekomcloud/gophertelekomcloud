package gpfs

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

// CreateFSInput is the input parameter of CreateBucket function
type CreateFSInput struct {
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
