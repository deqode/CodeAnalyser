package testing

import (
	"code-analyser/protos/pb/output/global"
	"reflect"
	"testing"
)

type DockerCaseOutput struct {
	DockerFile        *global.DockerFileOutput
	DockerComposeFile *global.DockerComposeFileOutput
}

func CheckProcfileEquality(got, output *global.ProcFileOutput, t *testing.T) {
	//for _, value := range got.Commands {
	//	sort.Strings(value.Args)
	//}
	//for _, value := range output.Commands {
	//	sort.Strings(value.Args)
	//}
	if !reflect.DeepEqual(got, output) {
		t.Error("expected this ", output, "\n got this ", got)
	}
}
