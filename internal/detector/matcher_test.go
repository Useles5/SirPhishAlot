package detector

import (
	"testing"

	"github.com/Useles5/sirphishalot/internal/config"
)

func TestNewMatcher(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr bool
	}{
		{
			name: "Valid Config",
			cfg: &config.Config{
				Brands: []config.Brand{
					{
						Name:    "Amazon",
						Domains: []string{"amazon.com"},
					},
					{
						Name:    "My Company",
						Domains: []string{"mycompany.com", "mycompany.co.uk"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Duplicate Domain",
			cfg: &config.Config{
				Brands: []config.Brand{
					{
						Name:    "Amazon",
						Domains: []string{"amazon.com"},
					},
					{
						Name:    "Fake Amazon",
						Domains: []string{"amazon.com"},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher, err := NewMatcher(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewMatcher() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				if matcher != nil {
					t.Fatalf("NewMatcher() matcher = %v, want nil", matcher)
				}
				return
			}

			if matcher == nil {
				t.Fatal("NewMatcher() matcher = nil, want non-nil")
			}

			if brand, ok := matcher.Match("amazon.com"); !ok || brand.Name != "Amazon" {
				t.Fatalf(`Match("amazon.com") = (%v, %v), want (Amazon, true)`, brand, ok)
			}

			if brand, ok := matcher.Match("mycompany.com"); !ok || brand.Name != "My Company" {
				t.Fatalf(`Match("mycompany.com") = (%v, %v), want ("My Company", true)`, brand, ok)
			}

			if brand, ok := matcher.Match("mycompany.co.uk"); !ok || brand.Name != "My Company" {
				t.Fatalf(`Match("mycompany.co.uk") = (%v, %v), want ("My Company", true)`, brand, ok)
			}

			if brand, ok := matcher.Match("unknown.com"); ok || brand != nil {
				t.Fatalf(`Match("unknown.com") = (%v, %v), want (nil, false)`, brand, ok)
			}
		})
	}
}
