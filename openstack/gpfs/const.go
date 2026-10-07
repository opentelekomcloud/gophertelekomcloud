package gpfs

const (
	userAgent               = "gophertelekomcloud-gpfs"
	headerPrefixOBS         = "x-obs-"
	headerDateOBS           = "x-obs-date"
	headerSecurityTokenOBS  = "x-obs-security-token"
	headerAZRedundancy      = "az-redundancy"
	headerBucketType        = "bucket-type"
	headerRequestID         = "request-id"
	headerContentType       = "content-type"
	headerHost              = "Host"
	headerAuthorization     = "Authorization"
	headerContentLength     = "Content-Length"
	headerLocation          = "Location"
	headerUserAgent         = "User-Agent"
	defaultConnectTimeout   = 60
	defaultSocketTimeout    = 60
	defaultHeaderTimeout    = 60
	defaultIdleConnTimeout  = 30
	defaultMaxRetryCount    = 3
	defaultMaxRedirectCount = 3
	defaultMaxConnsPerHost  = 1000
	rfc1123Format           = "Mon, 02 Jan 2006 15:04:05 GMT"
	httpGet                 = "GET"
	httpPut                 = "PUT"
	httpDelete              = "DELETE"
	subResourceSFSACL       = "sfsacl"
)

// AccessRuleAction defines the access granted to a VPC by a file system access rule.
type AccessRuleAction string

const (
	// AccessRuleActionFullControl grants read and write access.
	AccessRuleActionFullControl AccessRuleAction = "FullControl"
	// AccessRuleActionRead grants read-only access.
	AccessRuleActionRead AccessRuleAction = "Read"
)

// AccessRuleEffect defines whether a file system access rule grants access.
type AccessRuleEffect string

const (
	// AccessRuleEffectAllow grants the access specified by the rule action.
	AccessRuleEffectAllow AccessRuleEffect = "Allow"
)
