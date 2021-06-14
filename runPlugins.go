package main

import (
	decisionmakerPB "code-analyser/protos/pb"
	pb "code-analyser/protos/pb/plugin"
	"code-analyser/runners"
	"context"
)

// LoadGlobalPlugins executed all plugins related to globalFiles files (like docker-compose.yaml,DockerFile MakeFile etc)
func LoadGlobalPlugins(ctx context.Context, globalDetection *decisionmakerPB.GlobalDetections, globalPlugin *versionsPB.GlobalPlugin, path string) {

	globalDetection.DockerFile, globalDetection.DockerComposeFile = runners.ExecuteDockerAndComposePlugin(ctx, path, globalPlugin.DockerFile)

	globalDetection.ProcFile = runners.ExecuteProcFileDetectionPlugin(ctx, path, globalPlugin.ProcFile)

	globalDetection.Makefile = runners.ExecuteMakeFileDetectionPlugin(ctx, path, globalPlugin.MakeFile)
}

//LoadLanguagePlugins it runs all detectors of dependencies ex. orm,framework etc ....
func LoadLanguagePlugins(ctx context.Context, languageSpecificDetections *decisionmakerPB.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, pluginDetails *versionsPB.LanguageVersion, runtimeVersion string, projectRootPath string) {
	pluginInput := &pb.Input{
		RuntimeVersion: runtimeVersion,
		RootPath:       projectRootPath,
	}

	languageSpecificDetections.Commands = runners.ExecuteCommandsDetectionPlugin(ctx, projectRootPath, pluginDetails.CommandsPluginPath)

	languageSpecificDetections.Orm = runners.ExecuteOrmPlugins(ctx, allDependencies[runners.ORM], runtimeVersion, projectRootPath)

	languageSpecificDetections.Db = runners.ExecuteDbPlugins(ctx, allDependencies[runners.DB], runtimeVersion, projectRootPath)

	languageSpecificDetections.Framework = runners.ExecuteFrameworkPlugins(ctx, allDependencies[runners.Framework], runtimeVersion, projectRootPath)

	languageSpecificDetections.Env = runners.ExecuteEnvsDetectionPlugin(ctx, pluginDetails.EnvPluginPath, runtimeVersion, projectRootPath)

	languageSpecificDetections.StaticAssets = runners.ExecuteStaticAssetsPlugin(ctx, pluginInput, pluginDetails.StaticAssetsPluginPath)

	languageSpecificDetections.Libraries = runners.ExecuteLibraryPlugins(ctx, allDependencies[runners.Library], runtimeVersion, projectRootPath)

	languageSpecificDetections.BuildDirectory = runners.ExecuteBuildDirectoryPlugin(ctx, pluginInput, pluginDetails.BuildDirectoryPluginPath)

	languageSpecificDetections.TestCases = runners.ExecuteTestCommandDetectionPlugin(ctx, pluginInput, pluginDetails.TestCasesCommandPluginPath)
}
