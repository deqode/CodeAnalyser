package V_1_X

type Postgres_V_1_X struct {

}

func (receiver *Postgres_V_1_X) GetVersionName() string  {
	return "v1.x"
}

func (receiver *Postgres_V_1_X) GetSemver() string {
	return ">=1.x,<2.x"
}

func (receiver *Postgres_V_1_X) IsDbUsed(runtimeVersion ,root string) bool  {
	return true
}

func (receiver *Postgres_V_1_X) IsDbFound(runtimeVersion ,root string) bool {
	return true
}

func (receiver *Postgres_V_1_X ) GetDbName() string  {
	return "postgres"
}

func (receiver *Postgres_V_1_X) PercentOfDbUsed(runtimeVersion ,root string) float64  {
	return 90.67
}

// it returns port as well
func (receiver *Postgres_V_1_X) Detect(runtimeVersion, root string) (bool, int64) {
  return  true,8765
}
