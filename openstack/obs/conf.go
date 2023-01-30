package obs

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type securityProvider struct {
	ak            string
	sk            string
	securityToken string
}

type urlHolder struct {
	scheme string
	host   string
	port   int
}

type config struct {
	securityProvider *securityProvider
	urlHolder        *urlHolder
	pathStyle        bool
	cname            bool
	sslVerify        bool
	endpoint         string
	signature        SignatureType
	region           string
	connectTimeout   int
	socketTimeout    int
	headerTimeout    int
	idleConnTimeout  int
	finalTimeout     int
	maxRetryCount    int
	proxyUrl         string
	maxConnsPerHost  int
	pemCerts         []byte
	transport        *http.Transport
	ctx              context.Context
	maxRedirectCount int
}

func (conf config) String() string {
	return fmt.Sprintf("[endpoint:%s, signature:%s, pathStyle:%v, region:%s"+
		"\nconnectTimeout:%d, socketTimeout:%dheaderTimeout:%d, idleConnTimeout:%d"+
		"\nmaxRetryCount:%d, maxConnsPerHost:%d, sslVerify:%v, proxyUrl:%s, maxRedirectCount:%d]",
		conf.endpoint, conf.signature, conf.pathStyle, conf.region,
		conf.connectTimeout, conf.socketTimeout, conf.headerTimeout, conf.idleConnTimeout,
		conf.maxRetryCount, conf.maxConnsPerHost, conf.sslVerify, conf.proxyUrl, conf.maxRedirectCount,
	)
}

type Configurer func(conf *config)

// WithSslVerify is a wrapper for WithSslVerifyAndPemCerts.
func WithSslVerify(sslVerify bool) Configurer {
	return WithSslVerifyAndPemCerts(sslVerify, nil)
}

// WithSslVerifyAndPemCerts is a configurer for ObsClient to set conf.sslVerify and conf.pemCerts.
func WithSslVerifyAndPemCerts(sslVerify bool, pemCerts []byte) Configurer {
	return func(conf *config) {
		conf.sslVerify = sslVerify
		conf.pemCerts = pemCerts
	}
}

// WithHeaderTimeout is a configurer for ObsClient to set the timeout period of obtaining the response headers.
func WithHeaderTimeout(headerTimeout int) Configurer {
	return func(conf *config) {
		conf.headerTimeout = headerTimeout
	}
}

// WithProxyUrl is a configurer for ObsClient to set HTTP proxy.
func WithProxyUrl(proxyUrl string) Configurer {
	return func(conf *config) {
		conf.proxyUrl = proxyUrl
	}
}

// WithMaxConnections is a configurer for ObsClient to set the maximum number of idle HTTP connections.
func WithMaxConnections(maxConnsPerHost int) Configurer {
	return func(conf *config) {
		conf.maxConnsPerHost = maxConnsPerHost
	}
}

// WithPathStyle is a configurer for ObsClient.
func WithPathStyle(pathStyle bool) Configurer {
	return func(conf *config) {
		conf.pathStyle = pathStyle
	}
}

// WithSignature is a configurer for ObsClient.
func WithSignature(signature SignatureType) Configurer {
	return func(conf *config) {
		conf.signature = signature
	}
}

// WithRegion is a configurer for ObsClient.
func WithRegion(region string) Configurer {
	return func(conf *config) {
		conf.region = region
	}
}

// WithConnectTimeout is a configurer for ObsClient to set timeout period for establishing
// a http/https connection, in seconds.
func WithConnectTimeout(connectTimeout int) Configurer {
	return func(conf *config) {
		conf.connectTimeout = connectTimeout
	}
}

// WithSocketTimeout is a configurer for ObsClient to set the timeout duration for transmitting data at
// the socket layer, in seconds.
func WithSocketTimeout(socketTimeout int) Configurer {
	return func(conf *config) {
		conf.socketTimeout = socketTimeout
	}
}

// WithIdleConnTimeout is a configurer for ObsClient to set the timeout period of an idle HTTP connection
// in the connection pool, in seconds.
func WithIdleConnTimeout(idleConnTimeout int) Configurer {
	return func(conf *config) {
		conf.idleConnTimeout = idleConnTimeout
	}
}

// WithMaxRetryCount is a configurer for ObsClient to set the maximum number of retries when an HTTP/HTTPS connection is abnormal.
func WithMaxRetryCount(maxRetryCount int) Configurer {
	return func(conf *config) {
		conf.maxRetryCount = maxRetryCount
	}
}

// WithSecurityToken is a configurer for ObsClient to set the security token in the temporary access keys.
func WithSecurityToken(securityToken string) Configurer {
	return func(conf *config) {
		conf.securityProvider.securityToken = securityToken
	}
}

// WithHttpTransport is a configurer for ObsClient to set the customized http Transport.
func WithHttpTransport(transport *http.Transport) Configurer {
	return func(conf *config) {
		conf.transport = transport
	}
}

// WithRequestContext is a configurer for ObsClient to set the context for each HTTP request.
func WithRequestContext(ctx context.Context) Configurer {
	return func(conf *config) {
		conf.ctx = ctx
	}
}

// WithCustomDomainName is a configurer for ObsClient.
func WithCustomDomainName(cname bool) Configurer {
	return func(conf *config) {
		conf.cname = cname
	}
}

// WithMaxRedirectCount is a configurer for ObsClient to set the maximum number of times that the request is redirected.
func WithMaxRedirectCount(maxRedirectCount int) Configurer {
	return func(conf *config) {
		conf.maxRedirectCount = maxRedirectCount
	}
}

func (conf *config) prepareConfig() {
	if conf.connectTimeout <= 0 {
		conf.connectTimeout = DEFAULT_CONNECT_TIMEOUT
	}

	if conf.socketTimeout <= 0 {
		conf.socketTimeout = DEFAULT_SOCKET_TIMEOUT
	}

	conf.finalTimeout = conf.socketTimeout * 10

	if conf.headerTimeout <= 0 {
		conf.headerTimeout = DEFAULT_HEADER_TIMEOUT
	}

	if conf.idleConnTimeout < 0 {
		conf.idleConnTimeout = DEFAULT_IDLE_CONN_TIMEOUT
	}

	if conf.maxRetryCount < 0 {
		conf.maxRetryCount = DEFAULT_MAX_RETRY_COUNT
	}

	if conf.maxConnsPerHost <= 0 {
		conf.maxConnsPerHost = DEFAULT_MAX_CONN_PER_HOST
	}

	if conf.maxRedirectCount < 0 {
		conf.maxRedirectCount = DEFAULT_MAX_REDIRECT_COUNT
	}

	if conf.pathStyle && conf.signature == SignatureObs {
		conf.signature = SignatureV2
	}
}

func (conf *config) initConfigWithDefault() error {
	conf.securityProvider.ak = strings.TrimSpace(conf.securityProvider.ak)
	conf.securityProvider.sk = strings.TrimSpace(conf.securityProvider.sk)
	conf.securityProvider.securityToken = strings.TrimSpace(conf.securityProvider.securityToken)
	conf.endpoint = strings.TrimSpace(conf.endpoint)
	if conf.endpoint == "" {
		return errors.New("endpoint is not set")
	}

	if index := strings.Index(conf.endpoint, "?"); index > 0 {
		conf.endpoint = conf.endpoint[:index]
	}

	for strings.LastIndex(conf.endpoint, "/") == len(conf.endpoint)-1 {
		conf.endpoint = conf.endpoint[:len(conf.endpoint)-1]
	}

	if conf.signature == "" {
		conf.signature = DEFAULT_SIGNATURE
	}

	urlHolder := &urlHolder{}
	var address string
	if strings.HasPrefix(conf.endpoint, "https://") {
		urlHolder.scheme = "https"
		address = conf.endpoint[len("https://"):]
	} else if strings.HasPrefix(conf.endpoint, "http://") {
		urlHolder.scheme = "http"
		address = conf.endpoint[len("http://"):]
	} else {
		urlHolder.scheme = "https"
		address = conf.endpoint
	}

	addr := strings.Split(address, ":")
	if len(addr) == 2 {
		if port, err := strconv.Atoi(addr[1]); err == nil {
			urlHolder.port = port
		}
	}
	urlHolder.host = addr[0]
	if urlHolder.port == 0 {
		if urlHolder.scheme == "https" {
			urlHolder.port = 443
		} else {
			urlHolder.port = 80
		}
	}

	if IsIP(urlHolder.host) {
		conf.pathStyle = true
	}

	conf.urlHolder = urlHolder

	conf.region = strings.TrimSpace(conf.region)
	if conf.region == "" {
		conf.region = DEFAULT_REGION
	}

	conf.prepareConfig()
	conf.proxyUrl = strings.TrimSpace(conf.proxyUrl)
	return nil
}

func (conf *config) getTransport() error {
	if conf.transport == nil {
		conf.transport = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(network, addr, time.Second*time.Duration(conf.connectTimeout))
				if err != nil {
					return nil, err
				}
				return getConnDelegate(conn, conf.socketTimeout, conf.finalTimeout), nil
			},
			MaxIdleConns:          conf.maxConnsPerHost,
			MaxIdleConnsPerHost:   conf.maxConnsPerHost,
			ResponseHeaderTimeout: time.Second * time.Duration(conf.headerTimeout),
			IdleConnTimeout:       time.Second * time.Duration(conf.idleConnTimeout),
		}

		if conf.proxyUrl != "" {
			proxyUrl, err := url.Parse(conf.proxyUrl)
			if err != nil {
				return err
			}
			conf.transport.Proxy = http.ProxyURL(proxyUrl)
		}

		tlsConfig := &tls.Config{InsecureSkipVerify: !conf.sslVerify}
		if conf.sslVerify && conf.pemCerts != nil {
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(conf.pemCerts)
			tlsConfig.RootCAs = pool
		}

		conf.transport.TLSClientConfig = tlsConfig
	}

	return nil
}

func checkRedirectFunc(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}

func DummyQueryEscape(s string) string {
	return s
}

func (conf *config) prepareBaseURL(bucketName string) (requestURL string, canonicalizedURL string) {
	urlHolder := conf.urlHolder
	if conf.cname {
		requestURL = fmt.Sprintf("%s://%s:%d", urlHolder.scheme, urlHolder.host, urlHolder.port)
		if conf.signature == "v4" {
			canonicalizedURL = "/"
		} else {
			canonicalizedURL = "/" + urlHolder.host + "/"
		}
	} else {
		if bucketName == "" {
			requestURL = fmt.Sprintf("%s://%s:%d", urlHolder.scheme, urlHolder.host, urlHolder.port)
			canonicalizedURL = "/"
		} else {
			if conf.pathStyle {
				requestURL = fmt.Sprintf("%s://%s:%d/%s", urlHolder.scheme, urlHolder.host, urlHolder.port, bucketName)
				canonicalizedURL = "/" + bucketName
			} else {
				requestURL = fmt.Sprintf("%s://%s.%s:%d", urlHolder.scheme, bucketName, urlHolder.host, urlHolder.port)
				if conf.signature == "v2" || conf.signature == "OBS" {
					canonicalizedURL = "/" + bucketName + "/"
				} else {
					canonicalizedURL = "/"
				}
			}
		}
	}
	return
}

func (conf *config) prepareObjectKey(escape bool, objectKey string, escapeFunc func(s string) string) (encodeObjectKey string) {
	if escape {
		tempKey := []rune(objectKey)
		result := make([]string, 0, len(tempKey))
		for _, value := range tempKey {
			if string(value) == "/" {
				result = append(result, string(value))
			} else {
				if string(value) == " " {
					result = append(result, url.PathEscape(string(value)))
				} else {
					result = append(result, url.QueryEscape(string(value)))
				}
			}
		}
		encodeObjectKey = strings.Join(result, "")
	} else {
		encodeObjectKey = escapeFunc(objectKey)
	}
	return
}

func (conf *config) prepareEscapeFunc(escape bool) (escapeFunc func(s string) string) {
	if escape {
		return url.QueryEscape
	}
	return DummyQueryEscape
}

func (conf *config) formatUrls(bucketName, objectKey string, params map[string]string, escape bool) (requestUrl string, canonicalizedUrl string) {
	requestUrl, canonicalizedUrl = conf.prepareBaseURL(bucketName)

	if objectKey != "" {
		encodeObjectKey := conf.prepareObjectKey(escape, objectKey, conf.prepareEscapeFunc(escape))
		requestUrl += "/" + encodeObjectKey
		if !strings.HasSuffix(canonicalizedUrl, "/") {
			canonicalizedUrl += "/"
		}
		canonicalizedUrl += encodeObjectKey
	}

	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, strings.TrimSpace(key))
	}
	sort.Strings(keys)
	i := 0

	for index, key := range keys {
		if index == 0 {
			requestUrl += "?"
		} else {
			requestUrl += "&"
		}
		_key := url.QueryEscape(key)
		requestUrl += _key

		_value := params[key]
		if conf.signature == "v4" {
			requestUrl += "=" + url.QueryEscape(_value)
		} else {
			if _value != "" {
				requestUrl += "=" + url.QueryEscape(_value)
				_value = "=" + _value
			} else {
				_value = ""
			}
			lowerKey := strings.ToLower(key)
			_, ok := allowedResourceParameterNames[lowerKey]
			prefixHeader := HEADER_PREFIX
			isObs := conf.signature == SignatureObs
			if isObs {
				prefixHeader = HEADER_PREFIX_OBS
			}
			ok = ok || strings.HasPrefix(lowerKey, prefixHeader)
			if ok {
				if i == 0 {
					canonicalizedUrl += "?"
				} else {
					canonicalizedUrl += "&"
				}
				canonicalizedUrl += getQueryUrl(_key, _value)
				i++
			}
		}
	}
	return
}

func getQueryUrl(key, value string) string {
	queryUrl := ""
	queryUrl += key
	queryUrl += value
	return queryUrl
}
