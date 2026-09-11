package detector

import (
	"fmt"

	"github.com/Useles5/sirphishalot/internal/config"
)

// Matcher builds a domain index once at startup for fast O(1) lookups.
type Matcher struct {
	domains map[string]*config.Brand
}

// NewMatcher builds the lookup map once at startup.
func NewMatcher(cfg *config.Config) (*Matcher, error) {
	m := &Matcher{
		domains: make(map[string]*config.Brand),
	}

	for i := range cfg.Brands {
		brand := &cfg.Brands[i]

		for _, domain := range brand.Domains {
			existing, exists := m.domains[domain]
			if exists {
				return nil, fmt.Errorf("domain %q already belongs to brand %q; duplicate entry in brand %q", domain, existing.Name, brand.Name)
			}
			m.domains[domain] = brand
		}

	}
	return m, nil
}

func (m *Matcher) Match(registeredDomain string) (*config.Brand, bool) {
	brand, ok := m.domains[registeredDomain]
	return brand, ok
}
