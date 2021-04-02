package _go

import (
	"code-analyser/language_detectors/go/databases"
	"code-analyser/language_detectors/go/frameworks"
	"code-analyser/language_detectors/go/orms"
	"code-analyser/protos/protos"
	"code-analyser/runners"
	"code-analyser/versions"
	"context"
)

type GoDetector struct {
}

func (d *GoDetector) DetectRuntime(c context.Context, dir string) (string, error) {
	return "1.6", nil
}
func (d *GoDetector) RunPreDetect(ctx context.Context, version string, dir string) (bool, error) {
	return false, nil
}

func (d *GoDetector) RunParsers(context.Context, string, string) (interface{}, error) {
	return nil, nil
}
func (d *GoDetector) ParseENVs(context.Context, string) ([]*protos.EnvOutput, error) {
	return nil, nil

}
func (d *GoDetector) DetectFrameworks(ctx context.Context, runtimeVersion, path string) ([]*protos.FrameworkOutput, error) {
	//TODO:handle errors
	return runners.FrameworkRunner(&frameworks.GOFrameworkDetector{}, runtimeVersion, versions.GO, path), nil
}
func (d *GoDetector) DetectDBs(ctx context.Context, runtimeVersion, path string) (*protos.DBOutput, error) {
	//TODO:handle errors

	return runners.DbRunner(&databases.GODbDetector{}, runtimeVersion, versions.GO, path), nil
}
func (d *GoDetector) DetectORMs(ctx context.Context, runtimeVersion, path string) (*protos.OrmOutput, error) {
	return runners.OrmRunner(&orms.GoORMDetector{}, runtimeVersion, versions.GO, path), nil
}
func (d *GoDetector) DetectDependencies(context.Context, string, string) ([]*protos.DependenciesOutput, error) {
	return nil, nil

}
func (d *GoDetector) DetectLibraries(context.Context, string, string) ([]*protos.LibrariesOutput, error) {
	return nil, nil
}
func (d *GoDetector) GetStaticAssets(context.Context, string, string) ([]*protos.StaticAssetsOutput, error) {
	return nil, nil
}
func (d *GoDetector) GetStack(context.Context, string, string) ([]*protos.StackOutput, error) {
	return nil, nil
}
func (d *GoDetector) DetectAppserver(context.Context, string, string) ([]*protos.AppserverOutput, error) {
	return nil, nil
}
func (d *GoDetector) DetectBuildDirectory(context.Context, string, string) ([]*protos.BuildDirectoryOutput, error) {
	return nil, nil
}
func (d *GoDetector) DetectTestCasesRunCommands(context.Context, string, string) ([]*protos.BuildDirectoryOutput, error) {
	return nil, nil
}

func (d *GoDetector) DetectBuildCommands(context.Context, string) ([]*protos.BuildCommandsOutput, error) {
	return nil, nil
}
func (d *GoDetector) DetectStartUpCommands(context.Context, string) ([]*protos.StartUpCommandsOutput, error) {
	return nil, nil
}
func (d *GoDetector) DetectSeedCommands(context.Context, string) ([]*protos.SeedCommandsOutput, error) {
	return nil, nil
}
func (d *GoDetector) DetectMigrationCommands(context.Context, string) ([]*protos.MigrationCommandsOutput, error) {
	return nil, nil
}
func (d *GoDetector) DetectAdHocScripts(context.Context, string) (interface{}, error) {
	return nil, nil
}
