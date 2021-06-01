package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb"
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
func DetectRuntime(ctx context.Context, path string, pluginDetails *versionsPB.LanguageVersion) string {
	runtimeResponse, client := pluginClient.DetectRuntimePluginCall(utils.CallPluginCommand(pluginDetails.Detectruntimecommand))

	runtimeVersion, err := runtimeResponse.DetectRuntime(&pluginPb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return ""
	}
	if runtimeVersion.Error != nil {
		utils.Logger(runtimeVersion.Error)
		return ""
	}

	for client.Exited() {
		client.Kill()
	}
	return runtimeVersion.Value
}

func DetectAndRunBuildDirectory(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) map[string]string {
	res, client := pluginClient.BuildDirectoryPluginCall(utils.CallPluginCommand(pluginDetails.BuildDirectoryCommand))
	for client.Exited() {
		client.Kill()
	}
	detection, err := res.Detect(input)
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if detection.Error != nil {
		utils.Logger(detection.Error)
		return nil
	}
	return detection.Value
}

func DetectTestCasesCommand(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) *languageSpecific.TestCasesCommandOutput {
	res, client := pluginClient.TestCaseCommandPluginCall(utils.CallPluginCommand(pluginDetails.DetectTestCasesCommand))
	for client.Exited() {
		client.Kill()
	}
	detection, err := res.Detect(input)
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if detection.Error != nil {
		utils.Logger(detection.Error)
		return nil
	}
	testCasesCommandOutput := languageSpecific.TestCasesCommandOutput{
		Commands: detection.Commands,
	}
	if len(detection.Commands) > 0 {
		testCasesCommandOutput.Used = true
	}
	return &testCasesCommandOutput
}

func RunStaticAssetsCommand(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) *languageSpecific.StaticAssetsOutput {
	res, client := pluginClient.StaticAssetsPluginCall(utils.CallPluginCommand(pluginDetails.StaticAssetsCommand))
	for client.Exited() {
		client.Kill()
	}
	detection, err := res.Detect(input)
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if detection.Error != nil {
		utils.Logger(detection.Error)
		return nil
	}
	staticAssetsOutput := languageSpecific.StaticAssetsOutput{
		Assets: detection.StaticAsset,
	}
	if len(detection.StaticAsset) > 0 {
		staticAssetsOutput.Used = true
	}
	return &staticAssetsOutput
}

func GetCommands(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) *pb.Commands {
	response, client := pluginClient.CommandsPluginCall(utils.CallPluginCommand(pluginDetails.Commands))
	defer func() {
		for client.Exited() {
			client.Kill()
		}
	}()
	var err error
	serviceCommandsInput := &pluginPb.ServiceCommandsInput{
		Root:     input.Root,
		Language: input.RuntimeVersion,
	}
	commands := pb.Commands{
		BuildCommands:      nil,
		StartUpCommands:    nil,
		SeedCommands:       nil,
		MigrationCommands:  nil,
		AdHocScriptsOutput: nil,
	}
	detectAdHocScript, err := response.DetectAdHocScripts(serviceCommandsInput)
	if err != nil || detectAdHocScript.Error != nil {
		utils.Logger(err, detectAdHocScript.Error)
		return &commands
	}
	commands.AdHocScriptsOutput = detectAdHocScript.AdHocScripts

	detectSeedCommand, err := response.DetectSeedCommands(serviceCommandsInput)
	if err != nil || detectSeedCommand.Error != nil {
		utils.Logger(err, detectSeedCommand.Error)
		return &commands
	}
	commands.SeedCommands = detectSeedCommand.SeedCommands

	detectBuildCommands, err := response.DetectBuildCommands(serviceCommandsInput)
	if err != nil || detectBuildCommands.Error != nil {
		utils.Logger(err, detectBuildCommands.Error)
		return &commands
	}
	commands.BuildCommands = detectBuildCommands.BuildCommands

	detectMigrationCommands, err := response.DetectMigrationCommands(serviceCommandsInput)
	if err != nil || detectMigrationCommands.Error != nil {
		utils.Logger(err, detectMigrationCommands.Error)
		return &commands
	}
	commands.MigrationCommands = detectMigrationCommands.MigrationCommands

	detectStartUpCommands, err := response.DetectStartUpCommands(serviceCommandsInput)
	if err != nil || detectStartUpCommands.Error != nil {
		utils.Logger(err, detectStartUpCommands.Error)
		return &commands
	}
	commands.StartUpCommands = detectStartUpCommands.StartUpCommands

	return &commands
}

// RunPreDetectCommand this will run before detection for formatting, filtration, cleanup and all such similar commands
func RunPreDetectCommand(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) {
	runtimeResponse, client := pluginClient.PreDetectCommandPluginCall(utils.CallPluginCommand(pluginDetails.PreDetectCommand))
	for client.Exited() {
		client.Kill()
	}
	runtimeVersion, err := runtimeResponse.RunPreDetect(input)
	if err != nil {
		utils.Logger(err)
	}
	if runtimeVersion.Error != nil {
		utils.Logger(runtimeVersion.Error)
	}
	log.Println("pre detect commands found and executed successfully ")
}

//GetParsedDependencies get map of parsed dependencies for example beego is a framework
func GetParsedDependencies(ctx context.Context, languageVersion, path string, pluginDetails *versionsPB.LanguageVersion) map[string]map[string]DependencyDetail {
	AllDependencies := map[string]map[string]DependencyDetail{}
	var dependenciesCommand *versionsPB.DependencyVersionDetails
	var runtimeVersion string
	for rt, supportedRuntimeVersions := range pluginDetails.Runtimeversions {
		if helpers.SemverValidateFromArray(supportedRuntimeVersions.Semver, languageVersion) {
			dependenciesCommand = supportedRuntimeVersions
			runtimeVersion = rt
			break
		}
	}
	if dependenciesCommand != nil {
		dependenciesResponse, client := pluginClient.DependenciesPluginCall(utils.CallPluginCommand(dependenciesCommand.Plugincommand))
		defer func() {
			for client.Exited() {
				client.Kill()
			}
		}()
		getdependenciesFound, err := dependenciesResponse.GetDependencies(&pluginPb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           path,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if getdependenciesFound.Error != nil {
			utils.Logger(getdependenciesFound.Error)
			return nil
		}
		dependenciesList := getdependenciesFound.Value
		AllDependencies[Framework] = ParseFrameworkFromDependencies(dependenciesList, pluginDetails)
		AllDependencies[DB] = ParseDbFromDependencies(dependenciesList, pluginDetails)
		AllDependencies[ORM] = ParseOrmFromDependencies(dependenciesList, pluginDetails)
		AllDependencies[Library] = ParseLibraryFromDependencies(dependenciesList, pluginDetails)
		return AllDependencies
	}

	return nil

}
