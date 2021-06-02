package runners

import (
	"reflect"
	"strconv"
	"testing"
)

func TestEnv(t *testing.T) {
	for i, element := range EnvCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := EnvDetectAndRunner(input.pluginDetails, input.RuntimeVersion, input.Root)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}
