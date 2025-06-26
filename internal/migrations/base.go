/*
 * @Description: 迁移基础
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 20:05:47
 * @LastEditTime: 2025-06-23 20:41:46
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/migrations/base.go
 */
package migrations

import (
	"database/sql"
	"go-dora-api/internal/logger"
	"go-dora-api/utility/common"
	"go-dora-api/utility/migrate"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

// 迁移
type migration struct {
	sqlDB        *sql.DB
	dbName       string
	lockFilePath string
}

// 迁移实例
var Migration *migration

// 初始化迁移
func init() {
	db, err := g.DB().Open(g.DB().GetConfig())
	if err != nil {
		logger.SystemLogger.Errorf("数据库连接失败: %s", err)
		panic(err)
	}
	dbName := g.DB().GetConfig().Name
	if dbName == "" {
		dbName = "go-dora-api"
	}
	// 初始化迁移
	Migration = NewMigration(db, dbName, common.GetAbsPath("install.lock"))
}

// 创建迁移实例
func NewMigration(sqlDB *sql.DB, dbName string, lockFilePath string) *migration {
	Migration = &migration{
		sqlDB:        sqlDB,
		dbName:       dbName,
		lockFilePath: lockFilePath,
	}
	err := Migration.migrateTable()
	if err != nil {
		logger.SystemLogger.Errorf("数据库表初始化失败: %s", err)
		panic(err)
	}
	return Migration
}

// 判断是否存在锁文件
func (m *migration) isExistsLockFile() bool {
	return gfile.Exists(m.lockFilePath)
}

// migrateTable函数用于迁移表
func (m *migration) migrateTable() error {
	// 如果不存在锁文件，则返回nil
	if !m.isExistsLockFile() {
		return nil
	}

	// 创建一个新的GoMigration对象
	obj, err := migrate.NewGoMigration(m.sqlDB, m.dbName)
	if err != nil {
		return err
	}

	// 添加迁移函数
	err = obj.AddMigration(m.migrateV1()...)
	if err != nil {
		return err
	}

	// 执行迁移
	_, err = obj.Migrate()
	if err != nil {
		return err
	}

	// 迁移数据
	err = m.migrateData()
	if err != nil {
		return err
	}

	// 删除锁文件
	m.deleteLockFile()
	return nil
}

// 删除迁移锁文件
func (m *migration) deleteLockFile() error {
	// 使用os包中的Remove函数删除迁移锁文件
	return os.Remove(m.lockFilePath)
}

// 填充默认数据
func (m *migration) migrateData() (err error) {
	// ctx := gi18n.WithLanguage(context.TODO(), service.Translate().GetDefaultLanguage())
	// service.Area().InitData(ctx)
	// service.VoiceLibrary().InitData(ctx)
	// service.Role().InitRole(ctx)
	// service.User().InitUser(ctx)
	// service.Access().InitAccess(ctx)
	// service.Access().OrdinaryAccessInit(ctx)
	// service.TextTree().InitData(ctx)
	// service.TextLibrary().InitData(ctx)
	// service.SystemSetting().InitDefaultSystemSetting(ctx)
	// service.PrioritySetting().InitDefaultPrioritySetting(ctx)
	// service.Placeholder().InitData(ctx)
	// service.BroadcastRule().InitData(ctx)
	// service.Language().InitData(ctx)
	return
}
