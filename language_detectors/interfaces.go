package language_detectors

import "code-analyser/protos"

type LanguageSpecificDetector interface {
	DetectRuntime(string) string                        //root
	RunPreDetect(string, string) bool                   //root,version
	RunParsers(string, string) interface{}              //root version TODO: interface ?
	ParseENVs(string) []*protos.Output                  //root
	DetectFrameworks(string, string) []*protos.Output   //root,version
	DetectDBs(string, string) []*protos.Output          //root,version
	DetectORMs(string, string) []*protos.Output         //root,version
	DetectDependencies(string, string) []*protos.Output //root,version
	DetectLibraries(string, string) []*protos.Output    //root,version
	GetStaticAssets(string, string) []*protos.Output    //root,version
	GetStack(string, string) []*protos.Output           //root,version
	// build cmds
	// startup cmd
	// build directory

}
