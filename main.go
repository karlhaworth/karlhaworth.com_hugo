// Command pdf is a chromedp example demonstrating how to capture a pdf of a
// page.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// get local file
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filepath := path + "/public/index.html"

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF("file://"+filepath, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("./public/karl_haworth-resume.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote karl_haworth-resume.pdf")
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {

			return emulation.SetDefaultBackgroundColorOverride().
				// green
				WithColor(&cdp.RGBA{R: 0, G: 255, B: 0, A: 0}).
				Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithMarginTop(.15).WithMarginBottom(.15).WithMarginRight(.15).WithMarginLeft(.15).WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
