// Copyright 2019 Huawei Technologies Co.,Ltd.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License.  You may obtain a copy of the
// License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.

package obs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type IReadCloser interface {
	setReadCloser(body io.ReadCloser)
}

func (output *GetObjectOutput) setReadCloser(body io.ReadCloser) {
	output.Body = body
}

func setHeaders(headers map[string][]string, header string, headerValue []string, isObs bool) {
	if isObs {
		header = HeaderPrefixObs + header
		headers[header] = headerValue
	} else {
		header = HeaderPrefix + header
		headers[header] = headerValue
	}
}

func setHeadersNext(headers map[string][]string, header string, headerNext string, headerValue []string, isObs bool) {
	if isObs {
		headers[header] = headerValue
	} else {
		headers[headerNext] = headerValue
	}
}

type IBaseModel interface {
	setStatusCode(statusCode int)

	setRequestId(requestId string)

	setResponseHeaders(responseHeaders map[string][]string)
}

type ISerializable interface {
	trans(isObs bool) (map[string]string, map[string][]string, interface{}, error)
}

type DefaultSerializable struct {
	params  map[string]string
	headers map[string][]string
	data    interface{}
}

func (input DefaultSerializable) trans(isObs bool) (map[string]string, map[string][]string, interface{}, error) {
	return input.params, input.headers, input.data, nil
}

var defaultSerializable = &DefaultSerializable{}

func newSubResourceSerial(subResource SubResourceType) *DefaultSerializable {
	return &DefaultSerializable{map[string]string{string(subResource): ""}, nil, nil}
}

func trans(subResource SubResourceType, input interface{}) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(subResource): ""}
	data, err = ConvertRequestToIoReader(input)
	return
}

func (baseModel *BaseModel) setStatusCode(statusCode int) {
	baseModel.StatusCode = statusCode
}

func (baseModel *BaseModel) setRequestId(requestId string) {
	baseModel.RequestId = requestId
}

func (baseModel *BaseModel) setResponseHeaders(responseHeaders map[string][]string) {
	baseModel.ResponseHeaders = responseHeaders
}

func (input ListBucketsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	headers = make(map[string][]string)
	if input.QueryLocation && !isObs {
		setHeaders(headers, HeaderLocationAmz, []string{"true"}, isObs)
	}
	return
}

func (input CreateBucketInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	headers = make(map[string][]string)
	if acl := string(input.ACL); acl != "" {
		setHeaders(headers, HeaderAcl, []string{acl}, isObs)
	}
	if storageClass := string(input.StorageClass); storageClass != "" {
		if !isObs {
			if storageClass == "WARM" {
				storageClass = "STANDARD_IA"
			} else if storageClass == "COLD" {
				storageClass = "GLACIER"
			}
		}
		setHeadersNext(headers, HeaderStorageClassObs, HeaderStorageClass, []string{storageClass}, isObs)
		if epid := string(input.Epid); epid != "" {
			setHeaders(headers, HeaderEpidHeaders, []string{epid}, isObs)
		}
	}
	if grantReadId := string(input.GrantReadId); grantReadId != "" {
		setHeaders(headers, HeaderGrantReadObs, []string{grantReadId}, isObs)
	}
	if grantWriteId := string(input.GrantWriteId); grantWriteId != "" {
		setHeaders(headers, HeaderGrantWriteObs, []string{grantWriteId}, isObs)
	}
	if grantReadAcpId := string(input.GrantReadAcpId); grantReadAcpId != "" {
		setHeaders(headers, HeaderGrantReadAcpObs, []string{grantReadAcpId}, isObs)
	}
	if grantWriteAcpId := string(input.GrantWriteAcpId); grantWriteAcpId != "" {
		setHeaders(headers, HeaderGrantWriteAcpObs, []string{grantWriteAcpId}, isObs)
	}
	if grantFullControlId := string(input.GrantFullControlId); grantFullControlId != "" {
		setHeaders(headers, HeaderGrantFullControlObs, []string{grantFullControlId}, isObs)
	}
	if grantReadDeliveredId := string(input.GrantReadDeliveredId); grantReadDeliveredId != "" {
		setHeaders(headers, HeaderGrantReadDeliveredObs, []string{grantReadDeliveredId}, true)
	}
	if grantFullControlDeliveredId := string(input.GrantFullControlDeliveredId); grantFullControlDeliveredId != "" {
		setHeaders(headers, HeaderGrantFullControlDeliveredObs, []string{grantFullControlDeliveredId}, true)
	}
	if location := strings.TrimSpace(input.Location); location != "" {
		input.Location = location

		xml := make([]string, 0, 3)
		xml = append(xml, "<CreateBucketConfiguration>")
		if isObs {
			xml = append(xml, fmt.Sprintf("<Location>%s</Location>", input.Location))
		} else {
			xml = append(xml, fmt.Sprintf("<LocationConstraint>%s</LocationConstraint>", input.Location))
		}
		xml = append(xml, "</CreateBucketConfiguration>")

		data = strings.Join(xml, "")
	}
	return
}

func (input SetBucketStoragePolicyInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	xml := make([]string, 0, 1)
	if !isObs {
		storageClass := "STANDARD"
		if input.StorageClass == "WARM" {
			storageClass = "STANDARD_IA"
		} else if input.StorageClass == "COLD" {
			storageClass = "GLACIER"
		}
		params = map[string]string{string(SubResourceStoragePolicy): ""}
		xml = append(xml, fmt.Sprintf("<StoragePolicy><DefaultStorageClass>%s</DefaultStorageClass></StoragePolicy>", storageClass))
	} else {
		if input.StorageClass != "WARM" && input.StorageClass != "COLD" {
			input.StorageClass = StorageClassStandard
		}
		params = map[string]string{string(SubResourceStorageClass): ""}
		xml = append(xml, fmt.Sprintf("<StorageClass>%s</StorageClass>", input.StorageClass))
	}
	data = strings.Join(xml, "")
	return
}

func (input ListObjsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = make(map[string]string)
	if input.Prefix != "" {
		params["prefix"] = input.Prefix
	}
	if input.Delimiter != "" {
		params["delimiter"] = input.Delimiter
	}
	if input.MaxKeys > 0 {
		params["max-keys"] = IntToString(input.MaxKeys)
	}
	headers = make(map[string][]string)
	if origin := strings.TrimSpace(input.Origin); origin != "" {
		headers[HeaderOriginCamel] = []string{origin}
	}
	if requestHeader := strings.TrimSpace(input.RequestHeader); requestHeader != "" {
		headers[HeaderAccessControlRequestHeaderCamel] = []string{requestHeader}
	}
	return
}

func (input ListObjectsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params, headers, data, err = input.ListObjsInput.trans(isObs)
	if err != nil {
		return
	}
	if input.Marker != "" {
		params["marker"] = input.Marker
	}
	return
}

func (input ListVersionsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params, headers, data, err = input.ListObjsInput.trans(isObs)
	if err != nil {
		return
	}
	params[string(SubResourceVersions)] = ""
	if input.KeyMarker != "" {
		params["key-marker"] = input.KeyMarker
	}
	if input.VersionIdMarker != "" {
		params["version-id-marker"] = input.VersionIdMarker
	}
	return
}

func (input ListMultipartUploadsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceUploads): ""}
	if input.Prefix != "" {
		params["prefix"] = input.Prefix
	}
	if input.Delimiter != "" {
		params["delimiter"] = input.Delimiter
	}
	if input.MaxUploads > 0 {
		params["max-uploads"] = IntToString(input.MaxUploads)
	}
	if input.KeyMarker != "" {
		params["key-marker"] = input.KeyMarker
	}
	if input.UploadIdMarker != "" {
		params["upload-id-marker"] = input.UploadIdMarker
	}
	return
}

func (input SetBucketQuotaInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	return trans(SubResourceQuota, input)
}

func (input SetBucketAclInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceAcl): ""}
	headers = make(map[string][]string)

	if acl := string(input.ACL); acl != "" {
		setHeaders(headers, HeaderAcl, []string{acl}, isObs)
	} else {
		data, _ = convertBucketAclToXml(input.AccessControlPolicy, false, isObs)
	}
	return
}

func (input SetBucketPolicyInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourcePolicy): ""}
	data = strings.NewReader(input.Policy)
	return
}

func (input SetBucketCorsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceCors): ""}
	data, md5, err := ConvertRequestToIoReaderV2(input)
	if err != nil {
		return
	}
	headers = map[string][]string{HeaderMd5Camel: []string{md5}}
	return
}

func (input SetBucketVersioningInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	return trans(SubResourceVersioning, input)
}

func (input SetBucketWebsiteConfigurationInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceWebsite): ""}
	data, _ = ConvertWebsiteConfigurationToXml(input.BucketWebsiteConfiguration, false)
	return
}

func (input GetBucketMetadataInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	headers = make(map[string][]string)
	if origin := strings.TrimSpace(input.Origin); origin != "" {
		headers[HeaderOriginCamel] = []string{origin}
	}
	if requestHeader := strings.TrimSpace(input.RequestHeader); requestHeader != "" {
		headers[HeaderAccessControlRequestHeaderCamel] = []string{requestHeader}
	}
	return
}

func (input SetBucketLoggingConfigurationInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceLogging): ""}
	data, _ = ConvertLoggingStatusToXml(input.BucketLoggingStatus, false, isObs)
	return
}

func (input SetBucketLifecycleConfigurationInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceLifecycle): ""}
	data, md5 := ConvertLifecyleConfigurationToXml(input.BucketLifecyleConfiguration, true, isObs)
	headers = map[string][]string{HeaderMd5Camel: []string{md5}}
	return
}

func (input SetBucketTaggingInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceTagging): ""}
	data, md5, err := ConvertRequestToIoReaderV2(input)
	if err != nil {
		return
	}
	headers = map[string][]string{HeaderMd5Camel: []string{md5}}
	return
}

func (input SetBucketNotificationInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceNotification): ""}
	data, _ = ConvertNotificationToXml(input.BucketNotification, false, isObs)
	return
}

func (input DeleteObjectInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = make(map[string]string)
	if input.VersionId != "" {
		params[ParamVersionId] = input.VersionId
	}
	return
}

func (input DeleteObjectsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceDelete): ""}
	data, md5, err := ConvertRequestToIoReaderV2(input)
	if err != nil {
		return
	}
	headers = map[string][]string{HeaderMd5Camel: []string{md5}}
	return
}

func (input SetObjectAclInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceAcl): ""}
	if input.VersionId != "" {
		params[ParamVersionId] = input.VersionId
	}
	headers = make(map[string][]string)
	if acl := string(input.ACL); acl != "" {
		setHeaders(headers, HeaderAcl, []string{acl}, isObs)
	} else {
		data, _ = ConvertAclToXml(input.AccessControlPolicy, false, isObs)
	}
	return
}

func (input GetObjectAclInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceAcl): ""}
	if input.VersionId != "" {
		params[ParamVersionId] = input.VersionId
	}
	return
}

func (input RestoreObjectInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{string(SubResourceRestore): ""}
	if input.VersionId != "" {
		params[ParamVersionId] = input.VersionId
	}
	if !isObs {
		data, err = ConvertRequestToIoReader(input)
	} else {
		data = ConverntObsRestoreToXml(input)
	}
	return
}

func (header SseKmsHeader) GetEncryption() string {
	if header.Encryption != "" {
		return header.Encryption
	}
	if !header.isObs {
		return DefaultSseKmsEncryption
	} else {
		return DefaultSseKmsEncryptionObs
	}
}

func (header SseKmsHeader) GetKey() string {
	return header.Key
}

func (header SseCHeader) GetEncryption() string {
	if header.Encryption != "" {
		return header.Encryption
	}
	return DefaultSseCEncryption
}

func (header SseCHeader) GetKey() string {
	return header.Key
}

func (header SseCHeader) GetKeyMD5() string {
	if header.KeyMD5 != "" {
		return header.KeyMD5
	}

	if ret, err := Base64Decode(header.GetKey()); err == nil {
		return Base64Md5(ret)
	}
	return ""
}

func setSseHeader(headers map[string][]string, sseHeader ISseHeader, sseCOnly bool, isObs bool) {
	if sseHeader != nil {
		if sseCHeader, ok := sseHeader.(SseCHeader); ok {
			setHeaders(headers, HeaderSsecEncryption, []string{sseCHeader.GetEncryption()}, isObs)
			setHeaders(headers, HeaderSsecKey, []string{sseCHeader.GetKey()}, isObs)
			setHeaders(headers, HeaderSsecKeyMd5, []string{sseCHeader.GetEncryption()}, isObs)
		} else if sseKmsHeader, ok := sseHeader.(SseKmsHeader); !sseCOnly && ok {
			sseKmsHeader.isObs = isObs
			setHeaders(headers, HeaderSsekmsEncryption, []string{sseKmsHeader.GetEncryption()}, isObs)
			setHeadersNext(headers, HeaderSsekmsKeyObs, HeaderSsekmsKeyAmz, []string{sseKmsHeader.GetKey()}, isObs)
		}
	}
}

func (input GetObjectMetadataInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = make(map[string]string)
	if input.VersionId != "" {
		params[ParamVersionId] = input.VersionId
	}
	headers = make(map[string][]string)

	if input.Origin != "" {
		headers[HeaderOriginCamel] = []string{input.Origin}
	}

	if input.RequestHeader != "" {
		headers[HeaderAccessControlRequestHeaderCamel] = []string{input.RequestHeader}
	}
	setSseHeader(headers, input.SseHeader, true, isObs)
	return
}

func (input SetObjectMetadataInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = make(map[string]string)
	params = map[string]string{string(SubResourceMetadata): ""}
	if input.VersionId != "" {
		params[ParamVersionId] = input.VersionId
	}
	headers = make(map[string][]string)

	if directive := string(input.MetadataDirective); directive != "" {
		setHeaders(headers, HeaderMetadataDirective, []string{string(input.MetadataDirective)}, isObs)
	} else {
		setHeaders(headers, HeaderMetadataDirective, []string{string(ReplaceNew)}, isObs)
	}
	if input.CacheControl != "" {
		headers[HeaderCacheControlCamel] = []string{input.CacheControl}
	}
	if input.ContentDisposition != "" {
		headers[HeaderContentDispositionCamel] = []string{input.ContentDisposition}
	}
	if input.ContentEncoding != "" {
		headers[HeaderContentEncodingCamel] = []string{input.ContentEncoding}
	}
	if input.ContentLanguage != "" {
		headers[HeaderContentLanguageCamel] = []string{input.ContentLanguage}
	}

	if input.ContentType != "" {
		headers[HeaderContentTypeCaml] = []string{input.ContentType}
	}
	if input.Expires != "" {
		headers[HeaderExpiresCamel] = []string{input.Expires}
	}
	if input.WebsiteRedirectLocation != "" {
		setHeaders(headers, HeaderWebsiteRedirectLocation, []string{input.WebsiteRedirectLocation}, isObs)
	}
	if storageClass := string(input.StorageClass); storageClass != "" {
		if !isObs {
			if storageClass == "WARM" {
				storageClass = "STANDARD_IA"
			} else if storageClass == "COLD" {
				storageClass = "GLACIER"
			}
		}
		setHeaders(headers, HeaderStorageClass2, []string{storageClass}, isObs)
	}
	if input.Metadata != nil {
		for key, value := range input.Metadata {
			key = strings.TrimSpace(key)
			setHeadersNext(headers, HeaderPrefixMetaObs+key, HeaderPrefixMeta+key, []string{value}, isObs)
		}
	}
	return
}

func (input GetObjectInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params, headers, data, err = input.GetObjectMetadataInput.trans(isObs)
	if err != nil {
		return
	}
	if input.ResponseCacheControl != "" {
		params[ParamResponseCacheControl] = input.ResponseCacheControl
	}
	if input.ResponseContentDisposition != "" {
		params[ParamResponseContentDisposition] = input.ResponseContentDisposition
	}
	if input.ResponseContentEncoding != "" {
		params[ParamResponseContentEncoding] = input.ResponseContentEncoding
	}
	if input.ResponseContentLanguage != "" {
		params[ParamResponseContentLanguage] = input.ResponseContentLanguage
	}
	if input.ResponseContentType != "" {
		params[ParamResponseContentType] = input.ResponseContentType
	}
	if input.ResponseExpires != "" {
		params[ParamResponseExpires] = input.ResponseExpires
	}
	if input.ImageProcess != "" {
		params[ParamImageProcess] = input.ImageProcess
	}
	if input.RangeStart >= 0 && input.RangeEnd > input.RangeStart {
		headers[HeaderRange] = []string{fmt.Sprintf("bytes=%d-%d", input.RangeStart, input.RangeEnd)}
	}

	if input.IfMatch != "" {
		headers[HeaderIfMatch] = []string{input.IfMatch}
	}
	if input.IfNoneMatch != "" {
		headers[HeaderIfNoneMatch] = []string{input.IfNoneMatch}
	}
	if !input.IfModifiedSince.IsZero() {
		headers[HeaderIfModifiedSince] = []string{FormatUtcToRfc1123(input.IfModifiedSince)}
	}
	if !input.IfUnmodifiedSince.IsZero() {
		headers[HeaderIfUnmodifiedSince] = []string{FormatUtcToRfc1123(input.IfUnmodifiedSince)}
	}
	return
}

func (input ObjectOperationInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	headers = make(map[string][]string)
	params = make(map[string]string)
	if acl := string(input.ACL); acl != "" {
		setHeaders(headers, HeaderAcl, []string{acl}, isObs)
	}
	if GrantReadId := string(input.GrantReadId); GrantReadId != "" {
		setHeaders(headers, HeaderGrantReadObs, []string{GrantReadId}, true)
	}
	if GrantReadAcpId := string(input.GrantReadAcpId); GrantReadAcpId != "" {
		setHeaders(headers, HeaderGrantReadAcpObs, []string{GrantReadAcpId}, true)
	}
	if GrantWriteAcpId := string(input.GrantWriteAcpId); GrantWriteAcpId != "" {
		setHeaders(headers, HeaderGrantWriteAcpObs, []string{GrantWriteAcpId}, true)
	}
	if GrantFullControlId := string(input.GrantFullControlId); GrantFullControlId != "" {
		setHeaders(headers, HeaderGrantFullControlObs, []string{GrantFullControlId}, true)
	}
	if storageClass := string(input.StorageClass); storageClass != "" {
		if !isObs {
			if storageClass == "WARM" {
				storageClass = "STANDARD_IA"
			} else if storageClass == "COLD" {
				storageClass = "GLACIER"
			}
		}
		setHeaders(headers, HeaderStorageClass2, []string{storageClass}, isObs)
	}
	if input.WebsiteRedirectLocation != "" {
		setHeaders(headers, HeaderWebsiteRedirectLocation, []string{input.WebsiteRedirectLocation}, isObs)

	}
	setSseHeader(headers, input.SseHeader, false, isObs)
	if input.Expires != 0 {
		setHeaders(headers, HeaderExpires, []string{Int64ToString(input.Expires)}, true)
	}
	if input.Metadata != nil {
		for key, value := range input.Metadata {
			key = strings.TrimSpace(key)
			setHeadersNext(headers, HeaderPrefixMetaObs+key, HeaderPrefixMeta+key, []string{value}, isObs)
		}
	}
	return
}

func (input PutObjectBasicInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params, headers, data, err = input.ObjectOperationInput.trans(isObs)
	if err != nil {
		return
	}

	if input.ContentMD5 != "" {
		headers[HeaderMd5Camel] = []string{input.ContentMD5}
	}

	if input.ContentLength > 0 {
		headers[HeaderContentLengthCamel] = []string{Int64ToString(input.ContentLength)}
	}
	if input.ContentType != "" {
		headers[HeaderContentTypeCaml] = []string{input.ContentType}
	}

	return
}

func (input PutObjectInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params, headers, data, err = input.PutObjectBasicInput.trans(isObs)
	if err != nil {
		return
	}
	if input.Body != nil {
		data = input.Body
	}
	return
}

func (input CopyObjectInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params, headers, data, err = input.ObjectOperationInput.trans(isObs)
	if err != nil {
		return
	}

	var copySource string
	if input.CopySourceVersionId != "" {
		copySource = fmt.Sprintf("%s/%s?versionId=%s", input.CopySourceBucket, UrlEncode(input.CopySourceKey, false), input.CopySourceVersionId)
	} else {
		copySource = fmt.Sprintf("%s/%s", input.CopySourceBucket, UrlEncode(input.CopySourceKey, false))
	}
	setHeaders(headers, HeaderCopySource, []string{copySource}, isObs)

	if directive := string(input.MetadataDirective); directive != "" {
		setHeaders(headers, HeaderMetadataDirective, []string{directive}, isObs)
	}

	if input.MetadataDirective == ReplaceMetadata {
		if input.CacheControl != "" {
			headers[HeaderCacheControl] = []string{input.CacheControl}
		}
		if input.ContentDisposition != "" {
			headers[HeaderContentDisposition] = []string{input.ContentDisposition}
		}
		if input.ContentEncoding != "" {
			headers[HeaderContentEncoding] = []string{input.ContentEncoding}
		}
		if input.ContentLanguage != "" {
			headers[HeaderContentLanguage] = []string{input.ContentLanguage}
		}
		if input.ContentType != "" {
			headers[HeaderContentType] = []string{input.ContentType}
		}
		if input.Expires != "" {
			headers[HeaderExpires] = []string{input.Expires}
		}
	}

	if input.CopySourceIfMatch != "" {
		setHeaders(headers, HeaderCopySourceIfMatch, []string{input.CopySourceIfMatch}, isObs)
	}
	if input.CopySourceIfNoneMatch != "" {
		setHeaders(headers, HeaderCopySourceIfNoneMatch, []string{input.CopySourceIfNoneMatch}, isObs)
	}
	if !input.CopySourceIfModifiedSince.IsZero() {
		setHeaders(headers, HeaderCopySourceIfModifiedSince, []string{FormatUtcToRfc1123(input.CopySourceIfModifiedSince)}, isObs)
	}
	if !input.CopySourceIfUnmodifiedSince.IsZero() {
		setHeaders(headers, HeaderCopySourceIfUnmodifiedSince, []string{FormatUtcToRfc1123(input.CopySourceIfUnmodifiedSince)}, isObs)
	}
	if input.SourceSseHeader != nil {
		if sseCHeader, ok := input.SourceSseHeader.(SseCHeader); ok {
			setHeaders(headers, HeaderSsecCopySourceEncryption, []string{sseCHeader.GetEncryption()}, isObs)
			setHeaders(headers, HeaderSsecCopySourceKey, []string{sseCHeader.GetKey()}, isObs)
			setHeaders(headers, HeaderSsecCopySourceKeyMd5, []string{sseCHeader.GetKeyMD5()}, isObs)
		}
	}
	if input.SuccessActionRedirect != "" {
		headers[HeaderSuccessActionRedirect] = []string{input.SuccessActionRedirect}
	}
	return
}

func (input AbortMultipartUploadInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{"uploadId": input.UploadId}
	return
}

func (input InitiateMultipartUploadInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params, headers, data, err = input.ObjectOperationInput.trans(isObs)
	if err != nil {
		return
	}
	if input.ContentType != "" {
		headers[HeaderContentTypeCaml] = []string{input.ContentType}
	}
	params[string(SubResourceUploads)] = ""
	return
}

func (input UploadPartInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{"uploadId": input.UploadId, "partNumber": IntToString(input.PartNumber)}
	headers = make(map[string][]string)
	setSseHeader(headers, input.SseHeader, true, isObs)
	if input.ContentMD5 != "" {
		headers[HeaderMd5Camel] = []string{input.ContentMD5}
	}
	if input.Body != nil {
		data = input.Body
	}
	return
}

func (input CompleteMultipartUploadInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{"uploadId": input.UploadId}
	data, _ = ConvertCompleteMultipartUploadInputToXml(input, false)
	return
}

func (input ListPartsInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{"uploadId": input.UploadId}
	if input.MaxParts > 0 {
		params["max-parts"] = IntToString(input.MaxParts)
	}
	if input.PartNumberMarker > 0 {
		params["part-number-marker"] = IntToString(input.PartNumberMarker)
	}
	return
}

func (input CopyPartInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	params = map[string]string{"uploadId": input.UploadId, "partNumber": IntToString(input.PartNumber)}
	headers = make(map[string][]string, 1)
	var copySource string
	if input.CopySourceVersionId != "" {
		copySource = fmt.Sprintf("%s/%s?versionId=%s", input.CopySourceBucket, UrlEncode(input.CopySourceKey, false), input.CopySourceVersionId)
	} else {
		copySource = fmt.Sprintf("%s/%s", input.CopySourceBucket, UrlEncode(input.CopySourceKey, false))
	}
	setHeaders(headers, HeaderCopySource, []string{copySource}, isObs)
	if input.CopySourceRangeStart >= 0 && input.CopySourceRangeEnd > input.CopySourceRangeStart {
		setHeaders(headers, HeaderCopySourceRange, []string{fmt.Sprintf("bytes=%d-%d", input.CopySourceRangeStart, input.CopySourceRangeEnd)}, isObs)
	}

	setSseHeader(headers, input.SseHeader, true, isObs)
	if input.SourceSseHeader != nil {
		if sseCHeader, ok := input.SourceSseHeader.(SseCHeader); ok {
			setHeaders(headers, HeaderSsecCopySourceEncryption, []string{sseCHeader.GetEncryption()}, isObs)
			setHeaders(headers, HeaderSsecCopySourceKey, []string{sseCHeader.GetKey()}, isObs)
			setHeaders(headers, HeaderSsecCopySourceKeyMd5, []string{sseCHeader.GetKeyMD5()}, isObs)
		}

	}
	return
}

type partSlice []Part

func (parts partSlice) Len() int {
	return len(parts)
}

func (parts partSlice) Less(i, j int) bool {
	return parts[i].PartNumber < parts[j].PartNumber
}

func (parts partSlice) Swap(i, j int) {
	parts[i], parts[j] = parts[j], parts[i]
}

type readerWrapper struct {
	reader      io.Reader
	mark        int64
	totalCount  int64
	readedCount int64
}

func (rw *readerWrapper) seek(offset int64, whence int) (int64, error) {
	if r, ok := rw.reader.(*strings.Reader); ok {
		return r.Seek(offset, whence)
	} else if r, ok := rw.reader.(*bytes.Reader); ok {
		return r.Seek(offset, whence)
	} else if r, ok := rw.reader.(*os.File); ok {
		return r.Seek(offset, whence)
	}
	return offset, nil
}

func (rw *readerWrapper) Read(p []byte) (n int, err error) {
	if rw.totalCount == 0 {
		return 0, io.EOF
	}
	if rw.totalCount > 0 {
		n, err = rw.reader.Read(p)
		readedOnce := int64(n)
		if remainCount := rw.totalCount - rw.readedCount; remainCount > readedOnce {
			rw.readedCount += readedOnce
			return n, err
		} else {
			rw.readedCount += remainCount
			return int(remainCount), io.EOF
		}
	}
	return rw.reader.Read(p)
}

type fileReaderWrapper struct {
	readerWrapper
	filePath string
}
