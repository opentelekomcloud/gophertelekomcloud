package obs

import (
	"strings"
)

// Refresh refreshes ak, sk and securityToken for obsClient.
func (obsClient ObsClient) Refresh(ak, sk, securityToken string) {
	for _, sp := range obsClient.conf.securityProviders {
		if bsp, ok := sp.(*BasicSecurityProvider); ok {
			bsp.refresh(strings.TrimSpace(ak), strings.TrimSpace(sk), strings.TrimSpace(securityToken))
			break
		}
	}
}

func (obsClient ObsClient) getSecurity() securityHolder {
	if obsClient.conf.securityProviders != nil {
		for _, sp := range obsClient.conf.securityProviders {
			if sp == nil {
				continue
			}
			sh := sp.getSecurity()
			if sh.ak != "" && sh.sk != "" {
				return sh
			}
		}
	}
	return emptySecurityHolder
}

// Close closes ObsClient.
func (obsClient ObsClient) Close() {
	obsClient.httpClient = nil
	obsClient.conf.transport.CloseIdleConnections()
	obsClient.conf = nil
}
