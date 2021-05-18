package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/global"
	"code-analyser/protos/pb/output/languageSpecific"
	pluginPb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"context"
	"log"
	"os/exec"
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

//DetectRuntime It will detect language and its version
func DetectRuntime(ctx context.Context, path string, yamlLangObject *versionsPB.LanguageVersion) string {
	runtimeResponse, client := pluginClient.DetectRuntimePluginCall(exec.Command("sh", "-c", yamlLangObject.Detectruntimecommand))
	for client.Exited() {
		client.Kill()
	}
	runtimeVersion, err := runtimeResponse.DetectRuntime(&pluginPb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return ""
	}
	if runtimeVersion.Error != nil {
		utils.Logger(runtimeVersion.Error)
		return ""
	}
	return runtimeVersion.Value
}

func DetectAndRunBuildDirectory(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) map[string]string {
	res, client := pluginClient.BuildDirectoryPluginCall(exec.Command("sh", "-c", pluginDetails.BuildDirectoryCommand))
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
	res, client := pluginClient.TestCaseCommandPluginCall(exec.Command("sh", "-c", pluginDetails.DetectTestCasesCommand))
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
	res, client := pluginClient.StaticAssetsPluginCall(exec.Command("sh", "-c", pluginDetails.StaticAssetsCommand))
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

func GetCommands(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) (*global.SeedCommandsOutput, *global.BuildCommandsOutput, *global.MigrationCommandsOutput, *global.StartUpCommandsOutput, *global.AdHocScriptsOutput) {
	response, client := pluginClient.CommandsPluginCall(exec.Command("sh", "-c", pluginDetails.Commands))
	defer func() {
		for client.Exited() {
			client.Kill()
			log.Println("client kill")
		}
	}()
	var err error
	serviceCommandsInput := &pluginPb.ServiceCommandsInput{
		Root:     input.Root,
		Language: input.RuntimeVersion,
	}
	detectAdHocScript, err := response.DetectAdHocScripts(serviceCommandsInput)
	if err != nil {
		utils.Logger(err)
		return nil, nil, nil, nil, nil
	}
	if detectAdHocScript.Error != nil {
		utils.Logger(detectAdHocScript.Error)
		detectAdHocScript.AdHocScripts = nil
	}
	detectSeedCommand, err := response.DetectSeedCommands(serviceCommandsInput)
	if err != nil {
		utils.Logger(err)
		return nil, nil, nil, nil, detectAdHocScript.AdHocScripts
	}
	if detectSeedCommand.Error != nil {
		utils.Logger(detectSeedCommand.Error)
		detectSeedCommand.SeedCommands = nil
	}
	detectBuildCommands, err := response.DetectBuildCommands(serviceCommandsInput)
	if err != nil {
		utils.Logger(err)
		return detectSeedCommand.SeedCommands, nil, nil, nil, detectAdHocScript.AdHocScripts
	}
	if detectBuildCommands.Error != nil {
		utils.Logger(detectBuildCommands.Error)
		detectBuildCommands.BuildCommands = nil
	}
	detectMigrationCommands, err := response.DetectMigrationCommands(serviceCommandsInput)
	if err != nil {
		utils.Logger(err)
		return detectSeedCommand.SeedCommands, detectBuildCommands.BuildCommands, nil, nil, detectAdHocScript.AdHocScripts
	}
	if detectMigrationCommands.Error != nil {
		utils.Logger(detectMigrationCommands.Error)
		detectMigrationCommands.MigrationCommands = nil
	}
	detectStartUpCommands, err := response.DetectStartUpCommands(serviceCommandsInput)
	if err != nil {
		utils.Logger(err)
		detectStartUpCommands.StartUpCommands = nil
	}
	if detectStartUpCommands.Error != nil {
		utils.Logger(detectStartUpCommands.Error)
		detectStartUpCommands.StartUpCommands = nil
	}
	return detectSeedCommand.SeedCommands, detectBuildCommands.BuildCommands, detectMigrationCommands.MigrationCommands, detectStartUpCommands.StartUpCommands, detectAdHocScript.AdHocScripts
}

func RunPreDetectCommand(ctx context.Context, input *pluginPb.ServiceInput, pluginDetails *versionsPB.LanguageVersion) {
	runtimeResponse, client := pluginClient.PreDetectCommandPluginCall(exec.Command("sh", "-c", pluginDetails.PreDetectCommand))
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
		dependenciesResponse, client := pluginClient.DependenciesPluginCall(exec.Command("sh", "-c", dependenciesCommand.Plugincommand))
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
