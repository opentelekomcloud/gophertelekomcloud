package cloud_structuring

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type getOpts struct {
	LogGroupId  string `q:"logGroupId"`
	LogStreamId string `q:"logStreamId"`
}

func Get(client *golangsdk.ServiceClient, logGroupId, logStreamId string) (*StructuringResponse, error) {
	opts := getOpts{
		LogGroupId:  logGroupId,
		LogStreamId: logStreamId,
	}
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("lts", "struct", "template").
		WithQueryParams(opts).Build()
	if err != nil {
		return nil, err
	}
	// GET /v2/{project_id}/lts/struct/template
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}
	// Read raw.Body into bytes
	byteBody, err := io.ReadAll(raw.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}
	defer func() { _ = raw.Body.Close() }()

	// First unmarshal into a string (because it's a double-encoded JSON string)
	var rawJsonString string
	if err := json.Unmarshal(byteBody, &rawJsonString); err != nil {
		return nil, fmt.Errorf("error decoding outer JSON string: %w", err)
	}

	// Now turn that inner JSON string into a new io.Reader
	bodyReader := strings.NewReader(rawJsonString)

	var res StructuringResponse
	err = extract.Into(bodyReader, &res)
	return &res, err
}

type StructuringResponse struct {
	CustomTimeInfo struct {
		Enable bool   `json:"enable"`
		Key    string `json:"key"`
		Value  string `json:"value"`
	} `json:"custom_time_info"`
	// Structured field.
	DemoFields []FieldResponse `json:"demoFields"`
	// Keyword details.
	TagFields []FieldResponse `json:"tagFields"`
	// Sample log event.
	DemoLog string `json:"demoLog"`
	// Attributes of the sample log event.
	DemoLabel string `json:"demoLabel"`
	// Structuring rule ID.
	ID string `json:"id"`
	// Log group ID.
	LogGroupId string `json:"logGroupId"`
	// Log stream ID.
	LogStreamId string `json:"logStreamId"`
	// Project ID.
	ProjectId         string `json:"projectId"`
	SQLAnalysisEnable bool   `json:"sql_analysis_enable"`
	// Template name.
	Name string `json:"templateName"`
	// Regular expression.
	Regex string `json:"regex"`
	// Structuring method.
	Rule *RuleResponse `json:"rule"`
}

type RuleResponse struct {
	// Structuring parameter.
	Param string `json:"param"`
	// Structuring type.
	Type string `json:"type"`
}
type ClusterResponse struct {
	// Kafka cluster name.
	ClusterName string `json:"cluster_name"`
	// Kafka cluster server address.
	Address string `json:"kafka_bootstrap_servers"`
	// Whether SSL encrypted authentication is enabled for Kafka.
	SslEnable bool `json:"kafka_ssl_enable"`
}

type FieldResponse struct {
	Content      string `json:"content"`
	Name         string `json:"fieldName"`
	Index        int    `json:"index"`
	IsAnalysis   bool   `json:"isAnalysis"`
	IsLabelField bool   `json:"isLabelField"`
	IsModified   bool   `json:"isModified"`
	Type         string `json:"type"`
}
