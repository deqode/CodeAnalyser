package main

import (
	decisionmakerPB "code-analyser/protos/pb"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/runners"
	"context"
)

// RunAllGlobalPlugins executed all plugins related to global files (like docker-compose.yaml,DockerFile MakeFile etc)
func RunAllGlobalPlugins(globalDetection *decisionmakerPB.GlobalDetections, globalPlugin *versionsPB.GlobalPlugin, ctx context.Context, path string) {

	globalDetection.DockerFile, globalDetection.DockerComposeFile = runners.DetectDockerAndComposeFile(ctx, path, globalPlugin)

	globalDetection.ProcFile = runners.DetectProcFile(ctx, path, globalPlugin)

	globalDetection.Makefile = runners.DetectMakeFile(ctx, path, globalPlugin)
}

//RunAllLanguageSpecificPlugins it runs all detectors of dependencies ex. orm,framework etc ....
func RunAllLanguageSpecificPlugins(ctx context.Context, languageSpecificDetections *decisionmakerPB.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, pluginDetails *versionsPB.LanguageVersion, runtimeVersion string, path string) {
	if pluginDetails.Commands != "" {
		languageSpecificDetections.Commands = runners.GetCommands(ctx, &pb.Input{
			RuntimeVersion: runtimeVersion,
			RootPath:           path,
		}, pluginDetails)
	}

	languageSpecificDetections.Orm = runners.OrmRunner(allDependencies[runners.ORM], runtimeVersion, path)

	languageSpecificDetections.Db = runners.DbRunner(allDependencies[runners.DB], runtimeVersion, path)

	languageSpecificDetections.Framework = runners.FrameworkRunner(allDependencies[runners.Framework], runtimeVersion, path)
	if pluginDetails.DetectEnvCommand != "" {
		languageSpecificDetections.Env = runners.EnvDetectAndRunner(pluginDetails, runtimeVersion, path)
	}

	if pluginDetails.StaticAssetsCommand != "" {
		languageSpecificDetections.StaticAssets = runners.RunStaticAssetsCommand(ctx, &pb.Input{
			RuntimeVersion: runtimeVersion,
			RootPath:           path,
		}, pluginDetails)
	}

	languageSpecificDetections.Libraries = runners.LibraryRunner(allDependencies[runners.Library], runtimeVersion, path)

	if pluginDetails.BuildDirectoryCommand != "" {
		languageSpecificDetections.BuildDirectory = runners.DetectAndRunBuildDirectory(ctx, &pb.Input{
			RuntimeVersion: runtimeVersion,
			RootPath:           path,
		}, pluginDetails)
	}

	if pluginDetails.DetectTestCasesCommand != "" {
		languageSpecificDetections.TestCases = runners.DetectTestCasesCommand(ctx, &pb.Input{
			RuntimeVersion: runtimeVersion,
			RootPath:           path,
		}, pluginDetails)
	}
}

