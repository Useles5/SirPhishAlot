package detector

import (
	"strings"

	"github.com/Useles5/sirphishalot/internal/config"
)

func keywordMatch(registeredDomain string, matcher *Matcher) (*config.Brand, bool) {

	// remove tld
	base, _, _ := strings.Cut(registeredDomain, ".")

	tokens := strings.FieldsFunc(base, func(r rune) bool { return r == '-' })
	for _, token := range tokens {
		brand, found := matcher.MatchBrand(token)
		if found {
			return brand, found
		}
	}
	return nil, false

}
