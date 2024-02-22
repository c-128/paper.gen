package cli

import (
	"io"

	"github.com/c-128/paper.gen/paper_generator"
)

func generateBlank(outFile io.WriteCloser) error {
	pdf, err := paper_generator.GenerateBlank(
		format,
		orientation,
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
