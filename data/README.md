# Data Layer
本目录包含了所有数据相关的核心实现，采用领域驱动设计(DDD)的思想组织代码。

## 目录结构
```text
/data
├── repository/           # 仓储层：负责数据持久化
│   ├── model/           # 数据模型定义
│   │   ├── crm/        # 业务库模型
│   │   └── chdata/     # 数仓模型
│   └── query/          # 查询实现
├── infrastructure/      # 基础设施层
│   ├── database/       # 数据库配置与连接
│   └── cache/         # 缓存实现
└── types/             # 共享类型定义
```

## 领域说明

### 认证授权领域 (auth)

处理所有认证和授权相关的业务逻辑，包括：

 - 社交应用认证（企业微信等）
 - OIDC认证
 - 权限管理

### 用户领域 (customer)
 - 管理用户相关的所有信息和行为：
 - 用户基础信息
 - 客户画像
 - 客户分组

### 社交CRM领域 (scrm)
管理社交媒体集成和客户关系：


## 开发规范

1. 领域模型规范
    - 领域模型使用 snake_case 命名
    - 每个领域模型必须实现相应的接口
    - 领域服务必须包含完整的单元测试

2. 仓储层规范

   - 仓储接口定义在领域层
   - 实现放在 `repository` 目录
   - 查询逻辑统一在 query 包中实现
   
3. 依赖规则
   - 领域层不依赖基础设施层
   - 仓储层依赖领域层
   - 避免跨领域直接依赖
   
4. 代码复用与查询层规则
   - 对于需要复用的代码，应优先放在对应的领域层或仓储层中
   - 对于几乎不考虑复用的代码，允许将 GORM DB 直接当作查询层使用，无需额外抽象
   - 简单查询可直接使用 GORM，复杂查询应封装在 repository 层

### 使用示例

```go
// 初始化服务
services := &Services{
   CustomerService: customer.NewService(
       repository.NewCustomerRepository(db),
       cache.NewCustomerCache(redis),
   ),
   // ... 其他服务初始化
}

// 使用服务
profile, err := services.CustomerService.GetProfile(customerID)
```

## 注意事项

1. 保持领域边界清晰，避免跨领域调用
2. 领域模型应该是充血模型，包含业务行为
3. 使用依赖注入管理服务依赖
4. 确保每个领域服务都有完整的测试覆盖
5. 定期审查和重构领域边界
6. 在决定是否抽象查询逻辑时，考虑代码复用性和维护成本

data/biz && data/query 目录将会逐步废弃, 改为主要由 repository 和 domain 提供的较复杂化代码以支撑更高并发, 更快的请求速度, 如异步/并发等尽量封装在这两层内
