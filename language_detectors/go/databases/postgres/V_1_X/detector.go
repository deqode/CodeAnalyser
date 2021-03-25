package V_1_X

type Postgres_V_1_X struct {

}

func (receiver *Postgres_V_1_X) GetVersionName() (string,error)  {
	return "v1.x",nil
}

func (receiver *Postgres_V_1_X) GetSemver() (string,error) {
	return ">=1.x,<2.x",nil
}

func (receiver *Postgres_V_1_X) IsDbUsed(runtimeVersion ,root string) (bool,error)  {
	return true,nil
}

func (receiver *Postgres_V_1_X) IsDbFound(runtimeVersion ,root string) (bool,error) {
	return true,nil
}

func (receiver *Postgres_V_1_X ) GetDbName() (string,error)  {
	return "postgres",nil
}

func (receiver *Postgres_V_1_X) PercentOfDbUsed(runtimeVersion ,root string) (float64,error)  {
	return 90.67,nil
}

// it returns port as well
func (receiver *Postgres_V_1_X) Detect(runtimeVersion, root string) (bool, int64,error) {
  return  true,8765,nil
}
