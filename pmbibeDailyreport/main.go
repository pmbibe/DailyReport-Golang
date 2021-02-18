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
	router := gin.Default()
	router.Static("/img", "./img")
	router.LoadHTMLGlob("./index.html")
	router.GET("/", getting)
	router.GET("/lastest", lastest)
	router.Run(":80")
	s := gocron.NewScheduler()
	s.Every(3).Minutes().Do(Caturescreen)
	<-s.Start()
}

func lastest(c *gin.Context) {
	Caturescreen()
	c.Redirect(http.StatusMovedPermanently, "/")

}
func getting(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}
func Caturescreen() {
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()
	var report []byte
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 2000),
		chromedp.Navigate(`https://time.is/vi/`),
		// chromedp.Navigate(`https://www.joesnewbalanceoutlet.com/`),
		// chromedp.Click(`//a[@href='javascript:;']`),
		// chromedp.WaitNotPresent(`#Modals`, chromedp.BySearch),
		chromedp.CaptureScreenshot(&report),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("./img/dailyreport.png", report, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("OK")
}
