package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

//ExecuteEnvsDetectionPlugin will run to find Frameworks & returns its detectors.
func ExecuteEnvsDetectionPlugin(ctx context.Context, pluginPath, runtimeVersion, projectRootPath string) *languageSpecific.Envs {
	pluginCall, _ := pluginClient.CreateEnvClient(utils.CallPluginCommand(pluginPath))

	envVariables, err := pluginCall.Detect(&pb.Input{
		RuntimeVersion: runtimeVersion,
		RootPath:       projectRootPath,
	})
	if err != nil || envVariables.Error != nil {
		utils.Logger(err, envVariables)
		return nil
	}

	envOutput := languageSpecific.Envs{
		Used: false,
	}
	if envVariables.Value.EnvKeyValues != nil || envVariables.Value.Keys != nil {
		envOutput.Used = true
		envOutput.EnvKeyValues = envVariables.Value.EnvKeyValues
		envOutput.Keys = envVariables.Value.Keys
	}
	return &envOutput
}
