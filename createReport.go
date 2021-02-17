// Command emulate is a chromedp example demonstrating how to emulate a
// specific device such as an iPhone.
package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run
	var b2 []byte
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 2000),

		chromedp.Navigate(`https://www.joesnewbalanceoutlet.com/dailydeal`),
		chromedp.WaitVisible(`#Modals`),
		chromedp.Click(`//*[@id="Modals"]/div/a`, chromedp.NodeVisible),
		chromedp.CaptureScreenshot(&b2),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("screenshot2.png", b2, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("OK")
}
