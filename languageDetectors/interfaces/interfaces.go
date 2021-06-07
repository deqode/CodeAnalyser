package interfaces

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"code-analyser/protos/pb/plugin"
	"context"
)

//LanguageSpecificDetector It contains methods for detection of different dependencies
type LanguageSpecificDetector interface {
	//DetectRuntime will give runtime version of the language
	DetectRuntime(context.Context, string) (string, error)
	//RunPreDetect will run to modify libraries and format code
	RunPreDetect(context.Context, string, string) (bool, error)
	// TODO: interface ?
	RunParsers(context.Context, string, string) (interface{}, error)
	// ParseENVs  will return ENVs
//	ParseENVs(context.Context, string) ([]*languageSpecificPB.EnvOutput, error)
	////DetectFrameworks  will return framework detected in Dir
	DetectFrameworks(ctx context.Context, runtimeVersion string, root string) ([]*languageSpecificPB.FrameworkOutput, error)
	//DetectDBs will return Dbs detected in string
	DetectDBs(ctx context.Context, runtimeVersion, root string) (*languageSpecificPB.DBOutput, error)
	//DetectORMs will return orm detected in string
	DetectORMs(ctx context.Context, runtimeVersion, root string) (*languageSpecificPB.OrmOutput, error)
	//DetectDependencies will return getDependencies detected in string
	DetectDependencies(context.Context, string, string) ([]*languageSpecificPB.DependenciesOutput, error)
	//DetectLibraries will return libraries detected in string
	DetectLibraries(context.Context, string, string) ([]*languageSpecificPB.LibraryOutput, error)
	//GetStaticAssets will return type of staticAssets detected in string
	GetStaticAssets(context.Context, string, string) ([]*languageSpecificPB.StaticAssetsOutput, error)
	GetStack(context.Context, string, string) ([]*languageSpecificPB.StackOutput, error)
	DetectAppserver(context.Context, string, string) ([]*languageSpecificPB.AppserverOutput, error)
	DetectBuildDirectory(context.Context, string, string) (*languageSpecificPB.BuildDirectoryOutput, error)
	DetectTestCasesRunCommands(context.Context, string, string) ([]*languageSpecificPB.BuildDirectoryOutput, error) // Todo: proto
}

//TODO need to segregate and decouple our code

//DetectRunTime This is for DetectDockerFile version and its language
type DetectRunTime interface {
	Detect(inputString *plugin.StringInput) (*plugin.StringOutput, error)
}

//Dependencies It is for all dependencies for example beego,gin,postgres
type Dependencies interface {
	GetDependencies(inputString *plugin.Input) (*plugin.StringMapOutput, error)
}

type PreDetectCommands interface {
	RunPreDetect(input *plugin.Input) (*plugin.EmptyOutput, error)
}

type StaticAssets interface {
	Detect(input *plugin.Input) (*plugin.StaticAssetsOutput, error)
}

type BuildDirectory interface {
	Detect(input *plugin.Input) (*plugin.StringMapOutput, error)
}

type TestCasesRunCommands interface {
	Detect(input *plugin.Input) (*plugin.TestCommandOutput, error)
}
