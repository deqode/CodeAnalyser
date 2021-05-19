package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"os/exec"
)

//ParseOrmFromDependencies It takes dependenciesList and filter out orms in map format
func ParseOrmFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
	orm := map[string]DependencyDetail{}
	for key, supportedOrm := range langYamlObject.Orms {
		if versionUsed, ok := dependenciesList[key]; ok {
			for versionName, v := range supportedOrm.Version {
				if helpers.SemverValidateFromArray(v.Semver, versionUsed) {
					orm[key] = DependencyDetail{
						Version: versionName,
						Command: v.Plugincommand,
					}
				}
			}
		}
	}
	return orm
}

//OrmRunner it append list of ORMS in ormoutput object
func OrmRunner(ormList map[string]DependencyDetail, runtimeVersion, root string) *languageSpecificPB.OrmOutput {
	ormOutputs := languageSpecificPB.OrmOutput{
		Used: false,
		Orms: []*languageSpecificPB.ORM{},
	}
	for ormUsed, ormDetails := range ormList {
		usedOrm := OrmDetectorRunner(ormUsed, ormDetails, runtimeVersion, root)
		if usedOrm != nil {
			ormOutputs.Used = true
			ormOutputs.Orms = append(ormOutputs.Orms, usedOrm)
		}
	}
	return &ormOutputs
}

//OrmDetectorRunner it run plugin file of ORM
func OrmDetectorRunner(name string, ormDetails DependencyDetail, runTimeVersion, root string) *languageSpecificPB.ORM {
	ormResponse, client := pluginClient.OrmPluginCall(exec.Command("sh", "-c", ormDetails.Command))
	for client.Exited() {
		client.Kill()
	}
	isUsed, err := ormResponse.IsORMUsed(&pb.ServiceInput{
		RuntimeVersion: runTimeVersion,
		Root:           root,
	})
	if err != nil || isUsed.Error != nil {
		utils.Logger(err, isUsed.Error)
		return nil
	}

	if isUsed.Value {
		detection, err := ormResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runTimeVersion,
			Root:           root,
		})
		if err != nil || detection.Error != nil {
			utils.Logger(err, detection.Error)
			return nil
		}
		if detection.Used {
			return &languageSpecificPB.ORM{
				Name:    name,
				Version: ormDetails.Version,
			}
		}
	}
	return nil
}
