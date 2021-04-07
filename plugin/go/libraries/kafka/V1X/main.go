package V1X

import (
	"code-analyser/pluginClient/pb"
	languageSpecificPB "code-analyser/protos/protos/outputs/languageSpecific"
)

type KafkaV1x struct {}

func (k *KafkaV1x) GetLibraries(input *pb.ServiceInput) (*languageSpecificPB.LibrariesOutput, error) {
	return &languageSpecificPB.LibrariesOutput{},nil
}
