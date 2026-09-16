package analyzer

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Useles5/sirphishalot/internal/config"
)

func TestAnalyze(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.toml")

	cfg := `
[stream]
url = "ws://localhost:9000"

[[brands]]
name = "Apple"
domains = ["apple.com"]

[output]
level = "info" `

	// create config file
	if err := os.WriteFile(cfgPath, []byte(cfg), 0644); err != nil {
		t.Fatal(err)
	}

	analyzer, err := NewAnalyzer(cfgPath)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		hostname  string
		wantBrand *config.Brand
		wantErr   bool
	}{
		{
			name:      "Valid hostname",
			hostname:  "www.apple.com",
			wantBrand: &config.Brand{Name: "Apple", Domains: []string{"apple.com"}},
			wantErr:   false,
		},
		{
			name:      "Invalid hostname",
			hostname:  "",
			wantBrand: nil,
			wantErr:   true,
		},
		{
			name:      "Unknown hostname",
			hostname:  "www.google.com",
			wantBrand: nil,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brand, err := analyzer.Analyze(tt.hostname)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Analyze() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				if brand != nil {
					t.Fatalf("Analyze() brand = %#v, want nil", brand)
				}
			}

			if !reflect.DeepEqual(brand, tt.wantBrand) {
				t.Fatalf("Analyze() brand = %#v, want %#v", brand, tt.wantBrand)
			}

		})
	}
}
