package gpfs

const (
	obs_sdk_version        = "3.21.12"
	USER_AGENT             = "obs-sdk-go/" + obs_sdk_version
	HEADER_PREFIX          = "x-amz-"
	HEADER_PREFIX_META     = "x-amz-meta-"
	HEADER_PREFIX_OBS      = "x-obs-"
	HEADER_PREFIX_META_OBS = "x-obs-meta-"
	HEADER_DATE_AMZ        = "x-amz-date"
	HEADER_DATE_OBS        = "x-obs-date"
	HEADER_STS_TOKEN_AMZ   = "x-amz-security-token"
	HEADER_STS_TOKEN_OBS   = "x-obs-security-token"
	HEADER_ACCESSS_KEY_AMZ = "AWSAccessKeyId"
	PREFIX_META            = "meta-"

	HEADER_AZ_REDUNDANCY = "x-obs-az-redundancy"
	HEADER_BUCKET_TYPE   = "x-obs-bucket-type"

	HEADER_CONTENT_SHA256_AMZ = "x-amz-content-sha256"
	HEADER_REQUEST_ID         = "request-id"
	HEADER_CONTENT_LENGTH     = "content-length"
	HEADER_CONTENT_TYPE       = "content-type"

	HEADER_SSEC_ENCRYPTION = "server-side-encryption-customer-algorithm"
	HEADER_SSEC_KEY        = "server-side-encryption-customer-key"
	HEADER_SSEC_KEY_MD5    = "server-side-encryption-customer-key-MD5"

	HEADER_SSEKMS_ENCRYPTION      = "server-side-encryption"
	HEADER_SSEKMS_KEY             = "server-side-encryption-aws-kms-key-id"
	HEADER_SSEKMS_ENCRYPT_KEY_OBS = "server-side-encryption-kms-key-id"

	HEADER_SSEC_COPY_SOURCE_ENCRYPTION = "copy-source-server-side-encryption-customer-algorithm"
	HEADER_SSEC_COPY_SOURCE_KEY        = "copy-source-server-side-encryption-customer-key"
	HEADER_SSEC_COPY_SOURCE_KEY_MD5    = "copy-source-server-side-encryption-customer-key-MD5"

	HEADER_SSEKMS_KEY_AMZ = "x-amz-server-side-encryption-aws-kms-key-id"

	HEADER_SSEKMS_KEY_OBS = "x-obs-server-side-encryption-kms-key-id"

	HEADER_SUCCESS_ACTION_REDIRECT = "success_action_redirect"

	HEADER_FS_FILE_INTERFACE   = "fs-file-interface"
	HEADER_OBJECT_LOCK_ENABLED = "bucket-object-lock-enabled"

	HEADER_DATE_CAMEL                          = "Date"
	HEADER_HOST_CAMEL                          = "Host"
	HEADER_HOST                                = "host"
	HEADER_AUTH_CAMEL                          = "Authorization"
	HEADER_MD5_CAMEL                           = "Content-MD5"
	HEADER_LOCATION_CAMEL                      = "Location"
	HEADER_CONTENT_LENGTH_CAMEL                = "Content-Length"
	HEADER_CONTENT_TYPE_CAML                   = "Content-Type"
	HEADER_USER_AGENT_CAMEL                    = "User-Agent"
	HEADER_ORIGIN_CAMEL                        = "Origin"
	HEADER_ACCESS_CONTROL_REQUEST_HEADER_CAMEL = "Access-Control-Request-Headers"
	HEADER_CACHE_CONTROL_CAMEL                 = "Cache-Control"
	HEADER_CONTENT_DISPOSITION_CAMEL           = "Content-Disposition"
	HEADER_CONTENT_ENCODING_CAMEL              = "Content-Encoding"
	HEADER_CONTENT_LANGUAGE_CAMEL              = "Content-Language"
	HEADER_EXPIRES_CAMEL                       = "Expires"

	PARAM_VERSION_ID                   = "versionId"
	PARAM_RESPONSE_CONTENT_TYPE        = "response-content-type"
	PARAM_RESPONSE_CONTENT_LANGUAGE    = "response-content-language"
	PARAM_RESPONSE_EXPIRES             = "response-expires"
	PARAM_RESPONSE_CACHE_CONTROL       = "response-cache-control"
	PARAM_RESPONSE_CONTENT_DISPOSITION = "response-content-disposition"
	PARAM_RESPONSE_CONTENT_ENCODING    = "response-content-encoding"

	PARAM_ALGORITHM_AMZ_CAMEL     = "X-Amz-Algorithm"
	PARAM_CREDENTIAL_AMZ_CAMEL    = "X-Amz-Credential"
	PARAM_DATE_AMZ_CAMEL          = "X-Amz-Date"
	PARAM_DATE_OBS_CAMEL          = "X-Obs-Date"
	PARAM_EXPIRES_AMZ_CAMEL       = "X-Amz-Expires"
	PARAM_SIGNEDHEADERS_AMZ_CAMEL = "X-Amz-SignedHeaders"
	PARAM_SIGNATURE_AMZ_CAMEL     = "X-Amz-Signature"

	DEFAULT_SIGNATURE            = SignatureV2
	DEFAULT_REGION               = "region"
	DEFAULT_CONNECT_TIMEOUT      = 60
	DEFAULT_SOCKET_TIMEOUT       = 60
	DEFAULT_HEADER_TIMEOUT       = 60
	DEFAULT_IDLE_CONN_TIMEOUT    = 30
	DEFAULT_MAX_RETRY_COUNT      = 3
	DEFAULT_MAX_REDIRECT_COUNT   = 3
	DEFAULT_MAX_CONN_PER_HOST    = 1000
	EMPTY_CONTENT_SHA256         = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	UNSIGNED_PAYLOAD             = "UNSIGNED-PAYLOAD"
	LONG_DATE_FORMAT             = "20060102T150405Z"
	SHORT_DATE_FORMAT            = "20060102"
	ISO8601_DATE_FORMAT          = "2006-01-02T15:04:05Z"
	ISO8601_MIDNIGHT_DATE_FORMAT = "2006-01-02T00:00:00Z"
	RFC1123_FORMAT               = "Mon, 02 Jan 2006 15:04:05 GMT"

	V4_SERVICE_NAME   = "s3"
	V4_SERVICE_SUFFIX = "aws4_request"

	V2_HASH_PREFIX  = "AWS"
	OBS_HASH_PREFIX = "OBS"

	V4_HASH_PREFIX = "AWS4-HMAC-SHA256"
	V4_HASH_PRE    = "AWS4"

	HTTP_GET     = "GET"
	HTTP_POST    = "POST"
	HTTP_PUT     = "PUT"
	HTTP_DELETE  = "DELETE"
	HTTP_HEAD    = "HEAD"
	HTTP_OPTIONS = "OPTIONS"
)

type SignatureType string

const (
	SignatureV2  SignatureType = "v2"
	SignatureV4  SignatureType = "v4"
	SignatureObs SignatureType = "OBS"
)

var (
	interestedHeaders = []string{"content-md5", "content-type", "date"}

	allowedRequestHttpHeaderMetadataNames = map[string]bool{
		"content-type":                  true,
		"content-md5":                   true,
		"content-length":                true,
		"content-language":              true,
		"expires":                       true,
		"origin":                        true,
		"cache-control":                 true,
		"content-disposition":           true,
		"content-encoding":              true,
		"x-default-storage-class":       true,
		"location":                      true,
		"date":                          true,
		"etag":                          true,
		"host":                          true,
		"last-modified":                 true,
		"content-range":                 true,
		"x-reserved":                    true,
		"x-reserved-indicator":          true,
		"access-control-allow-origin":   true,
		"access-control-allow-headers":  true,
		"access-control-max-age":        true,
		"access-control-allow-methods":  true,
		"access-control-expose-headers": true,
		"connection":                    true,
	}

	allowedResourceParameterNames = map[string]bool{
		"acl":                          true,
		"backtosource":                 true,
		"policy":                       true,
		"torrent":                      true,
		"logging":                      true,
		"location":                     true,
		"storageinfo":                  true,
		"quota":                        true,
		"storageclass":                 true,
		"storagepolicy":                true,
		"requestpayment":               true,
		"versions":                     true,
		"versioning":                   true,
		"versionid":                    true,
		"uploads":                      true,
		"uploadid":                     true,
		"partnumber":                   true,
		"website":                      true,
		"notification":                 true,
		"lifecycle":                    true,
		"deletebucket":                 true,
		"delete":                       true,
		"cors":                         true,
		"object-lock":                  true,
		"restore":                      true,
		"encryption":                   true,
		"inventory":                    true,
		"tagging":                      true,
		"append":                       true,
		"position":                     true,
		"replication":                  true,
		"response-content-type":        true,
		"response-content-language":    true,
		"response-expires":             true,
		"response-cache-control":       true,
		"response-content-disposition": true,
		"response-content-encoding":    true,
		"x-image-process":              true,
		"x-oss-process":                true,
		"x-image-save-bucket":          true,
		"x-image-save-object":          true,
	}
)

type HttpMethodType string

const (
	HttpMethodGet    HttpMethodType = HTTP_GET
	HttpMethodPut    HttpMethodType = HTTP_PUT
	HttpMethodPost   HttpMethodType = HTTP_POST
	HttpMethodDelete HttpMethodType = HTTP_DELETE
)

type StorageClassType string

const (
	StorageClassStandard StorageClassType = "STANDARD"
	StorageClassWarm     StorageClassType = "WARM"
	StorageClassCold     StorageClassType = "COLD"
)
