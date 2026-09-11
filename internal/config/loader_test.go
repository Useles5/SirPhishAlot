package config

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		config  string
		wantErr bool
	}{
		{
			name: "Valid Config",
			config: `[stream]
url = "ws://localhost:9000"

[[brands]]
name = "Amazon"
domains = ["amazon.com"]

[[brands]]
name = "PayPal"
domains = ["paypal.com"]`,
			wantErr: false,
		},
		{
			name: "Invalid Config",
			config: `[stream
url = "ws://loclhost:9000"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "config.toml")

			//0, 6 -> read+write access to user, 4 -> read access to group, 4 -> read access to others
			if err := os.WriteFile(path, []byte(tt.config), 0644); err != nil {
				t.Fatal(err)
			}

			cfg, err := Load(path)
			if (err != nil) != tt.wantErr {
				t.Fatalf("got error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if cfg.Stream.URL != "ws://localhost:9000" {
				t.Errorf("got URL = %s, want ws://localhost:9000", cfg.Stream.URL)
			}

			if len(cfg.Brands) != 2 {
				t.Fatalf("got %d brands, want 2", len(cfg.Brands))
			}

			if cfg.Brands[0].Name != "Amazon" {
				t.Errorf("got brand = %s, want Amazon", cfg.Brands[0].Name)
			}

			if cfg.Brands[0].Domains[0] != "amazon.com" {
				t.Errorf("got brand = %s, want amazon.com", cfg.Brands[0].Domains[0])
			}

			if cfg.Brands[1].Name != "PayPal" {
				t.Errorf("got brand = %s, want PayPal", cfg.Brands[1].Name)
			}

			if cfg.Brands[1].Domains[0] != "paypal.com" {
				t.Errorf("got brand = %s, want paypal.com", cfg.Brands[1].Domains[0])
			}

		})
	}
}

func TestLoad_MissingFile(t *testing.T) {
	_, err := Load("does-not-exist.toml")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}
