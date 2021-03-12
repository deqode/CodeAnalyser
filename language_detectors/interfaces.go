package language_detectors

import (
	"code-analyser/protos/pb"
	"context"
)

type LanguageSpecificDetector interface {
	DetectRuntime(context.Context, string) (string, error)                                    //root
	RunPreDetect(context.Context, string, string) (bool, error)                               //root,version
	RunParsers(context.Context, string, string) (interface{}, error)                          //root version TODO: interface ?
	ParseENVs(context.Context, string) ([]*protos.EnvOutput, error)                           //root
	DetectFrameworks(context.Context, string, string) ([]*protos.FrameworkOutput, error)      //root,version
	DetectDBs(context.Context, string, string) ([]*protos.DBOutput, error)                    //root,version
	DetectORMs(context.Context, string, string) ([]*protos.OrmOutput, error)                  //root,version
	DetectDependencies(context.Context, string, string) ([]*protos.DependenciesOutput, error) //root,version
	DetectLibraries(context.Context, string, string) ([]*protos.LibrariesOutput, error)       //root,version
	GetStaticAssets(context.Context, string, string) ([]*protos.StaticAssetsOutput, error)    //root,version
	GetStack(context.Context, string, string) ([]*protos.StackOutput, error)                  //root,version
	DetectAppserver(context.Context, string, string) ([]*protos.AppserverOutput, error)
	DetectBuildDirectory(context.Context, string, string) ([]*protos.BuildDirectoryOutput, error)

	DetectBuildCommands(context.Context, string, string) ([]*protos.BuildCommandsOutput, error)
	DetectStartUpCommands(context.Context, string, string) ([]*protos.StartUpCommandsOutput, error)
	DetectSeedCommands(context.Context, string, string) ([]*protos.SeedCommandsOutput, error)
	DetectMigrationCommands(context.Context, string, string) ([]*protos.MigrationCommandsOutput, error)
}
