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

```
# 1.1 前端暴露ip地址，供局域网内访问本web应用（# 可选项：部署时用，现可用localhost:5173）
# 操作路径 ./auxTool-frontEnd-main/vite.config.js  
server:{
     host:'192.168.1.110'  # 前端所在ip
   }
# 1.2 前端配置发送数据请求的目标地址（后端地址）
# 操作路径 ./auxTool-frontEnd-main/src/api/myAios.js
const baseUrl = 'http://192.168.0.103:2345' # 后端ip地址，端口默认2345
```

```
# 2.1 后端暴露ip地址，供前端请求数据
# 操作路径 ./main.go
err := engine.Run("192.168.0.103:2345")' # 后端ip地址，端口默认2345
```

```
# 3.1 数据库链接
# 操作路径 ./common/databaseInit.go
	host := "192.168.0.102"  # 在3016局域网内可默认，否则依据本机mysql设置
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
