package normalizer

import (
	"fmt"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func RegisteredDomain(domain string) (string, error) {
	domain = strings.ToLower(domain)         // normalize domain name , eg. "PayPal" -> "paypal"
	domain = strings.TrimSuffix(domain, ".") // remove suffix ".", eg. "google.com."

	registeredDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return "", fmt.Errorf("registered domain: %w", err)
	}
	return registeredDomain, nil
}
