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

/*
ExtractDbsFromProjectDependencies It will filter out frameworks from dependencies list where
dependenciesList : list of dependencies of any project ,
langYamlObject : dependencies supported by us
*/
func ExtractDbsFromProjectDependencies(ctx context.Context, projectDependencies map[string]string, dbPlugins map[string]*versionsPB.DependencyDetails) map[string]DependencyDetail {
	db := map[string]DependencyDetail{}
	for name, details := range dbPlugins {
		for dbVersion, versionDetails := range details.Version {
			for _, library := range versionDetails.Libraries {
				if usedLibraryVersion, ok := projectDependencies[library.Name]; ok {
					if helpers.SemverValidateFromArray(library.Semver, usedLibraryVersion) {
						db[name] = DependencyDetail{
							Version: dbVersion,
							Command: versionDetails.PluginPath,
						}
					}
				}
			}
		}
	}
	return db
}

//ExecuteDbPlugins will run to detect its dbs and return its detectors
func ExecuteDbPlugins(ctx context.Context,dbPlugins map[string]DependencyDetail, runtimeVersion, projectRootPath string) *languageSpecificPB.DBOutput {
	dbOutput := languageSpecificPB.DBOutput{
		Used:      false,
		Databases: []*languageSpecificPB.DB{},
	}
	for name, details := range dbPlugins {
		dbPluginResponse := ExecuteDbPlugin(ctx,name, details, runtimeVersion, projectRootPath)
		if dbPluginResponse != nil {
			dbOutput.Used = true
			dbOutput.Databases = append(dbOutput.Databases, dbPluginResponse)
		}
	}
	return &dbOutput
}

//TODO handle errors on every method calls


//ExecuteDbPlugin It will execute db plugin which is located at dbPluginDetail.Command ,
//name :- db name,
//dbPluginDetail :- db plugin info for e.g. version path
func ExecuteDbPlugin(ctx context.Context,name string, dbPluginDetail DependencyDetail, runTimeVersion, projectRootPath string) *languageSpecificPB.DB {
	pluginCall, _ := pluginClient.CreateDbClient(utils.CallPluginCommand(dbPluginDetail.Command))

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
	if response.Value {
		return &languageSpecificPB.DB{
			Name:    name,
			Version: dbPluginDetail.Version,
			Port:    response.IntValue,
		}
	}
	return nil
}
