package charts

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UploadChartOpts struct {
	Parameters string // Chart Params
	FilePath   string // Path to the chart (.tgz)
}

// UploadChart uploads a Helm chart to the CCE
func UploadChart(client *golangsdk.ServiceClient, opts UploadChartOpts) (*ChartResponse, error) {
	// Open chart file
	file, err := os.Open(opts.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Pipe for streaming multipart content
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	// Start goroutine to write multipart data
	go func() {
		defer pw.Close()
		defer writer.Close()

		// File part, must be named "content"
		part, err := writer.CreateFormFile("content", opts.FilePath)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		if _, err := io.Copy(part, file); err != nil {
			pw.CloseWithError(err)
			return
		}

		_ = writer.WriteField("parameters", opts.Parameters)
	}()

	raw, err := client.Post(client.ServiceURL("charts"), pr, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": writer.FormDataContentType(),
		},
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload chart: %w", err)
	}

	var res ChartResponse
	return &res, extract.Into(raw.Body, &res)
}

type ChartResponse struct {
	// Chart ID.
	ID string `json:"id"`
	// Chart name.
	Name string `json:"name"`
	// Chart value.
	Values string `json:"values"`
	// Chart translation resources.
	Translate string `json:"translate"`
	// Chart description (instruction).
	Instruction string `json:"instruction"`
	// Chart version.
	Version string `json:"version"`
	// Chart description.
	Description string `json:"description"`
	// Chart source.
	Source string `json:"source"`
	// URL to chart icons.
	IconURL string `json:"icon_url"`
	// Whether the chart is public.
	Public bool `json:"public"`
	// URL to the chart.
	ChartURL string `json:"chart_url"`
	// Created at.
	CreateAt string `json:"create_at"`
	// Updated at.
	UpdateAt string `json:"update_at"`
}
