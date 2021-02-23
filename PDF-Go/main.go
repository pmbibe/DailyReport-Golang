package main

import (
	"context"
	"strconv"

	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/chromedp/chromedp"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	url := "https://www.joesnewbalanceoutlet.com/men/shoes/under-45"
	fileName := "hello.pdf"
	saveAsPdf(url, fileName)

}
func captureScreen(url string, fileName string) {
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()
	var report []byte
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1134, 3469),
		chromedp.Navigate(url),
		chromedp.Click(`//a[@href='javascript:;']`),
		chromedp.WaitNotPresent(`#Modals`, chromedp.BySearch),
		chromedp.CaptureScreenshot(&report), //Chụp màn hình #Paging > div.pagingWrapper > span.pagingPages
	); err != nil {
		log.Fatal(err)
	}
	fullPath := "./" + fileName + ".png"
	if err := ioutil.WriteFile(fullPath, report, 0o644); err != nil {
		log.Fatal(err)
	}

}
func saveAsPdf(url string, fileName string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	captureScreen(url, "dailyreport")
	pdf.Image("dailyreport.png", 10, 10, 30, 0, false, "", 0, "")
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		panic(err)
	}
}
