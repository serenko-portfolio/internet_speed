package main

type ProviderInterface interface {
	RunTest() error
	GetUploadData() (string, string)
	GetDownloadData() (string, string)
}
