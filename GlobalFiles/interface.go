package GlobalFiles

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/globalFiles"
)

type ProcFile interface {
	Detect(path *helpers.StringInput) (*globalFiles.ProcFile, error)
}

type Makefile interface {
	Detect(path *helpers.StringInput) (*globalFiles.MakeFile, error)
}

type DockerFile interface {
	DetectDockerFile(path *helpers.StringInput) (*globalFiles.DockerFile, error)
	DetectDockerComposeFile(path *helpers.StringInput) (*globalFiles.DockerCompose, error)
}
