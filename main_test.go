package main

import (
	"testing"

	publicmanifest "github.com/Silo-Server/silo-plugin-sdk/pkg/pluginsdk/manifest"
)

func TestManifestLoads(t *testing.T) {
	m, err := publicmanifest.Load(manifestJSON)
	if err != nil {
		t.Fatalf("failed to load manifest: %v", err)
	}
	if m.PluginId != "silo.sportarr" {
		t.Errorf("expected plugin_id silo.sportarr, got %s", m.PluginId)
	}
	if len(m.Capabilities) != 1 {
		t.Fatalf("expected 1 capability, got %d", len(m.Capabilities))
	}
	if m.Capabilities[0].Id != "sportarr" {
		t.Errorf("expected capability id sportarr, got %s", m.Capabilities[0].Id)
	}
	if m.Capabilities[0].Type != "metadata_provider.v1" {
		t.Errorf("expected capability type metadata_provider.v1, got %s", m.Capabilities[0].Type)
	}
}

func TestSportarrCanonicalPath(t *testing.T) {
	base := "https://sportarr.net"

	tests := []struct {
		name     string
		imageURL string
		want     string
	}{
		{"empty", "", ""},
		{"full url", "https://sportarr.net/api/images/abc123", "sportarr:///api/images/abc123"},
		{"external url", "https://example.com/image.jpg", "https://example.com/image.jpg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sportarrCanonicalPath(base, tt.imageURL)
			if got != tt.want {
				t.Errorf("sportarrCanonicalPath(%q) = %q, want %q", tt.imageURL, got, tt.want)
			}
		})
	}
}

func TestResolveOneSportarrPath(t *testing.T) {
	base := "https://sportarr.net"

	tests := []struct {
		name string
		path string
		want string
	}{
		{"empty", "", ""},
		{"canonical", "sportarr:///api/images/abc123", "https://sportarr.net/api/images/abc123"},
		{"full url passthrough", "https://example.com/image.jpg", "https://example.com/image.jpg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveOneSportarrPath(base, tt.path, "")
			if got != tt.want {
				t.Errorf("resolveOneSportarrPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
