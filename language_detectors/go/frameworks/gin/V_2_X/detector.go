package V_2_X

type Gin_V_2_X struct {

}

func (receiver *Gin_V_2_X) GetVersionName() string  {
	return "v2.x"
}

func (receiver *Gin_V_2_X) GetSemver() string  {
	return ">=2.x,<3.x"
}

func (receiver *Gin_V_2_X) Detect(runtimeVersion,root string) bool {
	//TODO implement
	return true
}

func (receiver *Gin_V_2_X) GetFrameworkName() string {
	return "gin"
}

func (receiver *Gin_V_2_X) IsFrameworkFound(runtimeVersion, root string) bool {
	return true
}

func (receiver *Gin_V_2_X) IsFrameworkUsed(runtimeVersion, root string) bool {
	return true
}

func (receiver *Gin_V_2_X) PercentOfFrameworkUsed(runtimeVersion, root string) float64 {
	return 78.9
}