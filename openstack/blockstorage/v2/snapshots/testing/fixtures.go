package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	fake "github.com/opentelekomcloud/gophertelekomcloud/testhelper/client"
)

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		expected := map[string]string{
			"all_tenants": "true",
			"limit":       "2",
			"name":        "snapshot-001",
			"offset":      "1",
			"project_id":  "0c2eba2c5af04d3f9e9d0d410b371fde",
			"status":      "available",
			"volume_id":   "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
		}
		expected["marker"] = r.URL.Query().Get("marker")
		th.TestFormValues(t, r, expected)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if r.URL.Query().Get("marker") == "96c3bda7-c82a-4f50-be73-ca7621794835" {
			_, _ = fmt.Fprint(w, `
    {
      "snapshots": [
        {
          "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
          "name": "snapshot-003",
          "volume_id": "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
          "description": "Monthly Backup",
          "status": "available",
          "size": 50,
          "metadata": {},
          "created_at": "2017-06-01T03:35:03.000000",
          "updated_at": null
        }
      ],
      "snapshots_links": null
    }
    `)
			return
		}

		_, _ = fmt.Fprintf(w, `
    {
      "snapshots": [
        {
          "id": "289da7f8-6440-407c-9fb4-7db01ec49164",
          "name": "snapshot-001",
          "volume_id": "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
          "description": "Daily Backup",
          "status": "available",
          "size": 0,
          "metadata": {},
		  "created_at": "2017-05-30T03:35:03.000000",
          "updated_at": null
        },
        {
          "id": "96c3bda7-c82a-4f50-be73-ca7621794835",
          "name": "snapshot-002",
          "volume_id": "76b8950a-8594-4e5b-8dce-0dfa9c696358",
          "description": "Weekly Backup",
          "status": "available",
          "size": 25,
          "metadata": {
            "environment": "test"
          },
		  "created_at": "2017-05-30T03:35:03.000000",
          "updated_at": "2017-05-31T03:35:03.000000"
        }
      ],
      "snapshots_links": [
        {
          "href": "%s/snapshots?all_tenants=true&limit=2&marker=96c3bda7-c82a-4f50-be73-ca7621794835&name=snapshot-001&offset=1&project_id=0c2eba2c5af04d3f9e9d0d410b371fde&status=available&volume_id=521752a6-acf6-4b2d-bc7a-119f9148cd8c",
          "rel": "next"
        }
      ]
    }
    `, th.Server.URL)
	})
}

func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `
{
    "snapshot": {
        "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
        "name": "snapshot-001",
        "description": "Daily backup",
        "volume_id": "521752a6-acf6-4b2d-bc7a-119f9148cd8c",
        "status": "available",
        "size": 0,
        "metadata": {},
		"created_at": "2017-05-30T03:35:03.000000",
        "updated_at": null
    }
}
      `)
	})
}

func MockCreateResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "snapshot": {
        "volume_id": "1234",
        "force": true,
        "name": "snapshot-001",
        "description": "Daily backup",
        "metadata": {
            "environment": "test"
        }
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		_, _ = fmt.Fprint(w, `
{
    "snapshot": {
        "volume_id": "1234",
        "name": "snapshot-001",
        "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
        "description": "Daily backup",
        "status": "creating",
        "size": 0,
        "metadata": {
            "environment": "test"
        },
        "created_at": "2017-05-30T03:35:03.000000"
  }
}
    `)
	})
}

func MockUpdateMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/123/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
    {
      "metadata": {
        "empty": "",
        "key": "v1"
      }
    }
    `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `
      {
        "metadata": {
          "empty": "",
          "key": "v1"
        }
      }
    `)
	})
}

func MockDeleteResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}
