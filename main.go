package main

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

func main() {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--disable-gpu",
			"--window-size=1280,800",
			"--disable-dev-shm-usage",
			"--no-sandbox",
		})

	driver := agouti.ChromeDriver(options)
	defer driver.Stop()
	driver.Start()

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page.Navigate("https://www.google.com/")
	page.Screenshot("./screenshot.jpg")
	fmt.Println(page.Title())
}
