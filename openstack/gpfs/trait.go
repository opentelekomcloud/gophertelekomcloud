package gpfs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// IReadCloser defines interface with function: setReadCloser
type IReadCloser interface {
	setReadCloser(body io.ReadCloser)
}

func (output *GetObjectOutput) setReadCloser(body io.ReadCloser) {
	output.Body = body
}

func setHeaders(headers map[string][]string, header string, headerValue []string, isObs bool) {
	if isObs {
		header = HEADER_PREFIX_OBS + header
		headers[header] = headerValue
	} else {
		header = HEADER_PREFIX + header
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

// IBaseModel defines interface for base response model
type IBaseModel interface {
	setStatusCode(statusCode int)

	setRequestId(requestId string)

	setResponseHeaders(responseHeaders map[string][]string)
}

// ISerializable defines interface with function: trans
type ISerializable interface {
	trans(isObs bool) (map[string]string, map[string][]string, interface{}, error)
}

// DefaultSerializable defines default serializable struct
type DefaultSerializable struct {
	params  map[string]string
	headers map[string][]string
	data    interface{}
}

func (input DefaultSerializable) trans(_ bool) (map[string]string, map[string][]string, interface{}, error) {
	return input.params, input.headers, input.data, nil
}

var defaultSerializable = &DefaultSerializable{}

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
		setHeaders(headers, HEADER_LOCATION_AMZ, []string{"true"}, isObs)
	}
	if input.BucketType != "" {
		setHeaders(headers, HEADER_BUCKET_TYPE, []string{string(input.BucketType)}, true)
	}
	return
}

func (input CreateFSInput) prepareGrantHeaders(headers map[string][]string, isObs bool) {
	if grantReadID := input.GrantReadId; grantReadID != "" {
		setHeaders(headers, HEADER_GRANT_READ_OBS, []string{grantReadID}, isObs)
	}
	if grantWriteID := input.GrantWriteId; grantWriteID != "" {
		setHeaders(headers, HEADER_GRANT_WRITE_OBS, []string{grantWriteID}, isObs)
	}
	if grantReadAcpID := input.GrantReadAcpId; grantReadAcpID != "" {
		setHeaders(headers, HEADER_GRANT_READ_ACP_OBS, []string{grantReadAcpID}, isObs)
	}
	if grantWriteAcpID := input.GrantWriteAcpId; grantWriteAcpID != "" {
		setHeaders(headers, HEADER_GRANT_WRITE_ACP_OBS, []string{grantWriteAcpID}, isObs)
	}
	if grantFullControlID := input.GrantFullControlId; grantFullControlID != "" {
		setHeaders(headers, HEADER_GRANT_FULL_CONTROL_OBS, []string{grantFullControlID}, isObs)
	}
	if grantReadDeliveredID := input.GrantReadDeliveredId; grantReadDeliveredID != "" {
		setHeaders(headers, HEADER_GRANT_READ_DELIVERED_OBS, []string{grantReadDeliveredID}, true)
	}
	if grantFullControlDeliveredID := input.GrantFullControlDeliveredId; grantFullControlDeliveredID != "" {
		setHeaders(headers, HEADER_GRANT_FULL_CONTROL_DELIVERED_OBS, []string{grantFullControlDeliveredID}, true)
	}
}

func (input CreateFSInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	headers = make(map[string][]string)
	if acl := string(input.ACL); acl != "" {
		setHeaders(headers, HEADER_ACL, []string{acl}, isObs)
	}
	if storageClass := string(input.StorageClass); storageClass != "" {
		if !isObs {
			if storageClass == "WARM" {
				storageClass = "STANDARD_IA"
			} else if storageClass == "COLD" {
				storageClass = "GLACIER"
			}
		}
		setHeadersNext(headers, HEADER_STORAGE_CLASS_OBS, HEADER_STORAGE_CLASS, []string{storageClass}, isObs)
	}
	if epid := input.Epid; epid != "" {
		setHeaders(headers, HEADER_EPID_HEADERS, []string{epid}, isObs)
	}

	input.prepareGrantHeaders(headers, isObs)
	if input.IsFSFileInterface {
		setHeaders(headers, HEADER_FS_FILE_INTERFACE, []string{"Enabled"}, true)
	}

	if input.ObjectLockEnabled {
		setHeaders(headers, HEADER_OBJECT_LOCK_ENABLED, []string{"true"}, true)
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

type readerWrapper struct {
	reader      io.Reader
	mark        int64 // nolint: structcheck
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
