# 媒体文件系统使用文档

## 概述

本系统提供了完整的媒体文件管理功能，支持文件上传、秒传、列表查询、详情查看、删除等操作。系统通过MD5哈希实现秒传功能，避免重复文件的上传。

## 数据库表结构

### media 表

```sql
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
```

## API 接口

### 1. 上传媒体文件

**接口地址：** `POST /v1/media/upload`

**请求参数：**
- `file`: 文件（multipart/form-data）
- `uploaderId`: 上传者ID（必填）
- `uploaderType`: 上传者类型（必填）
- `isPublic`: 是否公开（0-私有，1-公开，默认0）

**响应示例：**
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "media": {
            "id": 1,
            "fileName": "20250629_123456_abc123.jpg",
            "originalName": "photo.jpg",
            "filePath": "/uploads/2025-06-29/20250629_123456_abc123.jpg",
            "fileUrl": "http://localhost:8000/files/2025-06-29/20250629_123456_abc123.jpg",
            "fileSize": 1024000,
            "fileType": ".jpg",
            "mimeType": "image/jpeg",
            "md5Hash": "d41d8cd98f00b204e9800998ecf8427e",
            "sha1Hash": "da39a3ee5e6b4b0d3255bfef95601890afd80709",
            "uploaderId": 1,
            "uploaderType": "admin",
            "status": 1,
            "isPublic": 0,
            "downloadCount": 0,
            "viewCount": 0,
            "createTime": "2025-06-29 12:34:56",
            "updateTime": "2025-06-29 12:34:56"
        },
        "isQuickUpload": false
    }
}
```

### 2. 获取媒体文件列表

**接口地址：** `GET /v1/media/list`

**请求参数：**
- `page`: 页码（默认1）
- `pageSize`: 每页数量（默认10）
- `uploaderId`: 上传者ID（可选）
- `fileType`: 文件类型（可选）
- `status`: 状态（可选）
- `isPublic`: 是否公开（可选）
- `keyword`: 关键词搜索（可选）

**响应示例：**
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "list": [
            {
                "id": 1,
                "fileName": "20250629_123456_abc123.jpg",
                "originalName": "photo.jpg",
                "fileUrl": "http://localhost:8000/files/2025-06-29/20250629_123456_abc123.jpg",
                "fileSize": 1024000,
                "fileType": ".jpg",
                "mimeType": "image/jpeg",
                "uploaderId": 1,
                "uploaderType": "admin",
                "status": 1,
                "isPublic": 0,
                "downloadCount": 0,
                "viewCount": 0,
                "createTime": "2025-06-29 12:34:56"
            }
        ],
        "total": 1,
        "page": 1,
        "size": 10
    }
}
```

### 3. 获取媒体文件详情

**接口地址：** `GET /v1/media/detail`

**请求参数：**
- `id`: 媒体文件ID（必填）

**响应示例：**
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "id": 1,
        "fileName": "20250629_123456_abc123.jpg",
        "originalName": "photo.jpg",
        "filePath": "/uploads/2025-06-29/20250629_123456_abc123.jpg",
        "fileUrl": "http://localhost:8000/files/2025-06-29/20250629_123456_abc123.jpg",
        "fileSize": 1024000,
        "fileType": ".jpg",
        "mimeType": "image/jpeg",
        "md5Hash": "d41d8cd98f00b204e9800998ecf8427e",
        "sha1Hash": "da39a3ee5e6b4b0d3255bfef95601890afd80709",
        "uploaderId": 1,
        "uploaderType": "admin",
        "status": 1,
        "isPublic": 0,
        "downloadCount": 0,
        "viewCount": 0,
        "createTime": "2025-06-29 12:34:56",
        "updateTime": "2025-06-29 12:34:56"
    }
}
```

### 4. 删除媒体文件

**接口地址：** `DELETE /v1/media/delete`

**请求参数：**
- `id`: 媒体文件ID（必填）

**响应示例：**
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "success": true
    }
}
```

### 5. 记录媒体文件查看

**接口地址：** `POST /v1/media/view`

**请求参数：**
- `id`: 媒体文件ID（必填）

**响应示例：**
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "success": true
    }
}
```

### 6. 记录媒体文件下载

**接口地址：** `POST /v1/media/download`

**请求参数：**
- `id`: 媒体文件ID（必填）

**响应示例：**
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "success": true
    }
}
```

### 7. 检查文件MD5是否存在

**接口地址：** `POST /v1/media/check-md5`

**请求参数：**
- `md5Hash`: 文件MD5哈希（必填）

**响应示例：**
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "exists": true,
        "media": {
            "id": 1,
            "fileName": "20250629_123456_abc123.jpg",
            "originalName": "photo.jpg",
            "fileUrl": "http://localhost:8000/files/2025-06-29/20250629_123456_abc123.jpg",
            "fileSize": 1024000,
            "fileType": ".jpg",
            "mimeType": "image/jpeg",
            "uploaderId": 1,
            "uploaderType": "admin",
            "status": 1,
            "isPublic": 0,
            "downloadCount": 0,
            "viewCount": 0,
            "createTime": "2025-06-29 12:34:56"
        }
    }
}
```

## 秒传功能

系统通过MD5哈希实现秒传功能：

1. **上传前检查**：在上传文件前，可以先调用 `check-md5` 接口检查文件是否已存在
2. **自动秒传**：如果文件已存在，系统会自动返回已存在的文件信息，无需重复上传
3. **节省存储**：相同内容的文件只存储一份，节省存储空间

## 支持的文件类型

系统支持以下文件类型：
- 图片：`.jpg`, `.jpeg`, `.png`, `.gif`, `.bmp`, `.webp`
- 视频：`.mp4`, `.avi`, `.mov`, `.wmv`, `.flv`
- 音频：`.mp3`, `.wav`, `.flac`
- 文档：`.pdf`, `.doc`, `.docx`, `.xls`, `.xlsx`, `.ppt`, `.pptx`, `.txt`
- 压缩包：`.zip`, `.rar`, `.7z`

## 文件大小限制

- 默认最大文件大小：100MB
- 可在控制器中修改 `maxFileSize` 变量来调整限制

## 文件存储结构

文件按日期存储在以下结构中：
```
uploads/
├── 2025-06-29/
│   ├── 20250629_123456_abc123.jpg
│   └── 20250629_123457_def456.png
└── 2025-06-30/
    └── 20250630_120000_ghi789.mp4
```

## 使用示例

### 前端上传示例（JavaScript）

```javascript
// 上传文件
async function uploadMedia(file, uploaderId, uploaderType) {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('uploaderId', uploaderId);
    formData.append('uploaderType', uploaderType);
    formData.append('isPublic', 0);

    const response = await fetch('/v1/media/upload', {
        method: 'POST',
        body: formData
    });

    const result = await response.json();
    
    if (result.data.isQuickUpload) {
        console.log('文件秒传成功！');
    } else {
        console.log('文件上传成功！');
    }
    
    return result.data.media;
}

// 检查文件是否存在
async function checkFileExists(md5Hash) {
    const response = await fetch('/v1/media/check-md5', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ md5Hash })
    });

    const result = await response.json();
    return result.data.exists;
}
```

### 获取文件列表示例

```javascript
// 获取文件列表
async function getMediaList(page = 1, pageSize = 10) {
    const params = new URLSearchParams({
        page: page,
        pageSize: pageSize
    });

    const response = await fetch(`/v1/media/list?${params}`);
    const result = await response.json();
    
    return result.data;
}
```

## 注意事项

1. **权限控制**：所有接口都需要API密钥认证
2. **文件安全**：系统会验证文件类型和大小
3. **存储空间**：注意监控上传目录的磁盘空间使用情况
4. **备份策略**：建议定期备份上传的文件和数据库
5. **CDN配置**：生产环境建议配置CDN来加速文件访问

## 错误码说明

- `file_required`: 文件不能为空
- `file_too_large`: 文件过大
- `file_type_not_allowed`: 文件类型不允许
- `create_dir_failed`: 创建目录失败
- `save_file_failed`: 保存文件失败
- `calculate_md5_failed`: 计算MD5失败
- `media_not_found`: 媒体文件不存在
- `save_media_record_failed`: 保存媒体记录失败 