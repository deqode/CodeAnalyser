package runners

import (
	"code-analyser/protos/pb/helpers"
	global "code-analyser/protos/pb/output/globalFiles"
	"code-analyser/protos/pb/pluginDetails"
	commonUtils "code-analyser/utils"
	"code-analyser/utils/testing"
	"golang.org/x/net/context"
)

type DockerCase struct {
	Input  GlobalRunnerInput
	Output testing.DockerCaseOutput
}

type GlobalRunnerInput struct {
	Ctx          context.Context
	Path         string
	GlobalPlugin *pluginDetails.GlobalPlugins
}

type ProcFileCase struct {
	Input  GlobalRunnerInput
	Output *global.ProcFile
}

type MakeFileCase struct {
	Input  GlobalRunnerInput
	Output *global.MakeFile
}

var DockerCases = []DockerCase{
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.RootDirPath() + "/testingRepos/detectDockerFile/repo1",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: testing.DockerCaseOutput{
			DockerFile: &global.DockerFile{
				Used:     true,
				FilePath: commonUtils.RootDirPath() + "/testingRepos/detectDockerFile/repo1/Dockerfile",
			},
			DockerComposeFile: &global.DockerCompose{
				Used:     true,
				FilePath: commonUtils.RootDirPath() + "/testingRepos/detectDockerFile/repo1/docker-compose.yml",
			},
		},
	},
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.RootDirPath() + "/testingRepos/detectDockerFile/repo2",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: testing.DockerCaseOutput{
			DockerFile: nil,
			DockerComposeFile: &global.DockerCompose{
				Used:     true,
				FilePath: commonUtils.RootDirPath() + "/testingRepos/detectDockerFile/repo2/docker-compose.yaml",
			},
		},
	},
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.RootDirPath() + "/testingRepos/emptyRepo",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: testing.DockerCaseOutput{},
	},
}

var ProcFileCases = []ProcFileCase{
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.RootDirPath() + "/testingRepos/emptyRepo",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: nil,
	},
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.RootDirPath() + "/testingRepos/detectProcfile/repo1",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: &global.ProcFile{
			Used:     true,
			FilePath: commonUtils.RootDirPath() + "/testingRepos/detectProcfile/repo1/Procfile",
			Commands: map[string]*helpers.Command{
				"web": {
					Command: "bundle",
					Args:    []string{"exec", "rackup"},
				},
				"worker": {
					Command: "rake",
					Args:    []string{"resque:work"},
				},
			},
		},
	},
}

var MakeFileCases = []MakeFileCase{
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.RootDirPath() + "/testingRepos/emptyRepo",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: nil,
	},
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.RootDirPath() + "/testingRepos/detectMakefile/repo1",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: &global.MakeFile{
			Used:     true,
			FilePath: commonUtils.RootDirPath() + "/testingRepos/detectMakefile/repo1/Makefile",
		},
	},
}
