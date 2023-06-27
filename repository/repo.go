package repository

import (
	logmanager "github.com/1996Paul-Wen/SafetyBox/infrastructure/log_manager"
	"github.com/1996Paul-Wen/SafetyBox/util/callstack"
)

func LogError(traceID string, err error) {
	caller := callstack.GetCallerNameBySkip(3)
	logmanager.RepoLog().Errorf("TraceID: %v, caller_func: %v, err: %v", traceID, caller, err)
}

func LogInfo(traceID string, info string) {
	caller := callstack.GetCallerNameBySkip(3)
	logmanager.RepoLog().Infof("TraceID: %v, caller_func: %v, info: %v", traceID, caller, info)
}
