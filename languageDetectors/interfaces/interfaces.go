package interfaces

import (
	language_detectors "code-analyser/deployementFiles"
	"code-analyser/pluginClient/pb"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
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
	////ParseENVs will return ENVs
	ParseENVs(context.Context, string) ([]*languageSpecificPB.EnvOutput, error)
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
	DetectBuildDirectory(context.Context, string, string) ([]*languageSpecificPB.BuildDirectoryOutput, error)
	DetectTestCasesRunCommands(context.Context, string, string) ([]*languageSpecificPB.BuildDirectoryOutput, error) // Todo: proto
	language_detectors.Commands
}

//TODO need to segregate and decouple our code

//DetectRunTime This is for Detect version and its language
type DetectRunTime interface {
	DetectRuntime(inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error)
}

//Dependencies It is for all dependencies for example beego,gin,postgres
type Dependencies interface {
	GetDependencies(inputString *pb.ServiceInput) (*pb.ServiceOutputStringMap, error)
}
