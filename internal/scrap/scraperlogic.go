package scrap

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func Scraping(url string) {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	products := make([]Product, 0)

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach("td", func(i int, h *colly.HTMLElement) {
			item := Product{}
			item.Code = h.Text
			item.Letter_code = h.Text
			item.Units = h.Text
			item.Currency = h.Text
			item.Currency_rate = h.Text
			products = append(products, item)
		})

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

	c.Visit(url)

}
