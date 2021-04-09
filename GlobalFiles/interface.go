package GlobalFiles

import (
	globalPB "code-analyser/protos/pb/output/global"
	"code-analyser/protos/pb/plugin"
	"context"
)

//GlobalFile to be implemented for global file detector
type GlobalFile interface {
	Commands
	DetectProcFile(context.Context, string) (globalPB.ProcFileOutput, error)
	DetectMakefiles(context.Context, string) ([]*globalPB.MakefileOutput, error)
	DetectDockerFiles(context.Context, string) ([]*globalPB.DockerFileOutput, error)
	DetectDockerComposeFiles(context.Context, string) ([]*globalPB.DockerComposeFileOutput, error)
}

// Commands Various different commands
type Commands interface {
	DetectBuildCommands(context.Context, string) ([]*globalPB.BuildCommandsOutput, error)
	DetectStartUpCommands(context.Context, string) ([]*globalPB.StartUpCommandsOutput, error)
	DetectSeedCommands(context.Context, string) ([]*globalPB.SeedCommandsOutput, error)
	DetectMigrationCommands(context.Context, string) ([]*globalPB.MigrationCommandsOutput, error)
	DetectAdHocScripts(context.Context, string) (interface{}, error) //Todo: output proto
}

type ProcFile interface {
	Detect(inputString *plugin.ServiceInputString) (*plugin.ServiceOutputProcFile, error)
}
type Makefiles interface {
	Detect(inputString *plugin.ServiceInputString) (*plugin.ServiceOutputMakeFile, error)
}

type DockerFile interface {
	DetectDockerFiles(inputString *plugin.ServiceInputString) (*plugin.ServiceOutputDockerFile, error)
	DetectDockerComposeFiles(inputString *plugin.ServiceInputString) (*plugin.ServiceOutputDockerComposeFile, error)
}
