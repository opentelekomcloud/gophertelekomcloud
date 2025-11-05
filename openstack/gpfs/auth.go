package gpfs

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

func (obsClient ObsClient) doAuth(method, bucketName, objectKey string, params map[string]string,
	headers map[string][]string, hostName string) (requestUrl string, err error) {
	isAkSkEmpty := obsClient.conf.securityProvider == nil || obsClient.conf.securityProvider.ak == "" || obsClient.conf.securityProvider.sk == ""
	if !isAkSkEmpty && obsClient.conf.securityProvider.securityToken != "" {
		if obsClient.conf.signature == SignatureObs {
			headers[HEADER_STS_TOKEN_OBS] = []string{obsClient.conf.securityProvider.securityToken}
		} else {
			headers[HEADER_STS_TOKEN_AMZ] = []string{obsClient.conf.securityProvider.securityToken}
		}
	}
	isObs := obsClient.conf.signature == SignatureObs
	requestUrl, canonicalizedUrl := obsClient.conf.formatUrls(bucketName, objectKey, params, true)
	parsedRequestUrl, err := url.Parse(requestUrl)
	if err != nil {
		return "", err
	}
	encodeHeaders(headers)

	if hostName == "" {
		hostName = parsedRequestUrl.Host
	}

	isV4 := obsClient.conf.signature == SignatureV4
	prepareHostAndDate(headers, hostName, isV4)

	if isAkSkEmpty {
		doLog(LEVEL_WARN, "No ak/sk provided, skip to construct authorization")
	} else {
		ak := obsClient.conf.securityProvider.ak
		sk := obsClient.conf.securityProvider.sk
		var authorization string
		if isV4 {
			headers[HEADER_CONTENT_SHA256_AMZ] = []string{UNSIGNED_PAYLOAD}
			ret := v4Auth(ak, sk, obsClient.conf.region, method, canonicalizedUrl, parsedRequestUrl.RawQuery, headers)
			authorization = fmt.Sprintf("%s Credential=%s,SignedHeaders=%s,Signature=%s", V4_HASH_PREFIX, ret["Credential"], ret["SignedHeaders"], ret["Signature"])
		} else {
			ret := v2Auth(ak, sk, method, canonicalizedUrl, headers, isObs)
			hashPrefix := V2_HASH_PREFIX
			if isObs {
				hashPrefix = OBS_HASH_PREFIX
			}
			authorization = fmt.Sprintf("%s %s:%s", hashPrefix, ak, ret["Signature"])
		}
		headers[HEADER_AUTH_CAMEL] = []string{authorization}
	}
	return
}

func prepareHostAndDate(headers map[string][]string, hostName string, isV4 bool) {
	headers[HEADER_HOST_CAMEL] = []string{hostName}
	if date, ok := headers[HEADER_DATE_AMZ]; ok {
		flag := false
		if len(date) == 1 {
			if isV4 {
				if t, err := time.Parse(LONG_DATE_FORMAT, date[0]); err == nil {
					headers[HEADER_DATE_CAMEL] = []string{FormatUtcToRfc1123(t)}
					flag = true
				}
			} else {
				if strings.HasSuffix(date[0], "GMT") {
					headers[HEADER_DATE_CAMEL] = []string{date[0]}
					flag = true
				}
			}
		}
		if !flag {
			delete(headers, HEADER_DATE_AMZ)
		}
	}
	if _, ok := headers[HEADER_DATE_CAMEL]; !ok {
		headers[HEADER_DATE_CAMEL] = []string{FormatUtcToRfc1123(time.Now().UTC())}
	}
}

func encodeHeaders(headers map[string][]string) {
	for key, values := range headers {
		for index, value := range values {
			values[index] = UrlEncode(value, true)
		}
		headers[key] = values
	}
}

func prepareDateHeader(dataHeader, dateCamelHeader string, headers, _headers map[string][]string) {
	if _, ok := _headers[HEADER_DATE_CAMEL]; ok {
		if _, ok := _headers[dataHeader]; ok {
			_headers[HEADER_DATE_CAMEL] = []string{""}
		} else if _, ok := headers[dateCamelHeader]; ok {
			_headers[HEADER_DATE_CAMEL] = []string{""}
		}
	} else if _, ok := _headers[strings.ToLower(HEADER_DATE_CAMEL)]; ok {
		if _, ok := _headers[dataHeader]; ok {
			_headers[HEADER_DATE_CAMEL] = []string{""}
		} else if _, ok := headers[dateCamelHeader]; ok {
			_headers[HEADER_DATE_CAMEL] = []string{""}
		}
	}
}

func getStringToSign(keys []string, isObs bool, _headers map[string][]string) []string {
	stringToSign := make([]string, 0, len(keys))
	for _, key := range keys {
		var value string
		prefixHeader := HEADER_PREFIX
		prefixMetaHeader := HEADER_PREFIX_META
		if isObs {
			prefixHeader = HEADER_PREFIX_OBS
			prefixMetaHeader = HEADER_PREFIX_META_OBS
		}
		if strings.HasPrefix(key, prefixHeader) {
			if strings.HasPrefix(key, prefixMetaHeader) {
				for index, v := range _headers[key] {
					value += strings.TrimSpace(v)
					if index != len(_headers[key])-1 {
						value += ","
					}
				}
			} else {
				value = strings.Join(_headers[key], ",")
			}
			value = fmt.Sprintf("%s:%s", key, value)
		} else {
			value = strings.Join(_headers[key], ",")
		}
		stringToSign = append(stringToSign, value)
	}
	return stringToSign
}

func attachHeaders(headers map[string][]string, isObs bool) string {
	length := len(headers)
	_headers := make(map[string][]string, length)
	keys := make([]string, 0, length)

	for key, value := range headers {
		_key := strings.ToLower(strings.TrimSpace(key))
		if _key != "" {
			prefixheader := HEADER_PREFIX
			if isObs {
				prefixheader = HEADER_PREFIX_OBS
			}
			if _key == "content-md5" || _key == "content-type" || _key == "date" || strings.HasPrefix(_key, prefixheader) {
				keys = append(keys, _key)
				_headers[_key] = value
			}
		} else {
			delete(headers, key)
		}
	}

	for _, interestedHeader := range interestedHeaders {
		if _, ok := _headers[interestedHeader]; !ok {
			_headers[interestedHeader] = []string{""}
			keys = append(keys, interestedHeader)
		}
	}
	dateCamelHeader := PARAM_DATE_AMZ_CAMEL
	dataHeader := HEADER_DATE_AMZ
	if isObs {
		dateCamelHeader = PARAM_DATE_OBS_CAMEL
		dataHeader = HEADER_DATE_OBS
	}
	prepareDateHeader(dataHeader, dateCamelHeader, headers, _headers)

	sort.Strings(keys)
	stringToSign := getStringToSign(keys, isObs, _headers)
	return strings.Join(stringToSign, "\n")
}

func getCredential(ak, region, shortDate string) (string, string) {
	scope := getScope(region, shortDate)
	return fmt.Sprintf("%s/%s", ak, scope), scope
}

func getScope(region, shortDate string) string {
	return fmt.Sprintf("%s/%s/%s/%s", shortDate, region, V4_SERVICE_NAME, V4_SERVICE_SUFFIX)
}

func getSignedHeaders(headers map[string][]string) ([]string, map[string][]string) {
	length := len(headers)
	_headers := make(map[string][]string, length)
	signedHeaders := make([]string, 0, length)
	for key, value := range headers {
		_key := strings.ToLower(strings.TrimSpace(key))
		if _key != "" {
			signedHeaders = append(signedHeaders, _key)
			_headers[_key] = value
		} else {
			delete(headers, key)
		}
	}
	sort.Strings(signedHeaders)
	return signedHeaders, _headers
}

func getSignature(stringToSign, sk, region, shortDate string) string {
	key := HmacSha256([]byte(V4_HASH_PRE+sk), []byte(shortDate))
	key = HmacSha256(key, []byte(region))
	key = HmacSha256(key, []byte(V4_SERVICE_NAME))
	key = HmacSha256(key, []byte(V4_SERVICE_SUFFIX))
	return Hex(HmacSha256(key, []byte(stringToSign)))
}
