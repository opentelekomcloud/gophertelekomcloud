package gpfs

import (
	"errors"
)

// ListBuckets lists buckets.
//
// You can use this API to obtain the bucket list. In the list, bucket names are displayed in lexicographical order.
func (obsClient ObsClient) ListBuckets(input *ListBucketsInput) (output *ListBucketsOutput, err error) {
	if input == nil {
		input = &ListBucketsInput{}
	}
	output = &ListBucketsOutput{}
	err = obsClient.doActionWithoutBucket("ListBuckets", HTTP_GET, input, output)
	if err != nil {
		output = nil
	}
	return
}

// CreateFS creates a GPFS.
//
// You can use this API to create a bucket and name it as you specify. The created bucket name must be unique in OBS.
func (obsClient ObsClient) CreateFS(input *CreateFSInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("CreateFSInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithFS("CreateFS", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteFS deletes a GPFS.
//
// You can use this API to delete a storage FS.
func (obsClient ObsClient) DeleteFS(fsName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithFS("DeleteFS", HTTP_DELETE, fsName, defaultSerializable, output)
	if err != nil {
		output = nil
	}
	return
}
