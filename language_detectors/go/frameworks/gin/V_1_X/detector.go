package V_1_X

type Gin_V_1_X struct {
	
}

func (receiver *Gin_V_1_X) GetVersionName() string  {
	return "v1.x"
}

func (receiver *Gin_V_1_X) GetSemver() string  {
	return ">=1.x,<2.x"
}

func (receiver *Gin_V_1_X) Detect(runtimeVersion,root string) bool {
	//TODO implement
	return true
}

func (receiver *Gin_V_1_X) GetFrameworkName() string {
	return "gin"
}

func (receiver *Gin_V_1_X) IsFrameworkFound(runtimeVersion, root string) bool {
	return true
}

func (receiver *Gin_V_1_X) IsFrameworkUsed(runtimeVersion, root string) bool {
	return true
}

func (receiver *Gin_V_1_X) PercentOfFrameworkUsed(runtimeVersion, root string) float64 {
	return 78.9
}