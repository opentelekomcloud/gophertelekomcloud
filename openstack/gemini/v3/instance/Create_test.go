package instance_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/gemini/v3/instance"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCreateBackupStrategy(t *testing.T) {
	for _, tc := range []struct {
		name      string
		backup    string
		expected  int
		startTime string
		wantError bool
	}{
		{name: "numeric", backup: `{"start_time":"02:00-03:00","keep_days":7}`, expected: 7, startTime: "02:00-03:00"},
		{name: "quoted", backup: `{"start_time":"02:00-03:00","keep_days":"8"}`, wantError: true},
		{name: "numeric_zero", backup: `{"start_time":"02:00-03:00","keep_days":0}`, startTime: "02:00-03:00"},
		{name: "quoted_zero", backup: `{"start_time":"02:00-03:00","keep_days":"0"}`, wantError: true},
		{name: "missing_days", backup: `{"start_time":"02:00-03:00"}`, startTime: "02:00-03:00"},
		{name: "null_days", backup: `{"start_time":"02:00-03:00","keep_days":null}`, startTime: "02:00-03:00"},
		{name: "missing_strategy"},
		{name: "null_strategy", backup: `null`},
		{name: "boolean", backup: `{"keep_days":true}`, wantError: true},
		{name: "object", backup: `{"keep_days":{"days":7}}`, wantError: true},
		{name: "array", backup: `{"keep_days":[7]}`, wantError: true},
		{name: "invalid_string", backup: `{"keep_days":"invalid"}`, wantError: true},
		{name: "malformed_json", backup: `{`, wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			opts := createBackupStrategyOpts()
			expectedRequest, err := json.Marshal(opts)
			if err != nil {
				t.Fatal(err)
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				th.TestMethod(t, r, http.MethodPost)
				if r.URL.Path != "/instances" {
					t.Errorf("unexpected create path: %s", r.URL.Path)
				}
				th.TestJSONRequest(t, r, string(expectedRequest))
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusAccepted)
				response := `{"id":"fixture-instance","name":"fixture-instance","status":"creating","job_id":"fixture-job"`
				if tc.backup != "" {
					response += `,"backup_strategy":` + tc.backup
				}
				_, _ = fmt.Fprint(w, response+"}")
			}))
			defer server.Close()
			client := &golangsdk.ServiceClient{
				ProviderClient: &golangsdk.ProviderClient{},
				Endpoint:       server.URL + "/",
			}
			result, err := instance.Create(client, opts)
			if tc.wantError {
				if err == nil {
					t.Fatalf("malformed backup strategy must return a decode error: %#v", result)
				}
				return
			}
			if err != nil {
				t.Fatalf("create response decode failed: %v", err)
			}
			if result.Id != "fixture-instance" || result.JobId != "fixture-job" || result.Status != "creating" {
				t.Fatalf("create response metadata was lost: %#v", result)
			}
			if result.BackupStrategy.KeepDays != tc.expected || result.BackupStrategy.StartTime != tc.startTime {
				t.Fatalf("unexpected backup strategy: %#v", result.BackupStrategy)
			}
		})
	}
}

func TestBackupStrategyMarshal(t *testing.T) {
	value := instance.BackupStrategy{StartTime: "02:00-03:00", KeepDays: 7}
	th.AssertJSONEquals(t, `{"start_time":"02:00-03:00","keep_days":7}`, value)
}

func TestCreateBackupStrategyRequest(t *testing.T) {
	zeroDays, sevenDays, maxDays := 0, 7, 35
	for _, tc := range []struct {
		name       string
		strategy   *instance.BackupStrategyOpt
		wantBackup string
	}{
		{
			name:       "numeric_retention",
			strategy:   &instance.BackupStrategyOpt{StartTime: "02:00-03:00", KeepDays: &sevenDays},
			wantBackup: `{"start_time":"02:00-03:00","keep_days":7}`,
		},
		{
			name:       "disable_automated_backup",
			strategy:   &instance.BackupStrategyOpt{StartTime: "02:00-03:00", KeepDays: &zeroDays},
			wantBackup: `{"start_time":"02:00-03:00","keep_days":0}`,
		},
		{
			name:       "maximum_retention",
			strategy:   &instance.BackupStrategyOpt{StartTime: "02:00-03:00", KeepDays: &maxDays},
			wantBackup: `{"start_time":"02:00-03:00","keep_days":35}`,
		},
		{
			name:       "default_retention",
			strategy:   &instance.BackupStrategyOpt{StartTime: "02:00-03:00"},
			wantBackup: `{"start_time":"02:00-03:00"}`,
		},
		{name: "default_strategy"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			opts := createBackupStrategyOpts()
			opts.BackupStrategy = tc.strategy
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				th.TestMethod(t, r, http.MethodPost)
				if r.URL.Path != "/instances" {
					t.Errorf("unexpected create path: %s", r.URL.Path)
				}
				var request map[string]json.RawMessage
				th.AssertNoErr(t, json.NewDecoder(r.Body).Decode(&request))
				backup, present := request["backup_strategy"]
				if tc.strategy == nil {
					if present {
						t.Errorf("omitted backup strategy was sent: %s", backup)
					}
				} else {
					th.AssertJSONEquals(t, tc.wantBackup, backup)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusAccepted)
				_, _ = fmt.Fprint(w, `{"id":"fixture-instance"}`)
			}))
			defer server.Close()
			client := &golangsdk.ServiceClient{
				ProviderClient: &golangsdk.ProviderClient{},
				Endpoint:       server.URL + "/",
			}
			result, err := instance.Create(client, opts)
			th.AssertNoErr(t, err)
			th.AssertEquals(t, "fixture-instance", result.Id)
		})
	}
}

func createBackupStrategyOpts() instance.CreateOpts {
	enterpriseProjectID, sslOption := "0", "1"
	keepDays := 7
	return instance.CreateOpts{
		Name: "fixture-instance",
		DataStore: instance.DataStoreOpt{
			Type: "cassandra", Version: "3.11", StorageEngine: "rocksDB",
		},
		Region: "eu-de", AvailabilityZone: "eu-de-01,eu-de-02,eu-de-03",
		VpcId: "fixture-vpc", SubnetId: "fixture-subnet", SecurityGroupId: "fixture-security-group",
		Password: "TestOnlyPassword9!", Mode: "Cluster",
		Flavor: []instance.FlavorOpt{{
			Num: "3", Size: "100", Storage: "ULTRAHIGH", SpecCode: "geminidb.cassandra.xlarge.8",
		}},
		ConfigurationId:     "fixture-configuration",
		EnterpriseProjectId: &enterpriseProjectID,
		SslOption:           &sslOption,
		BackupStrategy: &instance.BackupStrategyOpt{
			StartTime: "02:00-03:00", KeepDays: &keepDays,
		},
	}
}
