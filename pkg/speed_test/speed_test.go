package speed_test_test

import (
	"../../internal/providers"
	"errors"
)

func CheckConnection(providerName string) (float64, float64, error) {
	var p providers.ProviderInterface
	switch providerName {
	case "Ookla":
		p = &providers.ProviderOokla{}
	case "Fast":
		p = &providers.ProviderFast{}
	default:
		return 0.0, 0.0, errors.New("invalid argument error")
	}
	p.RunTest()
	dSpeed := p.GetDownloadData()
	uSpeed := p.GetUploadData()
	return dSpeed, uSpeed, nil
}
