package main

import (
	"bytes"
	"testing"
)

var mockPayload = []byte(`{"data":{"leaf_cert":{"all_domains":["apple.com", "www.apple.com"]}}}`)

func TestExtractDomains(t *testing.T) {
	tests := []struct {
		name        string
		payload     []byte
		wantDomains []byte
		wantErr     bool
	}{
		{
			name:        "Valid Payload",
			payload:     mockPayload,
			wantDomains: []byte(`"apple.com", "www.apple.com"`),
			wantErr:     false,
		},

		{
			name:        "Missing Domains",
			payload:     []byte(`{"data":{"leaf_cert":{}}}`),
			wantDomains: nil,
			wantErr:     false,
		},

		{
			name:        "Malformed Payload - Missing colon",
			payload:     []byte(`{"data":{"leaf_cert":{"all_domains" ["apple.com", "www.apple.com"]}}}`),
			wantDomains: nil,
			wantErr:     true,
		},

		{
			name:        "Malformed Payload - Cut off Payload",
			payload:     []byte(`{"data":{"leaf_cert":{"all_domains":[`),
			wantDomains: nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractDomains(tt.payload)

			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error = %v, got %v", tt.wantErr, err)
			}
			if !bytes.Equal(got, tt.wantDomains) {
				t.Errorf("Expected domains = %v, got %v", string(tt.wantDomains), string(got))
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
