package interfaces

import "code-analyser/protos/pb/plugin"

type Commands interface {
	DetectBuildCommands(path *plugin.StringInput) (*plugin.CommandsOutput, error)
	DetectStartUpCommands(path *plugin.StringInput) (*plugin.CommandsOutput, error)
	DetectSeedCommands(path *plugin.StringInput) (*plugin.CommandsOutput, error)
	DetectMigrationCommands(path *plugin.StringInput) (*plugin.CommandsOutput, error)
	DetectAdHocScripts(path *plugin.StringInput) (*plugin.CommandsOutput, error) //Todo:DetectAdHocScripts output proto
}