package db

import (
	"fmt"
	"testing"

	"github.com/1996Paul-Wen/SafetyBox/config"
	"github.com/1996Paul-Wen/SafetyBox/model"
)

func TestSyncDB(t *testing.T) {
	globalConf := config.GlobalConfig()
	globalConf.Init("../../bin/app_config.yml")
	fmt.Printf("%+v\n", globalConf)

	dbManager := DefaultDBManager()
	dbManager.Init(globalConf.Database, model.LoadModels())

	dbManager.DropTables()
	dbManager.MigrateTables()
	dbManager.CleanTables()
}
