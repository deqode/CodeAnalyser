package GlobalFiles

import (
	"code-analyser/protos/pb/plugin"
)

type ProcFile interface {
	Detect(path *plugin.StringInput) (*plugin.ProcFileOutput, error)
}

type Makefile interface {
	Detect(path *plugin.StringInput) (*plugin.MakeFileOutput, error)
}

type DockerFile interface {
	DetectDockerFile(path *plugin.StringInput) (*plugin.DockerFileOutput, error)
	DetectDockerComposeFile(path *plugin.StringInput) (*plugin.DockerComposeFileOutput, error)
}
