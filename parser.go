package main

import "encoding/json"

type CertStreamMessage struct {
	Data CertData `json:"data"`
}

type CertData struct {
	LeafCert LeafCert `json:"leaf_cert"`
}

type LeafCert struct {
	AllDomains []string `json:"all_domains"`
}

func extractDomains(payload []byte) ([]string, error) {
	var msg CertStreamMessage

	err := json.Unmarshal(payload, &msg)
	if err != nil {
		return nil, err

	}

	return msg.Data.LeafCert.AllDomains, nil
}
