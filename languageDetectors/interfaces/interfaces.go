package interfaces

import (
	"code-analyser/protos/pb/helpers"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
)

//TODO need to segregate and decouple our code

//DetectRunTime plugin methods
type DetectRunTime interface {
	Detect(inputString *helpers.StringInput) (*helpers.StringOutput, error)
}

//Dependencies plugin methods
type Dependencies interface {
	GetDependencies(inputString *helpers.Input) (*helpers.StringMapOutput, error)
}

//PreDetectCommands plugin methods
type PreDetectCommands interface {
	RunPreDetect(input *helpers.Input) (*helpers.EmptyOutput, error)
}

//StaticAssets plugin methods
type StaticAssets interface {
	Detect(input *helpers.Input) (*languageSpecificPB.StaticAssetsOutput, error)
}

//BuildDirectory plugin methods
type BuildDirectory interface {
	Detect(input *helpers.Input) (*languageSpecificPB.BuildDirectoryOutput, error)
}

//TestCasesRunCommands plugin methods
type TestCasesRunCommands interface {
	Detect(input *helpers.Input) (*languageSpecificPB.TestCasesCommand, error)
}
