package parser

import "bytes"

// DomainScanner tracks current cursor position while scanning the payload.
type DomainScanner struct {
	data []byte
	pos  int
}

// NewDomainScanner initializes the scanner starting at the beginning of the paylaod/data.
func NewDomainScanner(data []byte) DomainScanner {
	return DomainScanner{data: data, pos: 0}
}

// Next finds and returns the next quoted domain.
func (s *DomainScanner) Next() (domain []byte, hasNext bool) {
	// find opening quote
	first := bytes.IndexByte(s.data[s.pos:], '"')
	if first == -1 {
		return nil, false
	}

	start := s.pos + first + 1

	// find closing quote
	second := bytes.IndexByte(s.data[start:], '"')
	if second == -1 {
		return nil, false
	}

	end := start + second

	// move cursor past the closing quote
	s.pos = end + 1
	return s.data[start:end], true
}
