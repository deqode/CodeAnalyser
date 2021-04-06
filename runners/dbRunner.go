package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"os/exec"
)

func ParseDbFromDependencies(dependenciesList map[string]string, langYamlObject *protos.LanguageVersion) map[string]DependencyDetail {
	//framework
	db := map[string]DependencyDetail{}
	for key, supportedDb := range langYamlObject.Databases {
		if versionUsed, ok := dependenciesList[key]; ok {
			for versionName, v := range supportedDb.Version {
				if helpers.SeverValidateFromArray(v.Semver, versionUsed) {
					db[key] = DependencyDetail{
						Version: versionName,
						Command: v.Plugincommand,
					}
				}
			}
		}
	}
	return db
}

//DbRunner will run to detect its dbs and return its detectors
func DbRunner(dbList map[string]DependencyDetail, runtimeVersion, root string) *protos.DBOutput {
	dbOutput := protos.DBOutput{
		Used:      false,
		Databases: []*protos.DB{},
	}
	for dbUsed, dbDetails := range dbList {
		isUsed := DbDetectorRunner(dbUsed, dbDetails, runtimeVersion, root)
		if isUsed != nil {
			dbOutput.Used = true
			dbOutput.Databases = append(dbOutput.Databases, isUsed)
		}
	}
	return &dbOutput
}

////TODO handle errors on every method calls
func DbDetectorRunner(name string, dbDetails DependencyDetail, runTimeVersion, root string) *protos.DB {
	dbResponse, client := pluginClient.DbPluginCall(exec.Command("sh", "-c", dbDetails.Command))
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
				Version: dbDetails.Version,
				Port:    detection.IntValue,
			}
		}
	}
	return nil
}
