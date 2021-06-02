package runners

import (
	"code-analyser/pluginClient"
	pbGlobal "code-analyser/protos/pb"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

func DetectDockerAndComposeFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) (*global.DockerFileOutput, *global.DockerComposeFileOutput) {
	runtimeResponse, client := pluginClient.DockerFilePluginCall(utils.CallPluginCommand(globalPlugin.DockerFile))
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

func DetectProcFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) *global.ProcFileOutput {
	runtimeResponse, client := pluginClient.ProcFilePluginCall(utils.CallPluginCommand(globalPlugin.ProcFile))
	for client.Exited() {
		client.Kill()
	}
	procFileOutput := &global.ProcFileOutput{}
	detectProcFile, err := runtimeResponse.Detect(&pb.ServiceInputString{Value: path})
	if err != nil {
		utils.Logger(err)
		return procFileOutput
	}
	if detectProcFile.Error != nil {
		utils.Logger(detectProcFile.Error)
		return procFileOutput
	}
	procFileOutput = detectProcFile.ProcFile
	return procFileOutput
}

func DetectMakeFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) *global.MakefileOutput {
	runtimeResponse, client := pluginClient.MakeFilePluginCall(utils.CallPluginCommand(globalPlugin.MakeFile))
	for client.Exited() {
		client.Kill()
	}
	detectMakeFile, err := runtimeResponse.Detect(&pb.ServiceInputString{Value: path})
	makeFileOutput := &global.MakefileOutput{}
	if err != nil {
		utils.Logger(err)
		return makeFileOutput
	}
	if detectMakeFile.Error != nil {
		utils.Logger(detectMakeFile.Error)
		return makeFileOutput
	}
	makeFileOutput = detectMakeFile.MakeFile
	return makeFileOutput
}

func DetectAndRunCommands(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) *pbGlobal.Commands {
	response, client := pluginClient.CommandsPluginCall(utils.CallPluginCommand(globalPlugin.Commands))
	defer func() {
		for client.Exited() {
			client.Kill()
		}
	}()
	var err error
	serviceCommandsInput := &pb.ServiceCommandsInput{
		Root:     path,
		Language: "",
	}
	commands := pbGlobal.Commands{
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
