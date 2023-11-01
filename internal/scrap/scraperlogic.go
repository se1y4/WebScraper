package scrap

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func Scraping() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	products := make([]Product, 0)

	c.OnHTML(".grid-row products-list .col col-xxl-4 col-sm-6 col-xs-12", func(e *colly.HTMLElement) {

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(products, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("products.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}

	})

	c.Visit("https://novosib.24ff.ru/catalog/bady")

}
