package db

import (
	"code-analyser/language_detectors/interfaces"
	"github.com/hashicorp/go-plugin"
)

type DbGRPCPlugin struct {
	plugin.Plugin

	Impl interfaces.DbVersion
}

