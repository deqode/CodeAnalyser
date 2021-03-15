package _go

import (
	protos "code-analyser/protos/pb"
	"context"
	"github.com/spf13/afero"
	"log"
)

type GoDetector struct {
}

func (d *GoDetector) DetectRuntime(c context.Context, a afero.Afero) (string, error) {
	log.Println("runtime")
	log.Println(a.Name())
	return "", nil

}
func (d *GoDetector) RunPreDetect(context.Context, string, string) (bool, error) {
	log.Println("preDetect")
	return false, nil
}

//func (d *GoDetector) RunParsers(context.Context, string, ) (interface{}, error){}
//func (d *GoDetector) ParseENVs(context.Context, Dir) ([]*protos.EnvOutput, error){}
func (d *GoDetector) DetectFrameworks(context.Context, string, string) ([]*protos.FrameworkOutput, error){
	getDetector()
}
//func (d *GoDetector) DetectDBs(context.Context, string, Dir) ([]*protos.DBOutput, error){}
//func (d *GoDetector)DetectORMs(context.Context, string, Dir) ([]*protos.OrmOutput, error){}
//func (d *GoDetector)DetectDependencies(context.Context, string, Dir) ([]*protos.DependenciesOutput, error){}
//func (d *GoDetector)DetectLibraries(context.Context, string, Dir) ([]*protos.LibrariesOutput, error){}
//func (d *GoDetector)GetStaticAssets(context.Context, string, Dir) ([]*protos.StaticAssetsOutput, error){}
//func (d *GoDetector)GetStack(context.Context, string, Dir) ([]*protos.StackOutput, error){}
//func (d *GoDetector)DetectAppserver(context.Context, string, Dir) ([]*protos.AppserverOutput, error){}
//func (d *GoDetector)DetectBuildDirectory(context.Context, string, Dir) ([]*protos.BuildDirectoryOutput, error){}
//func (d *GoDetector)DetectTestCasesRunCommands(context.Context, string, Dir) ([]*protos.BuildDirectoryOutput, error){}
