package gpfs

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	regex   = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	ipRegex = regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
)

// StringContains replaces subStr in src with subTranscoding and returns the new string
func StringContains(src string, subStr string, subTranscoding string) string {
	return strings.Replace(src, subStr, subTranscoding, -1)
}

// XmlTranscoding replaces special characters with their escaped form
func XmlTranscoding(src string) string {
	srcTmp := StringContains(src, "&", "&amp;")
	srcTmp = StringContains(srcTmp, "<", "&lt;")
	srcTmp = StringContains(srcTmp, ">", "&gt;")
	srcTmp = StringContains(srcTmp, "'", "&apos;")
	srcTmp = StringContains(srcTmp, "\"", "&quot;")
	return srcTmp
}

// StringToInt converts string value to int value with default value
func StringToInt(value string, def int) int {
	ret, err := strconv.Atoi(value)
	if err != nil {
		ret = def
	}
	return ret
}

// StringToInt64 converts string value to int64 value with default value
func StringToInt64(value string, def int64) int64 {
	ret, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		ret = def
	}
	return ret
}

// IntToString converts int value to string value
func IntToString(value int) string {
	return strconv.Itoa(value)
}

// Int64ToString converts int64 value to string value
func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

// GetCurrentTimestamp gets unix time in milliseconds
func GetCurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}

// FormatUtcNow gets a textual representation of the UTC format time value
func FormatUtcNow(format string) string {
	return time.Now().UTC().Format(format)
}

// FormatUtcToRfc1123 gets a textual representation of the RFC1123 format time value
func FormatUtcToRfc1123(t time.Time) string {
	ret := t.UTC().Format(time.RFC1123)
	return ret[:strings.LastIndex(ret, "UTC")] + "GMT"
}

// Md5 gets the md5 value of input
func Md5(value []byte) []byte {
	m := md5.New()
	_, err := m.Write(value)
	if err != nil {
		doLog(LEVEL_WARN, "MD5 failed to write with reason: %v", err)
	}
	return m.Sum(nil)
}

// HmacSha1 gets hmac sha1 value of input
func HmacSha1(key, value []byte) []byte {
	mac := hmac.New(sha1.New, key)
	_, err := mac.Write(value)
	if err != nil {
		doLog(LEVEL_WARN, "HmacSha1 failed to write with reason: %v", err)
	}
	return mac.Sum(nil)
}

// HmacSha256 get hmac sha256 value if input
func HmacSha256(key, value []byte) []byte {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(value)
	if err != nil {
		doLog(LEVEL_WARN, "HmacSha256 failed to write with reason: %v", err)
	}
	return mac.Sum(nil)
}

// Base64Encode wrapper of base64.StdEncoding.EncodeToString
func Base64Encode(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

// Base64Decode wrapper of base64.StdEncoding.DecodeString
func Base64Decode(value string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(value)
}

// HexMd5 returns the md5 value of input in hexadecimal format
func HexMd5(value []byte) string {
	return Hex(Md5(value))
}

// Base64Md5 returns the md5 value of input with Base64Encode
func Base64Md5(value []byte) string {
	return Base64Encode(Md5(value))
}

// Sha256Hash returns sha256 checksum
func Sha256Hash(value []byte) []byte {
	hash := sha256.New()
	_, err := hash.Write(value)
	if err != nil {
		doLog(LEVEL_WARN, "Sha256Hash failed to write with reason: %v", err)
	}
	return hash.Sum(nil)
}

// ParseXml wrapper of xml.Unmarshal
func ParseXml(value []byte, result interface{}) error {
	if len(value) == 0 {
		return nil
	}
	return xml.Unmarshal(value, result)
}

// parseJSON wrapper of json.Unmarshal
func parseJSON(value []byte, result interface{}) error {
	if len(value) == 0 {
		return nil
	}
	return json.Unmarshal(value, result)
}

// TransToXml wrapper of xml.Marshal
func TransToXml(value interface{}) ([]byte, error) {
	if value == nil {
		return []byte{}, nil
	}
	return xml.Marshal(value)
}

// Hex wrapper of hex.EncodeToString
func Hex(value []byte) string {
	return hex.EncodeToString(value)
}

// HexSha256 returns the Sha256Hash value of input in hexadecimal format
func HexSha256(value []byte) string {
	return Hex(Sha256Hash(value))
}

// UrlDecode wrapper of url.QueryUnescape
func UrlDecode(value string) (string, error) {
	ret, err := url.QueryUnescape(value)
	if err == nil {
		return ret, nil
	}
	return "", err
}

// UrlDecodeWithoutError wrapper of UrlDecode
func UrlDecodeWithoutError(value string) string {
	ret, err := UrlDecode(value)
	if err == nil {
		return ret
	}
	if isErrorLogEnabled() {
		doLog(LEVEL_ERROR, "Url decode error: %v", err)
	}
	return ""
}

// IsIP checks whether the value matches ip address
func IsIP(value string) bool {
	return ipRegex.MatchString(value)
}

// UrlEncode encodes the input value
func UrlEncode(value string, chineseOnly bool) string {
	if chineseOnly {
		values := make([]string, 0, len(value))
		for _, val := range value {
			_value := string(val)
			if regex.MatchString(_value) {
				_value = url.QueryEscape(_value)
			}
			values = append(values, _value)
		}
		return strings.Join(values, "")
	}
	return url.QueryEscape(value)
}
