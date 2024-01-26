package obs

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

// CreateBucket creates a bucket.
//
// You can use this API to create a bucket and name it as you specify. The created bucket name must be unique in OBS.
func (obsClient ObsClient) CreateBucket(input *CreateBucketInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("CreateBucketInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("CreateBucket", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucket deletes a bucket.
//
// You can use this API to delete a bucket. The bucket to be deleted must be empty
// (containing no objects, noncurrent object versions, or part fragments).
func (obsClient ObsClient) DeleteBucket(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucket", HTTP_DELETE, bucketName, defaultSerializable, output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketStoragePolicy sets bucket storage class.
//
// You can use this API to set storage class for bucket.
func (obsClient ObsClient) SetBucketStoragePolicy(input *SetBucketStoragePolicyInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketStoragePolicyInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketStoragePolicy", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

func (obsClient ObsClient) getBucketStoragePolicyS3(bucketName string) (output *GetBucketStoragePolicyOutput, err error) {
	output = &GetBucketStoragePolicyOutput{}
	var outputS3 = &getBucketStoragePolicyOutputS3{}
	err = obsClient.doActionWithBucket("GetBucketStoragePolicy", HTTP_GET, bucketName, newSubResourceSerial(SubResourceStoragePolicy), outputS3)
	if err != nil {
		output = nil
		return
	}
	output.BaseModel = outputS3.BaseModel
	output.StorageClass = string(outputS3.StorageClass)
	return
}

func (obsClient ObsClient) getBucketStoragePolicyObs(bucketName string) (output *GetBucketStoragePolicyOutput, err error) {
	output = &GetBucketStoragePolicyOutput{}
	var outputObs = &getBucketStoragePolicyOutputObs{}
	err = obsClient.doActionWithBucket("GetBucketStoragePolicy", HTTP_GET, bucketName, newSubResourceSerial(SubResourceStorageClass), outputObs)
	if err != nil {
		output = nil
		return
	}
	output.BaseModel = outputObs.BaseModel
	output.StorageClass = outputObs.StorageClass
	return
}

// GetBucketStoragePolicy gets bucket storage class.
//
// You can use this API to obtain the storage class of a bucket.
func (obsClient ObsClient) GetBucketStoragePolicy(bucketName string) (output *GetBucketStoragePolicyOutput, err error) {
	if obsClient.conf.signature == SignatureObs {
		return obsClient.getBucketStoragePolicyObs(bucketName)
	}
	return obsClient.getBucketStoragePolicyS3(bucketName)
}

// SetBucketQuota sets the bucket quota.
//
// You can use this API to set the bucket quota. A bucket quota must be expressed in bytes and the maximum value is 2^63-1.
func (obsClient ObsClient) SetBucketQuota(input *SetBucketQuotaInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketQuotaInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketQuota", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketQuota gets the bucket quota.
//
// You can use this API to obtain the bucket quota. Value 0 indicates that no upper limit is set for the bucket quota.
func (obsClient ObsClient) GetBucketQuota(bucketName string) (output *GetBucketQuotaOutput, err error) {
	output = &GetBucketQuotaOutput{}
	err = obsClient.doActionWithBucket("GetBucketQuota", HTTP_GET, bucketName, newSubResourceSerial(SubResourceQuota), output)
	if err != nil {
		output = nil
	}
	return
}

// HeadBucket checks whether a bucket exists.
//
// You can use this API to check whether a bucket exists.
func (obsClient ObsClient) HeadBucket(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("HeadBucket", HTTP_HEAD, bucketName, defaultSerializable, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketMetadata gets the metadata of a bucket.
//
// You can use this API to send a HEAD request to a bucket to obtain the bucket
// metadata such as the storage class and CORS rules (if set).
func (obsClient ObsClient) GetBucketMetadata(input *GetBucketMetadataInput) (output *GetBucketMetadataOutput, err error) {
	output = &GetBucketMetadataOutput{}
	err = obsClient.doActionWithBucket("GetBucketMetadata", HTTP_HEAD, input.Bucket, input, output)
	if err != nil {
		output = nil
	} else {
		ParseGetBucketMetadataOutput(output)
	}
	return
}

// GetBucketStorageInfo gets storage information about a bucket.
//
// You can use this API to obtain storage information about a bucket, including the
// bucket size and number of objects in the bucket.
func (obsClient ObsClient) GetBucketStorageInfo(bucketName string) (output *GetBucketStorageInfoOutput, err error) {
	output = &GetBucketStorageInfoOutput{}
	err = obsClient.doActionWithBucket("GetBucketStorageInfo", HTTP_GET, bucketName, newSubResourceSerial(SubResourceStorageInfo), output)
	if err != nil {
		output = nil
	}
	return
}

func (obsClient ObsClient) getBucketLocationS3(bucketName string) (output *GetBucketLocationOutput, err error) {
	output = &GetBucketLocationOutput{}
	var outputS3 = &getBucketLocationOutputS3{}
	err = obsClient.doActionWithBucket("GetBucketLocation", HTTP_GET, bucketName, newSubResourceSerial(SubResourceLocation), outputS3)
	if err != nil {
		output = nil
	} else {
		output.BaseModel = outputS3.BaseModel
		output.Location = outputS3.Location
	}
	return
}

func (obsClient ObsClient) getBucketLocationObs(bucketName string) (output *GetBucketLocationOutput, err error) {
	output = &GetBucketLocationOutput{}
	var outputObs = &getBucketLocationOutputObs{}
	err = obsClient.doActionWithBucket("GetBucketLocation", HTTP_GET, bucketName, newSubResourceSerial(SubResourceLocation), outputObs)
	if err != nil {
		output = nil
	} else {
		output.BaseModel = outputObs.BaseModel
		output.Location = outputObs.Location
	}
	return
}

// GetBucketLocation gets the location of a bucket.
//
// You can use this API to obtain the bucket location.
func (obsClient ObsClient) GetBucketLocation(bucketName string) (output *GetBucketLocationOutput, err error) {
	if obsClient.conf.signature == SignatureObs {
		return obsClient.getBucketLocationObs(bucketName)
	}
	return obsClient.getBucketLocationS3(bucketName)
}

// SetBucketAcl sets the bucket ACL.
//
// You can use this API to set the ACL for a bucket.
func (obsClient ObsClient) SetBucketAcl(input *SetBucketAclInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketAclInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketAcl", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

func (obsClient ObsClient) getBucketACLObs(bucketName string) (output *GetBucketAclOutput, err error) {
	output = &GetBucketAclOutput{}
	var outputObs = &GetBucketAclOutput{}
	err = obsClient.doActionWithBucket("GetBucketAcl", HTTP_GET, bucketName, newSubResourceSerial(SubResourceAcl), outputObs)
	if err != nil {
		output = nil
	} else {
		output.BaseModel = outputObs.BaseModel
		output.Owner = outputObs.Owner
		output.Grants = make([]Grant, 0, len(outputObs.Grants))
		for _, valGrant := range outputObs.Grants {
			tempOutput := Grant{}
			tempOutput.Delivered = valGrant.Delivered
			tempOutput.Permission = valGrant.Permission
			tempOutput.Grantee.DisplayName = valGrant.Grantee.DisplayName
			tempOutput.Grantee.ID = valGrant.Grantee.ID
			tempOutput.Grantee.Type = valGrant.Grantee.Type
			tempOutput.Grantee.URI = GroupAllUsers

			output.Grants = append(output.Grants, tempOutput)
		}
	}
	return
}

// GetBucketAcl gets the bucket ACL.
//
// You can use this API to obtain a bucket ACL.
func (obsClient ObsClient) GetBucketAcl(bucketName string) (output *GetBucketAclOutput, err error) {
	output = &GetBucketAclOutput{}
	if obsClient.conf.signature == SignatureObs {
		return obsClient.getBucketACLObs(bucketName)
	}
	err = obsClient.doActionWithBucket("GetBucketAcl", HTTP_GET, bucketName, newSubResourceSerial(SubResourceAcl), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketPolicy sets the bucket policy.
//
// You can use this API to set a bucket policy. If the bucket already has a policy, the
// policy will be overwritten by the one specified in this request.
func (obsClient ObsClient) SetBucketPolicy(input *SetBucketPolicyInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketPolicy is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketPolicy", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketPolicy gets the bucket policy.
//
// You can use this API to obtain the policy of a bucket.
func (obsClient ObsClient) GetBucketPolicy(bucketName string) (output *GetBucketPolicyOutput, err error) {
	output = &GetBucketPolicyOutput{}
	err = obsClient.doActionWithBucketV2("GetBucketPolicy", HTTP_GET, bucketName, newSubResourceSerial(SubResourcePolicy), output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketPolicy deletes the bucket policy.
//
// You can use this API to delete the policy of a bucket.
func (obsClient ObsClient) DeleteBucketPolicy(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucketPolicy", HTTP_DELETE, bucketName, newSubResourceSerial(SubResourcePolicy), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketCors sets CORS rules for a bucket.
//
// You can use this API to set CORS rules for a bucket to allow client browsers to send cross-origin requests.
func (obsClient ObsClient) SetBucketCors(input *SetBucketCorsInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketCorsInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketCors", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketCors gets CORS rules of a bucket.
//
// You can use this API to obtain the CORS rules of a specified bucket.
func (obsClient ObsClient) GetBucketCors(bucketName string) (output *GetBucketCorsOutput, err error) {
	output = &GetBucketCorsOutput{}
	err = obsClient.doActionWithBucket("GetBucketCors", HTTP_GET, bucketName, newSubResourceSerial(SubResourceCors), output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketCors deletes CORS rules of a bucket.
//
// You can use this API to delete the CORS rules of a specified bucket.
func (obsClient ObsClient) DeleteBucketCors(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucketCors", HTTP_DELETE, bucketName, newSubResourceSerial(SubResourceCors), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketVersioning sets the versioning status for a bucket.
//
// You can use this API to set the versioning status for a bucket.
func (obsClient ObsClient) SetBucketVersioning(input *SetBucketVersioningInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketVersioningInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketVersioning", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketVersioning gets the versioning status of a bucket.
//
// You can use this API to obtain the versioning status of a bucket.
func (obsClient ObsClient) GetBucketVersioning(bucketName string) (output *GetBucketVersioningOutput, err error) {
	output = &GetBucketVersioningOutput{}
	err = obsClient.doActionWithBucket("GetBucketVersioning", HTTP_GET, bucketName, newSubResourceSerial(SubResourceVersioning), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketWebsiteConfiguration sets website hosting for a bucket.
//
// You can use this API to set website hosting for a bucket.
func (obsClient ObsClient) SetBucketWebsiteConfiguration(input *SetBucketWebsiteConfigurationInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketWebsiteConfigurationInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketWebsiteConfiguration", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketWebsiteConfiguration gets the website hosting settings of a bucket.
//
// You can use this API to obtain the website hosting settings of a bucket.
func (obsClient ObsClient) GetBucketWebsiteConfiguration(bucketName string) (output *GetBucketWebsiteConfigurationOutput, err error) {
	output = &GetBucketWebsiteConfigurationOutput{}
	err = obsClient.doActionWithBucket("GetBucketWebsiteConfiguration", HTTP_GET, bucketName, newSubResourceSerial(SubResourceWebsite), output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketWebsiteConfiguration deletes the website hosting settings of a bucket.
//
// You can use this API to delete the website hosting settings of a bucket.
func (obsClient ObsClient) DeleteBucketWebsiteConfiguration(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucketWebsiteConfiguration", HTTP_DELETE, bucketName, newSubResourceSerial(SubResourceWebsite), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketLoggingConfiguration sets the bucket logging.
//
// You can use this API to configure access logging for a bucket.
func (obsClient ObsClient) SetBucketLoggingConfiguration(input *SetBucketLoggingConfigurationInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketLoggingConfigurationInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketLoggingConfiguration", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketLoggingConfiguration gets the logging settings of a bucket.
//
// You can use this API to obtain the access logging settings of a bucket.
func (obsClient ObsClient) GetBucketLoggingConfiguration(bucketName string) (output *GetBucketLoggingConfigurationOutput, err error) {
	output = &GetBucketLoggingConfigurationOutput{}
	err = obsClient.doActionWithBucket("GetBucketLoggingConfiguration", HTTP_GET, bucketName, newSubResourceSerial(SubResourceLogging), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketLifecycleConfiguration sets lifecycle rules for a bucket.
//
// You can use this API to set lifecycle rules for a bucket, to periodically transit
// storage classes of objects and delete objects in the bucket.
func (obsClient ObsClient) SetBucketLifecycleConfiguration(input *SetBucketLifecycleConfigurationInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketLifecycleConfigurationInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketLifecycleConfiguration", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketLifecycleConfiguration gets lifecycle rules of a bucket.
//
// You can use this API to obtain the lifecycle rules of a bucket.
func (obsClient ObsClient) GetBucketLifecycleConfiguration(bucketName string) (output *GetBucketLifecycleConfigurationOutput, err error) {
	output = &GetBucketLifecycleConfigurationOutput{}
	err = obsClient.doActionWithBucket("GetBucketLifecycleConfiguration", HTTP_GET, bucketName, newSubResourceSerial(SubResourceLifecycle), output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketLifecycleConfiguration deletes lifecycle rules of a bucket.
//
// You can use this API to delete all lifecycle rules of a bucket.
func (obsClient ObsClient) DeleteBucketLifecycleConfiguration(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucketLifecycleConfiguration", HTTP_DELETE, bucketName, newSubResourceSerial(SubResourceLifecycle), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketEncryption sets the default server-side encryption for a bucket.
//
// You can use this API to create or update the default server-side encryption for a bucket.
func (obsClient ObsClient) SetBucketEncryption(input *SetBucketEncryptionInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketEncryptionInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketEncryption", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketEncryption gets the encryption configuration of a bucket.
//
// You can use this API to obtain the encryption configuration of a bucket.
func (obsClient ObsClient) GetBucketEncryption(bucketName string) (output *GetBucketEncryptionOutput, err error) {
	output = &GetBucketEncryptionOutput{}
	err = obsClient.doActionWithBucket("GetBucketEncryption", HTTP_GET, bucketName, newSubResourceSerial(SubResourceEncryption), output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketEncryption deletes the encryption configuration of a bucket.
//
// You can use this API to delete the encryption configuration of a bucket.
func (obsClient ObsClient) DeleteBucketEncryption(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucketEncryption", HTTP_DELETE, bucketName, newSubResourceSerial(SubResourceEncryption), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketTagging sets bucket tags.
//
// You can use this API to set bucket tags.
func (obsClient ObsClient) SetBucketTagging(input *SetBucketTaggingInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketTaggingInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketTagging", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketTagging gets bucket tags.
//
// You can use this API to obtain the tags of a specified bucket.
func (obsClient ObsClient) GetBucketTagging(bucketName string) (output *GetBucketTaggingOutput, err error) {
	output = &GetBucketTaggingOutput{}
	err = obsClient.doActionWithBucket("GetBucketTagging", HTTP_GET, bucketName, newSubResourceSerial(SubResourceTagging), output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketTagging deletes bucket tags.
//
// You can use this API to delete the tags of a specified bucket.
func (obsClient ObsClient) DeleteBucketTagging(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucketTagging", HTTP_DELETE, bucketName, newSubResourceSerial(SubResourceTagging), output)
	if err != nil {
		output = nil
	}
	return
}

// SetBucketNotification sets event notification for a bucket.
//
// You can use this API to configure event notification for a bucket. You will be notified of all
// specified operations performed on the bucket.
func (obsClient ObsClient) SetBucketNotification(input *SetBucketNotificationInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketNotificationInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketNotification", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketNotification gets event notification settings of a bucket.
//
// You can use this API to obtain the event notification configuration of a bucket.
func (obsClient ObsClient) GetBucketNotification(bucketName string) (output *GetBucketNotificationOutput, err error) {
	if obsClient.conf.signature != SignatureObs {
		return obsClient.getBucketNotificationS3(bucketName)
	}
	output = &GetBucketNotificationOutput{}
	err = obsClient.doActionWithBucket("GetBucketNotification", HTTP_GET, bucketName, newSubResourceSerial(SubResourceNotification), output)
	if err != nil {
		output = nil
	}
	return
}

func (obsClient ObsClient) getBucketNotificationS3(bucketName string) (output *GetBucketNotificationOutput, err error) {
	outputS3 := &getBucketNotificationOutputS3{}
	err = obsClient.doActionWithBucket("GetBucketNotification", HTTP_GET, bucketName, newSubResourceSerial(SubResourceNotification), outputS3)
	if err != nil {
		return nil, err
	}

	output = &GetBucketNotificationOutput{}
	output.BaseModel = outputS3.BaseModel
	topicConfigurations := make([]TopicConfiguration, 0, len(outputS3.TopicConfigurations))
	for _, topicConfigurationS3 := range outputS3.TopicConfigurations {
		topicConfiguration := TopicConfiguration{}
		topicConfiguration.ID = topicConfigurationS3.ID
		topicConfiguration.Topic = topicConfigurationS3.Topic
		topicConfiguration.FilterRules = topicConfigurationS3.FilterRules

		events := make([]EventType, 0, len(topicConfigurationS3.Events))
		for _, event := range topicConfigurationS3.Events {
			events = append(events, ParseStringToEventType(event))
		}
		topicConfiguration.Events = events
		topicConfigurations = append(topicConfigurations, topicConfiguration)
	}
	output.TopicConfigurations = topicConfigurations
	return
}

// SetBucketReplication sets the cross-region replication for a bucket.
//
// You can use this API to create or update the cross-region replication for a bucket.
func (obsClient ObsClient) SetBucketReplication(input *SetBucketReplicationInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketReplicationInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetBucketReplication", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetBucketReplication gets the cross-region replication configuration of a bucket.
//
// You can use this API to obtain the cross-region replication configuration of a bucket.
func (obsClient ObsClient) GetBucketReplication(bucketName string) (output *GetBucketReplicationOutput, err error) {
	output = &GetBucketReplicationOutput{}
	err = obsClient.doActionWithBucket("GetBucketReplication", HTTP_GET, bucketName, newSubResourceSerial(SubResourceReplication), output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteBucketReplication deletes the cross-region replication configuration of a bucket.
//
// You can use this API to delete the cross-region replication configuration of a bucket.
func (obsClient ObsClient) DeleteBucketReplication(bucketName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("DeleteBucketReplication", HTTP_DELETE, bucketName, newSubResourceSerial(SubResourceReplication), output)
	if err != nil {
		output = nil
	}
	return
}

// SetWORMPolicy sets a default WORM policy for a bucket.
//
// You can use this API to configure the default WORM policy and a retention period.
func (obsClient ObsClient) SetWORMPolicy(input *SetWORMPolicyInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("SetBucketWORMInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithBucket("SetWORMPolicy", HTTP_PUT, input.Bucket, input, output)
	if err != nil {
		output = nil
	}
	return
}

// GetWORMPolicy gets a WORM policy for a bucket.
//
// You can use this API to retrieve the default WORM policy and a retention period.
func (obsClient ObsClient) GetWORMPolicy(bucketName string) (output *GetBucketWORMPolicyOutput, err error) {
	output = &GetBucketWORMPolicyOutput{}
	err = obsClient.doActionWithBucket("GetWORMPolicy", HTTP_GET, bucketName, newSubResourceSerial(SubResourceObjectLock), output)
	if err != nil {
		output = nil
	}
	return
}
