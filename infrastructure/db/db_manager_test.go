package db

import (
	"testing"

	"github.com/1996Paul-Wen/SafetyBox/config"
	"github.com/1996Paul-Wen/SafetyBox/model"
)

func TestSyncDB(t *testing.T) {
	config.Init("../../bin/app_config.yml")
	Init(config.GlobalConfig().Database, model.LoadModels())

	dbManager := DefaultDBManager()
	dbManager.DropTables()
	dbManager.MigrateTables()
	dbManager.CleanTables()
}
