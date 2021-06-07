package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

//ExecuteEnvsDetectionPlugin will run to find Frameworks & returns its detectors.
func ExecuteEnvsDetectionPlugin(ctx context.Context, envPluginPath, runtimeVersion, projectRootPath string) *languageSpecific.Envs {
	pluginCall, _ := pluginClient.CreateEnvClient(utils.CallPluginCommand(envPluginPath))

	response, err := pluginCall.Detect(&pb.Input{
		RuntimeVersion: runtimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil || response.Error != nil {
		utils.Logger(err, response)
		return nil
	}

	envOutput := languageSpecific.Envs{
		Used: false,
	}
	if response.Envs.EnvKeyValues != nil || response.Envs.Keys != nil {
		envOutput.Used = true
		envOutput.EnvKeyValues = response.Envs.EnvKeyValues
		envOutput.Keys = response.Envs.Keys
	}
	return &envOutput
}
