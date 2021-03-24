package postgress

import "code-analyser/versions"

var PostgresVersions = map[string]*versions.DBVersionDetector{
	"v1.x": {
		Default:  true,
		Name:     "v12",
		Detector: Postgress_v_12,
	},
}
