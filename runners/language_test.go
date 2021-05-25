package runners

import (
	testingUtil "code-analyser/utils/testing"
	"reflect"
	"strconv"
	"testing"
)

func TestDependencies(t *testing.T) {
	for i, element := range DependencyCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := GetParsedDependencies(input.Ctx, input.LanguageVersion, input.Path, input.PluginDetails)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestDetectRuntime(t *testing.T) {
	for i, element := range DetectRunTimeCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := DetectRuntime(input.Ctx, input.Path, input.YamlLangObject)
			if got != output {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestCommands(t *testing.T) {
	for i, element := range DetectCommandCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := GetCommands(input.Ctx, input.Input, input.PluginDetails)
			testingUtil.AllCommandTesting(t, got, output)
		})

	}
}

func TestDetectTestCasesCommand(t *testing.T) {
	for i, element := range DetectTestCommandCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := DetectTestCasesCommand(input.Ctx, input.Input, input.PluginDetails)
			testingUtil.CheckTestCaseCommandEquality(got, output, t)
		})
	}
}
