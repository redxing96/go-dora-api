/*
 * @Description: 迁移
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 20:03:11
 * @LastEditTime: 2025-06-24 10:21:52
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/utility/migrate/migrate.go
 */
package migrate

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type MigrationUp func(tx *sql.Tx) error

type MigrationDown func(tx *sql.Tx) error

type Migration struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Up          MigrationUp
	Down        MigrationDown
}

type GoMigration struct {
	mux        sync.Mutex
	migrations []*Migration
	db         *sql.DB
	logger     io.Writer
}

func NewGoMigration(db *sql.DB, dbName string) (*GoMigration, error) {
	gm := GoMigration{
		migrations: make([]*Migration, 0),
		db:         db,
		logger:     os.Stdout,
	}

	if err := gm.initSchema(dbName); err != nil {
		return nil, err
	}

	return &gm, nil
}

func NewGoMigrationWithLogger(db *sql.DB, dbName string, logger io.Writer) (*GoMigration, error) {
	gm := GoMigration{
		migrations: make([]*Migration, 0),
		db:         db,
		logger:     logger,
	}

	if err := gm.initSchema(dbName); err != nil {
		return nil, err
	}

	return &gm, nil
}

/**
 * @description: 添加迁移
 * @param {...*Migration} m
 * @return {*}
 */
func (gm *GoMigration) AddMigration(m ...*Migration) error {
	gm.mux.Lock()
	defer gm.mux.Unlock()

	for _, v := range m {
		if v.Name == "" {
			return errors.New("the migration of name not allow empty")
		}
		if gm.hasMigration(v.Name) {
			return errors.New("the migration of name not allow empty")
		}
	}

	gm.migrations = append(gm.migrations, m...)
	return nil
}

/**
 * @description: 根据名称检测是否存在同名迁移
 * @param {string} name
 * @return {*}
 */
func (gm *GoMigration) hasMigration(name string) bool {
	for _, v := range gm.migrations {
		if v.Name == name {
			return true
		}
	}
	return false
}

/**
 * @description: 根据名称获取迁移
 * @param {string} name
 * @return {*}
 */
func (gm *GoMigration) getMigration(name string) *Migration {
	for i := len(gm.migrations) - 1; i >= 0; i-- {
		if gm.migrations[i].Name == name {
			return gm.migrations[i]
		}
	}
	return nil
}

/**
 * @description: 检测数据库是否存在同名迁移
 * @param {string} name
 * @return {*}
 */
func (gm *GoMigration) hasMigrationName(name string) bool {
	var count int
	res := gm.db.QueryRow("SELECT count(*) AS m_count FROM migrations WHERE name = ?", name)
	err := res.Scan(&count)
	if err == nil && count > 0 {
		return true
	}
	return false
}

/**
 * @description: 执行数据迁移
 * @return {*}
 */
func (gm *GoMigration) Migrate() (int, error) {
	num := 0
	maxBatchVersion := gm.getMaxBatchVersion()

	for _, migration := range gm.migrations {
		if gm.hasMigrationName(migration.Name) {
			continue
		}

		tx, _ := gm.db.Begin()
		if err := migration.Up(tx); err != nil {
			tx.Rollback()
			return num, err
		}

		sqlStr := "INSERT INTO migrations(`name`, `batch_version`, `description`, `create_time`)VALUES(?,?,?,?)"
		res, err := tx.Exec(sqlStr, migration.Name, maxBatchVersion+1, migration.Description, time.Now())
		if err != nil {
			tx.Rollback()
			return num, err
		}
		lastInsertId, err := res.LastInsertId()
		if lastInsertId <= 0 || err != nil {
			tx.Rollback()
			return num, err
		}
		tx.Commit()

		num++
		fmt.Fprintln(gm.logger, "| "+migration.Name+" | Migrate  Success")
	}

	if num == 0 {
		fmt.Fprintln(gm.logger, "Nothing Migrate")
	}
	return num, nil
}

/**
 * @description: 回滚最后一批迁移
 * @return {*}
 */
func (gm *GoMigration) RollBack() (int, error) {
	num := 0
	maxBatchVersion := gm.getMaxBatchVersion()

	rows, err := gm.db.Query("SELECT * FROM migrations WHERE batch_version = ? ORDER BY id DESC", maxBatchVersion)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		var id int
		var name string
		var description string
		var batchVersion int
		var createdAt time.Time

		rows.Scan(&id, &name, &description, &batchVersion, &createdAt)
		migration := gm.getMigration(name)
		if migration == nil {
			return num, errors.New("the migration of " + name + " is not exist")
		}

		tx, _ := gm.db.Begin()
		if err := migration.Down(tx); err != nil {
			tx.Rollback()
			return num, err
		}

		_, err := tx.Exec("DELETE FROM migrations WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			return num, err
		}
		tx.Commit()

		num++
		fmt.Fprintln(gm.logger, "| "+migration.Name+" | Rollback Success")
	}

	if num == 0 {
		fmt.Fprintln(gm.logger, "Nothing Rollback")
	}
	return num, nil
}

/**
 * @description: 全部回滚
 * @return {*}
 */
func (gm *GoMigration) RollBackAll() (int, error) {
	num := 0

	rows, err := gm.db.Query("SELECT * FROM migrations ORDER BY id DESC")
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		var id int
		var name string
		var description string
		var batchVersion int
		var createdAt time.Time

		rows.Scan(&id, &name, &description, &batchVersion, &createdAt)
		migration := gm.getMigration(name)
		if migration == nil {
			return num, errors.New("the migration of " + name + " is not exist")
		}

		tx, _ := gm.db.Begin()
		if err := migration.Down(tx); err != nil {
			tx.Rollback()
			return num, err
		}

		_, err := tx.Exec("DELETE FROM migrations WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			return num, err
		}
		tx.Commit()

		num++
		fmt.Fprintln(gm.logger, "| "+migration.Name+" | Rollback Success")
	}

	if num == 0 {
		fmt.Fprintln(gm.logger, "Nothing Rollback")
	}
	return num, nil
}

/**
 * @description: 重建迁移
 * @return {*}
 */
func (gm *GoMigration) ReBuild() (int, error) {
	numRollBack, err := gm.RollBackAll()
	if err != nil {
		return numRollBack, err
	}

	numMigrate, err := gm.Migrate()
	if err != nil {
		return numRollBack + numMigrate, err
	}

	if numMigrate+numRollBack == 0 {
		fmt.Fprintln(gm.logger, "Nothing Rebuild")
	}
	return numMigrate + numRollBack, nil
}

/**
 * @description: 初始化迁移表
 * @return {*}
 */
func (gm *GoMigration) initSchema(dbName string) error {
	// 检查数据库是否存在
	var exists bool
	err := gm.db.QueryRow("SELECT EXISTS (SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?)", dbName).Scan(&exists)
	// if err != nil {
	// 	return fmt.Errorf("检查数据库是否存在时出错: %v", err)
	// }
	// TODO 未完成，连接时指定了数据库，导致启动时找不到数据库，需要修改该bug
	// 如果数据库不存在，创建数据库
	if !exists || err != nil {
		_, err = gm.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName))
		if err != nil {
			return fmt.Errorf("创建数据库失败: %v", err)
		}

		// 选择新创建的数据库
		_, err = gm.db.Exec(fmt.Sprintf("USE `%s`", dbName))
		if err != nil {
			return fmt.Errorf("切换数据库失败: %v", err)
		}
	}

	// 创建迁移表
	_, err = gm.db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(255) NOT NULL COMMENT '名称',
			description varchar(255) NOT NULL COMMENT '描述',
			batch_version int(11) unsigned NOT NULL COMMENT '版本',
			create_time timestamp NULL DEFAULT NULL COMMENT '创建时间',
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='迁移表';
	`)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: 获取当前迁移的最大批次号
 * @return {*}
 */
func (gm *GoMigration) getMaxBatchVersion() int {
	var maxBatchVersion int
	row := gm.db.QueryRow("SELECT IFNULL(max(batch_version), 0) AS batch_version FROM migrations")

	row.Scan(&maxBatchVersion)
	return maxBatchVersion
}
