package detector

import (
	"fmt"
	"strings"

	"github.com/Useles5/sirphishalot/internal/config"
)

// Matcher builds a domain index once at startup for fast O(1) lookups.
type Matcher struct {
	brands  map[string]*config.Brand
	domains map[string]*config.Brand
}

// NewMatcher builds the lookup map once at startup.
func NewMatcher(cfg *config.Config) (*Matcher, error) {
	m := &Matcher{
		brands:  make(map[string]*config.Brand),
		domains: make(map[string]*config.Brand),
	}

	for i := range cfg.Brands {
		brand := &cfg.Brands[i]
		key := strings.ToLower(brand.Name)
		if _, exists := m.brands[key]; exists {
			return nil, fmt.Errorf("brand %q already exists", brand.Name)
		}
		m.brands[key] = brand

		for _, domain := range brand.Domains {
			if existing, exists := m.domains[domain]; exists {
				return nil, fmt.Errorf("domain %q already belongs to brand %q; duplicate entry in brand %q", domain, existing.Name, brand.Name)
			}
			m.domains[domain] = brand
		}

	}
	return m, nil
}

func (m *Matcher) MatchDomain(registeredDomain string) (*config.Brand, bool) {
	brand, ok := m.domains[registeredDomain]
	return brand, ok
}

func (m *Matcher) MatchBrand(nameOfBrand string) (*config.Brand, bool) {
	brand, ok := m.brands[nameOfBrand]
	return brand, ok
}
