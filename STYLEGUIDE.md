
## On Pull Requests

- Please make sure to read our [contributing guide](/.github/CONTRIBUTING.md).

- Before you start a PR there needs to be a Github issue and a discussion about it
  on that issue with a core contributor, even if it's just a 'SGTM'.

- A PR's description must reference the issue it closes with a `For <ISSUE NUMBER>` (e.g. For #293).

- A PR's description must contain link(s) to documentation on [docs portal](https://docs.otc.t-systems.com/) For example,
  a pull request implementing the creation of a Dedicated WAF instance might put the
  following link in the description:

  https://docs.otc.t-systems.com/web-application-firewall-dedicated/api-ref/apis/dedicated_instance_management/creating_a_dedicated_waf_engine.html#createinstance

  From that link, a reviewer (or user) can verify the fields in the request/response
  objects in the PR.

- A PR that is in-progress should have `[wip]` in front of the PR's title. When
  ready for review, remove the `[wip]` and ping a core contributor with an `@`.

- Forcing PRs to be small can have the effect of users submitting PRs in a hierarchical chain, with
  one depending on the next. If a PR depends on another one, it should have a [Pending #PRNUM]
  prefix in the PR title. In addition, it will be the PR submitter's responsibility to remove the
  [Pending #PRNUM] tag once the PR has been updated with the merged, dependent PR. That will
  let reviewers know it is ready to review.

- A PR should be small. Even if you intend on implementing an entire
  service, a PR should only be one route of that service
  (e.g. create server or get server, but not both).

- Unless explicitly asked, do not squash commits in the middle of a review; only
  append. It makes it difficult for the reviewer to see what's changed from one
  review to the next.

## On Code

- In re design: follow as closely as is reasonable the code already in the library.
  Most operations (e.g. create, delete) admit the same design.

- Unit tests and acceptance (integration) tests must be written to cover each PR.
  Tests for operations with several options (e.g. list, create) should include all
  the options in the tests. This will allow users to verify an operation on their
  own infrastructure and see an example of usage.

- If in doubt, ask in-line on the PR.

### File Structure

- New service packages should follow the operation-per-file layout used by
  newer services such as `apigw` and `fgs`.

  - Service code lives under `openstack/<service>/<version>/<resource>/`.
    The `<resource>` directory is the Go package name and should be lowercase.
    Use snake case when the resource name is made of several words, for example
    `app_auth`, `async_config`, or `dependency_version`.
  - Each public API operation should live in its own file named after the
    exported operation function, for example `Create.go`, `List.go`,
    `GetMetadata.go`, `UpdateEIP.go`, or `ListAPIBoundPolicy.go`.
  - Put the operation request options, response types, small helper types, and
    extraction helpers in the same operation file when they are only used by
    that operation.
  - Shared resource types may live in the operation file where they are first
    introduced if that matches the existing package style. Move them to a common
    file only when they are reused widely enough to make the operation files
    harder to read.
  - Acceptance tests live under
    `acceptance/openstack/<service>/<version>/` and are named after the tested
    resource or feature in snake case, for example `gateway_test.go`,
    `app_auth_test.go`, or `dependency_version_test.go`.

### Naming

- Operation functions and files should use Go exported PascalCase names:
  `Create`, `Delete`, `Get`, `List`, `Update`, or a more specific API action
  such as `CreateAlias`, `InvokeAsync`, `PublishVersion`, `UpdateStatus`, or
  `ListGatewayFeatures`.

- Request option structs should be named after the operation:
  `CreateOpts`, `ListOpts`, `UpdateOpts`, or `<Action>Opts` when the operation
  name is more specific, for example `CreateAliasOpts`.

- Response structs should use clear exported names that match the API resource
  or returned payload, for example `GatewayResp`, `FuncGraph`, or
  `FuncAliasesResp`.

- Follow the acronym casing already used by the package and API. Keep common
  API acronyms uppercase when that is the established operation name, for
  example `UpdateEIP.go`, `ListAPIBoundPolicy.go`, `VpcID`, and `ProjectID`;
  keep existing mixed-case names such as `GetLts.go` consistent within their
  package.

- For request body operations, build payloads with `build.RequestBody(opts, "")`
  and keep path-only fields out of the JSON body with `json:"-"`.

- For query string operations, use `q` struct tags and
  `golangsdk.BuildQueryString(&opts)`.

- Prefer short, consistent local names in operation functions:
  `b` for request bodies, `q` for query strings, `raw` for the SDK response,
  `res` for the decoded response value, and `err` for errors.
