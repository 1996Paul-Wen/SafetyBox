package logmanager

import log "github.com/InVisionApp/go-logger"

// defaultLogManager is a placeholder-like instance to be inited
var defaultLogManager *logManager = newLogManager()

func DefaultLogManager() *logManager {
	return defaultLogManager
}

// GatewayLog is a shorthand of defaultLogManager.GatewayLog
func GatewayLog() log.Logger {
	return defaultLogManager.gatewayLog
}

// RepoLog is a shorthand of defaultLogManager.RepoLog
func RepoLog() log.Logger {
	return defaultLogManager.repoLog
}
