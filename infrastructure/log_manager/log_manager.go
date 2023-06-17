package logmanager

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type logManager struct {
	logDir     string // 日志文件的目录
	gatewayLog *logrus.Logger
	repoLog    *logrus.Logger
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
func (l *logManager) createLogrusLogger(filename string, isDebug bool) (*logrus.Logger, error) {
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

	return logrusLogger, nil
}

func (l *logManager) GatewayLog() *logrus.Logger {
	return l.gatewayLog
}

func (l *logManager) RepoLog() *logrus.Logger {
	return l.repoLog
}
