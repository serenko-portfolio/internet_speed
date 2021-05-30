package main

import (
	test "test.go"
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
