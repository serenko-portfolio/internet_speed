package main

import (
	"github.com/serenko-portfolio/speed_test"
)

func main() {
	var p ProviderInterface = &ProviderFast{}
	p.RunTest()
	dSpeed, dUnits := p.GetDownloadData()
	uSpeed, uUnits := p.GetUploadData()
	print(dSpeed)
	print(dUnits)
	println("")
	print(uSpeed)
	print(uUnits)
}
