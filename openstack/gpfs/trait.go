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

func (input ListFSInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	headers = make(map[string][]string)
	if input.BucketType != "" {
		setHeaders(headers, HEADER_BUCKET_TYPE, []string{string(input.BucketType)}, true)
	} else {
		panic("BucketType is a required parameter.")
	}
	return
}

func (input CreateFSInput) trans(isObs bool) (params map[string]string, headers map[string][]string, data interface{}, err error) {
	headers = make(map[string][]string)
	if redundancy := input.Redundancy; redundancy != "" {
		setHeaders(headers, HEADER_AZ_REDUNDANCY, []string{redundancy}, isObs)
	} else {
		panic("Redundancy is a required parameter.")
	}
	if bucketType := input.BucketType; bucketType != "" {
		setHeaders(headers, HEADER_BUCKET_TYPE, []string{bucketType}, isObs)
	} else {
		panic("BucketType is a required parameter.")
	}

	if location := strings.TrimSpace(input.Location); location != "" {
		input.Location = location

		xml := make([]string, 0, 3)
		xml = append(xml, "<CreateBucketConfiguration>")
		if isObs {
			xml = append(xml, fmt.Sprintf("<Location>%s</Location>", input.Location))
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
