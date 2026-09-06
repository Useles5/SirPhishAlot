package main

import (
	"bytes"
	"errors"
)

// exact match for "all_domains" key in JSON from where we need to extract domains
var domainKey = []byte(`"all_domains"`)

// identify whitespace
func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}
func extractDomains(payload []byte) ([]byte, error) {
	// find the start -> `"`
	startIdx := bytes.Index(payload, domainKey)
	if startIdx == -1 {
		return nil, nil
	}

	// cursor starts after `"all_domains"`
	cursor := startIdx + len(domainKey)

	// skip any whitespaces before ':'
	for cursor < len(payload) && isWhitespace(payload[cursor]) {
		cursor++
	}

	// check if colon exists
	if cursor >= len(payload) || payload[cursor] != ':' {
		return nil, errors.New("invalid/malformed payload: missing colon(:)")
	}

	// move forward ':'
	cursor++

	// skip nay whitespace before '['
	for cursor < len(payload) && isWhitespace(payload[cursor]) {
		cursor++
	}

	// check if '[' exists
	if cursor >= len(payload) || payload[cursor] != '[' {
		return nil, errors.New("invalid/malformed payload: missing opening bracket([)")
	}
	// move forward '['
	cursor++

	// find the closing ']'
	offset := bytes.IndexByte(payload[cursor:], ']')
	if offset == -1 {
		return nil, errors.New("invalid/malformed payload: missing closing bracket(])")
	}

	// return
	return payload[cursor : cursor+offset], nil
}
