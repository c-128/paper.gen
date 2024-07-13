package paper_generator

import (
	"github.com/c-128/paper.gen/paper"
	"github.com/go-pdf/fpdf"
)

func GenerateDotted(
	format paper.Format,
	orientation paper.Orientation,
	backgroundColor paper.Color,
	dotColor paper.Color,
	dotDistance float64,
	dotRadius float64,
) (*fpdf.Fpdf, error) {
	pdf := format.NewFPDF(orientation)

	pdf.AddPage()

	width, height := pdf.GetPageSize()

	pdf.SetFillColor(
		backgroundColor.RedInt(),
		backgroundColor.GreenInt(),
		backgroundColor.BlueInt(),
	)

	pdf.Rect(0, 0, width, height, "F")

	pdf.SetFillColor(
		dotColor.RedInt(),
		dotColor.GreenInt(),
		dotColor.BlueInt(),
	)

	for x := float64(0); x < width; x += dotDistance {
		for y := float64(0); y < height; y += dotDistance {
			pdf.Circle(
				x,
				y,
				dotRadius,
				"F",
			)
		}
	}

	return pdf, nil
}
