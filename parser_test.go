package main

import (
	"slices"
	"testing"
)

var mockPayload = []byte(`{"data":{"leaf_cert":{"all_domains":["apple.com", "www.apple.com"]}}}`)

func TestExtractDomains(t *testing.T) {
	tests := []struct {
		name        string
		payload     []byte
		wantDomains []string
		wantErr     bool
	}{
		{
			name:        "Valid Payload",
			payload:     mockPayload,
			wantDomains: []string{"apple.com", "www.apple.com"},
			wantErr:     false,
		},

		{
			name:        "Missing Domains",
			payload:     []byte(`{"data":{"leaf_cert":{}}}`),
			wantDomains: nil,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractDomains(tt.payload)

			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractDomains() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !slices.Equal(got, tt.wantDomains) {
				t.Errorf("ExtractDomains() = %v, want %v", got, tt.wantDomains)
			}
		})
	}

}

func BenchmarkExtractDomains(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = extractDomains(mockPayload)
	}
}
