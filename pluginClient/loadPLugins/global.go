package loadPLugins

import (
	"code-analyser/protos/pb"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

//GlobalPlugin contains 1. all global plugin (ProcFile, MakeFile, DockerFile) and Setting
type GlobalPlugin struct {
	ProcFile   *ProcFilePlugin
	MakeFile   *MakeFilePlugin
	DockerFile *DockerFilePlugin
	Setting    *utils.Setting
}

//Load It calls load function on all global plugins
func (globalPlugin *GlobalPlugin) Load(ctx context.Context, pluginYamlFiles []utils.FileDetails) error {
	globalPlugin.Setting.Logger.Debug("global plugin loading started")

	globalPlugin.DockerFile = &DockerFilePlugin{Setting: globalPlugin.Setting}
	globalPlugin.ProcFile = &ProcFilePlugin{Setting: globalPlugin.Setting}
	globalPlugin.MakeFile = &MakeFilePlugin{Setting: globalPlugin.Setting}

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
			globalPlugin.Setting.Logger.Info("docker plugin client creation started")
			globalPlugin.DockerFile.Load(pluginYamlFile)
			globalPlugin.Setting.Logger.Info("docker plugin client created successfully")
		case "procFile":
			globalPlugin.Setting.Logger.Info("procfile plugin client creation started")
			globalPlugin.ProcFile.Load(pluginYamlFile)
			globalPlugin.Setting.Logger.Info("procfile plugin client created successfully")
		case "makeFile":
			globalPlugin.Setting.Logger.Info("makefile plugin client creation started")
			globalPlugin.MakeFile.Load(pluginYamlFile)
			globalPlugin.Setting.Logger.Info("makefile plugin client created successfully")
		}
	}
	globalPlugin.Setting.Logger.Info("global plugin yaml file reading and client initialization completed")

	globalPlugin.Setting.Logger.Debug("global plugin loading completed")
	return nil
}

//Run It calls run function on all global plugins
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
