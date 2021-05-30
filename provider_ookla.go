package main

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

type ProviderOokla struct {
	uploadSpeed   float64
	downloadSpeed float64
}

func (provider *ProviderOokla) RunTest() error {
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

func (provider *ProviderOokla) GetUploadData() (string, string) {
	return fmt.Sprintf("%v", provider.uploadSpeed), "Kbps"
}

func (provider *ProviderOokla) GetDownloadData() (string, string) {
	return fmt.Sprintf("%v", provider.downloadSpeed), "Kbps"
}