# Go HTTP Service

这是一个使用Go语言构建的HTTP服务项目，能够连接MySQL和Redis数据库，并对接OpenAI的接口。

## 项目结构

```
go-http-service
├── cmd
│   └── main.go          # 应用程序入口点
├── internal
│   ├── api
│   │   └── handler.go   # HTTP请求处理程序
│   ├── db
│   │   ├── mysql.go     # MySQL数据库连接和操作
│   │   └── redis.go     # Redis数据库连接和操作
│   ├── openai
│   │   └── client.go    # OpenAI API交互
│   └── service
│       └── service.go   # 业务逻辑处理
├── config
│   └── config.yaml      # 配置文件
├── go.mod               # Go模块配置文件
├── go.sum               # 依赖项版本信息
└── README.md            # 项目文档
```

## 功能

- 提供HTTP服务，处理API请求
- 连接MySQL和Redis数据库，进行数据存储和检索
- 与OpenAI API进行交互，处理AI相关请求

## 使用方法

1. 克隆项目到本地:
   ```
   git clone <repository-url>
   ```

2. 进入项目目录:
   ```
   cd go-http-service
   ```

3. 安装依赖:
   ```
   go mod tidy
   ```

4. 配置数据库和OpenAI API信息，编辑`config/config.yaml`文件。

5. 启动服务:
   ```
   go run cmd/main.go
   ```

6. 访问API接口进行测试。

## 贡献

欢迎提交问题和拉取请求！请确保遵循项目的贡献指南。