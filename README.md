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
