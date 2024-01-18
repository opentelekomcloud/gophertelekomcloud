package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf/v1/certificates"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const (
	testKey            = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCvmuH5ViGtGOle\nvJ8vOoN3Ak4pp3SescdAfQa/r4cOz/bmBqBcZJTX9HODhiQzdemyLLs9aOkQXYIc\n8OrcaIsjns92XITVDpFW0ThGyjhTZdELj9LsbIcVzNPPclTcebZBlzAyX0oLqpHK\n73OUYQY2E6l44U9G8Id763Bnws9NRn3cg0qufrlUgdim/pYZ8ubjvlDJ9eEIhcsu\n9zu8c8i2+8qLjEsonx5PrwzNlYP3JqAmZ2dcbQeSPfv5U6ZceKEZfegK+Cxv4rFd\n5F4Rdxl+SAIY+6mr7qu1dAlcVMLSQcLlJLRWQ5NmqL9xju7Fbj2VZt+L6nb512iK\naedPo2GfAgMBAAECggEAeMAvDS3uAEI2Dx/y8h3xUn9yUfBFH+6tTanrXxoK6+OT\nKj96O64qL4l3eQRflkdJiGx74FFomglCtDXxudflfXvxurkJ2hunUySQ5xScwLQt\nmB6w8kP6a8IqD+bVdbn32ohk6u5dU0JZ+ErJlklVZRAGJAoCYox5DXwrEh6CP+bJ\npItgjv71tEEnX5sScQwV7FMRbjsPzXoJp8vCQjlUdetM1fk9rs3R2WSeFbPgLLtC\nxY0+8Hexy0q6BLmyPZvFCaVIAzAHCYeCyzPK3xcm4odbrBmRL/amOg24CCny065N\nMU9RFhEjQsY1RaK7dgkvjsntUZvU+aDcL8o6djOTuQKBgQDlDN/j2ntpGCtbTWH0\ncVTW13Ze7U7iE3BfDO3m4VYP3Xi/v5FI8nHlmLrcl30H1dPKvMTec0dCBOqD1wzF\nKiqHy8ELowO2CbXMYJpjuPzXH40/AE3eOJVTJM8mOeuFdeFgYCd/9cB7o5jfTA5Y\n4zj8EmcRzsH1rNSnvo7/O9q6+wKBgQDERDSvP8RScEbzDKuN6uhzj1K2CAEnY6//\nrDA1so18UhAie9NcAvlKa46jQTOcYD77g5h0WSlNt9ZbK9Plq9CY9psI0KNqN3Fl\nYVKOKdD5m6Rifmg+lt8KLc/WocQ10DXpPTXzzuRlN/TaMDdN2pedEre/0AAMs8Ia\nMIUnu4oyrQKBgQC6b6BNdqi9Ak9IIdR5g0XrGbXfzolGu0vcEkoSg5fpkfuXF/bJ\nyY2rtIVkyGmc1w9tFfmol2yI8Ddy2LgsRAYaQl7/edCre3vev0LrqMck0ynE/hpj\npurkojF6i+qI10p7h8ie/wmNmbv1BZMoBst7Yf9DH2gA8IynfRQn7DA9wQKBgGaU\nM2kJDgX8UsjDbYKuLTIAzb0AMAIzUxBxIX1fRh2dEnvDdjOYBk1EK/fdoyjvENwJ\n6ouc8j6BgBKEtKpMg6j+8wbHbTGdqrHPDQPqjSN4mpEz+i4EUqySRxep0tBBc3vl\nFybHko3okhvbqXwSbL2Ww90HzI7XAPMJOv8KQO+9AoGBAJxxftNWvypBXGkPCdH2\nf3ikvT2Vef9QZjqkvtipCecAkjM6ReLshVsdqFSv/ZmsVUeNKoTHvX2GnhweJM44\nx7N2mFK4skBzVtMVbjAHVjG78UitVu+FrzqGreaJXHaduhgUH2iFWfw09joOotAM\nX7ioLbTeWGBqFM+C80PkdBNp\n-----END PRIVATE KEY-----"
	testCert           = "-----BEGIN CERTIFICATE-----\nMIIE8DCCA9igAwIBAgISBBvT73JHmxkCaGQKE1cneHh1MA0GCSqGSIb3DQEBCwUA\nMDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD\nEwJSMzAeFw0yMzExMTQxNTE5MDJaFw0yNDAyMTIxNTE5MDFaMBwxGjAYBgNVBAMT\nEXJlbGF4LmZlcmNoYXUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC\nAQEA0htQi7o/Jb2lISY7s6/TExCc01zY1E4hpfts5iqjMT0JRc45mHK6EoVr84J3\nAi6tX78RlfWDwC1XGbugysyvCRspyPNTZUtdCgoi6XoZu93PdtpJ8TTP/FGztNPw\nbK9gUd3ZdGM4oeTPpr0ECPWf+uJ1ucBAp4OZxivltAswvlO6/vYyUmyYxdb4F07Y\nboXcf+pttM0t7ClRHWNIDrFC2P8G1opy8pjYimcVd9GRUIZevhT1fz+cJrhcxF7M\njxQqgbSjWfD+wJ2XFhNfTlOUMQDZVULM59Pwyf0r7jq4bcY7xTvNp7SiFr+IGbgz\n0Gl4w/z8qqJV/jHtYALQZxXxwQIDAQABo4ICFDCCAhAwDgYDVR0PAQH/BAQDAgWg\nMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0G\nA1UdDgQWBBQemg3O3qJ5s59Yel34Lescd/synDAfBgNVHSMEGDAWgBQULrMXt1hW\ny65QCUDmH6+dixTCxjBVBggrBgEFBQcBAQRJMEcwIQYIKwYBBQUHMAGGFWh0dHA6\nLy9yMy5vLmxlbmNyLm9yZzAiBggrBgEFBQcwAoYWaHR0cDovL3IzLmkubGVuY3Iu\nb3JnLzAcBgNVHREEFTATghFyZWxheC5mZXJjaGF1LmNvbTATBgNVHSAEDDAKMAgG\nBmeBDAECATCCAQUGCisGAQQB1nkCBAIEgfYEgfMA8QB3ADtTd3U+LbmAToswWwb+\nQDtn2E/D9Me9AA0tcm/h+tQXAAABi86hDRIAAAQDAEgwRgIhAPtP8/xmpNftBgvm\nw8T50cV923nC83IwUR/UYCgMUhDRAiEAwUUU9z7Afs6jUAQ+7XptmG6+60YCcIWv\nwnMI82Pqr2MAdgB2/4g/Crb7lVHCYcz1h7o0tKTNuyncaEIKn+ZnTFo6dAAAAYvO\noQ1IAAAEAwBHMEUCIQD6hkTTj4/gqWiygdXdEb0ylG3rAD7mRkc9x1Z7+VGhEwIg\nYQfYYHXBEFN2aRiIZTWXgBYVjyKCGKHQaMvd+pDb2cEwDQYJKoZIhvcNAQELBQAD\nggEBAFxxw3cAJGg/u6v8qttFa04YgWIgwHfW5XNRskOr96cibDSaCP5Amq33/BRh\nytoF7al6Y+3/4vM/A+8IbIt0cUVfoqEY3IzjX9gVsmqemhsyaQB9BFnbCGySKBaS\n7ecOqbhMz3aHbb9qD1p7ne136mvv3vyup9dGBNbZ2TYc4/q8eHD/2yrL/xv8Zzdu\nOk8QTDwsAxK473IORgMoBolfM5qU+mY2eau6XtwbffFhaDdjTbN/AdwJylMzZlXs\nvM3NdeMJuFFZIzvJ5R+9ddWgeRXM0Ksl5k4QGcZyKRu9I0F/9suGYDKdTQcCVNmP\nWjvoz9oBJtp410PMWxmrLcsXLzI=\n-----END CERTIFICATE-----"
	testCertExpiration = 1707751141000
)

func TestCertificateLifecycle(t *testing.T) {
	client, err := clients.NewWafV1Client()
	th.AssertNoErr(t, err)

	name := tools.RandomString("test_cert-", 5)

	opts := certificates.CreateOpts{
		Name:    name,
		Content: testCert,
		Key:     testKey,
	}

	created, err := certificates.Create(client, opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, int64(testCertExpiration), created.ExpireTime)

	defer func() {
		err := certificates.Delete(client, created.Id).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	got, err := certificates.Get(client, created.Id).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, created.Name, got.Name)
	th.AssertEquals(t, created.ExpireTime, got.ExpireTime)

	updateOpts := certificates.UpdateOpts{Name: name + "_updated"}
	updated, err := certificates.Update(client, created.Id, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, created.ExpireTime, updated.ExpireTime)
	th.AssertEquals(t, created.Id, updated.Id)
	th.AssertEquals(t, updateOpts.Name, updated.Name)

	pages, err := certificates.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	certs, err := certificates.ExtractCertificates(pages)
	th.AssertNoErr(t, err)
	if len(certs) == 0 {
		t.Errorf("no certificates in the list")
	}

	pages2, err := certificates.List(client, certificates.ListOpts{
		Limit: -1,
	}).AllPages()
	th.AssertNoErr(t, err)

	certs2, err := certificates.ExtractCertificates(pages2)
	th.AssertNoErr(t, err)
	if len(certs2) == 0 {
		t.Errorf("no certificates in the list")
	}
}
