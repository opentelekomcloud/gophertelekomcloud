package gpfs

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
	"strings"
	"time"
)

type securityProvider struct {
	ak            string
	sk            string
	securityToken string
}

type config struct {
	securityProvider *securityProvider
	endpoint         *url.URL
	sslVerify        bool
	connectTimeout   int
	socketTimeout    int
	headerTimeout    int
	idleConnTimeout  int
	finalTimeout     int
	maxRetryCount    int
	proxyURL         string
	maxConnsPerHost  int
	pemCerts         []byte
	transport        *http.Transport
	ctx              context.Context
	maxRedirectCount int
}

func (conf config) String() string {
	endpoint := ""
	if conf.endpoint != nil {
		endpoint = conf.endpoint.String()
	}
	return fmt.Sprintf("[endpoint:%s, connectTimeout:%d, socketTimeout:%d, headerTimeout:%d, "+
		"idleConnTimeout:%d, maxRetryCount:%d, maxConnsPerHost:%d, sslVerify:%v, "+
		"proxyConfigured:%t, maxRedirectCount:%d]",
		endpoint, conf.connectTimeout, conf.socketTimeout, conf.headerTimeout,
		conf.idleConnTimeout, conf.maxRetryCount, conf.maxConnsPerHost, conf.sslVerify,
		conf.proxyURL != "", conf.maxRedirectCount,
	)
}

// Configurer configures a Client.
type Configurer func(conf *config)

// WithSslVerify controls TLS certificate verification.
func WithSslVerify(value bool) Configurer { return WithSslVerifyAndPemCerts(value, nil) }

// WithSslVerifyAndPemCerts controls TLS verification and supplies additional root certificates.
func WithSslVerifyAndPemCerts(value bool, certs []byte) Configurer {
	return func(conf *config) { conf.sslVerify, conf.pemCerts = value, certs }
}

// WithHeaderTimeout sets the response-header timeout in seconds.
func WithHeaderTimeout(value int) Configurer {
	return func(conf *config) { conf.headerTimeout = value }
}

// WithProxyUrl sets the HTTP proxy URL.
func WithProxyUrl(value string) Configurer { return func(conf *config) { conf.proxyURL = value } }

// WithMaxConnections sets the maximum number of idle connections per host.
func WithMaxConnections(value int) Configurer {
	return func(conf *config) { conf.maxConnsPerHost = value }
}

// WithConnectTimeout sets the connection timeout in seconds.
func WithConnectTimeout(value int) Configurer {
	return func(conf *config) { conf.connectTimeout = value }
}

// WithSocketTimeout sets the socket timeout in seconds.
func WithSocketTimeout(value int) Configurer {
	return func(conf *config) { conf.socketTimeout = value }
}

// WithIdleConnTimeout sets the idle-connection timeout in seconds.
func WithIdleConnTimeout(value int) Configurer {
	return func(conf *config) { conf.idleConnTimeout = value }
}

// WithMaxRetryCount sets the maximum number of retries for transient failures.
func WithMaxRetryCount(value int) Configurer {
	return func(conf *config) { conf.maxRetryCount = value }
}

// WithSecurityToken sets the token used with temporary access keys.
func WithSecurityToken(value string) Configurer {
	return func(conf *config) { conf.securityProvider.securityToken = value }
}

// WithHttpTransport sets a custom HTTP transport.
func WithHttpTransport(value *http.Transport) Configurer {
	return func(conf *config) { conf.transport = value }
}

// WithRequestContext sets the context used for requests.
func WithRequestContext(value context.Context) Configurer {
	return func(conf *config) { conf.ctx = value }
}

// WithMaxRedirectCount sets the maximum number of redirects.
func WithMaxRedirectCount(value int) Configurer {
	return func(conf *config) { conf.maxRedirectCount = value }
}

func (conf *config) prepareDefaults() {
	if conf.connectTimeout <= 0 {
		conf.connectTimeout = defaultConnectTimeout
	}
	if conf.socketTimeout <= 0 {
		conf.socketTimeout = defaultSocketTimeout
	}
	conf.finalTimeout = conf.socketTimeout * 10
	if conf.headerTimeout <= 0 {
		conf.headerTimeout = defaultHeaderTimeout
	}
	if conf.idleConnTimeout < 0 {
		conf.idleConnTimeout = defaultIdleConnTimeout
	}
	if conf.maxRetryCount < 0 {
		conf.maxRetryCount = defaultMaxRetryCount
	}
	if conf.maxConnsPerHost <= 0 {
		conf.maxConnsPerHost = defaultMaxConnsPerHost
	}
	if conf.maxRedirectCount < 0 {
		conf.maxRedirectCount = defaultMaxRedirectCount
	}
}

func (conf *config) initialize(endpoint string) error {
	conf.securityProvider.ak = strings.TrimSpace(conf.securityProvider.ak)
	conf.securityProvider.sk = strings.TrimSpace(conf.securityProvider.sk)
	conf.securityProvider.securityToken = strings.TrimSpace(conf.securityProvider.securityToken)
	endpoint = strings.TrimSpace(endpoint)
	if endpoint == "" {
		return errors.New("endpoint is not set")
	}
	if !strings.Contains(endpoint, "://") {
		endpoint = "https://" + endpoint
	}
	parsed, err := url.Parse(endpoint)
	if err != nil {
		return errors.New("invalid endpoint")
	}
	if parsed.Hostname() == "" {
		return errors.New("endpoint host is not set")
	}
	if parsed.User != nil {
		return errors.New("endpoint must not contain user information")
	}
	parsed.Path, parsed.RawPath, parsed.RawQuery, parsed.Fragment = "", "", "", ""
	conf.endpoint = parsed
	conf.proxyURL = strings.TrimSpace(conf.proxyURL)
	conf.prepareDefaults()
	return nil
}

func (conf *config) getTransport() error {
	if conf.transport != nil {
		return nil
	}
	conf.transport = &http.Transport{
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			conn, err := (&net.Dialer{Timeout: time.Duration(conf.connectTimeout) * time.Second}).DialContext(ctx, network, address)
			if err != nil {
				return nil, err
			}
			return getConnDelegate(conn, conf.socketTimeout, conf.finalTimeout), nil
		},
		MaxIdleConns: conf.maxConnsPerHost, MaxIdleConnsPerHost: conf.maxConnsPerHost,
		ResponseHeaderTimeout: time.Duration(conf.headerTimeout) * time.Second,
		IdleConnTimeout:       time.Duration(conf.idleConnTimeout) * time.Second,
	}
	if conf.proxyURL != "" {
		proxy, err := url.Parse(conf.proxyURL)
		if err != nil {
			return errors.New("invalid proxy URL")
		}
		conf.transport.Proxy = http.ProxyURL(proxy)
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: !conf.sslVerify} // #nosec G402 -- explicitly configurable for private endpoints.
	if conf.sslVerify && len(conf.pemCerts) > 0 {
		pool, err := x509.SystemCertPool()
		if err != nil {
			pool = x509.NewCertPool()
		}
		if !pool.AppendCertsFromPEM(conf.pemCerts) {
			return errors.New("failed to parse PEM certificates")
		}
		tlsConfig.RootCAs = pool
	}
	conf.transport.TLSClientConfig = tlsConfig
	return nil
}

func checkRedirectFunc(_ *http.Request, _ []*http.Request) error { return http.ErrUseLastResponse }

func (conf *config) formatURLs(fsName string, params map[string]string) (string, string) {
	requestURL := *conf.endpoint
	canonicalURL := "/"
	if fsName != "" {
		host := fsName + "." + conf.endpoint.Hostname()
		if port := conf.endpoint.Port(); port != "" {
			host = net.JoinHostPort(host, port)
		}
		requestURL.Host, requestURL.Path = host, "/"
		canonicalURL = "/" + fsName + "/"
	}
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, strings.TrimSpace(key))
	}
	sort.Strings(keys)
	query := make([]string, 0, len(keys))
	for _, key := range keys {
		part := url.QueryEscape(key)
		if value := params[key]; value != "" {
			part += "=" + url.QueryEscape(value)
		}
		query = append(query, part)
	}
	requestURL.RawQuery = strings.Join(query, "&")
	if len(query) > 0 {
		canonicalURL += "?" + strings.Join(query, "&")
	}
	return requestURL.String(), canonicalURL
}
