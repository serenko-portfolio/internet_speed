package providers

type ProviderInterface interface {
	RunTest() error
	GetUploadData() float64
	GetDownloadData() float64
}
