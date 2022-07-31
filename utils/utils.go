package utils

import (
	"log"
	"net/url"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func IsValidUrl(urlWebsite string) bool {
	_, err := url.ParseRequestURI(urlWebsite)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GeneratePDF(url string) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	page := wkhtmltopdf.NewPage(url)
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./output.pdf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}
