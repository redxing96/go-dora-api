/*
 * @Description: 迁移v1版本
 * @Author: redxing96@163.com
 * @Date: 2025-06-23 20:08:22
 * @LastEditTime: 2025-07-01 16:25:06
 * @LastEditors: front end cabbage
 * @FilePath: /go-dora-api/internal/migrations/v1.go
 */
package migrations

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"database/sql"
	"go-dora-api/utility/migrate"
)

// migrateV1函数用于迁移v1版本
func (m *migration) migrateV1() []*migrate.Migration {
	return []*migrate.Migration{
		{
			Name:        "20250624194800_create_manage_table",
			Description: "创建管理员列表",
			Up: func(tx *sql.Tx) error {
				_, err := tx.Exec(`
					 CREATE TABLE manage (
						id int(11) NOT NULL AUTO_INCREMENT,
						account varchar(255) NOT NULL COMMENT '用户账号',
						password varchar(255) NOT NULL COMMENT '用户密码',
						email varchar(255) DEFAULT NULL COMMENT '电子邮箱',
						phone varchar(255) DEFAULT NULL COMMENT '电话号码',
						avatar varchar(255) DEFAULT NULL COMMENT '头像',
						status tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0-初始化 1-正常 2-冻结',
						is_super tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否超管 1-是',
						create_time datetime NOT NULL COMMENT '创建时间',
						update_time datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
						PRIMARY KEY (id)
					) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员列表';
				`)
				return err
			},
			Down: func(tx *sql.Tx) error {
				_, err := tx.Exec(`DROP TABLE IF EXISTS manage`)
				return err
			},
		},
		{
			Name:        "20250624194900_create_manage_init_data",
			Description: "初始化管理员列表数据",
			Up: func(tx *sql.Tx) error {
				_, err := tx.Exec(`
					INSERT INTO manage (id, account, password, email, phone, avatar, status, is_super, create_time, update_time) VALUES (1, 'admin', '$2a$10$y6eCHHLBAjv/NdNUnZSAu.XemNWXevmO.Zj1j/NiBSuF.NtyFfSw6', '', '', '', 0, 1, '2025-06-24 19:49:51', NULL), (2, 'test', '$2a$10$y6eCHHLBAjv/NdNUnZSAu.XemNWXevmO.Zj1j/NiBSuF.NtyFfSw6', '', '', '', 0, 0, '2025-06-24 19:50:18', NULL);
				`)
				return err
			},
			Down: func(tx *sql.Tx) error {
				_, err := tx.Exec(`DELETE FROM manage WHERE id IN (1, 2);`)
				return err
			},
		},
		{
			Name:        "20250624205100_create_sys_menu_table",
			Description: "创建系统菜单表",
			Up: func(tx *sql.Tx) error {
				_, err := tx.Exec(`
					CREATE TABLE sys_menu (
						id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
						pid bigint(20) DEFAULT NULL COMMENT '父菜单ID',
						type tinyint(3) NOT NULL DEFAULT '1' COMMENT '权限类型(1菜单,2接口,3按钮,4目录)',
						path varchar(255) DEFAULT NULL COMMENT '路由地址',
						sort int(11) DEFAULT '0' COMMENT '排序',
						component varchar(255) DEFAULT NULL COMMENT '组件路径',
						title varchar(50) DEFAULT NULL COMMENT '菜单标题',
						name varchar(255) DEFAULT NULL COMMENT '菜单名称',
						icon varchar(100) DEFAULT NULL COMMENT '图标类名',
						icon_svg varchar(255) DEFAULT NULL COMMENT 'svg图标',
						is_hidden tinyint(1) DEFAULT '0' COMMENT '是否隐藏(2-否,1-是)',
						is_keep_alive tinyint(1) DEFAULT '0' COMMENT '是否缓存(2-否,1-是)',
						active_menu varchar(255) DEFAULT NULL COMMENT '激活菜单的path',
						is_large_screen tinyint(1) DEFAULT '0' COMMENT '是否仅在大屏显示(2-否,1-是)',
						link varchar(255) DEFAULT '0' COMMENT '外部链接',
						status tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态(1正常,2禁用)',
						create_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
						update_time datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
						remark varchar(500) DEFAULT NULL COMMENT '备注',
						PRIMARY KEY (id),
						KEY idx_pid (pid),
						KEY idx_path (path)
					) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统菜单表'; 
				`)
				return err
			},
			Down: func(tx *sql.Tx) error {
				_, err := tx.Exec(`DROP TABLE IF EXISTS sys_menu;`)
				return err
			},
		},
		{
			Name:        "20250625101405_create_sys_role_table",
			Description: "创建角色表",
			Up: func(tx *sql.Tx) error {
				_, err := tx.Exec(`
					CREATE TABLE sys_role (
						id bigint NOT NULL AUTO_INCREMENT COMMENT '角色ID',
						role_name varchar(50) NOT NULL COMMENT '角色名称',
						role_desc varchar(255) NOT NULL COMMENT '角色描述',
						status tinyint NOT NULL DEFAULT '1' COMMENT '状态(1正常,0禁用)',
						create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
						PRIMARY KEY (id)
					) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
				`)
				return err
			},
			Down: func(tx *sql.Tx) error {
				_, err := tx.Exec(`DROP TABLE IF EXISTS sys_role;`)
				return err
			},
		},
		{
			Name:        "20250625101705_create_sys_manager_role_table",
			Description: "创建管理员角色关联表",
			Up: func(tx *sql.Tx) error {
				_, err := tx.Exec(`
					CREATE TABLE sys_manager_role (
						id bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
						manager_id bigint NOT NULL COMMENT '用户ID',
						role_id bigint NOT NULL COMMENT '角色ID',
						PRIMARY KEY (id),
						UNIQUE KEY idx_manager_role (manager_id,role_id),
						KEY idx_role_id (role_id)
					) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员角色关联表';
				`)
				return err
			},
			Down: func(tx *sql.Tx) error {
				_, err := tx.Exec(`DROP TABLE IF EXISTS sys_manager_role;`)
				return err
			},
		},
		{
			Name:        "20250625101805_create_sys_role_menu_table",
			Description: "创建角色菜单关联表",
			Up: func(tx *sql.Tx) error {
				_, err := tx.Exec(`
					CREATE TABLE sys_role_menu (
						id bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
						role_id bigint NOT NULL COMMENT '角色ID',
						menu_id bigint NOT NULL COMMENT '菜单ID',
						PRIMARY KEY (id),
						UNIQUE KEY idx_role_menu (role_id,menu_id),
						KEY idx_menu_id (menu_id)
					) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单关联表';  
				`)
				return err
			},
			Down: func(tx *sql.Tx) error {
				_, err := tx.Exec(`DROP TABLE IF EXISTS sys_role_menu;`)
				return err
			},
		},
		{
			Name:        "20250629123100_create_media_table",
			Description: "创建媒体表",
			Up: func(tx *sql.Tx) error {
				_, err := tx.Exec(`
					CREATE TABLE media (
						id int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
						file_name varchar(255) NOT NULL COMMENT '文件名',
						original_name varchar(255) NOT NULL COMMENT '原始文件名',
						file_path varchar(500) NOT NULL COMMENT '文件路径',
						file_url varchar(500) NOT NULL COMMENT '文件访问URL',
						file_size bigint(20) NOT NULL COMMENT '文件大小(字节)',
						file_type varchar(100) NOT NULL COMMENT '文件类型',
						mime_type varchar(100) NOT NULL COMMENT 'MIME类型',
						md5_hash varchar(64) NOT NULL COMMENT '文件MD5哈希',
						sha1_hash varchar(64) NOT NULL COMMENT '文件SHA1哈希',
						uploader_id int(11) NOT NULL COMMENT '上传者ID',
						uploader_type varchar(50) NOT NULL COMMENT '上传者类型',
						status tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 0-待处理 1-正常 2-删除',
						is_public tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否公开 0-私有 1-公开',
						download_count int(11) NOT NULL DEFAULT '0' COMMENT '下载次数',
						view_count int(11) NOT NULL DEFAULT '0' COMMENT '查看次数',
						create_time datetime DEFAULT NULL COMMENT '创建时间',
						update_time datetime DEFAULT NULL COMMENT '修改时间',
						PRIMARY KEY (id),
						KEY idx_uploader (uploader_id,uploader_type),
						KEY idx_file_type (file_type)
					) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='媒体资源表'; 
				`)
				return err
			},
			Down: func(tx *sql.Tx) error {
				_, err := tx.Exec(`DROP TABLE IF EXISTS media;`)
				return err
			},
		},
	}
}
