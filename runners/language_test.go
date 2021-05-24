package runners

import (
	"code-analyser/protos/pb/output/global"
	"reflect"
	"strconv"
	"testing"
)

func TestDependencies(t *testing.T) {
	for i, element := range DependencyCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := GetParsedDependencies(input.Ctx, input.LanguageVersion, input.Path, input.PluginDetails)
			if !reflect.DeepEqual(got, output) {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestDetectRuntime(t *testing.T) {
	for i, element := range DetectRunTimeCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := DetectRuntime(input.Ctx, input.Path, input.YamlLangObject)
			if got != output {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}

func TestGetCommands(t *testing.T) {
	for i, element := range DetectCommandCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			var got DetectCommandOutput
			got.SeedCommand, got.BuildCommand, got.MigrationCommand, got.StartUpCommand, got.AdHocScripts = GetCommands(input.Ctx, input.Input, input.PluginDetails)
			t.Run("Build command", func(r *testing.T) {
				r.Parallel()
				RunTestCommand(got.BuildCommand.Used, output.BuildCommand.Used, got.BuildCommand.BuildCommands, output.BuildCommand.BuildCommands, t)
			})
			t.Run("seed command", func(r *testing.T) {
				r.Parallel()
				RunTestCommand(got.SeedCommand.Used, output.SeedCommand.Used, got.SeedCommand.SeedCommands, output.SeedCommand.SeedCommands, t)
			})
			t.Run("startup command", func(r *testing.T) {
				r.Parallel()
				RunTestCommand(got.StartUpCommand.Used, output.StartUpCommand.Used, got.StartUpCommand.StartUpCommands, output.StartUpCommand.StartUpCommands, t)
			})
			t.Run("migration command", func(r *testing.T) {
				r.Parallel()
				RunTestCommand(got.MigrationCommand.Used, output.MigrationCommand.Used, got.MigrationCommand.MigrationCommands, output.MigrationCommand.MigrationCommands, t)
			})
			t.Run("Adhoc script", func(r *testing.T) {
				r.Parallel()
				RunTestCommand(got.AdHocScripts.Used, output.AdHocScripts.Used, got.AdHocScripts.AdHocScripts, output.AdHocScripts.AdHocScripts, t)
			})
		})

	}
}

func RunTestCommand(gotUsed, outputUsed bool, gotCommands, outputCommands []*global.Command, t *testing.T) {
	if gotUsed != outputUsed || len(gotCommands) != len(outputCommands) {
		t.Error("expected this ", outputUsed, outputCommands, "\n got this ", gotUsed, gotCommands)
	}
	originalOutputCommands := outputCommands
	for j := 0; j < len(gotCommands); j++ {
		for k := 0; k < len(outputCommands); k++ {
			if reflect.DeepEqual(gotCommands[j], outputCommands[k]) {
				outputCommands[k] = outputCommands[len(outputCommands)-1]
				outputCommands = outputCommands[:len(outputCommands)-1]
				break
			}
		}
	}
	if len(outputCommands) != 0 {
		t.Error("expected this ", outputUsed, originalOutputCommands, "\n got this ", gotUsed, gotCommands)
	}
}

func TestDetectTestCasesCommand(t *testing.T) {
	for i, element := range DetectTestCommandCases {
		input := element.Input
		output := element.Output
		t.Run("case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			got := DetectTestCasesCommand(input.Ctx, input.Input, input.PluginDetails)
			if got.Used != output.Used || len(got.Commands) != len(output.Commands) {
				t.Error("expected this ", output, "\n got this ", got)
			}
			var outputCommands []*global.Command
			copy(outputCommands, output.Commands)
			for j := 0; j < len(got.Commands); j++ {
				for k := 0; k < len(outputCommands); k++ {
					if reflect.DeepEqual(got.Commands[j], outputCommands[k]) {
						outputCommands[k] = outputCommands[len(outputCommands)-1]
						outputCommands = outputCommands[:len(outputCommands)-1]
						break
					}
				}
			}
			if len(outputCommands) != 0 {
				t.Error("expected this ", output, "\n got this ", got)
			}
		})
	}
}
