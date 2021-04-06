package pluginClient

import (
	"code-analyser/language_detectors/interfaces"
	"code-analyser/pluginClient/db"
	"code-analyser/pluginClient/dependencies"
	"code-analyser/pluginClient/detectRuntime"
	"code-analyser/pluginClient/env"
	"code-analyser/pluginClient/framework"
	"code-analyser/pluginClient/orm"
	"code-analyser/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"log"
	"os"
	"os/exec"
)

const (
	//PluginDispenserFramework is a string framework
	PluginDispenserFramework = "framework"
	//PluginDispenserDB is a string db
	PluginDispenserDB = "db"
	//PluginDispenserOrm is a string orm
	PluginDispenserOrm = "orm"
	//PluginDispenserDetectRuntime  is a detectRuntime
	PluginDispenserDetectRuntime = "detectRuntime"
	//PluginDispenserDependencies  is a getDependencies
	PluginDispenserDependencies = "getDependencies"
	//PluginDispenserEnv is a string env
	PluginDispenserEnv = "env"
)

//PluginMap is a map of dispenser string and plugin
var PluginMap = map[string]plugin.Plugin{
	PluginDispenserFramework:     &framework.FrameworkGRPCPlugin{},
	PluginDispenserDB:            &db.DbGRPCPlugin{},
	PluginDispenserDetectRuntime: &detectRuntime.DetectRuntimeGRPCPlugin{},
	PluginDispenserDependencies:  &dependencies.DependenciesGRPCPlugin{},
	PluginDispenserOrm:           &orm.OrmGRPCPlugin{},
	PluginDispenserEnv:           &env.EnvGRPCPlugin{},
}

//HandshakeConfig stores the config for plugin
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion: 1,
	MagicCookieKey:  "BASIC_PLUGIN",
	//todo: need to make this in environment variable
	MagicCookieValue: "hello",
}

var logger = hclog.New(&hclog.LoggerOptions{
	Name:   "plugin",
	Output: os.Stdout,
	Level:  hclog.Error,
})

func makeClient(cmd *exec.Cmd, pluginDispenser string) (interface{}, *plugin.Client) {
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

//FrameworkPluginCall call th plugin related to DetectRuntime framework
func FrameworkPluginCall(cmd *exec.Cmd) (interfaces.FrameworkVersions, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserFramework)
	return raw.(interfaces.FrameworkVersions), client
}

//OrmPluginCall call th plugin related to DetectRuntime ORM
func OrmPluginCall(cmd *exec.Cmd) (interfaces.ORMVersion, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserOrm)
	return raw.(interfaces.ORMVersion), client
}

//DetectRuntimePluginCall call th plugin related to DetectRuntime
func DetectRuntimePluginCall(cmd *exec.Cmd) (interfaces.DetectRunTime, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserDetectRuntime)
	return raw.(interfaces.DetectRunTime), client
}

//DependenciesPluginCall call th plugin related to Dependencies
func DependenciesPluginCall(cmd *exec.Cmd) (interfaces.Dependencies, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserDependencies)
	return raw.(interfaces.Dependencies), client
}

//DbPluginCall call th plugin related to DB
func DbPluginCall(cmd *exec.Cmd) (interfaces.DbVersion, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserDB)
	return raw.(interfaces.DbVersion), client
}

func EnvPluginCall(cmd *exec.Cmd) (interfaces.Env, *plugin.Client) {
	raw, client := makeClient(cmd, PluginDispenserEnv)
	return raw.(interfaces.Env), client
}
