package gpfs

import "net/http"

// Client is a client for the SFS Turbo general-purpose file system API.
type Client struct {
	conf       *config
	httpClient *http.Client
}

// New creates a GPFS client. The endpoint is expected to be the SFS3 endpoint.
func New(ak, sk, endpoint string, configurers ...Configurer) (*Client, error) {
	conf := &config{
		securityProvider: &securityProvider{ak: ak, sk: sk},
		sslVerify:        true,
		maxRetryCount:    -1,
		maxRedirectCount: -1,
	}
	for _, configure := range configurers {
		configure(conf)
	}
	if err := conf.initialize(endpoint); err != nil {
		return nil, err
	}
	if err := conf.getTransport(); err != nil {
		return nil, err
	}
	return &Client{
		conf: conf,
		httpClient: &http.Client{
			Transport:     conf.transport,
			CheckRedirect: checkRedirectFunc,
		},
	}, nil
}
