package paper

import "github.com/go-pdf/fpdf"

var (
	Formats = map[string]Format{
		"a0": A0,
		"a1": A1,
		"a2": A2,
		"a3": A3,
		"a4": A4,
		"a5": A5,
		"a6": A6,
	}

	A0 = Format{841, 1189}
	A1 = Format{594, 841}
	A2 = Format{420, 594}
	A3 = Format{297, 420}
	A4 = Format{210, 297}
	A5 = Format{148, 210}
	A6 = Format{105, 148}
)

type Format struct {
	Short float64
	Long  float64
}

func (f Format) ToFPDFSizeType(orientation Orientation) fpdf.SizeType {
	if orientation.ShortSideUp {
		return fpdf.SizeType{
			Wd: f.Short,
			Ht: f.Long,
		}
	} else {
		return fpdf.SizeType{
			Wd: f.Long,
			Ht: f.Short,
		}
	}
}

func (f Format) NewFPDF(orientation Orientation) *fpdf.Fpdf {
	pdf := fpdf.NewCustom(&fpdf.InitType{
		OrientationStr: fpdf.OrientationPortrait,
		UnitStr:        fpdf.UnitMillimeter,
		SizeStr:        "custom",
		Size:           f.ToFPDFSizeType(orientation),
		FontDirStr:     ".",
	})

	return pdf
}
