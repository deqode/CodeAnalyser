package V_1_X

type Beego_V_1_X struct {
}

func (d *Beego_V_1_X) GetVersionName() (string, error) { return "v1.x", nil }
func (d *Beego_V_1_X) GetSemver() (string, error)      { return ">=1.x,<2.x", nil }
func (d *Beego_V_1_X) Detect(runtimeVersion, root string) (bool, error) {
	//TODO: write detection here
	return true, nil
}

func (f *Beego_V_1_X) IsFrameworkFound(runtimeVersion, root string) (bool, error) {
	return true, nil
}
func (f *Beego_V_1_X) IsFrameworkUsed(runtimeVersion, root string) (bool, error) {
	return true, nil
}
func (f *Beego_V_1_X) PercentOfFrameworkUsed(runtimeVersion, root string) (int64, error) {
	return 82, nil
}
func (f *Beego_V_1_X) GetFrameworkName() (string, error) {
	return "main", nil
}
