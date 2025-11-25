# 表结构调整说明

由于无法直接访问数据库查看 `oms_order` 表的实际结构，当前代码使用了常见的订单表字段。

## 如何查看实际表结构

你可以通过以下方式查看 `oms_order` 表的实际结构：

```bash
mysql -h  -P 3306 -u oms -p'' oms -e "DESCRIBE oms_order;"
```

或者使用 MySQL 客户端工具连接数据库后执行：

```sql
DESCRIBE oms_order;
-- 或者
SHOW CREATE TABLE oms_order;
```

## 如何调整代码

如果实际表结构与当前代码不匹配，需要修改以下文件：

### 1. 修改 `domain/order.go`

更新 `Order` 结构体，添加或删除字段以匹配实际表结构：

```go
type Order struct {
    ID          int64   `json:"id" db:"id"`
    // 根据实际表结构添加或修改字段
    OrderNo     *string `json:"order_no,omitempty" db:"order_no"`
    // ... 其他字段
}
```

### 2. 修改 `repository/order_repository.go`

更新 SQL 查询语句和 Scan 方法：

1. 修改 SELECT 语句中的字段列表
2. 添加或删除对应的变量声明
3. 更新 Scan 方法的参数
4. 添加或删除对应的字段赋值逻辑

示例：

```go
// 如果表中有新字段，例如 order_date
query := fmt.Sprintf(`
    SELECT id, order_no, customer_id, total_amount, status, order_date, created_at, updated_at
    FROM %s
    WHERE id = ?
`, domain.TableOrder)

// 添加对应的变量
var orderDate sql.NullString

// 在 Scan 中添加
err := or.db.QueryRowContext(c, query, id).Scan(
    &order.ID,
    &orderNo,
    &customerID,
    &totalAmount,
    &status,
    &orderDate,  // 新增
    &createdAt,
    &updatedAt,
)

// 添加字段赋值
if orderDate.Valid {
    order.OrderDate = &orderDate.String
}
```

## 当前假设的字段

当前代码假设 `oms_order` 表包含以下字段：
- `id` (主键，整数类型)
- `order_no` (订单号，字符串类型，可为空)
- `customer_id` (客户ID，整数类型，可为空)
- `total_amount` (总金额，字符串或小数类型，可为空)
- `status` (状态，字符串类型，可为空)
- `created_at` (创建时间，时间戳或字符串，可为空)
- `updated_at` (更新时间，时间戳或字符串，可为空)

如果你的表结构不同，请按照上述步骤进行调整。

