package detector

import (
	"reflect"
	"testing"

	"github.com/Useles5/sirphishalot/internal/config"
)

func TestDetect(t *testing.T) {
	cfg := &config.Config{
		Brands: []config.Brand{
			{
				Name:    "Google",
				Domains: []string{"google.com"},
			},
			{
				Name:    "Amazon",
				Domains: []string{"amazon.com"},
			},
		},
	}

	detector, err := NewDetector(cfg)
	if err != nil {
		t.Fatalf("NewDetector() error = %v, want nil", err)
	}

	tests := []struct {
		name      string
		hostname  string
		wantBrand *config.Brand
		wantErr   bool
	}{
		{
			name:     "Known Brand",
			hostname: "store.google.com",
			wantBrand: &config.Brand{
				Name:    "Google",
				Domains: []string{"google.com"},
			},
			wantErr: false,
		},
		{
			name:      "Unknown Brand",
			hostname:  "netflix.com",
			wantBrand: nil,
			wantErr:   false,
		},
		{
			name:      "Invalid hostname",
			hostname:  "",
			wantBrand: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brand, err := detector.Detect(tt.hostname)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Detect(%q) error = %v, wantErr %v", tt.hostname, err, tt.wantErr)
			}

			// if wanted error and got error, return
			// check for invalid hostname testcase
			if tt.wantErr {
				return
			}

			// check for unknown brand testcase
			if tt.wantBrand == nil {
				if brand != nil {
					t.Fatalf("Detect(%q) brand = %v, want nil", tt.hostname, brand)
				}
				return
			}

			if brand == nil {
				t.Fatalf("Detect(%q) brand = nil, want %v", tt.hostname, tt.wantBrand)
			}

			if !reflect.DeepEqual(brand, tt.wantBrand) {
				t.Fatalf("Detect(%q) brand = %#v, want %#v", tt.hostname, brand, tt.wantBrand)
			}
		})
	}
}
