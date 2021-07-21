package interfaces

import "code-analyser/protos/pb/helpers"
import "code-analyser/protos/pb/output/languageSpecific"

//Commands plugin methods
type Commands interface {
	DetectBuildCommands(path *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectStartUpCommands(path *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectSeedCommands(path *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectMigrationCommands(path *helpers.StringInput) (*languageSpecific.CommandOutput, error)
	DetectAdHocScripts(path *helpers.StringInput) (*languageSpecific.CommandOutput, error) //Todo:DetectAdHocScripts output proto
}
