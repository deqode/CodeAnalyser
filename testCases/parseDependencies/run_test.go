package parseDependencies

import (
	"code-analyser/runners"
	"reflect"
	"testing"
)

func TestDependencies(t *testing.T) {
	for _, element := range Cases {
		input := element.Input
		got := runners.GetParsedDependencis(input.Ctx, input.LanguageVersion, input.Path, input.PluginDetails)
		if !reflect.DeepEqual(got, element.Output) {
         t.Error("failed at __________")
		}
	}

}
