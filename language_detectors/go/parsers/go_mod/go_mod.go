package go_mod

type GoMod struct {
	ModuleName string
	GoVersion  string
	Required   []*Package
}

type Package struct {
	Name     string
	Version  string
	Indirect bool
}

func ParseGoMod(modFile string) (*GoMod, error) {
	return &GoMod{
		Required: []*Package{
			{"beego", "1.6", false},
		},
	}, nil
}
