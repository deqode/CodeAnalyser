package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"golang.org/x/net/context"
	"os/exec"
)

func DetectDockerAndComposeFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) (*global.DockerFileOutput, *global.DockerComposeFileOutput) {
	runtimeResponse, client := pluginClient.DockerFilePluginCall(exec.Command("sh", "-c", globalPlugin.DockerFile))
	for client.Exited() {
		client.Kill()
	}
	detectDockerFile, err := runtimeResponse.DetectDockerFiles(&pb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return nil, nil
	}
	if detectDockerFile.Error != nil {
		utils.Logger(detectDockerFile.Error)
		return nil, nil
	}
	detectDockerComposeFile, err := runtimeResponse.DetectDockerComposeFiles(&pb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return nil, nil
	}
	if detectDockerComposeFile.Error != nil {
		utils.Logger(detectDockerComposeFile.Error)
		return nil, nil
	}
	return detectDockerFile.DockerFile, detectDockerComposeFile.DockerComposeFile
}

func DetectAndRunProcFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) *global.ProcFileOutput {
	runtimeResponse, client := pluginClient.ProcFilePluginCall(exec.Command("sh", "-c", globalPlugin.ProcFile))
	for client.Exited() {
		client.Kill()
	}
	detectProcFile, err := runtimeResponse.Detect(&pb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if detectProcFile.Error != nil {
		utils.Logger(detectProcFile.Error)
		return nil
	}
	return detectProcFile.ProcFile
}

func DetectAndRunMakeFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) *global.MakefileOutput {
	runtimeResponse, client := pluginClient.MakeFilePluginCall(exec.Command("sh", "-c", globalPlugin.MakeFile))
	for client.Exited() {
		client.Kill()
	}
	detectMakeFile, err := runtimeResponse.Detect(&pb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if detectMakeFile.Error != nil {
		utils.Logger(detectMakeFile.Error)
		return nil
	}
	return detectMakeFile.MakeFile
}

func DetectAndRunCommands(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) (*global.SeedCommandsOutput, *global.BuildCommandsOutput, *global.MigrationCommandsOutput, *global.StartUpCommandsOutput, *global.AdHocScriptsOutput) {
	response, client := pluginClient.CommandsPluginCall(exec.Command("sh", "-c", globalPlugin.Commands))
	for client.Exited() {
		client.Kill()
	}
	var err error
	serviceCommandsInput := &pb.ServiceCommandsInput{
		Root:     path,
		Language: "",
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
