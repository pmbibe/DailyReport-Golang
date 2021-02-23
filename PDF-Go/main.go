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
	// url := `https://time.is/vi/`
	// fileName := "hello.pdf"
	// saveAsPdf(url, fileName)
	// captureScreen(newURL())
	var totalPage int
	url := "https://www.joesnewbalanceoutlet.com/men/shoes/under-45"
	totalPage, _ = strconv.Atoi(findTotalPage(url))
	for i := 1; i <= totalPage; i++ {
		captureScreen(newURL(url, i), "Page"+strconv.Itoa(i))
		log.Print(newURL(url, i))
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
	q.Add("Page", strconv.Itoa(page))
	q.Add("Sorting", "LowestPrice")
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}
func findTotalPage(url string) string {
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
	var totalPage string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(req.URL.String()),
		chromedp.Evaluate(`document.querySelector('#Paging > div.pagingWrapper > span.pagingPages').innerText`, &totalPage),
	); err != nil {
		log.Fatal(err)
	}
	var re = regexp.MustCompile(`(?m)\d$`)
	totalPageStr := re.FindString(totalPage)
	return totalPageStr

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
	fullPath := "./img/" + fileName + ".png"
	if err := ioutil.WriteFile(fullPath, report, 0o644); err != nil {
		log.Fatal(err)
	}

}
func saveAsPdf(url string, fileName string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	captureScreen(url, "Page")
	pdf.Image("Page.png", 10, 10, 30, 0, false, "", 0, "")
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		panic(err)
	}
}
