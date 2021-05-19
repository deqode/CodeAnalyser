package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"reflect"
	"strconv"
	"testing"
)

func TestOrmParsing(t *testing.T) {
	for i, element := range OrmParsingCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := ParseOrmFromDependencies(input.DependenciesList, input.LangYamlObject)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestOrmRunner(t *testing.T) {
	for i, element := range OrmRunnerCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := OrmRunner(input.OrmList, input.RuntimeVersion, input.Root)
			gotMap := map[string]*languageSpecificPB.ORM{}
			for _, value := range got.Orms {
				gotMap[value.Name] = value
			}
			outputMap := map[string]*languageSpecificPB.ORM{}
			for _, value := range output.Orms {
				outputMap[value.Name] = value
			}
			if !reflect.DeepEqual(gotMap, outputMap) || output.Used != got.Used {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}
