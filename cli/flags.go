package cli

import (
	"errors"
	"flag"
	"os"
	"strconv"

	"github.com/c-128/paper.gen/paper"
)

var format = paper.A4
var orientation = paper.Portrait
var backgroundColor = paper.White
var dotColor = paper.Gray
var dotDistance = 5.0
var dotRadius = 0.35

func Start() error {
	var outputFile = "output.pdf"
	var template string

	flagString("out", "PDF Output file", &outputFile)
	flagString("templ", "Paper template to use", &template)

	flagFormat("format", "Format of the paper", &format)
	flagOrientation("orientation", "Orientation of the paper", &orientation)
	flagColor("background-color", "Background color", &backgroundColor)
	flagColor("dot-color", "Dot color", &dotColor)
	flagFloat64("distance", "Distance between dots", &dotDistance)
	flagFloat64("radius", "Radius of a dot", &dotRadius)
	flag.Parse()

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	switch template {
	case "blank":
		err = generateBlank(file)
		if err != nil {
			return err
		}
	case "dotted":
		err = generateDotted(file)
		if err != nil {
			return err
		}
	default:
		return errors.New("unkown paper template")
	}

	return nil
}

func flagFormat(name string, useage string, out *paper.Format) {
	flag.Func(name, useage, func(value string) error {
		format, ok := paper.Formats[value]
		if !ok {
			return errors.New("unkown format")
		}

		*out = format
		return nil
	})
}

func flagOrientation(name string, useage string, out *paper.Orientation) {
	flag.Func(name, useage, func(value string) error {
		orientation, ok := paper.Orientations[value]
		if !ok {
			return errors.New("unkown orientation")
		}

		*out = orientation
		return nil
	})
}

func flagColor(name string, useage string, out *paper.Color) {
	flag.Func(name, useage, func(value string) error {
		color, ok := paper.Colors[value]
		if ok {
			*out = color
			return nil
		}

		color, ok = paper.ColorFromString(value)
		if ok {
			*out = color
			return nil
		}

		return errors.New("unkown color")
	})
}

func flagFloat64(name string, useage string, out *float64) {
	flag.Func(name, useage, func(value string) error {
		float, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}

		*out = float
		return nil
	})
}

func flagString(name string, useage string, out *string) {
	flag.Func(name, useage, func(value string) error {
		*out = value
		return nil
	})
}
