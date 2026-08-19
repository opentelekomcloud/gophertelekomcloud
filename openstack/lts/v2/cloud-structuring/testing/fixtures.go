package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

const listCustomResponse = `
{
  "results": [
    {
      "create_time": 1641258099551,
      "demo_fields": [
        {
          "content": "value",
          "field_name": "message",
          "index": 0,
          "is_analysis": false,
          "relation": "root",
          "type": "string",
          "user_defined_name": "alias"
        }
      ],
      "demo_label": "sample-label",
      "demo_log": "sample log",
      "id": "43a8cc7b-b632-4c36-a65d-8150e98219f1",
      "project_id": "2a473356cca5487f8373be89xxxxxxxx",
      "rule": {
        "param": "{\"layers\":1}",
        "type": "json"
      },
      "tag_fields": [
        {
          "content": "host-1",
          "field_name": "host",
          "index": 0,
          "is_analysis": true,
          "type": "string"
        }
      ],
      "template_name": "custom-json",
      "template_type": "json"
    }
  ]
}`

func handleListCustom(t *testing.T) {
	t.Helper()
	th.Mux.HandleFunc("/lts/struct/customtemplate", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, http.MethodGet)
		th.TestHeader(t, r, "content-type", "application/json")
		th.TestFormValues(t, r, map[string]string{"id": "template-id"})
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, listCustomResponse)
	})
}

func handleListCustomError(t *testing.T) {
	t.Helper()
	th.Mux.HandleFunc("/lts/struct/customtemplate", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, `{"error_code":"LTS.2017"}`, http.StatusInternalServerError)
	})
}
