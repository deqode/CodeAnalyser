package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

func DetectDockerAndComposeFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) (*global.DockerFile, *global.DockerCompose) {
	runtimeResponse, client := pluginClient.CreateDockerFileClient(utils.CallPluginCommand(globalPlugin.DockerFile))
	for client.Exited() {
		client.Kill()
	}
	detectDockerFile, err := runtimeResponse.DetectDockerFile(&pb.StringInput{Value: path})
	if err != nil {
		utils.Logger(err)
		return nil, nil
	}
	if detectDockerFile.Error != nil {
		utils.Logger(detectDockerFile.Error)
		return nil, nil
	}
	detectDockerComposeFile, err := runtimeResponse.DetectDockerComposeFile(&pb.StringInput{Value: path})
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

func DetectProcFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) *global.ProcFile {
	runtimeResponse, client := pluginClient.CreateProcFileClient(utils.CallPluginCommand(globalPlugin.ProcFile))
	for client.Exited() {
		client.Kill()
	}
	procFileOutput := &global.ProcFile{}
	detectProcFile, err := runtimeResponse.Detect(&pb.StringInput{Value: path})
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

func DetectMakeFile(ctx context.Context, path string, globalPlugin *versionsPB.GlobalPlugin) *global.MakeFile {
	runtimeResponse, client := pluginClient.CreateMakeFileClient(utils.CallPluginCommand(globalPlugin.MakeFile))
	for client.Exited() {
		client.Kill()
	}
	detectMakeFile, err := runtimeResponse.Detect(&pb.StringInput{Value: path})
	makeFileOutput := &global.MakeFile{}
	if err != nil {
		utils.Logger(err)
		return makeFileOutput
	}
	if detectMakeFile.Error != nil {
		utils.Logger(detectMakeFile.Error)
		return makeFileOutput
	}
	makeFileOutput = detectMakeFile.Makefile
	return makeFileOutput
}
