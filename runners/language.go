package runners

//todo: think good name

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"code-analyser/utils"
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)




type Dependency struct {
	Name    string
	Version string
}

func DetectRuntime(ctx context.Context, path string, yamlLangObject *protos.LanguageVersion) string {
	runtimeResponse, client := pluginClient.DetectRuntimePluginCall(exec.Command("sh", "-c", yamlLangObject.Runtimeversions.Detector))
	defer client.Kill()
	runtimeVersion, err := runtimeResponse.DetectRuntime(&pb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return ""
	}
	if runtimeVersion.Error != nil {
		utils.Logger(runtimeVersion.Error)
		return ""
	}
	return runtimeVersion.Value
}

func ParseLangaugeYML(filePath string) *protos.LanguageVersion {
	path,_ := filepath.Abs(filePath)

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Logger(err,"ERROR")
		return nil
	}
	var lang protos.LanguageVersion
	err = yaml.Unmarshal(yamlFile, &lang)
	if err != nil {
		utils.Logger(err)
		return nil
	}
	return &lang
}
//
//func GetParsedDependencis(ctx context.Context, runtimeVersion, path string, langYamlObject *protos.LanguageVersion) *protos.LanguageVersion {
//	//dependenciesResponse, client := pluginClient.GetDependenciesPluginCall(exec.Command("sh", "-c", yamlLangObject.Runtimeversions.Detector))
//	//defer client.Kill()
//	dependenciesFound:= map[string]string{
//		"beego":"2.4",
//		"gin":"1.3",
//	}
//
//
//	//framework
//	GetDetectors:=protos.LanguageVersion{}
//	//framework := []*protos.PluginSemver{}
//	 for _,supportedFramework:=range  langYamlObject.Framework{
//		 if versionUsed,ok:=dependenciesFound[supportedFramework.Name];ok{
//		 	for _,v:=range supportedFramework.Versions{
//				if helpers.SeverValidate(v.Semver,versionUsed){
//					GetDetectors.Framework=append(GetDetectors.Framework,v)
//				}
//		    }
//		 }
//	 }
//
//	 log.Println(framework)
//
//
//
//}
