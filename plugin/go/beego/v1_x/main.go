package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/framework"
	"github.com/hashicorp/go-plugin"
)

type Beego_V_1_x struct{}

func (*Beego_V_1_x) GetVersionName() (string, error) {
	return "1.4", nil
}

func (*Beego_V_1_x) GetSemver() (string, error)                       { return ">=2.x", nil }
func (*Beego_V_1_x) Detect(runtimeVersion, root string) (bool, error) { return true, nil }

func (g *Beego_V_1_x) IsFrameworkFound(runtimeVersion, root string) (bool, error) {
	return true, nil
}

func (g *Beego_V_1_x) IsFrameworkUsed(runtimeVersion, root string) (bool, error) {
	return true, nil
}

func (g *Beego_V_1_x) PercentOfFrameworkUsed(runtimeVersion, root string) (int64, error) {
	return 92, nil
}

func (g *Beego_V_1_x) GetFrameworkName() (string, error) {
	return "beego", nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserFramework: &framework.FrameworkGRPCPlugin{Impl: &Beego_V_1_x{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
