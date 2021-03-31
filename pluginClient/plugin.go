package pluginClient

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/detectRuntime"
	"code-analyser/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"log"
	"os"
	"os/exec"
)

var PluginDispenserFramework = "framework"
var PluginDispenserDB = "db"
var PluginDispenserOrm = "orm"
var PluginDispenserDetectRuntime = "detectRuntime"

var PluginMap = map[string]plugin.Plugin{
/*	PluginDispenserFramework: &framework.FrameworkGRPCPlugin{},
	PluginDispenserDB: &db.DbGRPCPlugin{},
*/	PluginDispenserDetectRuntime: &detectRuntime.DetectRuntimeGRPCPlugin{},
	//PluginDispenserOrm: &orm.OrmGRPCPlugin{},
}

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var logger = hclog.New(&hclog.LoggerOptions{
	Name:   "plugin",
	Output: os.Stdout,
	Level:  hclog.Error,
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

func DbPluginCall(cmd *exec.Cmd, ) (interfaces.DbVersion, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserDB)
	return raw.(interfaces.DbVersion), client
}

func OrmPluginCall(cmd *exec.Cmd, ) (interfaces.ORMVersion, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserOrm)
	return raw.(interfaces.ORMVersion), client
}

func DetectRuntimePluginCall(cmd *exec.Cmd) (interfaces.DetectRunTime, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserDetectRuntime)
	return raw.(interfaces.DetectRunTime), client
}