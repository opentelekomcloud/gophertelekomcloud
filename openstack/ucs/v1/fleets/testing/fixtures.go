package testing

const (
	fleetID   = "bffbb35b-7949-11ee-886c-0255ac100037"
	clusterID = "49077339-f1cd-11ec-a2be-0255ac1001c2"
)

const createRequest = `
{
  "metadata": {
    "name": "group02281605"
  },
  "spec": {
    "clusterIds": ["514c1a3c-8ec7-11ec-b384-0255ac100189"],
    "description": "test fleet"
  }
}`

const createResponse = `
{
  "uid": "bffbb35b-7949-11ee-886c-0255ac100037"
}`

const getResponse = `
{
  "kind": "ClusterGroup",
  "apiVersion": "v1",
  "metadata": {
    "name": "cluster-test",
    "uid": "bffbb35b-7949-11ee-886c-0255ac100037",
    "creationTimestamp": "2023-11-02 06:33:35.558128 +0000 UTC"
  },
  "spec": {
    "federationId": "e2f27cc6-82b5-11ee-84e3-0255ac100032",
    "federationVersion": "v1.7.0",
    "description": "test fleet",
    "dnsSuffix": ["www.oidc.com"]
  },
  "status": {
    "conditions": [
      {
        "type": "Federation",
        "status": "Unavailable",
        "reason": "FederationUnavailable"
      }
    ]
  }
}`

const listResponse = `
{
  "items": [
    {
      "kind": "ClusterGroup",
      "apiVersion": "v1",
      "metadata": {
        "name": "cluster-test",
        "uid": "bffbb35b-7949-11ee-886c-0255ac100037"
      },
      "spec": {
        "federationId": "e2f27cc6-82b5-11ee-84e3-0255ac100032"
      },
      "status": {}
    }
  ],
  "total": 1
}`

const updateRequest = `
{
  "description": "new description"
}`

const addClusterRequest = `
{
  "clusterGroupID": "bffbb35b-7949-11ee-886c-0255ac100037"
}`
