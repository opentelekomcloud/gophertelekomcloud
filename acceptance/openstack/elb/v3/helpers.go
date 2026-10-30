package v3

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/certificates"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/loadbalancers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/pools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func createLoadBalancer(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create ELBv3 LoadBalancer")
	lbName := tools.RandomString("create-lb-", 3)
	adminStateUp := true

	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	networkID := clients.EnvOS.GetEnv("NETWORK_ID")
	subnetID := clients.EnvOS.GetEnv("SUBNET_ID")
	if networkID == "" || vpcID == "" || subnetID == "" {
		t.Skip("OS_NETWORK_ID/OS_VPC_ID/OS_SUBNET_ID env vars are missing but LBv3 test requires")
	}

	az := clients.EnvOS.GetEnv("AVAILABILITY_ZONE")
	if az == "" {
		az = "eu-nl-01"
	}
	ipTargetEnable := true

	createOpts := loadbalancers.CreateOpts{
		Name:                 lbName,
		Description:          "some interesting loadbalancer",
		VipSubnetCidrID:      subnetID,
		VpcID:                vpcID,
		AvailabilityZoneList: []string{az},
		Tags: []tags.ResourceTag{
			{
				Key:   "gophertelekomcloud",
				Value: "loadbalancer",
			},
		},
		AdminStateUp:   &adminStateUp,
		ElbSubnetIDs:   []string{networkID},
		IpTargetEnable: &ipTargetEnable,
	}

	loadbalancer, err := loadbalancers.Create(client, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Name, loadbalancer.Name)
	th.AssertEquals(t, createOpts.Description, loadbalancer.Description)
	t.Logf("Created ELBv3 LoadBalancer: %s", loadbalancer.ID)

	return loadbalancer.ID
}

func deleteLoadbalancer(t *testing.T, client *golangsdk.ServiceClient, loadbalancerID string) {
	t.Logf("Attempting to delete ELBv3 LoadBalancer: %s", loadbalancerID)
	err := loadbalancers.Delete(client, loadbalancerID)
	th.AssertNoErr(t, err)
	t.Logf("Deleted ELBv3 LoadBalancer: %s", loadbalancerID)
}

func createCertificate(t *testing.T, client *golangsdk.ServiceClient) string {
	t.Logf("Attempting to create ELBv3 certificate")
	certName := tools.RandomString("create-cert-", 3)

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	th.AssertNoErr(t, err)
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	th.AssertNoErr(t, err)
	template := &x509.Certificate{
		SerialNumber: serial,
		Subject:      pkix.Name{CommonName: "example.com"},
		DNSNames:     []string{"example.com"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	th.AssertNoErr(t, err)
	privateKey := string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}))
	cert := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))

	createOpts := certificates.CreateOpts{
		Name:        certName,
		Description: "some interesting certificate",
		Type:        "server",
		PrivateKey:  privateKey,
		Certificate: cert,
	}

	certificate, err := certificates.Create(client, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Name, certificate.Name)
	th.AssertEquals(t, createOpts.Description, certificate.Description)
	t.Logf("Created ELBv3 certificate: %s", certificate.ID)

	return certificate.ID
}

func deleteCertificate(t *testing.T, client *golangsdk.ServiceClient, certificateID string) {
	t.Logf("Attempting to delete ELBv3 certificate: %s", certificateID)
	err := certificates.Delete(client, certificateID)
	th.AssertNoErr(t, err)
	err = golangsdk.WaitFor(600, func() (bool, error) {
		_, err := certificates.Get(client, certificateID)
		var notFound golangsdk.ErrDefault404
		if errors.As(err, &notFound) {
			return true, nil
		}
		return false, err
	})
	th.AssertNoErr(t, err)
	t.Logf("Deleted ELBv3 certificate: %s", certificateID)
}

func createPool(t *testing.T, client *golangsdk.ServiceClient, loadbalancerID string) string {
	t.Logf("Attempting to create ELBv3 Pool")
	vpcID := clients.EnvOS.GetEnv("VPC_ID")
	poolName := tools.RandomString("create-pool-", 3)
	createOpts := pools.CreateOpts{
		LBMethod:                 "LEAST_CONNECTIONS",
		Protocol:                 "HTTP",
		LoadbalancerID:           loadbalancerID,
		Name:                     poolName,
		Description:              "some interesting description",
		VpcId:                    vpcID,
		Type:                     "instance",
		DeletionProtectionEnable: pointerto.Bool(true),
	}

	pool, err := pools.Create(client, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Name, pool.Name)
	th.AssertEquals(t, createOpts.Description, pool.Description)
	th.AssertEquals(t, createOpts.LBMethod, pool.LBMethod)
	th.AssertEquals(t, true, pool.DeletionProtectionEnable)
	th.AssertEquals(t, createOpts.Type, pool.Type)
	th.AssertEquals(t, createOpts.VpcId, pool.VpcId)
	t.Logf("Created ELBv3 Pool: %s", pool.ID)

	return pool.ID
}

func deletePool(t *testing.T, client *golangsdk.ServiceClient, poolID string) {
	t.Logf("Attempting to delete ELBv3 Pool: %s", poolID)
	err := pools.Delete(client, poolID)
	th.AssertNoErr(t, err)
	t.Logf("Deleted ELBv3 Pool: %s", poolID)
}

func createListener(t *testing.T, client *golangsdk.ServiceClient, loadbalancerID string) string {
	listener, err := listeners.Create(client, listeners.CreateOpts{
		LoadbalancerID: loadbalancerID,
		Protocol:       listeners.ProtocolHTTP,
		ProtocolPort:   80,
	})
	th.AssertNoErr(t, err)
	return listener.ID
}

func deleteListener(t *testing.T, client *golangsdk.ServiceClient, listenerID string) {
	th.AssertNoErr(t, listeners.Delete(client, listenerID))
}
