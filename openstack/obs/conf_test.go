package obs

import (
	"strings"
	"testing"
)

// A client without WithSignature must default to the OBS signature.
func TestDefaultSignatureIsObs(t *testing.T) {
	client, err := New("ak", "sk", "https://obs.eu-de.otc.t-systems.com")
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	if client.conf.signature != SignatureObs {
		t.Errorf("expected default signature %q, got %q", SignatureObs, client.conf.signature)
	}
}

// path-style must not downgrade the OBS signature to V2 (SSE KMS needs x-obs-*).
func TestPathStyleKeepsObsSignature(t *testing.T) {
	client, err := New("ak", "sk", "https://obs.eu-de.otc.t-systems.com",
		WithPathStyle(true), WithSignature(SignatureObs))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	if !client.conf.pathStyle {
		t.Errorf("expected pathStyle to be true")
	}
	if client.conf.signature != SignatureObs {
		t.Errorf("expected signature %q, got %q", SignatureObs, client.conf.signature)
	}
}

// A dotted bucket name must use path-style automatically (its virtual-hosted
// URL isn't covered by the wildcard cert), while a dotless one stays virtual-hosted.
func TestDottedBucketUsesPathStyle(t *testing.T) {
	const host = "obs.eu-de.otc.t-systems.com"
	client, err := New("ak", "sk", "https://"+host, WithSignature(SignatureObs))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// pathStyle is off, but a dotted bucket must still go path-style.
	if client.conf.pathStyle {
		t.Fatalf("expected pathStyle to be false by default")
	}

	requestURL, canonicalizedURL := client.conf.formatUrls("department.service.development", "", nil, true)
	if !strings.HasPrefix(requestURL, "https://"+host) {
		t.Errorf("expected host-rooted path-style URL, got %q", requestURL)
	}
	if !strings.Contains(requestURL, host+":443/department.service.development") {
		t.Errorf("expected bucket in the path, got %q", requestURL)
	}
	if canonicalizedURL != "/department.service.development" {
		t.Errorf("unexpected canonicalized URL %q", canonicalizedURL)
	}

	// A dotless bucket stays virtual-hosted.
	plainURL, _ := client.conf.formatUrls("plainbucket", "", nil, true)
	if !strings.HasPrefix(plainURL, "https://plainbucket."+host) {
		t.Errorf("expected virtual-hosted URL for a dotless bucket, got %q", plainURL)
	}
}
