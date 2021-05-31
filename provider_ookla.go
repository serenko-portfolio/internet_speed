package internet_speed

import (
	"github.com/showwin/speedtest-go/speedtest"
)

// providerOokla implementation for Ookla provider
type providerOokla struct {
	uploadSpeed   float64
	downloadSpeed float64
}

// runTest tests your internet connection, returns error in case of any problems, uses 	"github.com/showwin/speedtest-go/speedtest" library to
// get download/upload speed in kilobytes per second
func (provider *providerOokla) runTest() error {
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})
	uploadSpeed := 0.0
	downloadSpeed := 0.0
	for _, s := range targets {
		err := s.PingTest()
		if err != nil {
			return err
		}
		err = s.DownloadTest(false)
		if err != nil {
			return err
		}
		err = s.UploadTest(false)
		if err != nil {
			return err
		}
		uploadSpeed += s.ULSpeed
		downloadSpeed += s.DLSpeed
	}
	provider.uploadSpeed = uploadSpeed / float64(targets.Len())
	provider.downloadSpeed = downloadSpeed / float64(targets.Len())
	return nil
}

// getUploadSpeed returns internet upload speed in Mbps
func (provider *providerOokla) getUploadSpeed() float64 {
	return provider.uploadSpeed / 1024.0
}

// getUploadSpeed returns internet upload speed in Mbps
func (provider *providerOokla) getDownloadSpeed() float64 {
	return provider.downloadSpeed / 1024.0
}
