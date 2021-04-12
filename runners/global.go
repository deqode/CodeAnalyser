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
	defer client.Kill()
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
	defer client.Kill()
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
	defer client.Kill()
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
