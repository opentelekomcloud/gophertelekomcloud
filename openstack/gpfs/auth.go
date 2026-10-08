package gpfs

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

func (client Client) authorize(method, fsName string, params map[string]string, headers map[string][]string, redirectHost string) (string, error) {
	credentials := client.conf.securityProvider
	if credentials != nil && credentials.securityToken != "" {
		headers[headerSecurityTokenOBS] = []string{credentials.securityToken}
	}
	requestURL, canonicalURL := client.conf.formatURLs(fsName, params)
	parsed, err := url.Parse(requestURL)
	if err != nil {
		return "", fmt.Errorf("invalid request URL: %w", err)
	}
	host := parsed.Host
	if redirectHost != "" {
		host = redirectHost
	}
	prepareHostAndDate(headers, host)
	if credentials == nil || credentials.ak == "" || credentials.sk == "" {
		return requestURL, nil
	}
	stringToSign := method + "\n" + canonicalHeaders(headers) + "\n" + canonicalURL
	signature := base64Encode(hmacSHA1([]byte(credentials.sk), []byte(stringToSign)))
	headers[headerAuthorization] = []string{fmt.Sprintf("OBS %s:%s", credentials.ak, signature)}
	return requestURL, nil
}

func prepareHostAndDate(headers map[string][]string, host string) {
	headers[headerHost] = []string{host}
	for key := range headers {
		if strings.EqualFold(key, headerDateOBS) {
			return
		}
	}
	headers[headerDateOBS] = []string{formatRFC1123(time.Now())}
}

func canonicalHeaders(headers map[string][]string) string {
	canonical := make(map[string][]string, len(headers)+3)
	keys := make([]string, 0, len(headers)+3)
	for key, values := range headers {
		name := strings.ToLower(strings.TrimSpace(key))
		if name == "content-md5" || name == "content-type" || name == "date" || strings.HasPrefix(name, headerPrefixOBS) {
			canonical[name] = values
			keys = append(keys, name)
		}
	}
	for _, name := range []string{"content-md5", "content-type", "date"} {
		if _, ok := canonical[name]; !ok {
			canonical[name] = []string{""}
			keys = append(keys, name)
		}
	}
	if _, ok := canonical[headerDateOBS]; ok {
		canonical["date"] = []string{""}
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, name := range keys {
		value := strings.Join(canonical[name], ",")
		if strings.HasPrefix(name, headerPrefixOBS) {
			value = name + ":" + value
		}
		parts = append(parts, value)
	}
	return strings.Join(parts, "\n")
}
