package main

import (
	decisionmakerPB "code-analyser/protos/pb"
	"code-analyser/protos/pb/output/global"
	commonUtils "code-analyser/utils"
	utilTesting "code-analyser/utils/testing"
	"reflect"
	"strconv"
	"testing"
)

type GlobalDetectionCase struct {
	Input  string
	Output *decisionmakerPB.GlobalDetections
}

var GlobalDetectionCases = []GlobalDetectionCase{
	{
		Input: commonUtils.RootDirPath() + "/testingRepos/globalDetection/repo1",
		Output: &decisionmakerPB.GlobalDetections{
			ProcFile: &global.ProcFileOutput{
				Used:     true,
				FilePath: commonUtils.RootDirPath() + "/testingRepos/globalDetection/repo1/Procfile",
				Commands: map[string]*global.Command{
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
			Makefile: &global.MakefileOutput{
				Used:     true,
				FilePath: commonUtils.RootDirPath() + "/testingRepos/globalDetection/repo1/Makefile",
			},
			DockerFile: &global.DockerFileOutput{
				Used:     true,
				FilePath: commonUtils.RootDirPath() + "/testingRepos/globalDetection/repo1/Dockerfile",
			},
			DockerComposeFile: &global.DockerComposeFileOutput{
				Used:     true,
				FilePath: commonUtils.RootDirPath() + "/testingRepos/globalDetection/repo1/docker-compose.yml",
			},
		},
	},
	{
		Input:  commonUtils.RootDirPath() + "/testingRepos/emptyRepo",
		Output: &decisionmakerPB.GlobalDetections{},
	},
}

func TestGlobalDetection(t *testing.T) {
	for i, element := range GlobalDetectionCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := Scrape(input).GloabalDetections
			t.Run("procfile", func(t *testing.T) {
				t.Parallel()
				utilTesting.CheckProcfileEquality(got.ProcFile, output.ProcFile, t)
			})
			t.Run("makefile", func(t *testing.T) {
				t.Parallel()
				if !reflect.DeepEqual(got.Makefile, output.Makefile) {
					t.Error("expected this ", output.Makefile, "\n got this ", got.Makefile)
				}
			})
			t.Run("dockerComposeFile", func(t *testing.T) {
				t.Parallel()
				if !reflect.DeepEqual(got.DockerComposeFile, output.DockerComposeFile) {
					t.Error("expected this ", output.DockerComposeFile, "\n got this ", got.DockerComposeFile)
				}
			})
			t.Run("dockerFile", func(t *testing.T) {
				t.Parallel()
				if !reflect.DeepEqual(got.DockerFile, output.DockerFile) {
					t.Error("expected this ", output.DockerFile, "\n got this ", got.DockerFile)
				}
			})
		})
	}
}
