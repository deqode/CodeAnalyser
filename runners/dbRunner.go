package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	languageSpecificPB "code-analyser/protos/protos/outputs/languageSpecific"
	versionsPB "code-analyser/protos/protos/versions"
	"code-analyser/utils"
	"os/exec"
)

//ParseDbFromDependencies It will filter out frameworks from dependencies list
func ParseDbFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
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
func DbRunner(dbList map[string]DependencyDetail, runtimeVersion, root string) *languageSpecificPB.DBOutput {
	dbOutput := languageSpecificPB.DBOutput{
		Used:      false,
		Databases: []*languageSpecificPB.DB{},
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

//TODO handle errors on every method calls

//DbDetectorRunner It will execute plugin from command fetched from yaml file of same plugin
func DbDetectorRunner(name string, dbDetails DependencyDetail, runTimeVersion, root string) *languageSpecificPB.DB {
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
			return &languageSpecificPB.DB{
				Name:    name,
				Version: dbDetails.Version,
				Port:    detection.IntValue,
			}
		}
	}
	return nil
}
