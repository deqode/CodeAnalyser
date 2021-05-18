package runners

import (
	"reflect"
	"strconv"
	"testing"
)

func TestOrmParsing(t *testing.T) {
	for i, element := range OrmCases {
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