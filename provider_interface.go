package main

type ProviderInterface interface {
	runTest() error
	getUploadData() float64
	getDownloadData() float64
}
