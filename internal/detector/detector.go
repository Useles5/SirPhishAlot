package detector

import (
	"github.com/Useles5/sirphishalot/internal/config"
	"github.com/Useles5/sirphishalot/internal/normalizer"
)

// Detector checks domains against configured brands.
type Detector struct {
	matcher *Matcher
}

// NewDetector creates a detector from the config.
func NewDetector(cfg *config.Config) (*Detector, error) {
	matcher, err := NewMatcher(cfg)
	if err != nil {
		return nil, err
	}
	return &Detector{
		matcher: matcher,
	}, nil
}

// Detect normalizes a hostname and returns its matching configured brand.
func (d *Detector) Detect(hostname string) (*config.Brand, error) {
	domain, err := normalizer.RegisteredDomain(hostname)
	if err != nil {
		return nil, err
	}

	brand, _ := d.matcher.Match(domain)
	// if ok -> false, it means the domain is not cof
	//if !ok {
	//	return nil, nil
	//}

	return brand, nil
}
