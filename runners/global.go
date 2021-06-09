package runners

import (
	"code-analyser/pluginClient"
	pb "code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

func ExecuteDockerAndComposePlugin(ctx context.Context, projectRootPath, pluginPath string) (*global.DockerFile, *global.DockerCompose) {
	pluginCall, _ := pluginClient.CreateDockerFileClient(utils.CallPluginCommand(pluginPath))

	dockerFile, err := pluginCall.DetectDockerFile(&pb.StringInput{Value: projectRootPath})
	if err != nil || dockerFile.Error != nil {
		utils.Logger(err, dockerFile)
		return nil, nil
	}

	dockerCompose, err := pluginCall.DetectDockerComposeFile(&pb.StringInput{Value: projectRootPath})
	if err != nil || dockerCompose.Error != nil {
		utils.Logger(err, dockerCompose)
		return dockerFile, nil
	}

	return dockerFile, dockerCompose
}

func ExecuteProcFileDetectionPlugin(ctx context.Context, projectRootPath, pluginPath string) *global.ProcFile {
	pluginCall, _ := pluginClient.CreateProcFileClient(utils.CallPluginCommand(pluginPath))

	procFile, err := pluginCall.Detect(&pb.StringInput{Value: projectRootPath})
	if err != nil || procFile.Error != nil {
		utils.Logger(err, procFile)
		return nil
	}

	return procFile
}

func ExecuteMakeFileDetectionPlugin(ctx context.Context, projectRootPath, pluginPath string) *global.MakeFile {
	pluginCall, _ := pluginClient.CreateMakeFileClient(utils.CallPluginCommand(pluginPath))

	makeFile, err := pluginCall.Detect(&pb.StringInput{Value: projectRootPath})
	if err != nil ||makeFile.Error != nil{
		utils.Logger(err,makeFile)
		return nil
	}

	return makeFile
}
