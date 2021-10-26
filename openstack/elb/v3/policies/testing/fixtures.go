package testing

const (
	createRequestBody = `
{
  "l7policy" : {
    "action" : "REDIRECT_TO_LISTENER",
    "listener_id" : "e2220d2a-3faf-44f3-8cd6-0c42952bd0ab",
    "redirect_listener_id" : "48a97732-449e-4aab-b561-828d29e45050"
  }
}
`
	createResponseBody = `
{
  "request_id" : "b60d1d9a-5263-45b0-b1d6-2810ac7c52a1",
  "l7policy" : {
    "description" : "",
    "admin_state_up" : true,
    "rules" : [ ],
    "project_id" : "99a3fff0d03c428eac3678da6a7d0f24",
    "listener_id" : "e2220d2a-3faf-44f3-8cd6-0c42952bd0ab",
    "redirect_listener_id" : "48a97732-449e-4aab-b561-828d29e45050",
    "action" : "REDIRECT_TO_LISTENER",
    "position" : 100,
    "provisioning_status" : "ACTIVE",
    "id" : "cf4360fd-8631-41ff-a6f5-b72c35da74be",
    "name" : ""
  }
}
`
	listResponseBody = `
{
  "request_id" : "d3c67339-be91-4813-bb24-85728a5d326a",
  "l7policies" : [ {
    "redirect_pool_id" : "3b34340d-59e8-4c70-9ef5-b41b12023dc9",
    "description" : "",
    "admin_state_up" : true,
    "rules" : [ {
      "id" : "1e5f17df-feec-427e-a162-8e4e05e91085"
    } ],
    "project_id" : "99a3fff0d03c428eac3678da6a7d0f24",
    "listener_id" : "e2220d2a-3faf-44f3-8cd6-0c42952bd0ab",
    "action" : "REDIRECT_TO_POOL",
    "position" : 100,
    "provisioning_status" : "ACTIVE",
    "id" : "0d7bf316-2e03-411f-bf29-c403c04e52bf",
    "name" : "elbv3"
  }, {
    "redirect_pool_id" : "3b34340d-59e8-4c70-9ef5-b41b12023dc9",
    "description" : "",
    "admin_state_up" : true,
    "rules" : [ {
      "id" : "0f5e8c34-09d1-4588-8459-f9b9add0be05"
    } ],
    "project_id" : "99a3fff0d03c428eac3678da6a7d0f24",
    "listener_id" : "e2220d2a-3faf-44f3-8cd6-0c42952bd0ab",
    "action" : "REDIRECT_TO_POOL",
    "position" : 100,
    "provisioning_status" : "ERROR",
    "id" : "2587d8b1-9e8d-459c-9081-7bccaa075d2b",
    "name" : "elbv3"
  } ],
  "page_info" : {
    "previous_marker" : "0d7bf316-2e03-411f-bf29-c403c04e52bf",
    "current_count" : 2
  }
}
`
)
