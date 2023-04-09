package scrape

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func ScrapeSearchProducts(text string) []Product {
	var listProducts []Product

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: scrapeme.live
		colly.AllowedDomains("scrapeme.live"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("li.product", func(e *colly.HTMLElement) {

		fmt.Println("li.product")

		product := Product{}
		product.Name = e.ChildText("h2")
		product.Price = e.ChildText(".price")

		listProducts = append(listProducts, product)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://scrapeme.live/shop/
	c.Visit("https://scrapeme.live/shop/")

	//---------------

	fmt.Println("==================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> END")

	return listProducts

}

//-------------------------------

func Scrape() {

	var pokemonProducts []PokemonProduct

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: scrapeme.live
		colly.AllowedDomains("scrapeme.live"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("li.product", func(e *colly.HTMLElement) {

		fmt.Println("li.product")

		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://scrapeme.live/shop/
	c.Visit("https://scrapeme.live/shop/")

	//---------------

	fmt.Println("==================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CVS")

	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, pokemonProduct := range pokemonProducts {
		// converting a PokemonProduct to an array of strings
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}
