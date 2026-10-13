package testing

const clusterID = "b0d1ecb5-7947-11ee-9467-0255ac1001bf"

const createRequest = `
{
  "kind": "Cluster",
  "apiVersion": "v1",
  "metadata": {
    "name": "cce-cluster"
  },
  "spec": {
    "category": "self",
    "type": "turbo",
    "provider": {
      "CCE": "cce"
    },
    "country": "DE",
    "manageType": "discrete"
  }
}`

const createResponse = `
{
  "uid": "b0d1ecb5-7947-11ee-9467-0255ac1001bf"
}`

const getResponse = `
{
  "kind": "Cluster",
  "apiVersion": "v1",
  "metadata": {
    "name": "test-cluster",
    "uid": "b0d1ecb5-7947-11ee-9467-0255ac1001bf",
    "creationTimestamp": "2023-11-02T06:36:14Z"
  },
  "spec": {
    "clusterGroupID": "bffbb35b-7949-11ee-886c-0255ac100037",
    "manageType": "grouped",
    "provider": "cce",
    "type": "cce",
    "category": "self",
    "country": "DE",
    "city": "150900",
    "projectID": "b6315dd3d0ff4be5b31a963256794989"
  },
  "status": {
    "kubernetesVersion": "v1.25",
    "phase": "Available",
    "conditions": [
      {
        "type": "Ready",
        "status": "True",
        "reason": "ClusterAvailable"
      }
    ]
  }
}`

const listResponse = `
{
  "items": [
    {
      "kind": "Cluster",
      "apiVersion": "v1",
      "metadata": {
        "name": "test-cluster",
        "uid": "b0d1ecb5-7947-11ee-9467-0255ac1001bf"
      },
      "spec": {
        "manageType": "grouped",
        "provider": "cce",
        "type": "cce",
        "category": "self"
      },
      "status": {
        "phase": "Available"
      }
    }
  ],
  "total": 1
}`

const updateRequest = `
{
  "kind": "Cluster",
  "apiVersion": "v1",
  "spec": {
    "country": "AL",
    "city": "AL"
  }
}`
