package services

import (
	"bytes"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// GeneratePDF converts HTML content to PDF bytes
func GeneratePDF(html string) ([]byte, error) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// Set global options
	pdfg.Dpi.Set(96) // Default standard screen DPI, significantly reduces size compared to 300
	pdfg.ImageQuality.Set(75) // Reduce image quality slightly to save space
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	// Set margins (in millimeters)
	pdfg.MarginTop.Set(10)
	pdfg.MarginBottom.Set(10)
	pdfg.MarginLeft.Set(10)
	pdfg.MarginRight.Set(10)

	// Create page from HTML string
	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(html)))

	// Enable local file access for CSS/images if needed
	page.EnableLocalFileAccess.Set(true)

	// Add page to generator
	pdfg.AddPage(page)

	// Generate PDF
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	// Return PDF bytes
	return pdfg.Bytes(), nil
}
