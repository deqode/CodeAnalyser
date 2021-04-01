package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"os/exec"
)

func ParseDbFromDependencies(dependenciesList map[string]string, langYamlObject *protos.LanguageVersion) map[string]*protos.PluginSemver {
	//framework
	db := map[string]*protos.PluginSemver{}
	for _, supportedDb := range langYamlObject.Databases {
		if versionUsed, ok := dependenciesList[supportedDb.Name]; ok {
			for _, v := range supportedDb.Versions {
				if helpers.SeverValidate(v.Semver, versionUsed) {
					db[supportedDb.Name] = v
				}
			}
		}
	}
	return db
}

//DbRunner will run to detect its dbs and return its detectors
func DbRunner(dbList map[string]*protos.PluginSemver, runtimeVersion, root string) protos.DBOutput {
	dbOutput := protos.DBOutput{
		Used:      false,
		Databases: []*protos.DB{},
	}
	for dbUsed, dbDetails := range dbList {
		dbOutput.Used = true
		isUsed := DbDetectorRunner(dbUsed, dbDetails, runtimeVersion, root)
		dbOutput.Databases = append(dbOutput.Databases, isUsed)
	}
	return dbOutput
}

////TODO handle errors on every method calls
func DbDetectorRunner(name string, dbDetails *protos.PluginSemver, runTimeVersion, root string) *protos.DB {
	dbResponse, client := pluginClient.DbPluginCall(exec.Command("sh", "-c", "go run plugin/go/db/postgres/V_1_X/main.go"))
	defer client.Kill()
	isUsed, err := dbResponse.IsDbUsed(&pb.ServiceInput{
		RuntimeVersion: runTimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}

	if isUsed.Value == true {
		detection, err := dbResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runTimeVersion,
			Root:           root,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if detection.Value == true {
			return &protos.DB{
				Name:    name,
				Version: dbDetails.Name,
				Port:    detection.IntValue,
			}
		}
	}
	return nil
}
