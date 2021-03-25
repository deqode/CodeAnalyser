package V_2_X

type Gin_V_2_X struct {

}

func (receiver *Gin_V_2_X) GetVersionName() (string,error)  {
	return "v2.x",nil
}

func (receiver *Gin_V_2_X) GetSemver() (string,error)  {
	return ">=2.x,<3.x",nil
}

func (receiver *Gin_V_2_X) Detect(runtimeVersion,root string) (bool,error) {
	//TODO implement
	return true,nil
}

func (receiver *Gin_V_2_X) GetFrameworkName() (string,error) {
	return "gin",nil
}

func (receiver *Gin_V_2_X) IsFrameworkFound(runtimeVersion, root string) (bool,error) {
	return true,nil
}

func (receiver *Gin_V_2_X) IsFrameworkUsed(runtimeVersion, root string) (bool,error) {
	return true,nil
}

func (receiver *Gin_V_2_X) PercentOfFrameworkUsed(runtimeVersion, root string) (int64,error) {
	return 78,nil
}