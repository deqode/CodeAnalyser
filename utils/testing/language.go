package testing

import (
	"code-analyser/protos/pb"
	"code-analyser/protos/pb/output/global"
	"code-analyser/protos/pb/output/languageSpecific"
	"reflect"
	"testing"
)

func AllCommandTesting(t *testing.T, got, output *pb.Commands) {
	t.Run("Build command", func(t *testing.T) {
		t.Parallel()
		CheckCommandEquality(got.BuildCommands.Used, output.BuildCommands.Used, got.BuildCommands.BuildCommands, output.BuildCommands.BuildCommands, t)
	})
	t.Run("seed command", func(t *testing.T) {
		t.Parallel()
		CheckCommandEquality(got.SeedCommands.Used, output.SeedCommands.Used, got.SeedCommands.SeedCommands, output.SeedCommands.SeedCommands, t)
	})
	t.Run("startup command", func(t *testing.T) {
		t.Parallel()
		CheckCommandEquality(got.StartUpCommands.Used, output.StartUpCommands.Used, got.StartUpCommands.StartUpCommands, output.StartUpCommands.StartUpCommands, t)
	})
	t.Run("migration command", func(t *testing.T) {
		t.Parallel()
		CheckCommandEquality(got.MigrationCommands.Used, output.MigrationCommands.Used, got.MigrationCommands.MigrationCommands, output.MigrationCommands.MigrationCommands, t)
	})
	t.Run("Adhoc script", func(t *testing.T) {
		t.Parallel()
		CheckCommandEquality(got.AdHocScriptsOutput.Used, output.AdHocScriptsOutput.Used, got.AdHocScriptsOutput.AdHocScripts, output.AdHocScriptsOutput.AdHocScripts, t)
	})
}

func CheckCommandEquality(gotUsed, outputUsed bool, gotCommands, outputCommands []*global.Command, t *testing.T) {
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

func CheckTestCaseCommandEquality(got, output *languageSpecific.TestCasesCommandOutput, t *testing.T) {
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
}
