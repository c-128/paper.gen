package paper

var (
	Orientations = map[string]Orientation{
		"portrait":  Portrait,
		"landscape": Landscape,
	}

	Portrait  = Orientation{true, false}
	Landscape = Orientation{false, true}
)

type Orientation struct {
	ShortSideUp bool
	LongSideUp  bool
}
