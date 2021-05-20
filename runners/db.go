package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"log"
	"os/exec"
)

/*
ParseDbFromDependencies It will filter out frameworks from dependencies list where
dependenciesList : list of dependencies of any project ,
langYamlObject : dependencies supported by us
*/
func ParseDbFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
	db := map[string]DependencyDetail{}
	for key, supportedDb := range langYamlObject.Databases {
		for dbVersion, dependencyVersionDetails := range supportedDb.Version {
			for _, library := range dependencyVersionDetails.Libraries {
				if libraryUsedVersion, ok := dependenciesList[library.Name]; ok {
					if helpers.SemverValidateFromArray(library.Semver, libraryUsedVersion) {
						db[key] = DependencyDetail{
							Version: dbVersion,
							Command: dependencyVersionDetails.Plugincommand,
						}
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
	log.Fatalln(dbDetails.Command)
	dbResponse, client := pluginClient.DbPluginCall(exec.Command("sh", "-c", dbDetails.Command))
	for client.Exited() {
		client.Kill()
	}
	isUsed, err := dbResponse.IsDbUsed(&pb.ServiceInput{
		RuntimeVersion: runTimeVersion,
		Root:           root,
	})
	if err != nil || isUsed.Error != nil {
		utils.Logger(err, isUsed.Error)
		return nil
	}

	if isUsed.Value {
		detection, err := dbResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runTimeVersion,
			Root:           root,
		})
		if err != nil || detection.Error != nil {
			utils.Logger(err, detection.Error)
			return nil
		}
		if detection.Value {
			return &languageSpecificPB.DB{
				Name:    name,
				Version: dbDetails.Version,
				Port:    detection.IntValue,
			}
		}
	}
	return nil
}
