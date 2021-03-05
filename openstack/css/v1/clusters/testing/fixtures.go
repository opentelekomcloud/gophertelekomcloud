package testing

import "fmt"

const (
	vpcID    = "fccd753c-91c3-40e2-852f-5ddf76d1a1b2"
	subnetID = "af1c65ae-c494-4e24-acd8-81d6b355c9f1"
	sgID     = "7e3fed21-1a44-4101-ab29-34e57124f614"
	cmkID    = "42546bb1-8025-4ad1-868f-600729c341ae"

	clusterID   = "ef683016-871e-48bc-bf93-74a29d60d214"
	clusterName = "ES-Test"
)

const (
	listResponseBody = `
{
    "clusters": [
        {
            "datastore": {
                "type": "elasticsearch",
                "version": "7.6.2"
            },
            "instances": [
                {
                    "status": "200",
                    "type": "ess",
                    "id": "c8c90973-924d-4201-b9ff-f32279c87d0e",
                    "name": "css-5492-ess-esn-1-1",
                    "specCode": "css.xlarge.2",
                    "azCode": "eu-de-01"
                }
            ],
            "updated": "2020-12-01T07:47:34",
            "name": "css-5492",
            "created": "2020-12-01T07:47:34",
            "id": "66ea1e42-4ee2-44ad-bd80-c86e6d8c6b9e",
            "status": "200",
            "endpoint": "10.16.0.151:9200",
            "vpcId": "e7daa617-3ee6-4ff1-b042-8cda4a006a46",
            "subnetId": "6253dc44-24cd-4c0a-90b3-f965e7f4dcd4",
            "securityGroupId": "d478041e-bcbe-4d69-a492-b6122d774b7f",
            "httpsEnable": false,
            "authorityEnable": false,
            "diskEncrypted": true,
            "cmkId": "00f05033-f8ac-4ceb-a1ce-4072fadb6b28",
            "actionProgress": {},
            "actions": [],
            "tags": []
        },
        {
            "datastore": {
                "type": "elasticsearch",
                "version": "6.2.3"
            },
            "instances": [
                {
                    "status": "200",
                    "type": "ess",
                    "id": "a24adddb-1553-4873-9978-9d064418f903",
                    "name": "css-1d01-ess-esn-1-1",
                    "specCode": "css.xlarge.2",
                    "azCode": "eu-de-01"
                }
            ],
            "updated": "2020-11-26T10:08:44",
            "name": "css-1d01",
            "created": "2020-11-26T10:08:44",
            "id": "af5fbac7-b386-4305-b201-820a0f51f4f1",
            "status": "200",
            "endpoint": "10.16.0.124:9200",
            "vpcId": "e7daa617-3ee6-4ff1-b042-8cda4a006a46",
            "subnetId": "6253dc44-24cd-4c0a-90b3-f965e7f4dcd4",
            "securityGroupId": "d478041e-bcbe-4d69-a492-b6122d774b7f",
            "httpsEnable": true,
            "authorityEnable": false,
            "diskEncrypted": false,
            "cmkId": "",
            "actionProgress": {},
            "actions": [],
            "tags": []
        },
        {
            "datastore": {
                "type": "elasticsearch",
                "version": "7.6.2"
            },
            "instances": [
                {
                    "status": "303",
                    "type": "ess",
                    "id": "071c7ecf-a11d-45bd-9564-201ceb7cfae3",
                    "name": "css-9b36-ess-esn-1-1",
                    "specCode": "css.xlarge.2",
                    "azCode": "eu-de-02"
                }
            ],
            "updated": "2020-11-13T14:33:24",
            "name": "css-9b36",
            "created": "2020-11-13T14:33:26",
            "id": "cdb26954-c743-47dd-b23a-b693205eb2da",
            "status": "303",
            "endpoint": null,
            "vpcId": "e7daa617-3ee6-4ff1-b042-8cda4a006a46",
            "subnetId": "6253dc44-24cd-4c0a-90b3-f965e7f4dcd4",
            "securityGroupId": "d478041e-bcbe-4d69-a492-b6122d774b7f",
            "httpsEnable": true,
            "authorityEnable": true,
            "diskEncrypted": false,
            "cmkId": "",
            "actionProgress": {},
            "actions": [],
            "tags": []
        }
    ]
}
`
)

var (
	createRequestBody = fmt.Sprintf(`
{
    "cluster": {
        "name": "ES-Test",
        "instanceNum": 4,
        "instance": {
            "flavorRef": "css.large.8",
            "volume": {
                "volume_type": "COMMON",
                "size": 100
            },
            "nics": {
                "vpcId": "%s",
                "netId": "%s",
                "securityGroupId": "%s"
            }
        },
        "httpsEnable": "false",
        "diskEncryption": {
            "systemEncrypted": "1",
            "systemCmkid": "%s"
        }
    }
}
`, vpcID, subnetID, sgID, cmkID)

	createResponseBody = fmt.Sprintf(`
{
  "cluster": {
    "id": "%s",
    "name": "%s"
  }
}
`, clusterID, clusterName)
	getResponseBody = fmt.Sprintf(`
{
    "datastore": {
        "type": "elasticsearch",
        "version": "7.6.2"
    },
    "instances": [
        {
            "status": "200",
            "type": "ess",
            "id": "c2f29369-1985-4028-8e72-89cbb96a299d",
            "name": "%s-ess-esn-1-1",
            "specCode": "css.xlarge.2",
            "azCode": "eu-de-01"
        }
    ],
    "updated": "2020-12-03T07:02:08",
    "name": "%[1]s",
    "created": "2020-12-03T07:02:08",
    "id": "%s",
    "status": "200",
    "endpoint": "10.16.0.88:9200",
    "vpcId": "%s",
    "subnetId": "%s",
    "securityGroupId": "%s",
    "httpsEnable": true,
    "authorityEnable": true,
    "diskEncrypted": false,
    "actionProgress": {},
    "actions": [],
    "tags": []
}
`, clusterName, clusterID, vpcID, subnetID, sgID)
)
