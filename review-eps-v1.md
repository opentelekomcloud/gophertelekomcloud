# Code Review: EPS (Enterprise Project Service) v1

## Summary

This change adds full support for the **Enterprise Project Service (EPS) v1** API to the `gophertelekomcloud` SDK. EPS is a global OTC service for managing enterprise projects — logical containers for grouping cloud resources.

## Key Characteristics of EPS

- **Global service** — not project-scoped
- **Requires domain-scoped token** — project-scoped tokens are rejected with `EPS.0003 Unauthorized`
- **Not in Keystone service catalog** — endpoint must be derived or constructed manually
- **Endpoint URL has no `project_id`** — unlike most OTC services

## Files Changed

### Resource Implementation

| File | Description |
|------|-------------|
| `openstack/eps/v1/projects/requests.go` | CRUD operations: List, Create, Get, Update, Action (enable/disable) |
| `openstack/eps/v1/projects/results.go` | Response types: `EnterpriseProject`, `ProjectListResult`, result wrappers |
| `openstack/eps/v1/projects/urls.go` | URL helpers: `enterprise-projects` base path |
| `openstack/eps/v1/versions/list.go` | Query available API versions at service root |
| `openstack/eps/v1/providers/list.go` | Query supported services (providers) with locale/pagination |
| `openstack/eps/v1/resources/requests.go` | Resource filter and cross-project migration |

### Client Construction

| File | Description |
|------|-------------|
| `openstack/client.go` (`NewEPSV1`) | Service client factory using `initCommonServiceClient` with project_id stripping |
| `acceptance/clients/clients.go` (`NewEPSV1Client`) | Acceptance test client with domain-scoped auth and direct endpoint construction |

### Tests

| File | Description |
|------|-------------|
| `openstack/eps/v1/projects/testing/requests_test.go` | Unit tests (5): TestList, TestGet, TestCreate, TestUpdate, TestAction |
| `openstack/eps/v1/versions/testing/list_test.go` | Unit test (1): TestListVersions |
| `openstack/eps/v1/providers/testing/list_test.go` | Unit tests (2): TestListProviders, TestListProvidersWithOpts |
| `openstack/eps/v1/resources/testing/requests_test.go` | Unit tests (2): TestFilter, TestMigrate |
| `acceptance/openstack/eps/v1/projects_test.go` | Acceptance tests (5): ListAndGet, Lifecycle, Versions, Providers, ResourcesFilter |

## Design Decisions to Review

### 1. Endpoint Construction (`NewEPSV1` in `openstack/client.go`)

```go
func NewEPSV1(client *golangsdk.ProviderClient, eo golangsdk.EndpointOpts) (*golangsdk.ServiceClient, error) {
    sc, err := initCommonServiceClient(client, eo, "eps", "v1.0")
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

**Context**: EPS is not in the service catalog. `initCommonServiceClient` derives the URL from `volumev2` (EVS): `https://evs.{region}.xxx/v2/{project_id}` → replaces `evs` with `eps` and `v2` with `v1.0`, then strips `project_id`.

**Guard**: `if client.ProjectID != ""` — prevents URL corruption when called with a domain-scoped token (empty ProjectID would cause `strings.Replace(url, "/", "", 1)` to strip the first `/` from `https://`).

**Limitation**: This function only works with project-scoped tokens (needs `volumev2` in catalog). With domain-scoped tokens, `initCommonServiceClient` fails. The acceptance test client works around this.

### 2. Acceptance Test Client (`NewEPSV1Client`)

```go
func NewEPSV1Client() (client *golangsdk.ServiceClient, err error) {
    // Domain-scoped auth (no TenantID/TenantName)
    opts := golangsdk.AuthOptions{
        IdentityEndpoint: cloud.AuthInfo.AuthURL,
        Username:         cloud.AuthInfo.Username,
        Password:         cloud.AuthInfo.Password,
        DomainName:       cloud.AuthInfo.UserDomainName,
    }
    pClient, err := openstack.AuthenticatedClient(opts)
    // ...
    // Derive endpoint: iam.{region}.xxx → eps.{region}.xxx/v1.0/
    epsHost := strings.Replace(u.Host, "iam.", "eps.", 1)
    epsEndpoint := fmt.Sprintf("https://%s/v1.0/", epsHost)
    // ...
}
```

**Decision**: Bypasses `NewEPSV1` entirely because domain-scoped tokens have no `volumev2` in catalog. Constructs URL directly from IAM auth URL by replacing `iam.` prefix with `eps.`.

**Trade-off**: Fragile URL assumption (depends on OTC's consistent `{service}.{region}.{domain}` naming) but practical and proven against real API.

### 3. `ExtractProjects` returns `*ProjectListResult` (not `[]EnterpriseProject`)

```go
type ProjectListResult struct {
    EnterpriseProjects []EnterpriseProject `json:"enterprise_projects"`
    TotalCount         int                 `json:"total_count"`
}
```

**Rationale**: The API returns `total_count` alongside the project list. Returning only the slice would discard this pagination metadata.

### 4. Versions Endpoint — URL Stripping

```go
func List(client *golangsdk.ServiceClient) ([]Version, error) {
    baseURL := client.Endpoint
    if idx := strings.Index(baseURL, "/v"); idx > 0 {
        baseURL = baseURL[:idx] + "/"
    }
    // GET https://eps.{region}.xxx/
}
```

**Context**: The versions API queries the service root (`/`) without any version prefix, but `ServiceClient.Endpoint` is `https://eps.xxx/v1.0/`. The function strips the version path to reach the root.

### 5. Resources Filter — `projects` Field Semantics

The `projects` field in `FilterOpts` refers to **OpenStack/IAM project UUIDs** (Keystone tenants), not enterprise project IDs. The enterprise project ID is passed in the URL path. This can be confusing and is documented in the acceptance test with a skip condition.

### 6. No Delete Operation

The API Reference does not document a DELETE endpoint. Enterprise projects can only be **disabled** via `POST /enterprise-projects/{id}/action` with `{"action": "disable"}`.

## Questions for Reviewer

1. Is the `initCommonServiceClient` approach for `NewEPSV1` acceptable given it only works with project-scoped tokens? Should there be a fallback for domain-scoped tokens within `NewEPSV1` itself?
2. Is direct endpoint derivation from the auth URL (`iam.` → `eps.`) acceptable for the acceptance test client, or should there be an `OS_EPS_ENDPOINT` env var override?
3. Should `Action` use typed constants (`ActionEnable = "enable"`, `ActionDisable = "disable"`) instead of raw strings?
4. Is `SinglePageBase` the right pagination choice? The API supports `offset`/`limit` but currently returns all results.
5. Should `FilterOpts.Projects` be renamed or documented to avoid confusion with enterprise project IDs?

## Test Results

```
Unit tests:       10/10 PASS
Acceptance tests:  4/5 PASS, 1 SKIP (ResourcesFilter requires OS_PROJECT_ID)
```

## Full API Coverage

| API Operation | Package | Implemented | Unit Test | Acceptance Test |
|---------------|---------|:-----------:|:---------:|:---------------:|
| List Enterprise Projects | `projects` | ✅ | ✅ | ✅ |
| Create Enterprise Project | `projects` | ✅ | ✅ | ✅ |
| Get Enterprise Project | `projects` | ✅ | ✅ | ✅ |
| Update Enterprise Project | `projects` | ✅ | ✅ | ✅ |
| Enable/Disable (Action) | `projects` | ✅ | ✅ | ✅ |
| Query API Versions | `versions` | ✅ | ✅ | ✅ |
| Query Supported Services | `providers` | ✅ | ✅ | ✅ |
| Query/Filter Resources | `resources` | ✅ | ✅ | ⏭ SKIP (needs OS_PROJECT_ID) |
| Migrate Resources | `resources` | ✅ | ✅ | — (destructive, needs real resources) |

## Package Structure

```
openstack/eps/v1/
├── projects/
│   ├── requests.go          # List, Create, Get, Update, Action
│   ├── results.go           # EnterpriseProject, ProjectListResult, result types
│   ├── urls.go              # rootURL, resourceURL
│   └── testing/
│       └── requests_test.go # 5 unit tests
├── versions/
│   ├── list.go              # List (GET /)
│   └── testing/
│       └── list_test.go     # 1 unit test
├── providers/
│   ├── list.go              # List (GET /enterprise-projects/providers)
│   └── testing/
│       └── list_test.go     # 2 unit tests
└── resources/
    ├── requests.go          # Filter, Migrate
    └── testing/
        └── requests_test.go # 2 unit tests

acceptance/openstack/eps/v1/
└── projects_test.go         # 5 acceptance tests
```
