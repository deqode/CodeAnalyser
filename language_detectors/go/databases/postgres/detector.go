package postgres

import (
	"code-analyser/language_detectors/go/databases/postgres/V_1_X"
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient"
	"github.com/hashicorp/go-plugin"
	"os/exec"
)

var postgresVersions = []*interfaces.DbVersionDetector{
	{
		Default:  true,
		Name:     "v1.x",
		Semver:   "",
		Detector: &V_1_X.Postgres_V_1_X{},
	},
}

type PostgresDb struct {
	Version []*interfaces.DbVersionDetector
}

func (p *PostgresDb) GetVersionDetector(runtimeVersion, dbVersionFile, root string) (interfaces.DbVersionDetector, *plugin.Client) {

	////TODO: need to add semver check and return correct version
	dbDetector, client := pluginClient.DbPluginCall(exec.Command("sh", "-c", "go run plugin/go/db/postgres/V_1_X/main.go"))
	x := interfaces.DbVersionDetector{
		Detector: dbDetector,
	}
	return x, client
}
