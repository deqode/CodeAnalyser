package testing

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"reflect"
	"testing"
)

func CheckDbEquality(got, output *languageSpecificPB.DBOutput, t *testing.T) {
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
}