package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/languageSpecific"
	pluginPb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"context"
	"log"
)

//todo: think good name

//DependencyDetail It stores version and execution command of dependency as a plugin for example
//beego = {
//           version : 1.5
//           command : "go plugin/beego/main.go"
//        }
type DependencyDetail struct {
	Version string
	Command string
}

const (
	//Framework is FRAMEWORK
	Framework = "framework"
	//DB is DB ( database)
	DB = "db"
	//ORM is ORM
	ORM     = "orm"
	Library = "library"
)

//ExecuteRuntimeDetectionPlugin It will detect app's runtime version
func ExecuteRuntimeDetectionPlugin(ctx context.Context, projectRootPath, pluginPath string) string {
	pluginCall, _ := pluginClient.CreateDetectRuntimeClient(utils.CallPluginCommand(pluginPath))

	languageVersion, err := pluginCall.Detect(&pluginPb.StringInput{Value: projectRootPath})
	if err != nil || languageVersion.Error != nil {
		utils.Logger(err, languageVersion)
		return ""
	}

	return languageVersion.Value
}

func ExecuteBuildDirectoryPlugin(ctx context.Context, input *pluginPb.Input, pluginPath string) map[string]string {
	pluginCall, _ := pluginClient.CreateBuildDirectoryClient(utils.CallPluginCommand(pluginPath))

	buildDirectory, err := pluginCall.Detect(input)
	if err != nil || buildDirectory.Error != nil {
		utils.Logger(err, buildDirectory)
		return nil
	}

	return buildDirectory.Value
}

func ExecuteTestCommandDetectionPlugin(ctx context.Context, input *pluginPb.Input, pluginPath string) *languageSpecific.TestCasesCommandOutput {
	pluginCall, _ := pluginClient.CreateTestCaseCommandClient(utils.CallPluginCommand(pluginPath))

	commands, err := pluginCall.Detect(input)
	if err != nil || commands.Error != nil {
		utils.Logger(err, commands)
		return nil
	}

	testCasesCommand := &languageSpecific.TestCasesCommandOutput{}
	if len(commands.Commands) > 0 {
		testCasesCommand.Commands = commands.Commands
		testCasesCommand.Used = true
	}

	return testCasesCommand
}

func ExecuteStaticAssetsPlugin(ctx context.Context, input *pluginPb.Input, pluginPath string) *languageSpecific.StaticAssetsOutput {
	pluginCall, _ := pluginClient.CreateStaticAssetsClient(utils.CallPluginCommand(pluginPath))

	staticAsset, err := pluginCall.Detect(input)
	if err != nil || staticAsset.Error != nil {
		utils.Logger(err, staticAsset)
		return nil
	}

	staticAssetsOutput := languageSpecific.StaticAssetsOutput{}
	if len(staticAsset.Value) > 0 {
		staticAssetsOutput.Used = true
		staticAssetsOutput.Assets = staticAsset.Value
	}

	return &staticAssetsOutput
}

func ExecuteCommandsDetectionPlugin(ctx context.Context, projectRootPath, pluginPath string) *languageSpecific.Commands {
	pluginCall, _ := pluginClient.CreateCommandsClient(utils.CallPluginCommand(pluginPath))

	commands := languageSpecific.Commands{}
	rootPathInput := &pluginPb.StringInput{
		Value: projectRootPath,
	}

	adHocScripts, err := pluginCall.DetectAdHocScripts(rootPathInput)
	if err != nil || adHocScripts.Error != nil {
		utils.Logger(err, adHocScripts)
		return &commands
	}
	commands.AdHocScripts = adHocScripts

	seedCommands, err := pluginCall.DetectSeedCommands(rootPathInput)
	if err != nil || seedCommands.Error != nil {
		utils.Logger(err, seedCommands)
		return &commands
	}
	commands.SeedCommands = seedCommands

	buildCommands, err := pluginCall.DetectBuildCommands(rootPathInput)
	if err != nil || buildCommands.Error != nil {
		utils.Logger(err, buildCommands)
		return &commands
	}
	commands.BuildCommands = buildCommands

	migrationCommands, err := pluginCall.DetectMigrationCommands(rootPathInput)
	if err != nil || migrationCommands.Error != nil {
		utils.Logger(err, migrationCommands)
		return &commands
	}
	commands.MigrationCommands = migrationCommands

	startUpCommands, err := pluginCall.DetectStartUpCommands(rootPathInput)
	if err != nil || startUpCommands.Error != nil {
		utils.Logger(err, startUpCommands)
		return &commands
	}
	commands.StartUpCommands = startUpCommands

	return &commands
}

// ExecutePreDetectionPlugin this will run before detection for formatting, filtration, cleanup and all such similar commands
func ExecutePreDetectionPlugin(ctx context.Context, input *pluginPb.Input, pluginPath string) {
	pluginCall, _ := pluginClient.CreatePreDetectCommandClient(utils.CallPluginCommand(pluginPath))

	response, err := pluginCall.RunPreDetect(input)
	if err != nil || response.Error != nil {
		utils.Logger(err, response)
	} else {
		log.Println("pre detect commands found and executed successfully ")
	}
}

//TODO discuss with rajaram , rename variables

//GetDependenciesFromProject get map of parsed dependencies for example beego is a framework
func GetDependenciesFromProject(ctx context.Context, languageVersion, projectRootPath string, pluginDetails *versionsPB.LanguageVersion) *map[string]map[string]DependencyDetail {
	dependencies := map[string]map[string]DependencyDetail{}
	var dependencyVersionDetails *versionsPB.DependencyVersionDetails
	var runtimeVersion string

	for rt, supportedRuntimeVersions := range pluginDetails.RuntimeVersions {
		if helpers.SemverValidateFromArray(supportedRuntimeVersions.Semver, languageVersion) {
			dependencyVersionDetails = supportedRuntimeVersions
			runtimeVersion = rt
			break
		}
	}
	if dependencyVersionDetails == nil {
		return nil
	}
	pluginCall, _ := pluginClient.CreateDependenciesClient(utils.CallPluginCommand(dependencyVersionDetails.PluginPath))

	response, err := pluginCall.GetDependencies(&pluginPb.Input{
		RuntimeVersion: runtimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil || response.Error != nil {
		utils.Logger(err, response)
		return nil
	}
	dependencyList := response.Value
	dependencies[Framework] = ExtractFrameworksFromProjectDependencies(ctx,dependencyList, pluginDetails.Frameworks)
	dependencies[DB] = ExtractDbsFromProjectDependencies(ctx,dependencyList, pluginDetails.Databases)
	dependencies[ORM] = ExtractOrmsFromProjectDependencies(ctx,dependencyList, pluginDetails.Orms)
	dependencies[Library] = ExtractLibraryFromProjectDependencies(ctx,dependencyList, pluginDetails.Libraries)

	return &dependencies
}
