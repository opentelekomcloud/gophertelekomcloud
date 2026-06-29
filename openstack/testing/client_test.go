package testing

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const ID = "0123456789"

func TestAuthenticatedClientV3(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, `
			{
				"versions": {
					"values": [
						{
							"status": "stable",
							"id": "v3.0",
							"links": [
								{ "href": "%s", "rel": "self" }
							]
						},
						{
							"status": "stable",
							"id": "v2.0",
							"links": [
								{ "href": "%s", "rel": "self" }
							]
						}
					]
				}
			}
		`, th.Endpoint()+"v3/", th.Endpoint()+"v2.0/")
	})

	th.Mux.HandleFunc("/v3/auth/tokens", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Subject-Token", ID)

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, `
			{
				"token": {
					"expires_at": "2013-02-02T18:30:59.000000Z",
					"project": {
						"id": ""
					}
				}
			}
		`)
	})

	options := golangsdk.AuthOptions{
		Username:         "me",
		Password:         "secret",
		DomainName:       "default",
		TenantName:       "project",
		IdentityEndpoint: th.Endpoint(),
	}
	client, err := openstack.AuthenticatedClient(options)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, ID, client.TokenID)
}

func TestIdentityAdminV3Client(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, `
			{
				"versions": {
					"values": [
						{
							"status": "stable",
							"id": "v3.0",
							"links": [
								{ "href": "%s", "rel": "self" }
							]
						},
						{
							"status": "stable",
							"id": "v2.0",
							"links": [
								{ "href": "%s", "rel": "self" }
							]
						}
					]
				}
			}
		`, th.Endpoint()+"v3/", th.Endpoint()+"v2.0/")
	})

	th.Mux.HandleFunc("/v3/auth/tokens", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Subject-Token", ID)

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, `
	{
    "token": {
        "audit_ids": ["VcxU2JYqT8OzfUVvrjEITQ", "qNUTIJntTzO1-XUk5STybw"],
        "catalog": [
            {
                "endpoints": [
                    {
                        "id": "39dc322ce86c4111b4f06c2eeae0841b",
                        "interface": "public",
                        "region": "RegionOne",
                        "url": "http://localhost:5000"
                    },
                    {
                        "id": "ec642f27474842e78bf059f6c48f4e99",
                        "interface": "internal",
                        "region": "RegionOne",
                        "url": "http://localhost:5000"
                    },
                    {
                        "id": "c609fc430175452290b62a4242e8a7e8",
                        "interface": "admin",
                        "region": "RegionOne",
                        "url": "http://localhost:35357"
                    }
                ],
                "id": "4363ae44bdf34a3981fde3b823cb9aa2",
                "type": "identity",
                "name": "keystone"
            }
        ],
        "expires_at": "2013-02-27T18:30:59.999999Z",
        "is_domain": false,
        "issued_at": "2013-02-27T16:30:59.999999Z",
        "methods": [
            "password"
        ],
        "project": {
            "domain": {
                "id": "1789d1",
                "name": "example.com"
            },
            "id": "263fd9",
            "name": "project-x"
        },
        "roles": [
            {
                "id": "76e72a",
                "name": "admin"
            },
            {
                "id": "f4f392",
                "name": "member"
            }
        ],
        "service_providers": [
            {
                "auth_url":"https://example.com:5000/v3/OS-FEDERATION/identity_providers/acme/protocols/saml2/auth",
                "id": "sp1",
                "sp_url": "https://example.com:5000/Shibboleth.sso/SAML2/ECP"
            },
            {
                "auth_url":"https://other.example.com:5000/v3/OS-FEDERATION/identity_providers/acme/protocols/saml2/auth",
                "id": "sp2",
                "sp_url": "https://other.example.com:5000/Shibboleth.sso/SAML2/ECP"
            }
        ],
        "user": {
            "domain": {
                "id": "1789d1",
                "name": "example.com"
            },
            "id": "0ca8f6",
            "name": "Joe",
            "password_expires_at": "2016-11-06T15:32:17.000000"
        }
    }
}
	`)
	})

	options := golangsdk.AuthOptions{
		Username:         "me",
		Password:         "secret",
		DomainID:         "12345",
		IdentityEndpoint: th.Endpoint(),
	}
	pc, err := openstack.AuthenticatedClient(options)
	th.AssertNoErr(t, err)
	sc, err := openstack.NewIdentityV3(pc, golangsdk.EndpointOpts{
		Region:       "RegionOne",
		Availability: golangsdk.AvailabilityAdmin,
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "http://localhost:35357/v3/", sc.Endpoint)
}

func testAuthenticatedClientFails(t *testing.T, endpoint string) {
	options := golangsdk.AuthOptions{
		Username:         "me",
		Password:         "secret",
		DomainName:       "default",
		TenantName:       "project",
		IdentityEndpoint: endpoint,
	}
	_, err := openstack.AuthenticatedClient(options)
	if err == nil {
		t.Fatal("expected error but call succeeded")
	}
}

func TestAuthenticatedClientV3Fails(t *testing.T) {
	testAuthenticatedClientFails(t, "http://bad-address.example.com/v3")
}

func TestAuthenticatedClientV2Fails(t *testing.T) {
	testAuthenticatedClientFails(t, "http://bad-address.example.com/v2.0")
}

func TestAuthenticatedClientV3WithAgencyAKSKUsesTemporaryCredentials(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	var temporaryCredentialRequests int

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, `
			{
				"versions": {
					"values": [
						{
							"status": "stable",
							"id": "v3.0",
							"links": [
								{ "href": "%s", "rel": "self" }
							]
						}
					]
				}
			}
		`, th.Endpoint()+"v3/")
	})

	th.Mux.HandleFunc("/v3/auth/catalog", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		if r.Header.Get("Authorization") == "" {
			t.Errorf("expected request to be signed")
		}

		_, _ = fmt.Fprintf(w, `
			{
				"catalog": [
					{
						"type": "identity",
						"name": "iam",
						"endpoints": [
							{
								"interface": "public",
								"region": "eu-de",
								"url": "%s"
							}
						]
					}
				],
				"links": {
					"next": null,
					"previous": null
				}
			}
		`, th.Endpoint()+"v3/")
	})

	th.Mux.HandleFunc("/v3.0/OS-CREDENTIAL/securitytokens", func(w http.ResponseWriter, r *http.Request) {
		temporaryCredentialRequests++
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Domain-Id", "source-domain-id")
		if r.Header.Get("Authorization") == "" {
			t.Errorf("expected request to be signed")
		}
		th.TestJSONRequest(t, r, `
			{
				"auth": {
					"identity": {
						"methods": ["assume_role"],
						"assume_role": {
							"agency_name": "target-agency",
							"domain_name": "target-domain"
						}
					}
				}
			}
		`)

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintf(w, `
			{
				"credential": {
					"access": "temporary-ak-%[1]d",
					"secret": "temporary-sk-%[1]d",
					"securitytoken": "temporary-security-token-%[1]d",
					"expires_at": "2030-01-01T00:00:00.000000Z"
				}
			}
		`, temporaryCredentialRequests)
	})

	th.Mux.HandleFunc("/v3/projects", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestFormValues(t, r, map[string]string{"name": "target-project"})
		if r.Header.Get("Authorization") == "" {
			t.Errorf("expected request to be signed")
		}
		th.TestHeader(t, r, "X-Security-Token", fmt.Sprintf("temporary-security-token-%d", temporaryCredentialRequests))
		if !strings.Contains(r.Header.Get("Authorization"), fmt.Sprintf("Credential=temporary-ak-%d/", temporaryCredentialRequests)) {
			t.Errorf("expected temporary AK in Authorization header, got %q", r.Header.Get("Authorization"))
		}

		_, _ = fmt.Fprint(w, `
			{
				"projects": [
					{
						"id": "target-project-id",
						"name": "target-project"
					}
				],
				"links": {
					"next": null,
					"previous": null
				}
			}
		`)
	})

	options := golangsdk.AKSKAuthOptions{
		IdentityEndpoint: th.Endpoint(),
		DomainID:         "source-domain-id",
		AccessKey:        "source-ak",
		SecretKey:        "source-sk",
		AgencyName:       "target-agency",
		AgencyDomainName: "target-domain",
		DelegatedProject: "target-project",
	}

	client, err := openstack.AuthenticatedClient(options)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, temporaryCredentialRequests)
	th.CheckEquals(t, "temporary-ak-1", client.AKSKAuthOptions.AccessKey)
	th.CheckEquals(t, "temporary-sk-1", client.AKSKAuthOptions.SecretKey)
	th.CheckEquals(t, "temporary-security-token-1", client.AKSKAuthOptions.SecurityToken)
	th.CheckEquals(t, "target-project-id", client.ProjectID)

	err = client.ReauthFunc()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, temporaryCredentialRequests)
	th.CheckEquals(t, "temporary-ak-2", client.AKSKAuthOptions.AccessKey)
	th.CheckEquals(t, "temporary-sk-2", client.AKSKAuthOptions.SecretKey)
	th.CheckEquals(t, "temporary-security-token-2", client.AKSKAuthOptions.SecurityToken)
}

func TestAuthenticatedClientV3WithAgencyAKSKWithoutDelegatedProjectKeepsDomainScope(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, `
			{
				"versions": {
					"values": [
						{
							"status": "stable",
							"id": "v3.0",
							"links": [
								{ "href": "%s", "rel": "self" }
							]
						}
					]
				}
			}
		`, th.Endpoint()+"v3/")
	})

	th.Mux.HandleFunc("/v3/auth/catalog", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, `
			{
				"catalog": [
					{
						"type": "identity",
						"name": "iam",
						"endpoints": [
							{
								"interface": "public",
								"region": "eu-de",
								"url": "%s"
							}
						]
					}
				],
				"links": {
					"next": null,
					"previous": null
				}
			}
		`, th.Endpoint()+"v3/")
	})

	th.Mux.HandleFunc("/v3.0/OS-CREDENTIAL/securitytokens", func(w http.ResponseWriter, r *http.Request) {
		th.TestJSONRequest(t, r, `
			{
				"auth": {
					"identity": {
						"methods": ["assume_role"],
						"assume_role": {
							"agency_name": "target-agency",
							"domain_name": "target-domain"
						}
					}
				}
			}
		`)

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, `
			{
				"credential": {
					"access": "temporary-ak",
					"secret": "temporary-sk",
					"securitytoken": "temporary-security-token",
					"expires_at": "2030-01-01T00:00:00.000000Z"
				}
			}
		`)
	})

	th.Mux.HandleFunc("/v3/auth/domains", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestFormValues(t, r, map[string]string{"name": "target-domain"})
		th.TestHeader(t, r, "X-Security-Token", "temporary-security-token")
		if !strings.Contains(r.Header.Get("Authorization"), "Credential=temporary-ak/") {
			t.Errorf("expected temporary AK in Authorization header, got %q", r.Header.Get("Authorization"))
		}

		_, _ = fmt.Fprint(w, `
			{
				"domains": [
					{
						"id": "target-domain-id",
						"name": "target-domain"
					}
				],
				"links": {
					"next": null,
					"previous": null
				}
			}
		`)
	})

	options := golangsdk.AKSKAuthOptions{
		IdentityEndpoint: th.Endpoint(),
		DomainID:         "source-domain-id",
		AccessKey:        "source-ak",
		SecretKey:        "source-sk",
		AgencyName:       "target-agency",
		AgencyDomainName: "target-domain",
	}

	client, err := openstack.AuthenticatedClient(options)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "target-domain-id", client.DomainID)
	th.CheckEquals(t, "target-domain-id", client.AKSKAuthOptions.DomainID)
	th.CheckEquals(t, "temporary-ak", client.AKSKAuthOptions.AccessKey)
	th.CheckEquals(t, "temporary-security-token", client.AKSKAuthOptions.SecurityToken)
}
