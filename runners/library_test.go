package runners

import (
	testingUtil "code-analyser/utils/testing"
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
			testingUtil.CheckLibraryEquality(got, output, t)
		})
	}
}


