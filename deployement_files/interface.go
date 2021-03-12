package language_detectors

import (
	"code-analyser/protos/pb"
	"context"
)

type GlobalFile interface {
	GlobalDetectBuildCommands(context.Context, string) ([]*protos.BuildCommandsOutput, error)
	GlobalDetectStartUpCommands(context.Context, string) ([]*protos.StartUpCommandsOutput, error)
	GlobalDetectSeedCommands(context.Context, string) ([]*protos.SeedCommandsOutput, error)
	GlobalDetectMigrationCommands(context.Context, string) ([]*protos.MigrationCommandsOutput, error)

	DetectProcFile(context.Context, string)([]*protos.ProcFileOutput, error)
	DetectMakefiles(context.Context, string)([]*protos.MakefileOutput ,error)
	DetectDockerFiles(context.Context, string)([]*protos.DockerFileOutput, error)
	DetectDockerComposeFiles(context.Context, string)([]*protos.DockerComposeFileOutput ,error)
}
