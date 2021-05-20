package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"os/exec"
)

//EnvDetectAndRunner will run to find Frameworks & returns its detectors.
func EnvDetectAndRunner(pluginDetails *versionsPB.LanguageVersion, runtimeVersion, root string) *languageSpecific.EnvOutput {
	res, client := pluginClient.EnvPluginCall(exec.Command("sh", "-c", pluginDetails.DetectEnvCommand))

	for client.Exited() {
		client.Kill()
	}

	detection, err := res.Detect(&pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	if err != nil || detection.Error != nil {
		utils.Logger(err, detection.Error)
		return nil
	}

	envOutput := languageSpecific.EnvOutput{
		EnvUsed: false,
	}
	if detection.EnvKeyValues != nil || detection.Keys != nil {
		envOutput.EnvUsed = true
		envOutput.EnvKeyValues = detection.EnvKeyValues
		envOutput.Variables = detection.Keys
	}
	return &envOutput
}
