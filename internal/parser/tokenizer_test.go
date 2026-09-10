package parser

import (
	"bytes"
	"testing"
)

func TestDomainScanner(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		wantDomains [][]byte
	}{
		{
			name:        "Single Domain",
			input:       []byte(`"apple.com"`),
			wantDomains: [][]byte{[]byte("apple.com")},
		},
		{
			name:  "Multiple Domains",
			input: []byte(`"apple.com","google.com","amazon.com"`),
			wantDomains: [][]byte{
				[]byte("apple.com"),
				[]byte("google.com"),
				[]byte("amazon.com"),
			},
		},
		{
			name:  "Multiple Domains with random spaces",
			input: []byte(`  "apple.com",    "google.com",      "amazon.com"`),
			wantDomains: [][]byte{
				[]byte("apple.com"),
				[]byte("google.com"),
				[]byte("amazon.com"),
			},
		},
		{
			name:  "Empty Domain",
			input: []byte(`"","google.com"`),
			wantDomains: [][]byte{
				[]byte(""),
				[]byte("google.com"),
			},
		},
		{
			name:        "Empty input",
			input:       []byte(""),
			wantDomains: nil,
		},
		{
			name:        "Missing closing quote",
			input:       []byte(`"apple.com`),
			wantDomains: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewDomainScanner(tt.input)
			var gotDomains [][]byte

			for {
				domain, ok := scanner.Next()
				if !ok {
					break
				}
				//append([]byte(nil), domain...) copies domain and appends to gotDomains
				gotDomains = append(gotDomains, append([]byte(nil), domain...))
			}

			if len(gotDomains) != len(tt.wantDomains) {
				t.Fatalf("got %d domains, want %d", len(gotDomains), len(tt.wantDomains))
			}

			for i, domain := range gotDomains {
				if !bytes.Equal(domain, tt.wantDomains[i]) {
					t.Errorf("got domain %q, want %q", domain, tt.wantDomains[i])
				}
			}

		})
	}
}

func BenchmarkDomainScanner(b *testing.B) {
	input := []byte(`"apple.com","google.com","paypal.com"`)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scanner := NewDomainScanner(input)
		for {
			_, ok := scanner.Next()
			if !ok {
				break
			}
		}
	}
}
