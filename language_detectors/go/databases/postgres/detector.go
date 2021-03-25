package postgres

import (
	"code-analyser/language_detectors/go/databases/postgres/V_1_X"
	"code-analyser/language_detectors/interfaces"
)

var postgresVersions=[]*interfaces.DbVersionDetector{
	{
		Default:true,
		Name:"v1.x",
		Semver: "",
		Detector: &V_1_X.Postgres_V_1_X{},
	},
}

type PostgresDb struct {
	Version[]*interfaces.DbVersionDetector
}

func (p *PostgresDb) GetVersionDetector(runtimeVersion, dbVersionFile, root string) interfaces.DbVersionDetector {
	return *postgresVersions[0]
}
