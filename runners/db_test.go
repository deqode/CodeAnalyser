package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"reflect"
	"strconv"
	"testing"
)

func TestDbParsing(t *testing.T) {
	for i, element := range DbCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := ParseDbFromDependencies(input.DependenciesList, input.LangYamlObject)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestDbRunner(t *testing.T) {
	for i, element := range DbRunnerCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := DbRunner(input.DbList, input.RuntimeVersion, input.Root)
			gotMap := map[string]*languageSpecificPB.DB{}
			for _, value := range got.Databases {
				gotMap[value.Name] = value
			}
			outputMap := map[string]*languageSpecificPB.DB{}
			for _, value := range output.Databases {
				outputMap[value.Name] = value
			}
			if !reflect.DeepEqual(gotMap, outputMap) || output.Used != got.Used {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}
