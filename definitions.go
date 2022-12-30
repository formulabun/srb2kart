package srb2kart

type Srb2kart struct {
	Name       string
	Port       int
	AddonGroup string
}

func DefaultSrb2kart() Srb2kart {
	return Srb2kart{
		"srb2kart",
		5029,
		"addons",
	}
}
