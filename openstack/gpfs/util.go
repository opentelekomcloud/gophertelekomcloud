package gpfs

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"strings"
	"time"
)

func formatRFC1123(value time.Time) string {
	return strings.Replace(value.UTC().Format(time.RFC1123), "UTC", "GMT", 1)
}

func hmacSHA1(key, value []byte) []byte {
	digest := hmac.New(sha1.New, key)
	_, _ = digest.Write(value)
	return digest.Sum(nil)
}

func base64Encode(value []byte) string { return base64.StdEncoding.EncodeToString(value) }

func parseXML(value []byte, result interface{}) error  { return xml.Unmarshal(value, result) }
func parseJSON(value []byte, result interface{}) error { return json.Unmarshal(value, result) }
func marshalXML(value interface{}) ([]byte, error)     { return xml.Marshal(value) }
