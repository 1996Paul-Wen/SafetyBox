package db

import (
	"fmt"
	"time"

	"github.com/1996Paul-Wen/SafetyBox/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	log "github.com/InVisionApp/go-logger"
)

// dbManager 基于Gorm的数据库连接抽象
type dbManager struct {
	config config.DatabaseConfig
	engine *gorm.DB
	tables []interface{}

	logger log.Logger
}

func new() *dbManager {
	return &dbManager{}
}

// Init should only be called once for a single dbManager instance
func Init(conf config.DatabaseConfig, tables []interface{}) {
	defaultDBManager.config = conf
	defaultDBManager.engine = newMySQLEngine(conf)
	defaultDBManager.tables = tables
	defaultDBManager.logger = log.NewSimple()
}

func (dbManager *dbManager) GetEngine() *gorm.DB {
	return dbManager.engine
}

func (dbManager *dbManager) SetLogger(logger log.Logger) {
	dbManager.logger = logger
}

func (dbManager *dbManager) MigrateTables() error {
	return dbManager.engine.AutoMigrate(dbManager.tables...)
}

func (dbManager *dbManager) CleanTables() error {
	for _, table := range dbManager.tables {
		result := dbManager.engine.Session(&gorm.Session{AllowGlobalUpdate: true}).Where("1=1").Delete(table)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (dbManager *dbManager) DropTables() error {
	return dbManager.engine.Migrator().DropTable(dbManager.tables...)
}

func (dbManager *dbManager) Reconnect() error {
	if err := dbManager.Close(); err != nil {
		return err
	}
	dbManager.engine = newMySQLEngine(dbManager.config)
	return nil
}

func (dbManager *dbManager) Close() error {
	if dbManager.engine != nil {
		sqlDB, err := dbManager.engine.DB()
		if err != nil {
			return err
		}
		dbManager.engine = nil
		return sqlDB.Close()
	}

	return nil
}

func (dbManager *dbManager) BeginTransaction() *gorm.DB {
	tx := dbManager.engine.Begin()
	return tx
}

// StartMonitor monitor DBState, should run in a goroutine
func (dbManager *dbManager) StartMonitor() {
	sqlDB, err := dbManager.engine.DB()
	if err != nil {
		dbManager.logger.Error("Failed to get DB connection: %s", err.Error())
		return
	}

	for {
		time.Sleep(5 * time.Second)
		stat := sqlDB.Stats()
		msg := fmt.Sprintf("maxopen : %d , opened:  %d , idle : %d ,  inuse: %d , wait : %d,  waitduration : %s",
			stat.MaxOpenConnections, stat.OpenConnections, stat.Idle, stat.InUse, stat.WaitCount,
			stat.WaitDuration.String(),
		)
		dbManager.logger.Info(msg)
	}
}

func newMySQLEngine(dbConfig config.DatabaseConfig) *gorm.DB {
	// `gorm.DB` provides a higher level of abstraction
	dbEngine, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dbConfig.DSN,
		DefaultStringSize: dbConfig.DefaultStringSize, // string 类型字段的默认长度
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// `sql.DB` is a lower-level tool that gives you more control over your database interactions, represents a pool of zero or more underlying connections
	db, err := dbEngine.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxIdleTime(dbConfig.ConnMaxIdleTime)

	return dbEngine
}
