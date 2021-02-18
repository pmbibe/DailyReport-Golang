package capturescreen

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

//Caturescreen for taking Screenshot
func Caturescreen() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var report []byte
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 2000),
		chromedp.Navigate(`https://www.joesnewbalanceoutlet.com/`),
		chromedp.Click(`//a[@href='javascript:;']`),
		chromedp.WaitNotPresent(`#Modals`, chromedp.BySearch),
		chromedp.CaptureScreenshot(&report),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("../img/dailyreport.png", report, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("OK")
}
