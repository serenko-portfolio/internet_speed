package main

import "errors"

func CheckConnection(providerName string) (float64, float64, error) {
	var p ProviderInterface
	switch providerName {
	case "Ookla":
		p = &ProviderOokla{}
	case "Fast":
		p = &ProviderFast{}
	default:
		return 0.0, 0.0, errors.New("invalid argument error")
	}
	p.runTest()
	dSpeed := p.getDownloadData()
	uSpeed := p.getUploadData()
	return dSpeed, uSpeed, nil
}
