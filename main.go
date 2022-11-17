// Package main ...
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	// get local file
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filepath := path + "/pdf_public/index.html"

	outfilePath := "./public/karl_haworth_resume.pdf"

	page := rod.New().MustConnect().MustPage("file://" + filepath).MustWaitLoad()
	page.MustPDF(outfilePath)
	// customized version
	pdf, _ := page.PDF(&proto.PagePrintToPDF{
		PaperWidth:   gson.Num(8.5),
		PaperHeight:  gson.Num(11),
		MarginTop:    gson.Num(0.15),
		MarginBottom: gson.Num(0.15),
		MarginRight:  gson.Num(0.15),
		MarginLeft:   gson.Num(0.15),
	})
	_ = utils.OutputFile(outfilePath, pdf)
	fmt.Println("wrote " + outfilePath)
}
