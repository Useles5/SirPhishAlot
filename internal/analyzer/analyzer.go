package analyzer

import (
	"github.com/Useles5/sirphishalot/internal/config"
	"github.com/Useles5/sirphishalot/internal/detector"
)

type Analyzer struct {
	detector *detector.Detector
}

func NewAnalyzer(cfgPath string) (*Analyzer, error) {
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return nil, err
	}

	dtc, err := detector.NewDetector(&cfg)
	if err != nil {
		return nil, err
	}

	return &Analyzer{
		detector: dtc,
	}, nil
}

func (a *Analyzer) Analyze(hostname string) (*config.Brand, error) {
	return a.detector.Detect(hostname)
}
