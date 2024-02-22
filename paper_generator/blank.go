package paper_generator

import (
	"github.com/c-128/paper.gen/paper"
	"github.com/go-pdf/fpdf"
)

func GenerateBlank(
	format paper.Format,
	orientation paper.Orientation,
) (*fpdf.Fpdf, error) {
	pdf := format.NewFPDF(orientation)
	pdf.AddPage()

	return pdf, nil
}
