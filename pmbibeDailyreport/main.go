package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

func main() {
	webDisplay()
}

func webDisplay() {
	go task()
	router := gin.Default()
	router.Static("/img", "./img")
	router.LoadHTMLGlob("./index.html")
	router.GET("/", getting)
	router.Run(":80")

}
func task() {
	s := gocron.NewScheduler()
	s.Every(15).Seconds().Do(catureScreen)
	<-s.Start()
}
func getting(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}
func catureScreen() {
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()
	var report []byte
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 2000),
		chromedp.Navigate(`https://time.is/vi/`),
// 		chromedp.Navigate(`https://www.joesnewbalanceoutlet.com/dailydeal`),
// 		chromedp.Click(`//a[@href='javascript:;']`),
// 		chromedp.WaitNotPresent(`#Modals`, chromedp.BySearch),
		chromedp.CaptureScreenshot(&report),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./img/dailyreport.png", report, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Updated")
}
