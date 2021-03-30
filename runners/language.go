package runners

//todo: think good name

import (
	"code-analyser/protos/protos"
	"context"
)

type Dependency struct {
	Name string
	Version string
}


func DetectRuntime(ctx context.Context, path, languagePath string) (string,) {
// this function will call plugin to get runtime version of language
	//by giving dir path and plugin detector file
	//this will also return yaml file object
	return ""
}

func ParseLangaugeYML(filePath string)*protos.LanguageVersion{return nil}

func GetParsedDependencis(ctx context.Context,path,runtimeVersion string){
//this will firstly fetch version of repo used
//	and ParseLangaugeYML
//	check and segregate dependencies accordingly

}
