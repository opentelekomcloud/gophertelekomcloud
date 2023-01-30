package obs

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func (obsClient ObsClient) CreateBrowserBasedSignature(input *CreateBrowserBasedSignatureInput) (output *CreateBrowserBasedSignatureOutput, err error) {
	if input == nil {
		return nil, errors.New("CreateBrowserBasedSignatureInput is nil")
	}

	params := make(map[string]string, len(input.FormParams))
	for key, value := range input.FormParams {
		params[key] = value
	}

	date := time.Now().UTC()
	shortDate := date.Format(SHORT_DATE_FORMAT)
	longDate := date.Format(LONG_DATE_FORMAT)

	credential, _ := getCredential(obsClient.conf.securityProvider.ak, obsClient.conf.region, shortDate)

	if input.Expires <= 0 {
		input.Expires = 300
	}

	expiration := date.Add(time.Second * time.Duration(input.Expires)).Format(ISO8601_DATE_FORMAT)
	params[PARAM_ALGORITHM_AMZ_CAMEL] = V4_HASH_PREFIX
	params[PARAM_CREDENTIAL_AMZ_CAMEL] = credential
	params[PARAM_DATE_AMZ_CAMEL] = longDate

	if obsClient.conf.securityProvider.securityToken != "" {
		if obsClient.conf.signature == SignatureObs {
			params[HEADER_STS_TOKEN_OBS] = obsClient.conf.securityProvider.securityToken
		} else {
			params[HEADER_STS_TOKEN_AMZ] = obsClient.conf.securityProvider.securityToken
		}
	}

	matchAnyBucket := true
	matchAnyKey := true
	count := 5
	if bucket := strings.TrimSpace(input.Bucket); bucket != "" {
		params["bucket"] = bucket
		matchAnyBucket = false
		count--
	}

	if key := strings.TrimSpace(input.Key); key != "" {
		params["key"] = key
		matchAnyKey = false
		count--
	}

	originPolicySlice := make([]string, 0, len(params)+count)
	originPolicySlice = append(originPolicySlice, fmt.Sprintf("{\"expiration\":\"%s\",", expiration))
	originPolicySlice = append(originPolicySlice, "\"conditions\":[")
	for key, value := range params {
		if _key := strings.TrimSpace(strings.ToLower(key)); _key != "" {
			originPolicySlice = append(originPolicySlice, fmt.Sprintf("{\"%s\":\"%s\"},", _key, value))
		}
	}

	if matchAnyBucket {
		originPolicySlice = append(originPolicySlice, "[\"starts-with\", \"$bucket\", \"\"],")
	}

	if matchAnyKey {
		originPolicySlice = append(originPolicySlice, "[\"starts-with\", \"$key\", \"\"],")
	}

	originPolicySlice = append(originPolicySlice, "]}")

	originPolicy := strings.Join(originPolicySlice, "")
	policy := Base64Encode([]byte(originPolicy))
	signature := getSignature(policy, obsClient.conf.securityProvider.sk, obsClient.conf.region, shortDate)

	output = &CreateBrowserBasedSignatureOutput{
		OriginPolicy: originPolicy,
		Policy:       policy,
		Algorithm:    params[PARAM_ALGORITHM_AMZ_CAMEL],
		Credential:   params[PARAM_CREDENTIAL_AMZ_CAMEL],
		Date:         params[PARAM_DATE_AMZ_CAMEL],
		Signature:    signature,
	}
	return
}
