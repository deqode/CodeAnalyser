package pluginClient

import (
	"code-analyser/GlobalFiles"
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient/buildDirectory"
	"code-analyser/pluginClient/commands"
	"code-analyser/pluginClient/db"
	"code-analyser/pluginClient/dependencies"
	"code-analyser/pluginClient/detectRuntime"
	dockerFile "code-analyser/pluginClient/docker"
	"code-analyser/pluginClient/env"
	"code-analyser/pluginClient/framework"
	"code-analyser/pluginClient/library"
	"code-analyser/pluginClient/makeFile"
	"code-analyser/pluginClient/orm"
	"code-analyser/pluginClient/preDetectCommand"
	"code-analyser/pluginClient/procFile"
	"code-analyser/pluginClient/staticAssets"
	"code-analyser/pluginClient/testCasesCommands"
	"code-analyser/utils"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"log"
	"os"
	"os/exec"
)

const (
	//Framework is a string framework
	Framework = "framework"
	//DB is a string db
	DB = "db"
	//Orm is a string orm
	Orm = "orm"
	//DetectRuntime  is a detectRuntime
	DetectRuntime = "detectRuntime"
	//Dependencies  is a getDependencies
	Dependencies = "getDependencies"
	Library      = "library"
	//Env is a string env
	Env              = "env"
	PreDetectCommand = "preDetectCommand"
	StaticAssets     = "staticAssets"
	BuildDirectory   = "buildDirectory"
	TestCaseCommand  = "testCaseCommand"
	ProcFile         = "procFile"
	MakeFile         = "makeFile"
	DockerFile       = "dockerFile"
	Command          = "command"
)

//PluginMap is a map of dispenser string and plugin
var PluginMap = map[string]plugin.Plugin{
	Library:          &library.GRPCPlugin{},
	Framework:        &framework.GRPCPlugin{},
	DB:               &db.GRPCPlugin{},
	DetectRuntime:    &detectRuntime.GRPCPlugin{},
	Dependencies:     &dependencies.GRPCPlugin{},
	Orm:              &orm.GRPCPlugin{},
	Env:              &env.GRPCPlugin{},
	PreDetectCommand: &preDetectCommand.GRPCPlugin{},
	StaticAssets:     &staticAssets.GRPCPlugin{},
	BuildDirectory:   &buildDirectory.GRPCPlugin{},
	TestCaseCommand:  &testCasesCommands.GRPCPlugin{},
	ProcFile:         &procFile.GRPCPlugin{},
	MakeFile:         &makeFile.GRPCPlugin{},
	DockerFile:       &dockerFile.GRPCPlugin{},
	Command:          &commands.GRPCPlugin{},
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
	Level:  hclog.Debug,
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

//CreateFrameworkClient generate client for framework plugin calls
func CreateFrameworkClient(cmd *exec.Cmd) (interfaces.Framework, *plugin.Client) {
	raw, client := makeClient(cmd, Framework)
	return raw.(interfaces.Framework), client
}

//CreateOrmClient generate client for orm plugin calls
func CreateOrmClient(cmd *exec.Cmd) (interfaces.ORMVersion, *plugin.Client) {
	raw, client := makeClient(cmd, Orm)
	return raw.(interfaces.ORMVersion), client
}

//CreateDetectRuntimeClient generate client for runtime detection plugin calls
func CreateDetectRuntimeClient(cmd *exec.Cmd) (interfaces.DetectRunTime, *plugin.Client) {
	raw, client := makeClient(cmd, DetectRuntime)
	return raw.(interfaces.DetectRunTime), client
}

//CreateDependenciesClient generate client for dependencies plugin calls
func CreateDependenciesClient(cmd *exec.Cmd) (interfaces.Dependencies, *plugin.Client) {
	raw, client := makeClient(cmd, Dependencies)
	return raw.(interfaces.Dependencies), client
}

//CreateDbClient generate client for db plugin calls
func CreateDbClient(cmd *exec.Cmd) (interfaces.Db, *plugin.Client) {
	raw, client := makeClient(cmd, DB)
	return raw.(interfaces.Db), client
}

//CreateEnvClient generate client for env variables plugin calls
func CreateEnvClient(cmd *exec.Cmd) (interfaces.Env, *plugin.Client) {
	raw, client := makeClient(cmd, Env)
	return raw.(interfaces.Env), client
}

//CreateLibraryClient generate client for library plugin calls
func CreateLibraryClient(cmd *exec.Cmd) (interfaces.Library, *plugin.Client) {
	raw, client := makeClient(cmd, Library)
	return raw.(interfaces.Library), client
}

//CreatePreDetectCommandClient generate client for framework plugin calls
func CreatePreDetectCommandClient(cmd *exec.Cmd) (interfaces.PreDetectCommands, *plugin.Client) {
	raw, client := makeClient(cmd, PreDetectCommand)
	return raw.(interfaces.PreDetectCommands), client
}

//CreateStaticAssetsClient generate client for static assets plugin calls
func CreateStaticAssetsClient(cmd *exec.Cmd) (interfaces.StaticAssets, *plugin.Client) {
	raw, client := makeClient(cmd, StaticAssets)
	return raw.(interfaces.StaticAssets), client
}

//CreateBuildDirectoryClient generate client for build directory plugin calls
func CreateBuildDirectoryClient(cmd *exec.Cmd) (interfaces.BuildDirectory, *plugin.Client) {
	raw, client := makeClient(cmd, BuildDirectory)
	return raw.(interfaces.BuildDirectory), client
}

//CreateTestCaseCommandClient generate client for test commands plugin calls
func CreateTestCaseCommandClient(cmd *exec.Cmd) (interfaces.TestCasesRunCommands, *plugin.Client) {
	raw, client := makeClient(cmd, TestCaseCommand)
	return raw.(interfaces.TestCasesRunCommands), client
}

//CreateProcFileClient generate client for procfile plugin calls
func CreateProcFileClient(cmd *exec.Cmd) (GlobalFiles.ProcFile, *plugin.Client) {
	raw, client := makeClient(cmd, ProcFile)
	return raw.(GlobalFiles.ProcFile), client
}

//CreateMakeFileClient generate client for makefile plugin calls
func CreateMakeFileClient(cmd *exec.Cmd) (GlobalFiles.Makefile, *plugin.Client) {
	raw, client := makeClient(cmd, MakeFile)
	return raw.(GlobalFiles.Makefile), client
}

//CreateDockerFileClient generate client for dockerfile plugin calls
func CreateDockerFileClient(cmd *exec.Cmd) (GlobalFiles.DockerFile, *plugin.Client) {
	raw, client := makeClient(cmd, DockerFile)
	return raw.(GlobalFiles.DockerFile), client
}

//CreateCommandsClient generate client for commands plugin calls
func CreateCommandsClient(cmd *exec.Cmd) (GlobalFiles.Commands, *plugin.Client) {
	raw, client := makeClient(cmd, Command)
	return raw.(GlobalFiles.Commands), client
}
