package main

import (
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/framework"
	"github.com/hashicorp/go-plugin"
)

type Greeter struct{}

func (*Greeter) GetVersionName() (string, error) {
	return "1.3", nil
}

func (*Greeter) GetSemver() (string, error)                       { return ">=2.x", nil }
func (*Greeter) Detect(runtimeVersion, root string) (bool, error) { return true, nil }

func (g *Greeter) IsFrameworkFound(runtimeVersion, root string) (bool, error) {
	return true, nil
}

func (g *Greeter) IsFrameworkUsed(runtimeVersion, root string) (bool, error) {
	return true, nil
}

func (g *Greeter) PercentOfFrameworkUsed(runtimeVersion, root string) (int64, error) {
	return 92, nil
}

func (g *Greeter) GetFrameworkName() (string, error) {
	return "beego", nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserFramework: &framework.GreeterGRPCPlugin{Impl: &Greeter{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
