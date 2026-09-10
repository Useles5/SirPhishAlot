package normalizer

import "testing"

func TestRegisteredDomain(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "Subdomain",
			input:   "resident.uidai.gov.in",
			want:    "uidai.gov.in",
			wantErr: false,
		},
		{
			name:    "IN Domain",
			input:   "example.gov.in",
			want:    "example.gov.in",
			wantErr: false,
		},
		{
			name:    "GitHub Pages",
			input:   "foo.bar.github.io",
			want:    "bar.github.io",
			wantErr: false,
		},
		{
			name:    "Invalid domain",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Only TLD",
			input:   "com",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Deep subdomain",
			input:   "a.b.c.amazon.com",
			want:    "amazon.com",
			wantErr: false,
		},
		{
			name:    "Trailing dot",
			input:   "google.com.",
			want:    "google.com",
			wantErr: false,
		},
		{
			name:    "Mixed case",
			input:   "LoGiN.PayPal.Com",
			want:    "paypal.com",
			wantErr: false,
		},
		{
			name:    "Single label",
			input:   "localhost",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisteredDomain(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error = %v, got = %v", tt.wantErr, err)
			}

			if got != tt.want {
				t.Errorf("got = %q, want = %q", got, tt.want)
			}
		})
	}
}

func BenchmarkRegisteredDomain(b *testing.B) {
	input := "resident.uidai.gov.in."
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := RegisteredDomain(input)
		if err != nil {
			b.Fatal(err)
		}
	}
}
