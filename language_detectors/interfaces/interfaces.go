package interfaces

import (
	language_detectors "code-analyser/deployement_files"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/protos"
	"context"
)

type LanguageSpecificDetector interface {
	//DetectRuntime will give runtime version of the language
	DetectRuntime(context.Context, string) (string, error)
	//RunPreDetect will run to modify libraries and format code
	RunPreDetect(context.Context, string, string) (bool, error)
	// TODO: interface ?
	RunParsers(context.Context, string, string) (interface{}, error)
	////ParseENVs will return ENVs
	ParseENVs(context.Context, string) ([]*protos.EnvOutput, error)
	////DetectFrameworks  will return framework detected in Dir
	DetectFrameworks(ctx context.Context, runtimeVersion string, root string) ([]*protos.FrameworkOutput, error)
	//DetectDBs will return Dbs detected in string
	DetectDBs(ctx context.Context, runtimeVersion, root string) (*protos.DBOutput, error)
	//DetectORMs will return orm detected in string
	DetectORMs(ctx context.Context, runtimeVersion, root string) (*protos.OrmOutput, error)
	//DetectDependencies will return dependencies detected in string
	DetectDependencies(context.Context, string, string) ([]*protos.DependenciesOutput, error)
	//DetectLibraries will return libraries detected in string
	DetectLibraries(context.Context, string, string) ([]*protos.LibrariesOutput, error)
	//GetStaticAssets will return type of staticAssets detected in string
	GetStaticAssets(context.Context, string, string) ([]*protos.StaticAssetsOutput, error)
	GetStack(context.Context, string, string) ([]*protos.StackOutput, error)
	DetectAppserver(context.Context, string, string) ([]*protos.AppserverOutput, error)
	DetectBuildDirectory(context.Context, string, string) ([]*protos.BuildDirectoryOutput, error)
	DetectTestCasesRunCommands(context.Context, string, string) ([]*protos.BuildDirectoryOutput, error) // Todo: proto
	language_detectors.Commands
}

//TODO need to segregate and decouple our code

type DetectRunTime interface {
	DetectRuntime(inputString *pb.ServiceInputString) (*pb.ServiceOutputString, error)
}

type Dependencies interface {
	GetDependencies(inputString *pb.ServiceInput) (*pb.ServiceOutputStringMap, error)
}

//count 14
