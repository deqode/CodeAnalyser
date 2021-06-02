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
	DetectBuildCommands(inputString *plugin.ServiceCommandsInput) (*plugin.ServiceOutputDetectBuildCommands, error)
	DetectStartUpCommands(inputString *plugin.ServiceCommandsInput) (*plugin.ServiceOutputStartUpCommands, error)
	DetectSeedCommands(inputString *plugin.ServiceCommandsInput) (*plugin.ServiceOutputSeedCommands, error)
	DetectMigrationCommands(inputString *plugin.ServiceCommandsInput) (*plugin.ServiceMigrationCommands, error)
	DetectAdHocScripts(inputString *plugin.ServiceCommandsInput) (*plugin.ServiceAdHocScripts, error) //Todo:DetectAdHocScripts output proto
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
