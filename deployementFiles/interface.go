package deployementFiles

import (
	"code-analyser/protos/protos"
	"context"
)

//GlobalFile to be implemented for global file detector
type GlobalFile interface {
	Commands
	DetectProcFile(context.Context, string) ([]*protos.ProcFileOutput, error)
	DetectMakefiles(context.Context, string) ([]*protos.MakefileOutput, error)
	DetectDockerFiles(context.Context, string) ([]*protos.DockerFileOutput, error)
	DetectDockerComposeFiles(context.Context, string) ([]*protos.DockerComposeFileOutput, error)
}

// Commands Various different commands
type Commands interface {
	DetectBuildCommands(context.Context, string) ([]*protos.BuildCommandsOutput, error)
	DetectStartUpCommands(context.Context, string) ([]*protos.StartUpCommandsOutput, error)
	DetectSeedCommands(context.Context, string) ([]*protos.SeedCommandsOutput, error)
	DetectMigrationCommands(context.Context, string) ([]*protos.MigrationCommandsOutput, error)
	DetectAdHocScripts(context.Context, string) (interface{}, error) //Todo: output proto
}
