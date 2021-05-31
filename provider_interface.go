package internet_speed

// providerInterface used to trigger several providers similar way
type providerInterface interface {
	runTest() error
	getUploadSpeed() float64
	getDownloadSpeed() float64
}
