

# go-dora-api

## 简介
`go-dora-api` 是一个使用 Go 语言开发的后端服务项目，主要为管理后台提供 RESTful API 接口。项目基于 [Gogf](https://goframe.org) 框架，提供管理员登录认证、菜单管理、角色管理等功能。

## 功能特性
- **权限管理**：通过 `token` 实现客户端和管理端的 JWT 认证
- **菜单管理**：支持菜单的增删改查
- **角色管理**：支持角色的增删改查
- **日志系统**：提供完善的日志记录功能
- **多语言支持**：通过 `i18n` 实现中英文切换
- **数据库迁移**：使用 `migrate` 包实现数据库结构版本控制
- **密码安全**：使用 `bcrypt` 实现安全的密码存储
- **Swagger API 文档**：提供完整的 OpenAPI 接口文档

## 接口文档
开发环境下可通过访问 `/swagger` 查看 API 文档

## 安装部署

### 依赖环境
- Go 1.20+
- MySQL 8.0+
- Redis (用于权限缓存)

### 安装步骤
1. 克隆仓库
2. 安装依赖: `go mod download`
3. 配置数据库: 修改 `hack/config.yaml` 中的数据库连接信息
4. 执行迁移: `make migrate`
5. 启动服务: `make run`

### 配置
- 语言设置: 修改 `hack/config.yaml` 中的 `server.develop` 字段控制开发模式
- 上传大小: 在 `internal/cmd/cmd.go` 中可修改 `Max_UPLOAD_SIZE` 调整最大上传文件大小
- 多语言: 支持 `zh-CN` 和 `en` 两种语言，通过请求头 `Lang` 设置

## 使用说明
### 登录认证
- 管理员登录接口: `/manage/v1/auth/login`
- 客户端 JWT 认证字段: `Authorization` (在 `internal/consts/consts.go` 中定义)
- 管理端 JWT 认证字段: `Authorization`

### 菜单管理
- 创建菜单: `/api/v1/menu/add`
- 删除菜单: `/api/v1/menu/delete`
- 更新菜单: `/api/v1/menu/update`
- 获取所有菜单: `/api/v1/menu/all`
- 获取用户权限菜单: `/manage/v1/menu/get_auth_menu`

### 角色管理
- 创建角色: `/api/v1/role/add`
- 删除角色: `/api/v1/role/delete`
- 更新角色: `/api/v1/role/update`
- 获取所有角色: `/api/v1/role/all`

## 中间件
- JWT 认证中间件: 验证请求中的 token
- 跨域(CORS)中间件: 处理跨域请求
- 访问日志中间件: 记录所有请求日志
- 异常处理中间件: 捕获并处理系统异常
- 权限验证中间件: 验证用户对特定接口的访问权限

## 开发规范
- 请求参数验证: 使用 `g.Meta` 字段进行接口级参数验证
- 全局响应结构: `JsonResponse` 结构体统一返回格式
- 日志记录: 使用 `logger.ILogger` 接口记录日志
- 数据访问: 所有数据库操作通过 DAO 模式实现
- 业务逻辑: 采用 Service 模式提供业务逻辑接口

## 目录结构
- `api/` - API 接口定义
- `internal/` - 核心代码
  - `cmd/` - 命令行处理
  - `common_return/` - 全局返回格式
  - `consts/` - 常量定义
  - `controller/` - 控制器
  - `dao/` - 数据访问对象
  - `logger/` - 日志处理模块
  - `logic/` - 业务逻辑
  - `migrations/` - 数据库迁移脚本
  - `model/` - 数据模型
  - `service/` - 服务接口
- `utility/` - 工具类
  - `common/` - 通用工具
  - `jwt/` - JWT 处理工具
  - `logger/` - 日志工具
  - `migrate/` - 数据库迁移工具
  - `common_return/` - 通用返回工具

## 版本管理
- 当前版本: v1.0.0 (可在 `internal/consts/consts.go` 中修改)
- 数据库迁移: 使用 `utility/migrate` 实现数据库结构版本控制

## 安全特性
- 密码安全: 使用 `bcrypt` 简化密码哈希处理
- token 管理: 提供 token 黑名单、刷新和失效功能

## 日志系统
- 日志配置: 默认日志配置在 `utility/logger/config.go`
- 日志字段: 使用 `LogField` 结构体记录日志上下文信息
- 日志级别: 支持 debug, info, warn, error 等多个日志级别

## 命令行参数
- 启动命令: `main` 函数执行服务启动
- 初始化定时任务: `sCrontab.InitData` 方法

## 许可证
该项目使用 MIT 许可证。详细信息请查看 `LICENSE` 文件。