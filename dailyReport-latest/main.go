package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"regexp"
	"strconv"

	"github.com/chromedp/chromedp"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	var totalPage int
	url := "https://www.joesnewbalanceoutlet.com/men/shoes/under-45"
	totalPage, _ = strconv.Atoi(findTotalPage(url))
	for i := 1; i <= totalPage; i++ {
		fullPath := user.HomeDir + `\Pictures\` + "Page" + strconv.Itoa(i) + ".png"
		if err := ioutil.WriteFile(fullPath, captureScreen(newURL(url, i)), 0o644); err != nil {
			log.Fatal(err)
		}
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

func captureScreen(url string) []byte {
	pageWidth, pageHeight := func() (float64, float64) {
		var pageWidth, pageHeight float64
		log.Print(url)
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()
		if err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.Click(`//a[@href='javascript:;']`),
			chromedp.WaitNotPresent(`#Modals`, chromedp.BySearch),
			chromedp.Evaluate(`document.body.offsetWidth`, &pageWidth),
			chromedp.Evaluate(`document.body.offsetHeight`, &pageHeight),
		); err != nil {
			log.Fatal(err)
		}
		return pageWidth, pageHeight
	}()
	log.Print(pageWidth)
	log.Print(pageHeight)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var report []byte
	const par float64 = 0.7092619096299325

	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(int64(pageWidth/par), int64(pageHeight*par)),
		chromedp.Navigate(url),
		chromedp.Click(`//a[@href='javascript:;']`),           //Đóng bảng
		chromedp.WaitNotPresent(`#Modals`, chromedp.BySearch), //Đợi đến lúc mất bảng hiện lên
		chromedp.CaptureScreenshot(&report),                   //Chụp màn hình
	); err != nil {
		log.Fatal(err)
	}
	return report
}
