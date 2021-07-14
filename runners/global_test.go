package runners

import (
	utilTesting "code-analyser/utils/testing"
	"reflect"
	"strconv"
	"testing"
)

func TestDetectDockerAndComposeFile(t *testing.T) {
	for i, element := range DockerCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := utilTesting.DockerCaseOutput{}
			got.DockerFile, got.DockerComposeFile = ExecuteDockerAndComposePlugin(input.Ctx, input.Path, input.GlobalPlugin)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestProcfile(t *testing.T) {
	for i, element := range ProcFileCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := ExecuteProcFileDetectionPlugin(input.Ctx, input.Path, input.GlobalPlugin)
			utilTesting.CheckProcfileEquality(got, output, t)
		})
	}
}

func TestMakefile(t *testing.T) {
	for i, element := range MakeFileCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := ExecuteMakeFileDetectionPlugin(input.Ctx, input.Path, input.GlobalPlugin)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}
