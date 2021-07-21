package GlobalFiles

import (
	"code-analyser/protos/pb/helpers"
	"code-analyser/protos/pb/output/globalFiles"
)

//ProcFile plugin methods
type ProcFile interface {
	Detect(path *helpers.StringInput) (*globalFiles.ProcFile, error)
}

//Makefile plugin methods
type Makefile interface {
	Detect(path *helpers.StringInput) (*globalFiles.MakeFile, error)
}

//DockerFile plugin methods
type DockerFile interface {
	DetectDockerFile(path *helpers.StringInput) (*globalFiles.DockerFile, error)
	DetectDockerComposeFile(path *helpers.StringInput) (*globalFiles.DockerCompose, error)
}
