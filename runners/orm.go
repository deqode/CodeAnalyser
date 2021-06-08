package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"golang.org/x/net/context"
)

//ExtractOrmsFromProjectDependencies It takes dependenciesList and filter out orms in map format
func ExtractOrmsFromProjectDependencies(ctx context.Context, projectDependencies map[string]string, ormPlugins map[string]*versionsPB.DependencyDetails) map[string]DependencyDetail {
	orm := map[string]DependencyDetail{}
	for name, details := range ormPlugins {
		if usedLibraryVersion, ok := projectDependencies[name]; ok {
			for version, versionDetails := range details.Version {
				if helpers.SemverValidateFromArray(versionDetails.Semver, usedLibraryVersion) {
					orm[name] = DependencyDetail{
						Version: version,
						Command: versionDetails.PluginPath,
					}
				}
			}
		}
	}
	return orm
}

//ExecuteOrmPlugins it append list of ORMS in ormoutput object
func ExecuteOrmPlugins(ctx context.Context, ormPlugins map[string]DependencyDetail, runtimeVersion, projectRootPath string) *languageSpecificPB.OrmOutput {
	ormOutput := languageSpecificPB.OrmOutput{
		Used: false,
		Orms: []*languageSpecificPB.ORM{},
	}
	for name, details := range ormPlugins {
		ormPluginResponse := ExecuteOrmPlugin(ctx, name, details, runtimeVersion, projectRootPath)
		if ormPluginResponse != nil {
			ormOutput.Used = true
			ormOutput.Orms = append(ormOutput.Orms, ormPluginResponse)
		}
	}
	return &ormOutput
}

//ExecuteOrmPlugin it run plugin file of ORM
func ExecuteOrmPlugin(ctx context.Context, name string, ormDetails DependencyDetail, runTimeVersion, projectRootPath string) *languageSpecificPB.ORM {
	pluginCall, _ := pluginClient.CreateOrmClient(utils.CallPluginCommand(ormDetails.Command))

	pluginInput := &pb.Input{
		RuntimeVersion: runTimeVersion,
		RootPath:       projectRootPath,
	}

	isUsed, err := pluginCall.IsUsed(pluginInput)
	if err != nil || isUsed.Error != nil {
		utils.Logger(err, isUsed)
		return nil
	}
	if isUsed.Value == false {
		return nil
	}

	response, err := pluginCall.Detect(pluginInput)
	if err != nil || response.Error != nil {
		utils.Logger(err, response)
		return nil
	}

	if response.Used {
		return &languageSpecificPB.ORM{
			Name:    name,
			Version: ormDetails.Version,
		}
	}
	return nil
}
