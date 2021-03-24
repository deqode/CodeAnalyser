package V_1_X

type Beego_V_1_X struct {
}

func (d *Beego_V_1_X) GetVersionName() string { return "v1.x" }
func (d *Beego_V_1_X) GetSemver() string      { return ">=1.x,<2.x" }
func (d *Beego_V_1_X) Detect(runtimeVersion, root string) bool {
	//TODO: write detection here
	return true
}

func (f *Beego_V_1_X) IsFrameworkFound(runtimeVersion, root string) bool {
	return true
}
func (f *Beego_V_1_X) IsFrameworkUsed(runtimeVersion, root string) bool {
	return true
}
func (f *Beego_V_1_X) PercentOfFrameworkUsed(runtimeVersion, root string) float64 {
	return 82.5
}
func (f *Beego_V_1_X) GetFrameworkName() string {
	return "beego"
}
