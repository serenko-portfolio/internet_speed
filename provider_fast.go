package main

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"log"
	"strconv"
	"time"
)

type ProviderFast struct {
	uploadSpeed   float64
	downloadSpeed float64
}

type Fast struct {
	Up       string
	Down     string
	UpUnit   string
	DownUnit string
}

func normalizeData(speed string, units string) float64 {
	parsedSpeed, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		return -1.0
	}
	switch units {
	case "Kbps":
		return parsedSpeed / 1024.0
	case "Gbps":
		return parsedSpeed * 1024.0
	default:
		return parsedSpeed
	}
}

func (provider *ProviderFast) runTest() error {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 180*time.Second)
	defer cancel()
	fast := new(Fast)
	actions := []chromedp.Action{
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &fast.Down, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#speed-units.succeeded`, &fast.DownUnit, chromedp.NodeVisible, chromedp.ByQuery),
	}
	actions = append(actions, chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#upload-value.succeeded`, &fast.Up, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#upload-units.succeeded`, &fast.UpUnit, chromedp.NodeVisible, chromedp.ByQuery),
	)
	err := chromedp.Run(ctx, actions...)
	if err != nil {
		return err
	}
	provider.uploadSpeed = normalizeData(fast.Up, fast.UpUnit)
	provider.downloadSpeed = normalizeData(fast.Down, fast.DownUnit)
	return nil
}

func (provider *ProviderFast) getUploadData() float64 {
	return provider.uploadSpeed
}

func (provider *ProviderFast) getDownloadData() float64 {
	return provider.downloadSpeed
}
