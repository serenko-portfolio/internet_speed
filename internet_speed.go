package internet_speed

import (
	"errors"
)

func GetInternetSpeed(providerName string) (float64, float64, error) {
	var p providerInterface
	switch providerName {
	case "Ookla":
		p = &providerOokla{}
	case "fastStruct":
		p = &providerFast{}
	default:
		return 0.0, 0.0, errors.New("invalid argument error")
	}
	err := p.runTest()
	if err != nil {
		return 0, 0, err
	}
	dSpeed := p.getDownloadData()
	uSpeed := p.getUploadData()
	return dSpeed, uSpeed, nil
}
