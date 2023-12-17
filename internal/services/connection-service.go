package services

import "github.com/gocolly/colly"

const Link string = "https://azoramoon.com"

func Connection() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("azoramoon.com", "https://azoramoon.com"),
	)

	return c
}
