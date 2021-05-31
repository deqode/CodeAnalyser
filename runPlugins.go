package main

import (
	decisionmakerPB "code-analyser/protos/pb"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/runners"
	"context"
	"sync"
)

// RunAllGlobalPlugins executed all plugins related to global files (like docker-compose.yaml,DockerFile MakeFile etc)
func RunAllGlobalPlugins(globalDetection *decisionmakerPB.GlobalDetections, globalPlugin *versionsPB.GlobalPlugin, ctx context.Context, path string) {
	var wg sync.WaitGroup
	wg.Add(3)
	var mutex = &sync.Mutex{}
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.DockerFile, globalDetection.DockerComposeFile = runners.DetectDockerAndComposeFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.ProcFile = runners.DetectProcFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.Makefile = runners.DetectMakeFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
}

//RunAllLanguageSpecificPlugins it runs all detectors of dependencies ex. orm,framework etc ....
func RunAllLanguageSpecificPlugins(ctx context.Context, languageSpecificDetections *decisionmakerPB.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, pluginDetails *versionsPB.LanguageVersion, runtimeVersion string, path string, globalPlugin *versionsPB.GlobalPlugin, commands *decisionmakerPB.Commands) {
	var wg sync.WaitGroup
	wg.Add(9)
	var mutex = &sync.Mutex{}
	go func() {
		defer wg.Done()
		if pluginDetails.Commands != "" {
			mutex.Lock()
			languageSpecificDetections.Commands = runners.GetCommands(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Orm = runners.OrmRunner(allDependencies[runners.ORM], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		commands = runners.DetectAndRunCommands(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Db = runners.DbRunner(allDependencies[runners.DB], runtimeVersion, path)
		mutex.Unlock()

	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Framework = runners.FrameworkRunner(allDependencies[runners.Framework], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		if pluginDetails.DetectEnvCommand != "" {
			mutex.Lock()
			languageSpecificDetections.Env = runners.EnvDetectAndRunner(pluginDetails, runtimeVersion, path)
			mutex.Unlock()
		}

	}()
	go func() {
		defer wg.Done()
		if pluginDetails.StaticAssetsCommand != "" {
			mutex.Lock()
			languageSpecificDetections.StaticAssets = runners.RunStaticAssetsCommand(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Libraries = runners.LibraryRunner(allDependencies[runners.Library], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		if pluginDetails.BuildDirectoryCommand != "" {
			mutex.Lock()
			languageSpecificDetections.BuildDirectory = runners.DetectAndRunBuildDirectory(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if pluginDetails.DetectTestCasesCommand != "" {
			mutex.Lock()
			languageSpecificDetections.TestCases = runners.DetectTestCasesCommand(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	wg.Wait()

}
