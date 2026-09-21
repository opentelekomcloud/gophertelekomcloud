package listeners_test

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func serviceClient() *golangsdk.ServiceClient {
	return client.ServiceClient()
}

const listenerJSON = `{
	"id":"listener-id","admin_state_up":true,"client_ca_tls_container_ref":"ca-id",
	"connection_limit":-1,"created_at":"2026-09-21T08:00:00Z","updated_at":"2026-09-21T08:01:00Z",
	"default_pool_id":"pool-id","default_tls_container_ref":"certificate-id","description":"listener",
	"http2_enable":true,"loadbalancers":[{"id":"loadbalancer-id"}],"name":"listener-test",
	"project_id":"project-id","protocol":"HTTPS","protocol_port":443,
	"sni_container_refs":["sni-id"],"sni_match_algo":"wildcard","tags":[{"key":"key","value":"value"}],
	"tls_ciphers_policy":"tls-1-2","security_policy_id":"policy-id","enable_member_retry":true,
	"keepalive_timeout":60,"client_timeout":60,"member_timeout":60,
	"ipgroup":{"ipgroup_id":"ipgroup-id","enable_ipgroup":true,"type":"white"},
	"insert_headers":{"X-Forwarded-ELB-IP":true},"transparent_client_ip_enable":true,
	"enhance_l7policy_enable":true,"quic_config":{"quic_listener_id":"quic-id","enable_quic_upgrade":true},
	"protection_status":"consoleProtection","protection_reason":"managed","gzip_enable":true,
	"cps":100,"connection":1000,"nat64_enable":true,"proxy_protocol_enable":true,
	"tracing_config":{"tracing_enable":true,"tracing_strategy":"ratio","tracing_sample":3000,"tracing_type":"W3CTraceContext"},
	"access_log_customized_headers_config":{"enable":true,"include_headers":["X-Test"],"exclude_headers":["X-Skip"]}
}`
