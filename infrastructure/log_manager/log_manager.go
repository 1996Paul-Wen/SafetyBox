package logmanager

import (
	"path"
	"time"

	log "github.com/InVisionApp/go-logger"
	lgrs "github.com/InVisionApp/go-logger/shims/logrus"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type logManager struct {
	logDir string // 日志文件的目录

	// 不同层次使用不同的log
	gatewayLog log.Logger
	repoLog    log.Logger
}

func newLogManager() *logManager {
	return &logManager{}
}

// Init should only be called once
func Init(logDir string, isDebug bool) {
	defaultLogManager.logDir = logDir

	gatewayLog, err := defaultLogManager.createLogrusLogger("gateway.log", isDebug)
	if err != nil {
		panic(err)
	}
	defaultLogManager.gatewayLog = gatewayLog

	repoLog, err := defaultLogManager.createLogrusLogger("repo.log", isDebug)
	if err != nil {
		panic(err)
	}
	defaultLogManager.repoLog = repoLog
}

// createLogrusLogger 创建Logrus日志句柄
func (l *logManager) createLogrusLogger(filename string, isDebug bool) (log.Logger, error) {
	writer, err := rotatelogs.New(
		path.Join(l.logDir, filename+".%Y-%m-%d"),
		rotatelogs.WithLinkName(path.Join(l.logDir, filename)),
		// 最大保留5天的日志数据
		rotatelogs.WithMaxAge(5*24*time.Hour),
		rotatelogs.WithRotationTime(1*time.Hour),
		rotatelogs.WithRotationSize(100*1024*1024),
	)
	if err != nil {
		return nil, err
	}

	logrusLogger := logrus.New()
	logrusLogger.Out = writer
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logrusLogger.SetReportCaller(true)
	if isDebug {
		logrusLogger.SetLevel(logrus.DebugLevel)
	}

	// logrusLogger是log接口的具体实现，再使用lgrs.New()将实现包装成interface
	return lgrs.New(logrusLogger), nil
}

func (l *logManager) GatewayLog() log.Logger {
	return l.gatewayLog
}

func (l *logManager) RepoLog() log.Logger {
	return l.repoLog
}
