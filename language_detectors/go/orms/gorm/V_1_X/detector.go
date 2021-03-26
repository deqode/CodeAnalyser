package V_1_X

type GORM_V_1_X struct {

}

func (reciever *GORM_V_1_X) GetVersionName() (string, error) {
	return "v1.x", nil
}

func (reciever *GORM_V_1_X) GetSemver() (string, error) {
	return ">=1.x,<2.x",nil
}

func (reciever *GORM_V_1_X) Detect(runtimeVersion, root string) (bool, error) {
	return true, nil
}

func (reciever *GORM_V_1_X) IsORMUsed(runtimeVersion, root string) (bool, error) {
	return true, nil
}

func (reciever *GORM_V_1_X) IsORMFound(runtimeVersion, root string) (bool, error) {
	return true, nil
}

func (reciever *GORM_V_1_X) GetORMName() (string, error) {
	return "gorm", nil
}

func (reciever *GORM_V_1_X) PercentOfORMUsed(runtimeVersion, root string) (float64, error) {
	return 0, nil
}
