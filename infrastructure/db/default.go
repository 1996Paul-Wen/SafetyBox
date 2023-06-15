package db

var defaultDBManager *dbManager = new()

func DefaultDBManager() *dbManager {
	return defaultDBManager
}
