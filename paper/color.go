package paper

import (
	"fmt"
	"log"
)

var (
	Colors = map[string]Color{
		"black":     Black,
		"darkgray":  DarkGray,
		"gray":      Gray,
		"lightgray": LightGray,
		"white":     White,
	}

	Black     = Color{0, 0, 0}
	DarkGray  = Color{85, 85, 85}
	Gray      = Color{128, 128, 128}
	LightGray = Color{204, 204, 204}
	White     = Color{255, 255, 255}
)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (c Color) RedInt() int {
	return int(c.Red)
}

func (c Color) GreenInt() int {
	return int(c.Green)
}

func (c Color) BlueInt() int {
	return int(c.Blue)
}

func ColorFromString(s string) (Color, bool) {
	c := Color{}

	_, err := fmt.Sscanf(s, "%02x%02x%02x", &c.Red, &c.Green, &c.Blue)
	if err != nil {
		log.Printf("[ERROR]: %s", err.Error())
		return c, false
	}

	return c, true
}
