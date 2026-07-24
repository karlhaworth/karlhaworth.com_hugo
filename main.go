package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
	"golang.org/x/net/html"
)

const (
	inputFilePath  = "/pdf_public/index.html"
	outputFilePath = "./public/karl_haworth_resume.pdf"
	pageWidthInch  = 8.5
	pageHeightInch = 11.0
	marginInch     = 0.25
)

type PDFRenderer struct {
	pdf      *gofpdf.Fpdf
	lineH    float64
	fontSize float64
	isHidden bool
}

// inchesToMM converts inches to millimeters (gofpdf uses mm)
func inchesToMM(inches float64) float64 {
	return inches * 25.4
}

func NewPDFRenderer() *PDFRenderer {
	// Page dimensions in mm
	pageW := inchesToMM(pageWidthInch)
	pageH := inchesToMM(pageHeightInch)
	marginMM := inchesToMM(marginInch)

	pdf := gofpdf.New("P", "mm", "", gofpdf.SizeType{Wd: pageW, Ht: pageH})
	pdf.AddPage()
	pdf.SetMargins(marginMM, marginMM, marginMM)
	pdf.SetAutoPageBreak(true, marginMM)
	pdf.SetFont("Helvetica", "", 11)

	return &PDFRenderer{
		pdf:      pdf,
		lineH:    5.0,
		fontSize: 11,
	}
}

// getClassAttribute extracts class attribute from HTML node
func getClassAttribute(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			return attr.Val
		}
	}
	return ""
}

// shouldHide checks if element should be hidden based on Tailwind classes
func shouldHide(classes string) bool {
	// Check for print:hidden, hidden, or display:none
	return strings.Contains(classes, "hidden") || strings.Contains(classes, "print:hidden")
}

// walkNode recursively walks the HTML DOM and renders to PDF
func (pr *PDFRenderer) walkNode(n *html.Node) {
	if n == nil {
		return
	}

	switch n.Type {
	case html.TextNode:
		if !pr.isHidden {
			text := strings.TrimSpace(n.Data)
			if text != "" {
				pr.pdf.Write(pr.lineH, text)
			}
		}

	case html.ElementNode:
		classes := getClassAttribute(n)

		// Check if element should be hidden
		wasHidden := pr.isHidden
		if shouldHide(classes) {
			pr.isHidden = true
		}

		if !pr.isHidden {
			switch n.Data {
			case "body", "html":
				pr.renderNodeChildren(n)

			case "head", "style", "script":
				// Skip these nodes entirely
				return

			case "h1":
				pr.pdf.SetFont("Helvetica", "B", 24)
				pr.pdf.Ln(3)
				pr.renderNodeText(n)
				pr.pdf.Ln(5)
				pr.pdf.SetFont("Helvetica", "", 11)

			case "h2":
				pr.pdf.SetFont("Helvetica", "B", 16)
				pr.pdf.Ln(2)
				pr.renderNodeText(n)
				pr.pdf.Ln(3)
				pr.pdf.SetFont("Helvetica", "", 11)

			case "h3":
				pr.pdf.SetFont("Helvetica", "B", 13)
				pr.pdf.Ln(1)
				pr.renderNodeText(n)
				pr.pdf.Ln(2)
				pr.pdf.SetFont("Helvetica", "", 11)

			case "h4":
				pr.pdf.SetFont("Helvetica", "B", 12)
				pr.pdf.Ln(1)
				pr.renderNodeText(n)
				pr.pdf.Ln(1)
				pr.pdf.SetFont("Helvetica", "", 11)

			case "p", "div", "section", "article":
				pr.renderNodeChildren(n)
				pr.pdf.Ln(2)

			case "span":
				pr.renderNodeChildren(n)

			case "strong", "b":
				oldSize := pr.fontSize
				pr.pdf.SetFont("Helvetica", "B", pr.fontSize)
				pr.renderNodeChildren(n)
				pr.pdf.SetFont("Helvetica", "", oldSize)

			case "em", "i":
				oldSize := pr.fontSize
				pr.pdf.SetFont("Helvetica", "I", pr.fontSize)
				pr.renderNodeChildren(n)
				pr.pdf.SetFont("Helvetica", "", oldSize)

			case "ul", "ol":
				pr.pdf.Ln(2)
				pr.renderList(n, n.Data == "ol")
				pr.pdf.Ln(2)

			case "li":
				pr.renderNodeChildren(n)

			case "br":
				pr.pdf.Ln(pr.lineH)

			case "hr":
				pr.pdf.Ln(2)
				pr.pdf.SetLineWidth(0.5)
				x, y := pr.pdf.GetXY()
				pr.pdf.Line(10, y, pr.pdf.GetPageWidth()-10, y)
				pr.pdf.Ln(3)

			case "a":
				// Extract href and render text
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						pr.pdf.SetFont("Helvetica", "U", pr.fontSize)
						text := extractNodeText(n)
						if text != "" {
							pr.pdf.WriteLinkString(pr.lineH, strings.TrimSpace(text), attr.Val)
						}
						pr.pdf.SetFont("Helvetica", "", pr.fontSize)
						return
					}
				}
				pr.renderNodeChildren(n)

			default:
				// Recursively handle unknown elements
				pr.renderNodeChildren(n)
			}
		} else {
			// Element is hidden, don't render but still traverse for non-hidden children
			pr.renderNodeChildren(n)
		}

		// Restore hidden state
		pr.isHidden = wasHidden
	}
}

// renderNodeChildren recursively renders child nodes
func (pr *PDFRenderer) renderNodeChildren(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pr.walkNode(c)
	}
}

// renderNodeText extracts and renders all text from a node and children
func (pr *PDFRenderer) renderNodeText(n *html.Node) {
	text := extractNodeText(n)
	text = strings.TrimSpace(text)
	if text != "" {
		pr.pdf.Multi(pr.lineH, text)
	}
}

// renderList handles ul/ol lists
func (pr *PDFRenderer) renderList(n *html.Node, isOrdered bool) {
	index := 1
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "li" {
			if isOrdered {
				pr.pdf.Write(pr.lineH, fmt.Sprintf("%d. ", index))
				index++
			} else {
				pr.pdf.Write(pr.lineH, "• ")
			}

			// Render list item content
			for child := c.FirstChild; child != nil; child = child.NextSibling {
				pr.walkNode(child)
			}
			pr.pdf.Ln(pr.lineH)
		}
	}
}

// extractNodeText recursively extracts all text from a node
func extractNodeText(n *html.Node) string {
	var buf bytes.Buffer
	walk := func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
	}
	forEachNode(n, walk)
	return buf.String()
}

// forEachNode visits every node in the tree rooted at n
func forEachNode(n *html.Node, f func(*html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f)
	}
}

// Render parses HTML and generates PDF
func (pr *PDFRenderer) Render(htmlPath string) error {
	file, err := os.Open(htmlPath)
	if err != nil {
		return fmt.Errorf("failed to open HTML: %w", err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Find body and render
	var body *html.Node
	var findBody func(*html.Node)
	findBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			body = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findBody(c)
		}
	}
	findBody(doc)

	if body == nil {
		return fmt.Errorf("body element not found in HTML")
	}

	pr.walkNode(body)
	return nil
}

func (pr *PDFRenderer) Save(outputPath string) error {
	return pr.pdf.OutputFileAndClose(outputPath)
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	filepath := path + inputFilePath
	renderer := NewPDFRenderer()

	if err := renderer.Render(filepath); err != nil {
		log.Fatalf("Failed to render PDF: %v", err)
	}

	if err := renderer.Save(outputFilePath); err != nil {
		log.Fatalf("Failed to save PDF: %v", err)
	}

	fmt.Printf("Wrote %s\n", outputFilePath)
}
