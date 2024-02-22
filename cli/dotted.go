package cli

import (
	"io"

	"github.com/c-128/paper.gen/paper_generator"
)

func generateDotted(outFile io.WriteCloser) error {
	pdf, err := paper_generator.GenerateDotted(
		format,
		orientation,
		dotColor,
		dotDistance,
		dotRadius,
	)
	if err != nil {
		return err
	}

	err = pdf.OutputAndClose(outFile)
	if err != nil {
		return err
	}

	return nil
}
