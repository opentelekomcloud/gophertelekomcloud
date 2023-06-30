package obs

import (
	"strings"
)

// Refresh refreshes ak, sk and securityToken for obsClient.
func (obsClient ObsClient) Refresh(ak, sk, securityToken string) {
	sp := &securityProvider{ak: strings.TrimSpace(ak), sk: strings.TrimSpace(sk), securityToken: strings.TrimSpace(securityToken)}
	obsClient.conf.securityProvider = sp
}

// Close closes ObsClient.
func (obsClient ObsClient) Close() {
	obsClient.conf.transport.CloseIdleConnections()
}
