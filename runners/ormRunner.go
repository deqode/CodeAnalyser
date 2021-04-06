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

//ParseOrmFromDependencies It takes dependenciesList and filter out orms in map format
func ParseOrmFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
	orm := map[string]DependencyDetail{}
	for key, supportedOrm := range langYamlObject.Orms {
		if versionUsed, ok := dependenciesList[key]; ok {
			for versionName, v := range supportedOrm.Version {
				if helpers.SeverValidateFromArray(v.Semver, versionUsed) {
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
	defer client.Kill()
	isUsed, err := ormResponse.IsORMUsed(&pb.ServiceInput{
		RuntimeVersion: runTimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}

	if isUsed.Value == true {
		detection, err := ormResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runTimeVersion,
			Root:           root,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if detection.Used == true {
			return &languageSpecificPB.ORM{
				Name:    name,
				Version: ormDetails.Version,
			}
		}
	}
	return nil
}
