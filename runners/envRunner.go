package runners

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	versionsPB "code-analyser/protos/protos/versions"
	"os/exec"
)

func DetectEnvs(yamlLangObject *versionsPB.LanguageVersion) *pb.ServiceOutputEnv{
	res, client := pluginClient.EnvPluginCall(exec.Command("sh", "-c", yamlLangObject.DetectenvCommand))
	defer client.Kill()
}
