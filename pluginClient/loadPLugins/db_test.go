package loadPLugins

import (
	"strconv"
	"testing"
)

func TestDbPlugin_Load(t *testing.T) {
	for i, element := range TestDbPluginLoadCases {
		input, output := element.Input, element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := &DbPlugin{Setting: &Set}
			got.Load(input)
			DeepEqualDbPluginLoad(output, got, t)
		})
	}
}

func DeepEqualDbPluginLoad(expected, got *DbPlugin, t *testing.T) {
	if len(got.Dbs) != len(expected.Dbs)||got.Setting!=expected.Setting {
		t.Error("db length mismatch, expected this ", expected, "\n got this ", got)
	}
	for name, gotVersion := range got.Dbs {
		if expectedVersion, ok := expected.Dbs[name]; ok {
			for gotVersionName, gotDetails := range gotVersion.Version {
				if expectedDetails, ok := expectedVersion.Version[gotVersionName]; ok {
					if len(gotDetails.Libraries) != len(expectedDetails.Libraries) || gotDetails.Client == nil || gotDetails.Methods == nil || gotDetails.Setting != expectedDetails.Setting {
						t.Error("expected this ", expected, "\n got this ", got)
					}
				} else {
					t.Error("expected this ", expected, "\n got this ", got)
				}
			}
		} else {
			t.Error("expected this ", expected, "\n got this ", got)
		}
	}
}

func TestDbPlugin_Extract(t *testing.T) {

}

func TestDbPlugin_Run(t *testing.T) {
}

func TestDbPluginDetails_Run(t *testing.T) {

}
