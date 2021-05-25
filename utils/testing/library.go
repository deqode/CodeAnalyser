package testing

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"reflect"
	"testing"
)

func CheckLibraryEquality(got, output []*languageSpecificPB.LibraryOutput, t *testing.T) {
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
}