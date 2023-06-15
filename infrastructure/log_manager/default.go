package logmanager

var defaultLogManager *logManager = newLogManager()

func DefaultLogManager() *logManager {
	return defaultLogManager
}
