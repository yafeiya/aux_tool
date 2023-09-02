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

```
# 1.前端
# 操作路径 ./auxTool-frontEnd-main/url_config.ys  
  const url_dict =
  {
    // 局域网内暴露本机地址
    "frontEndUrl":'192.168.0.103',  // 可选项：部署时用，开发时用localhost:5173

    // 向后端发送请求的地址+端口
    "backEndUrl": 'http://192.168.0.103:8080'  // 必填
}  
# 2. 后端
# 操作路径 ./congfig/application.yml
# 后端服务
server:
  host:             //不填，默认自动获取本机地址
 
# SQL数据库
datasource:    // # 在3016局域网内可默认，否则依据本机mysql设置

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
