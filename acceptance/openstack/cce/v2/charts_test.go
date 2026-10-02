package v2

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cce/v2/charts"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestChartLifecycle(t *testing.T) {
	t.Skip("Too long. Non reproducible in CI")
	client, err := clients.NewCceV2Client()
	th.AssertNoErr(t, err)

	// Step 1: Create minimal chart directory
	chartDir := "mychart"
	os.Mkdir(chartDir, 0755)
	t.Cleanup(func() {
		os.RemoveAll(chartDir)
	})

	os.WriteFile(filepath.Join(chartDir, "Chart.yaml"), []byte(`apiVersion: v2
name: mychart
version: 0.1.0
description: A minimal chart
`), 0644)

	os.WriteFile(filepath.Join(chartDir, "values.yaml"), []byte(`replicaCount: 1
`), 0644)

	// Step 2: Tar + gzip the chart folder
	tgzPath := "mychart-0.1.0.tgz"
	if err := createTarGz(tgzPath, chartDir); err != nil {
		t.Fatalf("Failed to create tgz: %v", err)
	}
	t.Cleanup(func() {
		os.Remove(tgzPath)
	})

	parameters := `{"skip_lint":true,"override":true,"source":"package"}`

	opts := charts.UploadChartOpts{
		FilePath:   tgzPath,
		Parameters: parameters,
	}

	res, err := charts.UploadChart(client, opts)
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		t.Logf("Deleting chart : %s", res.ID)
		charts.Delete(client, res.ID)
	})

	t.Logf("Chart uploaded successfully: %+v", res)
}

func createTarGz(tgzPath, srcDir string) error {
	f, err := os.Create(tgzPath)
	if err != nil {
		return err
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	return filepath.Walk(srcDir, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.Mode().IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(srcDir, file)
		if err != nil {
			return err
		}

		hdr := &tar.Header{
			Name:    relPath,
			Mode:    int64(fi.Mode().Perm()),
			Size:    fi.Size(),
			ModTime: fi.ModTime(),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		fh, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fh.Close()

		_, err = io.Copy(tw, fh)
		return err
	})
}
