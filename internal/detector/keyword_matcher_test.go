package detector

import (
	"reflect"
	"testing"

	"github.com/Useles5/sirphishalot/internal/config"
)

func TestKeywordMatch(t *testing.T) {
	cfg := &config.Config{
		Brands: []config.Brand{
			{Name: "Apple", Domains: []string{"apple.com"}},
			{Name: "PayPal", Domains: []string{"paypal.com"}},
			{Name: "Google", Domains: []string{"google.com"}},
		},
	}

	matcher, err := NewMatcher(cfg)
	if err != nil {
		t.Fatalf("NewMatcher() error = %v", err)
	}

	tests := []struct {
		name             string
		registeredDomain string
		wantBrand        *config.Brand
		wantBool         bool
	}{
		{
			name:             "matches hyphenated brand",
			registeredDomain: "apple-login.com",
			wantBrand: &config.Brand{
				Name:    "Apple",
				Domains: []string{"apple.com"},
			},
			wantBool: true,
		},
		{
			name:             "matches exact brand domain",
			registeredDomain: "google.com",
			wantBrand: &config.Brand{
				Name:    "Google",
				Domains: []string{"google.com"},
			},
			wantBool: true,
		},
		{
			name:             "matches multi-part public suffix",
			registeredDomain: "google-login.co.uk",
			wantBrand: &config.Brand{
				Name:    "Google",
				Domains: []string{"google.com"},
			},
			wantBool: true,
		},
		{
			name:             "does not match substring",
			registeredDomain: "pineapple.com",
			wantBrand:        nil,
			wantBool:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brand, found := keywordMatch(tt.registeredDomain, matcher)
			if found != tt.wantBool {
				t.Fatalf("KeywordMatch() got %v, want %v", found, tt.wantBool)
			}
			if !reflect.DeepEqual(brand, tt.wantBrand) {
				t.Fatalf("KeywordMatch() got %#v, want %#v", brand, tt.wantBrand)
			}
		})
	}
}

func TestDetectKeywordMatch(t *testing.T) {
	cfg := &config.Config{
		Brands: []config.Brand{
			{Name: "Apple", Domains: []string{"apple.com"}},
		},
	}

	detector, err := NewDetector(cfg)
	if err != nil {
		t.Fatalf("NewDetector() error = %v", err)
	}

	tests := []struct {
		name             string
		registeredDomain string
		wantBrand        *config.Brand
		wantErr          bool
	}{
		{
			name:             "matches brand",
			registeredDomain: "apple-login.com",
			wantBrand: &config.Brand{
				Name:    "Apple",
				Domains: []string{"apple.com"},
			},
			wantErr: false,
		},
		{
			name:             "does not match substring",
			registeredDomain: "pineapple.com",
			wantBrand:        nil,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brand, err := detector.Detect(tt.registeredDomain)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Detector.Detect() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(brand, tt.wantBrand) {
				t.Fatalf("Detector.Detect() got %#v, want %#v", brand, tt.wantBrand)
			}
		})
	}

}
