package v1

import (
	"log"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/load_balancer"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestGetCSSConfiguringLBListenersHTTP(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}

	configuringOpts := load_balancer.LoadBalancingListenerOpts{
		Protocol:     "HTTP",
		ProtocolPort: 80,
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	got, err := load_balancer.ConfigureLoadBalancingListeners(client, clusterID, configuringOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, got)
}

func TestCSSLoadBalancerFullLifecycle(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	agency := clients.EnvOS.GetEnv("AGENCY_NAME")
	if agency == "" {
		t.Skipf("OS_AGENCY_NAME is required for this test")
	}
	elbid := clients.EnvOS.GetEnv("ELB_ID")
	if elbid == "" {
		t.Skipf("OS_ELB_ID is required for this test")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	basicOptsEnable := load_balancer.EnableLoadBalancerOpts{
		ElbId:  elbid,
		Agency: agency,
	}
	elbID, err := load_balancer.EnableLoadBalancer(client, clusterID, basicOptsEnable)
	th.AssertNoErr(t, err)
	log.Println("CSS loadbalancer id")
	tools.PrintResource(t, elbID)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	clientELB, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	configuringOpts := load_balancer.LoadBalancingListenerOpts{
		Protocol:     "HTTPS",
		ProtocolPort: 81,
		ServerCertId: createCertificate(t, clientELB),
	}

	load_balancerID, err := load_balancer.ConfigureLoadBalancingListeners(client, clusterID, configuringOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, load_balancerID)

	log.Println("The load balancer listener for the CSS cluster is available.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))
	time.Sleep(8 * time.Second)

	elbDetails, err := load_balancer.Get(client, clusterID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, elbDetails)

	listenerID := elbDetails.Listener.Id

	log.Printf("The ID: %s", elbDetails.Listener.Id)

	time.Sleep(8 * time.Second)

	configuringUpdateOpts := load_balancer.UpdatingListenerOpts{
		Listener: load_balancer.EsListenerRequest{
			DefaultTlsContainerRef: createCertificate(t, clientELB),
		},
	}

	NewElbDetails, err := load_balancer.Get(client, clusterID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, NewElbDetails)

	updated, err := load_balancer.UpdatingLoadBalancingListeners(client, clusterID, listenerID, configuringUpdateOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, updated)

	time.Sleep(8 * time.Second)

	log.Println("The load balancer listener for the CSS cluster is updated.")

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	defer deleteCertificate(t, clientELB, configuringOpts.ServerCertId)
	defer deleteCertificate(t, clientELB, configuringUpdateOpts.Listener.DefaultTlsContainerRef)

	err = load_balancer.DisableLoadBalancer(client, clusterID)

	log.Println("The load balancer of the CSS cluster is disabled.")

	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

}

func TestGetCSSLoadBalancingConfiguration(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	got, err := load_balancer.Get(client, clusterID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, got)
}

func createCertificate(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create ELBv3 certificate")
	certName := tools.RandomString("create-cert-", 3)

	privateKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAwZ5UJULAjWr7p6FVwGRQRjFN2s8tZ/6LC3X82fajpVsYqF1x
qEuUDndDXVD09E4u83MS6HO6a3bIVQDp6/klnYldiE6Vp8HH5BSKaCWKVg8lGWg1
UM9wZFnlryi14KgmpIFmcu9nA8yV/6MZAe6RSDmb3iyNBmiZ8aZhGw2pI1YwR+15
MVqFFGB+7ExkziROi7L8CFCyCezK2/oOOvQsH1dzQ8z1JXWdg8/9Zx7Ktvgwu5PQ
M3cJtSHX6iBPOkMU8Z8TugLlTqQXKZOEgwajwvQ5mf2DPkVgM08XAgaLJcLigwD5
13koAdtJd5v+9irw+5LAuO3JclqwTvwy7u/YwwIDAQABAoIBACU9S5fjD9/jTMXA
DRs08A+gGgZUxLn0xk+NAPX3LyB1tfdkCaFB8BccLzO6h3KZuwQOBPv6jkdvEDbx
Nwyw3eA/9GJsIvKiHc0rejdvyPymaw9I8MA7NbXHaJrY7KpqDQyk6sx+aUTcy5jg
iMXLWdwXYHhJ/1HVOo603oZyiS6HZeYU089NDUcX+1SJi3e5Ke0gPVXEqCq1O11/
rh24bMxnwZo4PKBWdcMBN5Zf/4ij9vrZE+fFzW7vGBO48A5lvZxWU2U5t/OZQRtN
1uLOHmMFa0FIF2aWbTVfwdUWAFsvAOkHj9VV8BXOUwKOUuEktdkfAlvrxmsFrO/H
yDeYYPkCgYEA/S55CBbR0sMXpSZ56uRn8JHApZJhgkgvYr+FqDlJq/e92nAzf01P
RoEBUajwrnf1ycevN/SDfbtWzq2XJGqhWdJmtpO16b7KBsC6BdRcH6dnOYh31jgA
vABMIP3wzI4zSVTyxRE8LDuboytF1mSCeV5tHYPQTZNwrplDnLQhywcCgYEAw8Yc
Uk/eiFr3hfH/ZohMfV5p82Qp7DNIGRzw8YtVG/3+vNXrAXW1VhugNhQY6L+zLtJC
aKn84ooup0m3YCg0hvINqJuvzfsuzQgtjTXyaE0cEwsjUusOmiuj09vVx/3U7siK
Hdjd2ICPCvQ6Q8tdi8jV320gMs05AtaBkZdsiWUCgYEAtLw4Kk4f+xTKDFsrLUNf
75wcqhWVBiwBp7yQ7UX4EYsJPKZcHMRTk0EEcAbpyaJZE3I44vjp5ReXIHNLMfPs
uvI34J4Rfot0LN3n7cFrAi2+wpNo+MOBwrNzpRmijGP2uKKrq4JiMjFbKV/6utGF
Up7VxfwS904JYpqGaZctiIECgYA1A6nZtF0riY6ry/uAdXpZHL8ONNqRZtWoT0kD
79otSVu5ISiRbaGcXsDExC52oKrSDAgFtbqQUiEOFg09UcXfoR6HwRkba2CiDwve
yHQLQI5Qrdxz8Mk0gIrNrSM4FAmcW9vi9z4kCbQyoC5C+4gqeUlJRpDIkQBWP2Y4
2ct/bQKBgHv8qCsQTZphOxc31BJPa2xVhuv18cEU3XLUrVfUZ/1f43JhLp7gynS2
ep++LKUi9D0VGXY8bqvfJjbECoCeu85vl8NpCXwe/LoVoIn+7KaVIZMwqoGMfgNl
nEqm7HWkNxHhf8A6En/IjleuddS1sf9e/x+TJN1Xhnt9W6pe7Fk1
-----END RSA PRIVATE KEY-----`
	cert := `-----BEGIN CERTIFICATE-----
MIIDpTCCAo2gAwIBAgIJAKdmmOBYnFvoMA0GCSqGSIb3DQEBCwUAMGkxCzAJBgNV
BAYTAnh4MQswCQYDVQQIDAJ4eDELMAkGA1UEBwwCeHgxCzAJBgNVBAoMAnh4MQsw
CQYDVQQLDAJ4eDELMAkGA1UEAwwCeHgxGTAXBgkqhkiG9w0BCQEWCnh4QDE2My5j
b20wHhcNMTcxMjA0MDM0MjQ5WhcNMjAxMjAzMDM0MjQ5WjBpMQswCQYDVQQGEwJ4
eDELMAkGA1UECAwCeHgxCzAJBgNVBAcMAnh4MQswCQYDVQQKDAJ4eDELMAkGA1UE
CwwCeHgxCzAJBgNVBAMMAnh4MRkwFwYJKoZIhvcNAQkBFgp4eEAxNjMuY29tMIIB
IjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwZ5UJULAjWr7p6FVwGRQRjFN
2s8tZ/6LC3X82fajpVsYqF1xqEuUDndDXVD09E4u83MS6HO6a3bIVQDp6/klnYld
iE6Vp8HH5BSKaCWKVg8lGWg1UM9wZFnlryi14KgmpIFmcu9nA8yV/6MZAe6RSDmb
3iyNBmiZ8aZhGw2pI1YwR+15MVqFFGB+7ExkziROi7L8CFCyCezK2/oOOvQsH1dz
Q8z1JXWdg8/9Zx7Ktvgwu5PQM3cJtSHX6iBPOkMU8Z8TugLlTqQXKZOEgwajwvQ5
mf2DPkVgM08XAgaLJcLigwD513koAdtJd5v+9irw+5LAuO3JclqwTvwy7u/YwwID
AQABo1AwTjAdBgNVHQ4EFgQUo5A2tIu+bcUfvGTD7wmEkhXKFjcwHwYDVR0jBBgw
FoAUo5A2tIu+bcUfvGTD7wmEkhXKFjcwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B
AQsFAAOCAQEAWJ2rS6Mvlqk3GfEpboezx2J3X7l1z8Sxoqg6ntwB+rezvK3mc9H0
83qcVeUcoH+0A0lSHyFN4FvRQL6X1hEheHarYwJK4agb231vb5erasuGO463eYEG
r4SfTuOm7SyiV2xxbaBKrXJtpBp4WLL/s+LF+nklKjaOxkmxUX0sM4CTA7uFJypY
c8Tdr8lDDNqoUtMD8BrUCJi+7lmMXRcC3Qi3oZJW76ja+kZA5mKVFPd1ATih8TbA
i34R7EQDtFeiSvBdeKRsPp8c0KT8H1B4lXNkkCQs2WX5p4lm99+ZtLD4glw8x6Ic
i1YhgnQbn5E0hz55OLu5jvOkKQjPCW+8Kg==
-----END CERTIFICATE-----`

	createOpts := certificates.CreateOpts{
		Name:        certName,
		Description: "some interesting certificate",
		Type:        "server",
		PrivateKey:  privateKey,
		Certificate: cert,
	}

	certificate, err := certificates.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Name, certificate.Name)
	th.AssertEquals(t, createOpts.Description, certificate.Description)
	t.Logf("Created ELBv3 certificate: %s", certificate.ID)

	return certificate.ID
}

func deleteCertificate(t *testing.T, client *golangsdk.ServiceClient, certificateID string) {
	t.Logf("Attempting to delete ELBv3 certificate: %s", certificateID)
	err := certificates.Delete(client, certificateID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Deleted ELBv3 certificate: %s", certificateID)
}
