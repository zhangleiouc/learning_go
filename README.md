# Learning Go - Order API

基于 Clean Architecture 的 Go 项目，提供订单查询 API。

## 项目结构

```
.
├── api/
│   ├── controller/     # 控制器层
│   └── route/          # 路由层
├── bootstrap/          # 应用启动和配置
├── cmd/                # 应用入口
├── domain/             # 领域模型和接口定义
├── repository/         # 数据访问层
├── usecase/            # 业务逻辑层
├── go.mod
├── go.sum
└── .env                # 环境配置文件
```

## 功能

- 根据订单 ID 查询订单信息
- API 端点: `GET /order/:id`

## 环境要求

- Go 1.21.6 或更高版本
- MySQL 数据库

## 配置

创建 `.env` 文件（参考 `.env.example`）：

```env
APP_ENV=development
SERVER_ADDRESS=:8080
CONTEXT_TIMEOUT=10
DB_HOST=
DB_PORT=
DB_USER=
DB_PASS=
DB_NAME=
```

## 安装依赖

```bash
go mod tidy
```

## 运行项目

```bash
go run cmd/main.go
```

服务将在 `http://localhost:8080` 启动。

## API 使用示例

### 查询订单

```bash
curl http://localhost:8080/order/1
```

**成功响应 (200):**
```json
{
  "data": {
    "id": 1,
    "order_no": "ORD20240101001",
    "customer_id": 123,
    "total_amount": "100.00",
    "status": "completed",
    "created_at": "2024-01-01 10:00:00",
    "updated_at": "2024-01-01 10:00:00"
  }
}
```

**订单不存在 (404):**
```json
{
  "message": "Order not found"
}
```

**无效的订单ID (400):**
```json
{
  "message": "Invalid order ID"
}
```

## 注意事项

1. 当前 `Order` 结构体包含常见字段（id, order_no, customer_id, total_amount, status, created_at, updated_at）
2. 如果 `oms_order` 表的实际结构不同，请修改：
   - `domain/order.go` - 更新 Order 结构体字段
   - `repository/order_repository.go` - 更新 SQL 查询语句

## 架构说明

本项目采用 Clean Architecture 分层架构：

- **API Layer (api/)**: 处理 HTTP 请求和响应
- **Domain Layer (domain/)**: 定义业务实体和接口
- **UseCase Layer (usecase/)**: 实现业务逻辑
- **Repository Layer (repository/)**: 数据访问实现
- **Bootstrap Layer (bootstrap/)**: 应用初始化和配置

## 开发规范与目录职责

- `api/controller/`: 只做 HTTP 层与 Usecase 的数据编解码，不写业务逻辑。负责参数校验、HTTP 状态码、响应包装。
- `api/route/`: 所有路由注册集中于此，便于统一的中间件、版本控制和分组管理。
- `bootstrap/`: 负责环境变量读入、数据库连接、HTTP Server、依赖注入等启动流程。
- `cmd/`: 应用入口（`main.go`），仅调用 bootstrap 启动服务，不直接处理业务。
- `domain/`: 定义实体（Entity）、请求/响应 DTO 以及 Repository、Usecase 接口，是跨层共享的契约。
- `repository/`: 实现 Domain 中的 Repository 接口，专注于数据库/外部资源访问，禁止出现 HTTP 细节或业务判断。
- `usecase/`: 实现 Domain Usecase 接口，处理领域业务逻辑，组合和编排多个 Repository。

统一约定：

- 新增/修改接口时，先改 `domain/` 中的契约，再向外扩散，保持层与层之间依赖倒置。
- 所有对外 JSON 响应使用 `domain.SuccessResponse`、`domain.ErrorResponse` 包装。
- Context 传递到 Repository 层，Usecase 内通过 `context.WithTimeout` 控制。
- SQL 语句集中在 Repository，命名规范：`TableXxx` 常量来自 Domain 层。

## 新增接口的步骤指南

以下步骤以“新增一个订单接口”为例，可推广到所有资源类型：

1. **Domain 定义**  
   - 在 `domain/xxx.go` 中新增/更新实体字段、请求/响应结构体。  
   - 给 Repository、Usecase 接口增加方法签名，命名以动词开头（`Create`, `GetByID`, `List` 等）。

2. **Repository 实现**  
   - 在 `repository/` 下对应文件实现新的接口方法。  
   - 只接收 Domain Entity/DTO，返回原始错误时加语义（`fmt.Errorf("failed to ...: %w", err)`）。  
   - 统一使用 `db.QueryContext` 或 `db.ExecContext`，注意 Null 类型转换。

3. **Usecase 实现**  
   - 在 `usecase/` 中实现刚才在 Domain 定义的 Usecase 接口方法。  
   - 负责业务校验、拼装仓储输入、组合多个仓储调用，并处理超时 (`context.WithTimeout`)。

4. **Controller 层**  
   - 在 `api/controller/` 中添加 handler，解析路径/查询参数或 `ShouldBindJSON`。  
   - 调用 Usecase，转换 HTTP 状态码，返回标准响应结构。

5. **Route 注册**  
   - 在 `api/route/` 中将新的 handler 注册到对应的 HTTP method & path。  
   - 若需要中间件/版本号，统一在此层添加。

6. **Bootstrap 注入**  
   - 若引入了新的依赖（例如额外 Repository、Usecase），在 `bootstrap` 或路由初始化时完成注入。

7. **验证与文档**  
   - 使用 `curl`/Postman 自测，必要时补充单元测试。  
   - 更新 `README.md` 的“功能”与“API 使用示例”部分，保持文档同步。
