package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
    "github.com/go-rod/rod/lib/proto"
    "github.com/go-rod/rod/lib/utils"
    "github.com/ysmood/gson"
)

const (
    inputFilePath  = "/pdf_public/index.html"
    outputFilePath = "./public/karl_haworth_resume.pdf"
)

func main() {
	path, err := os.Getwd()
    if err != nil {
        log.Fatalf("Failed to get current working directory: %v", err)
    }

	filepath := path + inputFilePath

	l, err := launcher.New().
		Set("disable-web-security").
        Set("disable-setuid-sandbox").
        Set("no-sandbox").
        Set("no-first-run", "true").
        Set("disable-gpu").
        Headless(true).
        Launch()
	if err != nil {
		log.Fatalf("Failed to launch browser: %v", err)
	}
	
	page := rod.New().ControlURL(l).MustConnect().MustPage("file://" + filepath).MustWaitLoad()

	pdf, _ := page.PDF(&proto.PagePrintToPDF{
		PaperWidth:      gson.Num(8.5),
		PaperHeight:     gson.Num(11),
		MarginTop:       gson.Num(0.25),
		MarginBottom:    gson.Num(0.25),
		MarginRight:     gson.Num(0.25),
		MarginLeft:      gson.Num(0.25),
		Scale:           gson.Num(0.95),
		PrintBackground: true,
	})
    if err != nil {
        log.Fatalf("Failed to generate PDF: %v", err)
    }

    if err := utils.OutputFile(outputFilePath, pdf); err != nil {
        log.Fatalf("Failed to save PDF: %v", err)
    }

    fmt.Printf("Wrote %s\n", outputFilePath)
}
