package pluginClient

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/framework"
	"code-analyser/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"log"
	"os"
	"os/exec"
)

var PluginDispenserFramework = "framework"

var PluginMap = map[string]plugin.Plugin{
	PluginDispenserFramework: &framework.FrameworkGRPCPlugin{},
}

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}
var logger = hclog.New(&hclog.LoggerOptions{
	Name:   "plugin",
	Output: os.Stdout,
	Level:  hclog.Debug,
})

func makeClient(cmd *exec.Cmd, pluginDispenser string, ) (interface{}, *plugin.Client) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  HandshakeConfig,
		Plugins:          PluginMap,
		Cmd:              cmd,
		Logger:           logger,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	rpcClient, err := client.Client()
	if err != nil {
		utils.Logger(err)
	}
	raw, err := rpcClient.Dispense(pluginDispenser)
	if err != nil {
		log.Fatal(err)
	}
	return raw, client
}

func FrameworkPluginCall(cmd *exec.Cmd, ) (interfaces.FrameworkVersions, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserFramework)
	return raw.(interfaces.FrameworkVersions), client
}

func ShutDown(client *plugin.Client) {
	client.Kill()
}
