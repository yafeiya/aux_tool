# auxTool

## Introduction

auxTool是一款前后端分离的数据驱动建模辅助工具

**Tech Stack:**

Vue3+vite+JavaScript+View UI Plus+Go+MySql

## Install

node: [20.4.6](https://nodejs.cn/download/)

go: [1.20](https://go.dev/dl/).*

mySQL: [5.7.43](https://dev.mysql.com/downloads/installer/)

```
git clone https://github.com/yafeiya/aux_tool.git
```

## Configurarion

TODO：后续统一配置文件

1. 配置后端地址
   ./main.go

   ```
   err := engine.Run("x.x.x.x:2345") # 更改替换为自己ip
   ```
2. 配置SQL数据库地址
   ./common/databaseInit.go
   当前配置为实验室共享SQL数据库，如无法连接3016局域网，请用安装MySql时配置重设以下参数

   ```
   func InitDB() *gorm.DB {
   	// driverName := "mysql"
   	host := "192.168.0.102"  # 3016实验室共享数据库地址
   	port := "3306"
   	database := "auxtool"
   	username := "root"
   	password := "123456"
   ```

## Build & Run

```
# 开启后端服务器
go run main.go

# 进入前端目录
cd ./auxTool-frontEnd-main

# 安装前端依赖
npm install

# 开启json虚拟后端
npm run json  # TODO后续删除

# 开启前端服务 (./前端代码目录)
npm run dev

```

## Reference

[gorm](https://gorm.io/docs/)

[View UI Plus](https://www.iviewui.com/view-ui-plus/guide/introduce)
