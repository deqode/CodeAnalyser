package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/pb/output/languageSpecific"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"os/exec"
)
//EnvDetectAndRunner will run to find Frameworks & returns its detectors.
func EnvDetectAndRunner(pluginDetails *versionsPB.LanguageVersion, runtimeVersion, root string) *languageSpecific.EnvOutput {
	res, client := pluginClient.EnvPluginCall(exec.Command("sh", "-c", pluginDetails.DetectEnvCommand))
	defer client.Kill()

	detection, err := res.Detect(&pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if detection.EnvKeyValues != nil || detection.Keys != nil {

		return &languageSpecific.EnvOutput{
			EnvUsed:      true,
			EnvKeyValues: detection.EnvKeyValues,
			Variables:    detection.Keys,
		}
	}
	return nil
}
