package loadPLugins

import (
	"code-analyser/protos/pb"
	"code-analyser/utils"
	"golang.org/x/net/context"
	"log"
)

type GlobalPlugin struct {
	ProcFile   *ProcFilePlugin
	MakeFile   *MakeFilePlugin
	DockerFile *DockerFilePlugin
}

func (globalPlugin *GlobalPlugin) Load(ctx context.Context, pluginYamlFiles []utils.FileDetails) error {

	globalPlugin.DockerFile = &DockerFilePlugin{}
	globalPlugin.ProcFile = &ProcFilePlugin{}
	globalPlugin.MakeFile = &MakeFilePlugin{}

	for _, pluginFile := range pluginYamlFiles {
		parsedRawFile, err := utils.ReadPluginYamlFile(ctx, pluginFile)
		if err != nil {
			log.Println(err)
			return err
		}

		pluginYamlFile := parsedRawFile.PluginDetails

		switch pluginYamlFile.Type {
		case "dockerFile":
			globalPlugin.DockerFile.Load(pluginYamlFile)
		case "procFile":
			globalPlugin.ProcFile.Load(pluginYamlFile)
		case "makeFile":
			globalPlugin.MakeFile.Load(pluginYamlFile)
		}
	}

	return nil
}

func (globalPlugin *GlobalPlugin) Run(ctx context.Context, projectRootPath string) (*pb.GlobalDetections, error) {
	detectedDependencies := &pb.GlobalDetections{}

	dockerFile, dockerCompose, err := globalPlugin.DockerFile.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	detectedDependencies.DockerComposeFile = dockerCompose
	detectedDependencies.DockerFile = dockerFile

	makefile, err := globalPlugin.MakeFile.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	detectedDependencies.Makefile = makefile

	procfile, err := globalPlugin.ProcFile.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	detectedDependencies.ProcFile = procfile

	return detectedDependencies, nil
}
