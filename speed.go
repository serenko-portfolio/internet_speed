package main

import "errors"

type LoadStruct struct {
	loadType string
	speed    string
	units    string
}

func checkConnection(providerName string) (LoadStruct, LoadStruct, error) {
	var p ProviderInterface
	switch providerName {
	case "Ookla":
		p = &ProviderOokla{}
	case "Fast":
		p = &ProviderFast{}
	default:
		return LoadStruct{}, LoadStruct{}, errors.New("invalid argument error")
	}
	p.RunTest()
	dSpeed, dUnits := p.GetDownloadData()
	uSpeed, uUnits := p.GetUploadData()
	return LoadStruct{"down", dSpeed, dUnits},
		LoadStruct{"up", uSpeed, uUnits}, nil
}
