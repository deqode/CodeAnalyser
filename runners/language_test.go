package runners

import (
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

func TestGetCommands(t *testing.T) {

}

func TestDetectTestCasesCommand(t *testing.T) {

}
