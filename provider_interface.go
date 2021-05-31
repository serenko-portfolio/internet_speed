package internet_speed

type providerInterface interface {
	runTest() error
	getUploadData() float64
	getDownloadData() float64
}
