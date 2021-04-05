package runners

import (
	"code-analyser/helpers"
	"code-analyser/protos/protos"
)

func ParseOrmFromDependencies(dependenciesList map[string]string, langYamlObject *protos.LanguageVersion) map[string]DependencyDetail {
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

//func OrmRunner(ormList map[string]*protos.PluginSemver, runtimeVersion, root string) protos.OrmOutput {
//	ormOutputs := protos.OrmOutput{
//		Used: false,
//		Orms: []*protos.ORM{},
//	}
//	for ormUsed, ormDetails := range ormList {
//		ormOutputs.Used = true
//		usedOrm := OrmDetectorRunner(ormUsed, ormDetails, runtimeVersion, root)
//		ormOutputs.Orms = append(ormOutputs.Orms, usedOrm)
//	}
//	return ormOutputs
//}
//
//func OrmDetectorRunner(name string, ormDetails *protos.PluginSemver, runTimeVersion, root string) *protos.ORM {
//	ormResponse, client := pluginClient.OrmPluginCall(exec.Command("sh", "-c", "go run plugin/go/orm/gorm/V_1_X/main.go"))
//	defer client.Kill()
//	isUsed, err := ormResponse.IsORMUsed(&pb.ServiceInput{
//		RuntimeVersion: runTimeVersion,
//		Root:           root,
//	})
//	if err != nil {
//		utils.Logger(err)
//		return nil
//	}
//
//	if isUsed.Value == true {
//		detection, err := ormResponse.Detect(&pb.ServiceInput{
//			RuntimeVersion: runTimeVersion,
//			Root:           root,
//		})
//		if err != nil {
//			utils.Logger(err)
//			return nil
//		}
//		if detection.Value == true {
//			return &protos.ORM{
//				Name:    name,
//				Version: ormDetails.Name,
//			}
//		}
//	}
//	return nil
//}
