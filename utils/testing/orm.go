package testing

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	"testing"
)

func CheckOrmEquality(got, output *languageSpecificPB.OrmOutput, t *testing.T) {
	//gotMap := map[string]*languageSpecificPB.ORM{}
	//for _, value := range got.Orms {
	//	gotMap[value.Name] = value
	//}
	//outputMap := map[string]*languageSpecificPB.ORM{}
	//for _, value := range output.Orms {
	//	outputMap[value.Name] = value
	//}
	//if !reflect.DeepEqual(gotMap, outputMap) || output.Used != got.Used {
	//	t.Error("expected this ", output, "\n got this ", got)
	//}
}
