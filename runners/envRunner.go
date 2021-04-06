package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	"code-analyser/protos/pb/output/languageSpecific"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"os/exec"
)

func EnvDetectAndRunner(yamlLangObject *versionsPB.LanguageVersion, runtimeVersion, root string) *languageSpecific.EnvOutput {
	res, client := pluginClient.EnvPluginCall(exec.Command("sh", "-c", yamlLangObject.DetectEnvCommand))
	defer client.Kill()
	isUsed, err := res.IsUsed(&pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if isUsed.Value == true {
		detection, err := res.Detect(&pb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           root,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		return &languageSpecific.EnvOutput{
			EnvUsed:      isUsed.Value,
			EnvKeyValues: detection.EnvKeyValues,
			Variables:    detection.Keys,
		}
	}
	return nil
}
