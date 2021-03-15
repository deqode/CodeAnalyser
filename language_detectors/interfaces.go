package language_detectors

import (
	protos "code-analyser/protos/pb"
	"context"
	"github.com/spf13/afero"
)

type LanguageSpecificDetector interface {
	//DetectRuntime will give runtime version of the language
	DetectRuntime(context.Context, afero.Afero) (string, error)
	//RunPreDetect will run to modify libraries and format code
	RunPreDetect(context.Context, string, string) (bool, error)
	//// TODO: interface ?
	//RunParsers(context.Context, string, Dir) (interface{}, error)
	////ParseENVs will return ENVs
	//ParseENVs(context.Context, Dir) ([]*protos.EnvOutput, error)
	////DetectFrameworks  will return framework detected in Dir
	DetectFrameworks(context.Context, string, string) ([]*protos.FrameworkOutput, error)
	////DetectDBs will return Dbs detected in Dir
	//DetectDBs(context.Context, string, Dir) ([]*protos.DBOutput, error)
	////DetectORMs will return orm detected in Dir
	//DetectORMs(context.Context, string, Dir) ([]*protos.OrmOutput, error)
	////DetectDependencies will return dependencies detected in Dir
	//DetectDependencies(context.Context, string, Dir) ([]*protos.DependenciesOutput, error)
	////DetectLibraries will return libraries detected in Dir
	//DetectLibraries(context.Context, string, Dir) ([]*protos.LibrariesOutput, error)
	////GetStaticAssets will return type of staticAssets detected in Dir
	//GetStaticAssets(context.Context, string, Dir) ([]*protos.StaticAssetsOutput, error)
	//GetStack(context.Context, string, Dir) ([]*protos.StackOutput, error)
	//DetectAppserver(context.Context, string, Dir) ([]*protos.AppserverOutput, error)
	//DetectBuildDirectory(context.Context, string, Dir) ([]*protos.BuildDirectoryOutput, error)
	//DetectTestCasesRunCommands(context.Context, string, Dir) ([]*protos.BuildDirectoryOutput, error)
	//language_detectors.Commands
}

type Dir interface {
	Validate(string) bool
	Abs(string) string
	Join(string, ...string) (string, error)

	SetDir(string) (string, error)
	GetDir(string) (string, error)
	IsDir(string) bool

	SetFile(string) (string, error)
	GetFile(string) (string, error)
	IsFile(string) bool

	Ext(string) (string, error)
}
