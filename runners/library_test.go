package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"reflect"
	"strconv"
	"testing"
)

func TestLibraryParsing(t *testing.T) {
	for i, element := range LibrariesParsingCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := ParseLibraryFromDependencies(input.DependenciesList, input.LangYamlObject)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestLibraryRunnerRunner(t *testing.T) {
	for i, element := range LibraryRunnerCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := LibraryRunner(input.LibraryList, input.RuntimeVersion, input.Root)
			gotMap := map[string]*languageSpecificPB.LibraryOutput{}
			for _, value := range got {
				gotMap[value.Name] = value
			}
			outputMap := map[string]*languageSpecificPB.LibraryOutput{}
			for _, value := range output {
				outputMap[value.Name] = value
			}
			if !reflect.DeepEqual(gotMap, outputMap) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}
