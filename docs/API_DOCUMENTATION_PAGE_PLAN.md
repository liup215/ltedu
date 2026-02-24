# API 文档页面实现方案

## 需求概述

创建一个页面，展示系统所有 API 接口，包括：
- 接口路径
- 接口说明
- 请求方法（GET / POST）
- 认证要求（是否需要 JWT）
- 请求参数（Body / Query / Path）
- 返回值结构

---

## 推荐方案：swaggo/swag + Swagger UI（纯后端）

### 为什么选择这个方案？

| 方案 | 优点 | 缺点 |
|------|------|------|
| **✅ swaggo/swag**（推荐） | 自动生成、可交互测试、标准规范、永远与代码同步、无需前端改动 | 需为每个 handler 添加注释 |
| 手写静态 Markdown | 简单快速 | 维护困难、易与代码脱节 |
| 纯手写 HTML 页面 | 样式完全自由 | 同上 |

本项目使用 Gin 框架，`swaggo/swag` 是业界标准解决方案，支持从 Go 注释自动生成 OpenAPI 3.0 规范，并由 **Go 后端直接托管 Swagger UI 页面**，无需前端项目参与。

---

## 实现步骤（纯后端）

### 第一步：安装依赖

```bash
# 安装 swag CLI 工具
go install github.com/swaggo/swag/cmd/swag@latest

# 添加 Go 模块依赖
go get github.com/swaggo/swag
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

### 第二步：在 `main.go` 添加全局注释

```go
// @title           LTEdu API
// @version         1.0
// @description     LTEdu 在线教育平台后端接口文档
// @host            localhost:8080
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

### 第三步：为每个 handler 添加 Swagger 注释（示例）

```go
// @Summary      用户登录
// @Description  用户名密码登录，返回 JWT Token
// @Tags         认证
// @Accept       json
// @Produce      json
// @Param        body body LoginRequest true "登录信息"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  ErrorResponse
// @Router       /v1/login [post]
func (ctrl *AuthController) Login(c *gin.Context) { ... }
```

### 第四步：注册 Swagger UI 路由（无需认证）

在 `server/api/controller.go` 的 `noAuthRout` 中添加：

```go
import swaggerfiles "github.com/swaggo/files"
import ginSwagger "github.com/swaggo/gin-swagger"
import _ "edu/docs" // swag 生成的文档包

// Swagger UI 路由
r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
```

访问地址：`http://localhost:8080/api/docs/index.html`

### 第五步：生成文档

```bash
# 在项目根目录执行
swag init -g main.go --output docs/swagger
```

每次 handler 注释更新后重新执行此命令。

---

## API 分组规划

根据现有路由，Swagger Tags 分组建议如下：

| Tag（标签） | 包含接口 | 认证要求 |
|------------|---------|---------|
| 认证 | 登录、注册、修改密码 | 无 |
| 用户管理 | /v1/user/* | 需要 |
| 试卷管理 | /v1/paper/* | 部分 |
| 题目管理 | /v1/question/* | 部分 |
| 考纲管理 | /v1/organisation, qualification, syllabus, chapter | 部分 |
| 学习导航 | /v1/goal, task, attempt, knowledge-state | 需要 |
| 知识点 | /v1/knowledge-point/* | 需要 |
| 课程 | /v1/course*, courseVideo* | 部分 |
| 文档 | /v1/document*, documentCategory* | 部分 |
| 媒体 | /v1/mediaImage*, mediaVideo* | 部分 |
| 学校管理 | /v1/school/* | 需要 |
| 词汇 | /v1/vocabulary* | 部分 |
| 系统设置 | /v1/syssetting/* | 需要 |
| MCP | /v1/mcp/*, /mcp | 混合 |

---

## 工作量评估

| 步骤 | 工作内容 | 预估工时 |
|------|---------|---------|
| 环境配置 | 安装依赖、配置 main.go、注册路由 | 1h |
| Handler 注释 | 为约 130 个 handler 添加 Swagger 注释 | 6~8h |
| **合计** | | **~8h** |

---

## 快速验证方案（最小可行版本）

如果想快速看到效果，可以先只完成：

1. 安装依赖、配置路由（30 分钟）
2. 为 5~10 个核心接口添加注释（1 小时）
3. 运行 `swag init` 并访问 `/api/docs/index.html`

验证效果后再批量补全所有接口注释。

---

## 参考资源

- [swaggo/swag 官方文档](https://github.com/swaggo/swag)
- [gin-swagger 集成示例](https://github.com/swaggo/gin-swagger)
- [OpenAPI 3.0 规范](https://swagger.io/specification/)
- [Swagger UI](https://swagger.io/tools/swagger-ui/)
