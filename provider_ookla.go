package main

import (
	"github.com/showwin/speedtest-go/speedtest"
)

type ProviderOokla struct {
	uploadSpeed   float64
	downloadSpeed float64
}

func (provider *ProviderOokla) runTest() error {
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})
	uploadSpeed := 0.0
	downloadSpeed := 0.0
	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)
		uploadSpeed += s.ULSpeed
		downloadSpeed += s.DLSpeed
	}
	provider.uploadSpeed = uploadSpeed / float64(targets.Len())
	provider.downloadSpeed = downloadSpeed / float64(targets.Len())
	return nil
}

func (provider *ProviderOokla) getUploadData() float64 {
	return provider.uploadSpeed / 1024.0
}

func (provider *ProviderOokla) getDownloadData() float64 {
	return provider.downloadSpeed / 1024.0
}
