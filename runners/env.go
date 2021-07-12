package runners

import (
	"code-analyser/pluginClient"
	pb "code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/languageSpecific"
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

	return envVariables
}
