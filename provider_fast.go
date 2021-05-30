package main

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type ProviderFast struct {
	uploadSpeed       string
	uploadSpeedUnit   string
	downloadSpeed     string
	downloadSpeedUnit string
}

type Fast struct {
	Up       string
	Down     string
	UpUnit   string
	DownUnit string
}

func (provider *ProviderFast) RunTest() error {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 180*time.Second)
	defer cancel()
	fast := new(Fast)
	cmds := []chromedp.Action{
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &fast.Down, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#speed-units.succeeded`, &fast.DownUnit, chromedp.NodeVisible, chromedp.ByQuery),
	}
	cmds = append(cmds, chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#upload-value.succeeded`, &fast.Up, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#upload-units.succeeded`, &fast.UpUnit, chromedp.NodeVisible, chromedp.ByQuery),
	)
	err := chromedp.Run(ctx, cmds...)
	if err != nil {
		return err
	}
	provider.uploadSpeedUnit = fast.UpUnit
	provider.uploadSpeed = fast.Up
	provider.downloadSpeedUnit = fast.DownUnit
	provider.downloadSpeed = fast.Down
	return nil
}

func (provider *ProviderFast) GetUploadData() (string, string) {
	return provider.uploadSpeed, provider.uploadSpeedUnit
}

func (provider *ProviderFast) GetDownloadData() (string, string) {
	return provider.downloadSpeed, provider.downloadSpeedUnit
}
