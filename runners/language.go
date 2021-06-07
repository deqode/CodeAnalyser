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

//DetectRuntime It will detect app's runtime version
func DetectRuntime(ctx context.Context, projectRootPath, pluginPath string) string {
	pluginCall, _ := pluginClient.CreateDetectRuntimeClient(utils.CallPluginCommand(pluginPath))

	languageVersion, err := pluginCall.Detect(&pluginPb.StringInput{Value: projectRootPath})
	if err != nil || languageVersion.Error != nil {
		utils.Logger(err, languageVersion.Error)
		return ""
	}
	return languageVersion.Value
}

func DetectAndRunBuildDirectory(ctx context.Context, input *pluginPb.Input, pluginPath string) map[string]string {
	pluginCall, _ := pluginClient.CreateBuildDirectoryClient(utils.CallPluginCommand(pluginPath))
	buildDirectory, err := pluginCall.Detect(input)
	if err != nil || buildDirectory.Error != nil {
		utils.Logger(err, buildDirectory.Error)
		return nil
	}
	return buildDirectory.Value
}

func DetectTestCasesCommand(ctx context.Context, input *pluginPb.Input, pluginPath string) *languageSpecific.TestCasesCommandOutput {
	pluginCall, _ := pluginClient.CreateTestCaseCommandClient(utils.CallPluginCommand(pluginPath))

	commands, err := pluginCall.Detect(input)
	if err != nil || commands.Error != nil {
		utils.Logger(err, commands.Error)
		return nil
	}
	testCasesCommand := &languageSpecific.TestCasesCommandOutput{}
	if len(commands.Commands) > 0 {
		testCasesCommand.Commands = commands.Commands
		testCasesCommand.Used = true
	}
	return testCasesCommand
}

func RunStaticAssetsCommand(ctx context.Context, input *pluginPb.Input, pluginPath string) *languageSpecific.StaticAssetsOutput {
	pluginCall, _ := pluginClient.CreateStaticAssetsClient(utils.CallPluginCommand(pluginPath))

	respose, err := pluginCall.Detect(input)
	if err != nil || respose.Error != nil {
		utils.Logger(err, respose.Error)
		return nil
	}
	staticAssetsOutput := languageSpecific.StaticAssetsOutput{}
	if len(respose.StaticAsset) > 0 {
		staticAssetsOutput.Used = true
		staticAssetsOutput.Assets = respose.StaticAsset
	}
	return &staticAssetsOutput
}

func GetCommands(ctx context.Context, projectRootPath, pluginPath string) *languageSpecific.Commands {
	pluginCall, _ := pluginClient.CreateCommandsClient(utils.CallPluginCommand(pluginPath))

	var err error
	commands := languageSpecific.Commands{
		BuildCommands:     nil,
		StartUpCommands:   nil,
		SeedCommands:      nil,
		MigrationCommands: nil,
		AdHocScripts:      nil,
	}
	rootPathInput := &pluginPb.StringInput{
		Value: projectRootPath,
	}
	adHocScripts, err := pluginCall.DetectAdHocScripts(rootPathInput)
	if err != nil || adHocScripts.Error != nil {
		utils.Logger(err, adHocScripts.Error)
		return &commands
	}
	commands.AdHocScripts = adHocScripts

	seedCommands, err := pluginCall.DetectSeedCommands(rootPathInput)
	if err != nil || seedCommands.Error != nil {
		utils.Logger(err, seedCommands.Error)
		return &commands
	}
	commands.SeedCommands = seedCommands

	buildCommands, err := pluginCall.DetectBuildCommands(rootPathInput)
	if err != nil || buildCommands.Error != nil {
		utils.Logger(err, buildCommands.Error)
		return &commands
	}
	commands.BuildCommands = buildCommands

	migrationCommands, err := pluginCall.DetectMigrationCommands(rootPathInput)
	if err != nil || migrationCommands.Error != nil {
		utils.Logger(err, migrationCommands.Error)
		return &commands
	}
	commands.MigrationCommands = migrationCommands

	startUpCommands, err := pluginCall.DetectStartUpCommands(rootPathInput)
	if err != nil || startUpCommands.Error != nil {
		utils.Logger(err, startUpCommands.Error)
		return &commands
	}
	commands.StartUpCommands = startUpCommands

	return &commands
}

// RunPreDetectCommand this will run before detection for formatting, filtration, cleanup and all such similar commands
func RunPreDetectCommand(ctx context.Context, input *pluginPb.Input, pluginPath string) {
	pluginCall, _ := pluginClient.CreatePreDetectCommandClient(utils.CallPluginCommand(pluginPath))

	response, err := pluginCall.RunPreDetect(input)
	if err != nil || response.Error != nil {
		utils.Logger(err, response.Error)
	}
	log.Println("pre detect commands found and executed successfully ")
}

//TODO discuss with rajaram , rename variables

//GetParsedDependencies get map of parsed dependencies for example beego is a framework
func GetParsedDependencies(ctx context.Context, languageVersion, projectRootPath string, pluginDetails *versionsPB.LanguageVersion) *map[string]map[string]DependencyDetail {
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
	dependencies[Framework] = ExtractFrameworksFromProjectDependencies(dependencyList, pluginDetails)
	dependencies[DB] = ExtractDbsFromProjectDependencies(dependencyList, pluginDetails.Databases)
	dependencies[ORM] = ExtractOrmsFromProjectDependencies(dependencyList, pluginDetails)
	dependencies[Library] = ExtractLibraryFromProjectDependencies(dependencyList, pluginDetails)
	return &dependencies
}
