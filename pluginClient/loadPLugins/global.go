package loadPLugins

import (
	"code-analyser/utils"
	"golang.org/x/net/context"
	"log"
)

type GlobalPlugin struct {
	ProcFile   *ProcFileClient
	MakeFile   *MakeFileClient
	DockerFile *DockerFilePlugin
}

func (globalPlugin *GlobalPlugin) Load(ctx context.Context, pluginYamlFiles []utils.FileDetails) error {
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
