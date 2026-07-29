# UCS (Ubiquitous Cloud Native Service) Implementation Guide

This document is a guide for implementing the UCS service (`ucs_cluster` and `fleet`) in `gophertelekomcloud`.

## 1. Client Entry Point (openstack/client.go)
The initialization function `NewUCSV1` has already been added to `client.go`:

```go
// NewUCSV1 creates a ServiceClient that may be used to access the UCS service.
func NewUCSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
	sc, err := initCommonServiceClient(client, eo, "ucs", "v1")
	if err != nil {
		return nil, err
	}
	if client.ProjectID != "" {
		sc.Endpoint = strings.Replace(sc.Endpoint, client.ProjectID+"/", "", 1)
	}
	sc.ResourceBase = sc.Endpoint
	return sc, nil
}
```

---

## 2. Directory Structure (EPS-like Modern Pattern)
Instead of the old `requests.go`/`results.go` pattern, we use a single file per API operation. The base empty files have been generated for you:

```text
openstack/ucs/v1/
├── clusters/
│   ├── Cluster.go
│   ├── Create.go
│   ├── Get.go
│   ├── List.go
│   ├── Update.go
│   ├── Delete.go
│   ├── Activate.go
│   └── testing/
│       ├── requests_test.go
│       └── fixtures.go
└── fleets/ (or clustergroups)
    ├── Fleet.go
    ├── Create.go
    ├── Get.go
    ├── List.go
    ├── Update.go
    ├── Delete.go
    ├── AddCluster.go
    ├── RemoveCluster.go
    ├── EnableFederation.go
    ├── DisableFederation.go
    └── testing/
        ├── requests_test.go
        └── fixtures.go
```

---

## 3. Implementing ucs_cluster (openstack/ucs/v1/clusters)

### Defining the Model (`Cluster.go`)
Define the struct that represents a cluster based on the API docs.

```go
package clusters

type Cluster struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	// Add remaining fields according to documentation
}
```

### Implementing Requests (`Get.go`, `Create.go`, etc.)
Do not use `urls.go` or `GetResult`. Build the URL dynamically and use `extract.IntoStructPtr` directly in the method.

**Example: `Get.go`**
```go
package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Cluster, error) {
	// Root URL for clusters is `clusters`
	raw, err := client.Get(client.ServiceURL("clusters", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Cluster
	// Extract directly into the struct pointer. Use the JSON root key if applicable.
	err = extract.IntoStructPtr(raw.Body, &res, "cluster")
	return &res, err
}
```

---

## 4. Implementing Fleet (openstack/ucs/v1/fleets)

### Defining the Model (`Fleet.go`)
Similar to `Cluster.go`, define the `Fleet` (or `ClusterGroup`) struct here.

### Implementing Requests
The root URL for fleets is `clustergroups`.

**Example: `Delete.go`**
```go
package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) error {
	_, err := client.Delete(client.ServiceURL("clustergroups", id), nil)
	return err
}
```

---

## 5. General Tips and Code Patterns

1. **Strictly adhere to `gofmt` and `goimports` formatting.**
2. **Error Handling**: Never ignore errors. Use explicit `error` returns.
3. Options for requests (e.g., `CreateOpts`, `UpdateOpts`) must implement the appropriate interfaces via `To[Action]Map()` methods, using `golangsdk.BuildRequestBody()`.
4. **Tests**: Write unit tests (`requests_test.go` and `fixtures.go`) mocking responses via `th.SetupHTTP()` and `th.TeardownHTTP()`.
5. Add acceptance tests (`acceptance/openstack/ucs/v1/`) to verify real API calls against the infrastructure.
