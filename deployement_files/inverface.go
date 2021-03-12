package language_detectors

import (
	"code-analyser/protos/pb"
	"context"
)

type GlobalFile interface {
	GlobalDetectBuildCommands(context.Context, string) []*protos.BuildCommandsOutput
	GlobalDetectStartUpCommands(context.Context, string) []*protos.StartUpCommandsOutput
	GlobalDetectSeedCommands(context.Context, string) []*protos.SeedCommandsOutput
	GlobalDetectMigrationCommands(context.Context, string) []*protos.MigrationCommandsOutput

	DetectProcFile(context.Context, string)
	DetectMakefiles(context.Context, string)
	DetectDockerFiles(context.Context, string)
	DetectDockerComposeFiles(context.Context, string)
}
