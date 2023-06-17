// safety box application entry
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/1996Paul-Wen/SafetyBox/api"
	"github.com/1996Paul-Wen/SafetyBox/config"
	"github.com/1996Paul-Wen/SafetyBox/infrastructure/db"
	logmanager "github.com/1996Paul-Wen/SafetyBox/infrastructure/log_manager"
	"github.com/1996Paul-Wen/SafetyBox/model"
)

var (
	command string
	// 日志目录
	logDir string = "log"
	// 配置文件地址
	configFilePath string = "app_config.yml"
)

func main() {
	var err error

	// 可执行文件所在路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("failed to get executable path, err:", err)
		os.Exit(1)
	}
	exeDir := path.Dir(exePath)
	configFilePath = path.Join(exeDir, configFilePath)
	logDir = path.Join(exeDir, logDir)

	// 解析命令行参数
	parseOption()
	validateCMDLineOption()

	// 初始化项目依赖
	initDependenies(configFilePath, logDir)

	fmt.Println("command : ", command)
	switch command {
	case "migrate_tables":
		err = MigrateTables()
	case "drop_tables":
		err = DropTables()
	case "start_web_app":
		err = StartWebAPP()
	case "hello":
		err = Hello()

	default:
		err = fmt.Errorf("unspport command :%s", command)
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// parseOption 解析命令行参数
func parseOption() {
	flag.Usage = Usage
	flag.StringVar(&configFilePath, "f", configFilePath, "Config File Path")
	flag.StringVar(&logDir, "l", logDir, "Log File Directory")
	flag.Parse()
	command = flag.Arg(0)
}

// Usage 打印CommandLine Usage `xxx -h`
func Usage() {
	helpHeader := `safetybox cmdline 
Options:
safetybox command [ options ]
command : 
	migrate_tables : 同步数据库表结构
	drop_tables : 删除数据库表
	start_web_app : 启动web应用
`
	fmt.Println(helpHeader)
	flag.PrintDefaults()
}

func validateCMDLineOption() {
	validateConfigFile()
	validateLogDir()
}

func validateConfigFile() {
	configFile, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		fmt.Println("config file not exist, path:", configFilePath)
		os.Exit(1)
	} else {
		if configFile.IsDir() {
			fmt.Println("config file should not be a directory, path:", configFilePath)
			os.Exit(1)
		}
	}
}

func validateLogDir() {
	logDirInfo, err := os.Stat(logDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			fmt.Printf("make log dir %+v err: %+v\n", logDir, err.Error())
			os.Exit(1)
		}
	} else {
		if !logDirInfo.IsDir() {
			fmt.Println("logDir should be a directory, path:", logDir)
			os.Exit(1)
		}
	}
}

func initDependenies(configFilePath, logDir string) {
	// 初始化全局配置
	config.Init(configFilePath)
	globalConfig := config.GlobalConfig()
	fmt.Printf("using global config: %+v\n", globalConfig)

	// 初始化日志
	logmanager.Init(logDir, globalConfig.Debug)

	// 初始化数据库连接
	db.Init(globalConfig.Database, model.LoadModels())
}

func MigrateTables() error {
	return db.DefaultDBManager().MigrateTables()
}

func DropTables() error {
	return db.DefaultDBManager().DropTables()
}

func StartWebAPP() error {
	return api.StartWebServer()
}

func Hello() error {
	fmt.Println("hello safetybox")
	return nil
}
