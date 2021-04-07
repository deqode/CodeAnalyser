package interfaces

import (
	"code-analyser/pluginClient/pb"
	languagePB "code-analyser/protos/protos/outputs/languageSpecific"
)

type Libraries interface {
	GetLibraries(input *pb.ServiceInput) (*languagePB.LibrariesOutput, error)
}
