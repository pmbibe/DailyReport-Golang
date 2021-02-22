package main

import (
	"context"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/chromedp/chromedp"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	// url := `https://time.is/vi/`
	// fileName := "hello.pdf"
	// saveAsPdf(url, fileName)
	// catureScreen(newURL())
	url := "https://www.joesnewbalanceoutlet.com/men/shoes/under-45"
	log.Print(findTotalPage(url))

}
func saveAsPdf(url string, fileName string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	catureScreen(url)
	pdf.Image("dailyreport.png", 10, 10, 30, 0, false, "", 0, "")
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		panic(err)
	}
}

func newURL(url string, page int) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
	}
	q := req.URL.Query()
	q.Add("Filters[Size]", "9")
	q.Add("Filters[Size]", "9.5")
	q.Add("Filters[Size]", "10")
	q.Add("Page", int)
	q.Add("Sorting", "LowestPrice")
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}
func findTotalPage(url string) int {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
	}
	q := req.URL.Query()
	q.Add("Filters[Size]", "9")
	q.Add("Filters[Size]", "9.5")
	q.Add("Filters[Size]", "10")
	req.URL.RawQuery = q.Encode()
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// var value string
	var totalPage string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(req.URL.String()),
		chromedp.Evaluate(`document.querySelector('#Paging > div.pagingWrapper > span.pagingPages').innerText`, &totalPage),
	); err != nil {
		log.Fatal(err)
	}
	var re = regexp.MustCompile(`(?m)\d$`)
	totalPageStr := re.FindString(totalPage)
	totalPageInt, _ := strconv.Atoi(totalPageStr)
	return totalPageInt

}

// return 1

func catureScreen(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()
	var report []byte
	// var value string
	var total, totalPage string
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1134, 3469),
		chromedp.Navigate(url),
		chromedp.Click(`//a[@href='javascript:;']`),
		chromedp.WaitNotPresent(`#Modals`, chromedp.BySearch),
		chromedp.Evaluate(`document.querySelector('#Paging > div.pagingWrapper > span.pagingTotal').innerText`, &total),
		chromedp.CaptureScreenshot(&report), //Chụp màn hình #Paging > div.pagingWrapper > span.pagingPages
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./img/dailyreport.png", report, 0o644); err != nil {
		log.Fatal(err)
	}

	return totalPage

}
