package testing

import (
	global "code-analyser/protos/pb/output/globalFiles"
	"reflect"
	"sort"
	"testing"
)

type DockerCaseOutput struct {
	DockerFile        *global.DockerFile
	DockerComposeFile *global.DockerCompose
}

func CheckProcfileEquality(got, output *global.ProcFile, t *testing.T) {
	if got == nil && output == nil {
		return
	}
	for _, value := range got.Commands {
		sort.Strings(value.Args)
	}
	for _, value := range output.Commands {
		sort.Strings(value.Args)
	}
	commandsEqual := true
	for key, _ := range got.Commands {
		if got.Commands[key].Command != output.Commands[key].Command || !reflect.DeepEqual(got.Commands[key].Args, output.Commands[key].Args) {
			commandsEqual = false
			break
		}
	}
	if got.Used != output.Used || got.FilePath != output.FilePath || len(got.Commands) != len(output.Commands) || !commandsEqual {
		t.Error("expected this ", output, "\n got this ", got)
	}
}
