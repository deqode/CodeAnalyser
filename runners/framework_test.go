package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"reflect"
	"strconv"
	"testing"
)

func TestFrameworkParsing(t *testing.T) {
	for i, element := range FrameworkParsingCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := ParseFrameworkFromDependencies(input.DependenciesList, input.LangYamlObject)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestFrameworkRunner(t *testing.T) {
	for i, element := range FrameworkRunnerCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := FrameworkRunner(input.FrameworkList, input.RuntimeVersion, input.Root)
			gotMap := map[string]*languageSpecificPB.FrameworkOutput{}
			for _, value := range got {
				gotMap[value.Name] = value
			}
			outputMap := map[string]*languageSpecificPB.FrameworkOutput{}
			for _, value := range output {
				if value.Used == true {
					outputMap[value.Name] = value
				}
			}
			if !reflect.DeepEqual(gotMap, outputMap) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}
