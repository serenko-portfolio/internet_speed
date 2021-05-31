package internet_speed

import (
	"github.com/showwin/speedtest-go/speedtest"
)

type providerOokla struct {
	uploadSpeed   float64
	downloadSpeed float64
}

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

func (provider *providerOokla) getUploadData() float64 {
	return provider.uploadSpeed / 1024.0
}

func (provider *providerOokla) getDownloadData() float64 {
	return provider.downloadSpeed / 1024.0
}
