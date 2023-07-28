package v1

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	testKey = `-----BEGIN PRIVATE KEY-----
MIIJQwIBADANBgkqhkiG9w0BAQEFAASCCS0wggkpAgEAAoICAQDgurPwlRs8khOO
6KDHQdexG/phYbaWLUtkfyrFVNFt2KgBzylntkyiAATWFA/+4CZwwa6y2qNlgwkg
EKsY+vVU34AqGjzbeV4msaHohuQ2P84WO7UegUphQJPgYz0u2SZPXXf2MoCynyfQ
DKXcr/8O6gw59kWqiB1yoX6RWaBDqKmWoCqmBuYVSb1HvecwUmE6Ij0rt3encmIR
V65FNz2arCoqNW60/d90p8h2sS1Oeg9+n0Z80wxBMwJNnwCjCbmkt4vVudm4qyWh
hTjuFYkP1d14OOifOwkfXfGHyRziZMumv4OFg2dW29QKkEMDdY+FisVAG5tH9x9W
8NGGwWxU6jtAwNhZ/C36ZIPlp349AKR3U+kuCEyrP6mXP6HU6JI8ilYPpMYGXLlJ
Hv1qBQVaU5SHSfzOip+g4kCiGCdIm8RG/+LTpT97IVoyvCW6bab2q/i+EitgakkC
0JEVCWt9b7KeJh8ZjUKzkHDhSdGvZZmq2QuXcFHvXXfkbmaJ7YrzYSKBXfXXxUeP
w0JuWEuhwNXSOJZz2CfJ5+mkI/oihwYyL4bL75uYcg9A/bIjg29fHVq1kQHzAfVh
MiwNqAN9fd5/1aQp+1y2CfxPYvyt7ymFREaMI8w16rjJuyVdjgR+V044UhRayGRl
RWJXQrJwkhvongi4GkdPW1ABNwzbmQIDAQABAoICAEdw0vcuT4RH49PQfBwcAFeb
T1NZ3tOK/qaqDozA0/sZnv9EPiNsPpxZaTAtHJCn7VB3IfRVsQ/6QhJheiLs1MTw
cCvyP1p+EMI4QgJLr4zXZ8qFnKRf8adNAjWZFsAn5Bfi3Nn1YBhopB1th+TKRkkV
emGKusblkob4c+X9GgeoPJFXxXcWRlqKIJQH+NDRv3rdm5ikMHOY1zgwKYRzdTAQ
fy7/4XvEIR9Sn1WsKX0DLJ3SQHQ6G3E2qArI+0jZNJz6hIejF2Wvcr0QPvLhAbt4
/3jSjpDgEZxZHwlNk9Mcu+j8hPESvu1L4PKivcsBumh3nxEsNYcBNoNK9zDhmHAl
qAVConp4jD0V0ZadKGWHzzIEFX5ga4eKNvWGWoK6IJr0botKgT0tEtpVwWSxdWnc
vgzLQ9BQcKVH26gLTM7mNyumexF024sV3J7s3wEEgjbAHGljFAdTCuQn6oAM07ZS
IsDSOcghiIKHZ2m8m9itlzzrI4SHy0Y10ZDFB716PVsUpjvqQ9ooAZJzWHNBI2le
20vRAu62flueeS1I8erb66eMkQTd/5H/7HYzBbUBYSvhMMlEDrQY8K2UGZ8bXUx4
fiYW2ZzZxtx27Zg9Iw3eJCZJqi/0I1uqIrKPXLOC4YX/WeW1uzOv/OCaU5CFYg9b
BX53Q1NQnGZkwrd4M/sBAoIBAQD0GUJLEdHCWRohMSxXTgdxUdAgWgOh6rPkI+n+
UrYi+upLFwVsxGlZXXyokuWnPy6xwVs3tuWsEwHsZjmGcetCtVdYmHyILlss73Ol
Rij+1q+THBfrL5hVnBMBm+DN/XnzLguFBQVJRcbE+wOavppBk2NIrbpCQgpiFPTT
no4goehZYqKpr2t6kLuLoJn4/RfMUkiV2nsl3klQnurGn+CS0PaBuZAhmOy5fVb+
wV8nvUNLXWAVM5hGOTMgWrtIVoEVRYi7lXoA2OISXAhPG3AXYRej99lXxl1HlNuj
Jd6IPKo2h/nHW29iGaWNfLNSn3CQ0kIuLtFqoFo/XjL7LkTxAoIBAQDrr67zgWm9
cpqwmdCNXkE0us4Wl1GlybLemajdVi9BBBo6Nuivb3GiT4r8mZ9KTyk4yqAsULpY
Z+7+4qRyfMpArcq+7ERjXUJw6yq14DVbA/GlJcQ2yuWCOqv+2Mul91myt7O4V7Wc
Os36f4KPCJkEqucTtz+DiX4YwbDVCvHfTpJXYbDXOVG6UEVf+Tx6Wx16Emjfhi4B
0uzu0atniVidQSibF81AvymmjPkRGkx91Br0qDMkRQBsfrxNJu1O712bKms1GAFO
MnTdIh5ljU4MQDPaJtFKcjmw5eTHbMvGruxS6LwgM8DrQvOIXdPMdSTz7tOSgToF
fzT9jlKCzOEpAoIBAQClJgHIMIIub4JSOqa5Wr2GWcfqW3xhrB2RmQrTWrqH6CNk
MmslL63nHG0e0GQ4R3McKKnChCfXx/RhMLhy0dhOBcrW0jRPHq3pNQiVJWbPJAke
Cr/UCxuRsErbp87tDzXW5aw9jywIawEUfI/vvk03WLSvk3qVIYFM4sjR9FBMm75L
24QaMekRv6Jj0YDbCMF1J6acXHk9IauQtDQ7tieGrYJaOmXdlU10Ie0d506t4EsL
Tl2XepTnzgNdPIXBZ2VmMulToMoukI5DxaiJfRLVfoc0FJgj3r11lK0VMKXinsi6
pDzGOIKfaKKtm1Tn7Z+HG/pSrLJa5aqpfN4ZOzDBAoIBACVt+TLi0pArqzVwuBY7
ac+d+yzLS0QxDB8d+BtunIKOzDuCjOGPqVRFnaUQIKQEfl9ujpF7IJz5pJMGG2ez
Ocubzh8UFqhRH0QflODdgpu5vJ6lqMuq3VgZSUdn1q+84JnpYrlb9JOjIyMtLOba
TrLXEWuoJoYVR9lWqWasHk2AhO0rrpH/oGMebGYZhulHnx7L3avh+1x+yvICil4f
CduvhWtcFFS8BzlUGhoFOzCghsdkDvsrmi2g0vbNv9JRYWRLEEuWTF7G1Jhp2rn1
/vcjGxkCISrZiR/24qZpONOM5CsmmvniPjkeoN5/SCuoTv4OZ7tUmopU8W1zNNdh
AkECggEBANvWUz3Vd4xttKAOYGWpMG1btteQoFjTRIal2arW5pbs/oYkHboufxpH
o+R4J/QRAsZ+juRmao4D/JtyCamKKQpNRU+lXOduHIPa+BynCsRGnV/IjyeQaXKa
rgB6c1gENB3sKCEl9KRjSqkQVjLuiO2T9Ugvn0h2qYcvkbqsaAseOXj5E4FbWSQE
Sy4cUIi/4Rf7EVlJFDH4L4k1388uh7lqYB4wzMHEmPTD8uhOgaokF9pO3pYwSTM7
v2FKjteSuRcJ+oGbMoLvKHFnAGsGC0vYbrVQhw5krFjSPEjHdSj+P0cSGKOVax26
6Ty0X9L6AsYiuWgPiTkurNI9qz7lYwg=
-----END PRIVATE KEY-----`
	testCert = `-----BEGIN CERTIFICATE-----
MIIFRzCCAy8CFBKwjdhuzVswxRC4CExG3q6kRXZoMA0GCSqGSIb3DQEBCwUAMGAx
CzAJBgNVBAYTAkVTMRAwDgYDVQQIDAdHUkFOQURBMRAwDgYDVQQHDAdHUkFOQURB
MRAwDgYDVQQKDAdERUZBVUxUMQwwCgYDVQQLDANFQ08xDTALBgNVBAMMBEhPU1Qw
HhcNMjMwNzI4MTExNTEyWhcNMjQwNzI3MTExNTEyWjBgMQswCQYDVQQGEwJFUzEQ
MA4GA1UECAwHR1JBTkFEQTEQMA4GA1UEBwwHR1JBTkFEQTEQMA4GA1UECgwHREVG
QVVMVDEMMAoGA1UECwwDRUNPMQ0wCwYDVQQDDARIT1NUMIICIjANBgkqhkiG9w0B
AQEFAAOCAg8AMIICCgKCAgEA4Lqz8JUbPJITjuigx0HXsRv6YWG2li1LZH8qxVTR
bdioAc8pZ7ZMogAE1hQP/uAmcMGustqjZYMJIBCrGPr1VN+AKho823leJrGh6Ibk
Nj/OFju1HoFKYUCT4GM9LtkmT1139jKAsp8n0Ayl3K//DuoMOfZFqogdcqF+kVmg
Q6iplqAqpgbmFUm9R73nMFJhOiI9K7d3p3JiEVeuRTc9mqwqKjVutP3fdKfIdrEt
TnoPfp9GfNMMQTMCTZ8Aowm5pLeL1bnZuKsloYU47hWJD9XdeDjonzsJH13xh8kc
4mTLpr+DhYNnVtvUCpBDA3WPhYrFQBubR/cfVvDRhsFsVOo7QMDYWfwt+mSD5ad+
PQCkd1PpLghMqz+plz+h1OiSPIpWD6TGBly5SR79agUFWlOUh0n8zoqfoOJAohgn
SJvERv/i06U/eyFaMrwlum2m9qv4vhIrYGpJAtCRFQlrfW+yniYfGY1Cs5Bw4UnR
r2WZqtkLl3BR71135G5mie2K82EigV3118VHj8NCblhLocDV0jiWc9gnyefppCP6
IocGMi+Gy++bmHIPQP2yI4NvXx1atZEB8wH1YTIsDagDfX3ef9WkKftctgn8T2L8
re8phURGjCPMNeq4ybslXY4EfldOOFIUWshkZUViV0KycJIb6J4IuBpHT1tQATcM
25kCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEApAJOscxakTRp7ET7AY0CwaLbvgUO
2NZWbg/Pr+jrt55Sxo9exMtOAexxUSRCVcAMfz2DPfdv8TzW/eIMV84BN1RZivTO
g3LgbrbgwVe8q200GNoji8lQWePyBvUqMXaZ5ESN8K7aWiEphCoSp7W+2OvwNd1o
yY/ovKGEmpioqla64qxIRMO4JHJJXv3lTLh1jBrPFET6AyhQbv/urZwkm9rWNTG+
fJ11/k8W1cHCdid72YL94TqQ3AIq5swSizERDXHck0jONkA88bBYExdQoqfE+X8n
rFSrpW4HGPDnE8/FBvir4JWOlEymifgAqmfCQZfkr/XTOircNwMLiAwUwqn7cMYc
kWCq8JgJXVpAegCPy/rzB1+M4FxVL4HMFLAiTVvkdK52e3bJ4HRCX5k48fKSj8RX
wWiWnK6YSShkplerCZ50ng+SvSiWzcqsJRqgupwFJKM3+iZ6zsyIpOLqZ/XgSiRg
F1NrTkx+2qdNqp/F0PUEJMJQA0ZukviDqvHefr07ZjFL7qBJb19cEV0sgvnA93b4
KirQdAKxHYWrZl+rTTmlHPsfvylMFXDef5X66USvRdcK3xDpeAiXwJRg+HeDGk8w
9DAfnoBiV1K/gi4dWW9GqIzzqk7HdyzmjRyJJRoYV3nRssOGlcjUeu5GJgrKHv/0
nGCTZwjrFytcgdo=
-----END CERTIFICATE-----`
)

func waitForInstanceToBeCreated(client *golangsdk.ServiceClient, secs int, id string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		instance, err := instances.Get(client, id)
		if err != nil {
			return false, err
		}
		if instance.Status == 1 {
			return true, nil
		}
		if instance.Status == 4 {
			return false, fmt.Errorf("error creating instance")
		}

		return false, nil
	})
}

func waitForInstanceToBeDeleted(client *golangsdk.ServiceClient, secs int, id string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		_, err := instances.Get(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, fmt.Errorf("error retriving WAF instance status: %w", err)
		}
		return false, nil
	})
}

func getWafdClient(t *testing.T, region string) (*golangsdk.ServiceClient, error) {
	var client *golangsdk.ServiceClient
	var err error
	if region == "eu-ch2" {
		client, err = clients.NewWafdSwissV1Client()
		th.AssertNoErr(t, err)
	} else {
		client, err = clients.NewWafdV1Client()
		th.AssertNoErr(t, err)
	}
	return client, err
}
