package paper_generator

import (
	"github.com/c-128/paper.gen/paper"
	"github.com/go-pdf/fpdf"
)

func GenerateBlank(
	format paper.Format,
	orientation paper.Orientation,
	backgroundColor paper.Color,
) (*fpdf.Fpdf, error) {
	pdf := format.NewFPDF(orientation)
	pdf.AddPage()

	pdf.SetFillColor(
		backgroundColor.RedInt(),
		backgroundColor.GreenInt(),
		backgroundColor.BlueInt(),
	)

	width, height := pdf.GetPageSize()

	pdf.Rect(0, 0, width, height, "F")

	return pdf, nil
}
