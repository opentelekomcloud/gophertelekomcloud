package testing

import (
	"fmt"
)

const (
	mappingID = "ACME"

	listURI = "/OS-FEDERATION/mappings"
)

var (
	mappingURI = fmt.Sprintf("%s/%s", listURI, mappingID)

	mappingResponse = fmt.Sprintf(`
{
    "mapping": {
        "id": "%s",
        "rules":[
          {
            "local": [
              {
                "user": {
                  "name": "{0}"
                }
              },
              {
                "groups": "[\"admin\",\"manager\"]"
              }
            ],
            "remote": [
              {
                "type": "uid"
              }
            ]
          }
        ],
        "links": {
            "self": "https://example.com/v3/OS-FEDERATION/mappings/ACME"
        }
    }
}
`, mappingID)

	mappingListResponse = fmt.Sprintf(`
{
    "mappings": [
        {
            "id": "%s",
            "links": {
                "self": "https://example.com/v3/OS-FEDERATION/mappings/ACME"
            }
        },
        {
            "id": "ACME-contractors",
            "rules":[
              {
                "local": [
                  {
                    "user": {
                      "name": "{0}"
                    }
                  },
                  {
                    "groups": "[\"admin\",\"manager\"]"
                  }
                ],
                "remote": [
                  {
                    "type": "uid"
                  }
                ]
              }
            ],
            "links": {
                "self": "https://example.com/v3/OS-FEDERATION/identity_mappings/ACME"
            }
        }
    ],
    "links": {
        "self": "https://example.com/v3/OS-FEDERATION/mappings"
    }
}
`, mappingID)

	updatedMappingResponse = fmt.Sprintf(`
{
    "mapping": {
        "id": "%s",
        "rules": [
          {
            "local": [
              {
                "user": {
                  "name": "samltestid-{0}"
                }
              }
            ],
            "remote": [
              {
                "type": "uid"
              }
            ]
          }
        ],
        "links": {
            "self": "https://example.com/v3/OS-FEDERATION/mappings/ACME"
        }
    }
}
`, mappingID)
)
