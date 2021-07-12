package loadPLugins

import (
	"code-analyser/protos/pb"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

type GlobalPlugin struct {
	ProcFile   *ProcFilePlugin
	MakeFile   *MakeFilePlugin
	DockerFile *DockerFilePlugin
	Setting    *utils.Setting
}

func (globalPlugin *GlobalPlugin) Load(ctx context.Context, pluginYamlFiles []utils.FileDetails) error {
	globalPlugin.Setting.Logger.Debug("global plugin loading started")

	globalPlugin.DockerFile = &DockerFilePlugin{
		Setting: globalPlugin.Setting,
	}
	globalPlugin.ProcFile = &ProcFilePlugin{
		Setting: globalPlugin.Setting,
	}
	globalPlugin.MakeFile = &MakeFilePlugin{
		Setting: globalPlugin.Setting,
	}

	globalPlugin.Setting.Logger.Info("global plugin yaml file reading and client initialization started")
	for _, pluginFile := range pluginYamlFiles {
		parsedRawFile, err := utils.ReadPluginYamlFile(ctx, pluginFile)
		if err != nil {
			globalPlugin.Setting.Logger.Error(err.Error())
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
	globalPlugin.Setting.Logger.Info("global plugin yaml file reading and client initialization completed")

	globalPlugin.Setting.Logger.Debug("global plugin loading completed")
	return nil
}

func (globalPlugin *GlobalPlugin) Run(ctx context.Context, projectRootPath string) (*pb.GlobalDetections, error) {
	globalPlugin.Setting.Logger.Debug("global plugins execution started")
	detectedDependencies := &pb.GlobalDetections{}

	globalPlugin.Setting.Logger.Info("docker file plugin methods execution started")
	dockerFile, dockerCompose, err := globalPlugin.DockerFile.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	detectedDependencies.DockerComposeFile = dockerCompose
	detectedDependencies.DockerFile = dockerFile
	globalPlugin.Setting.Logger.Info("docker file plugin methods execution completed")

	globalPlugin.Setting.Logger.Info("makefile plugin methods execution started")
	makefile, err := globalPlugin.MakeFile.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	detectedDependencies.Makefile = makefile
	globalPlugin.Setting.Logger.Info("makefile plugin methods execution completed")

	globalPlugin.Setting.Logger.Info("procfile plugin methods execution started")
	procfile, err := globalPlugin.ProcFile.Run(ctx, projectRootPath)
	if err != nil {
		return detectedDependencies, err
	}
	detectedDependencies.ProcFile = procfile
	globalPlugin.Setting.Logger.Info("procfile plugin methods execution completed")

	globalPlugin.Setting.Logger.Debug("global plugins execution completed")
	return detectedDependencies, nil
}
