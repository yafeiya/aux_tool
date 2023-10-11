# auxTool

## Introduction

auxTool是一款前后端分离的数据驱动建模辅助工具

**Tech Stack:**

Vue3+vite+JavaScript+View UI Plus+Gin+MySql

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
# 路径 ./auxTool-frontEnd-main/url_config.ys  
    # 局域网内暴露本机地址
    "frontEndUrl":'127.0.0.1',
    "frontpost":"5173",
    # 向后端发送请求的地址+端口
    "backEndUrl": 'http://127.0.0.1:8080',
    "fileUrl": 'http://127.0.0.1:5173'
# 2. 后端
# 路径 ./congfig/application.yml
  # 后端服务
   server:        # 默认自动获取本机地址
  
  # SQL数据库
   datasource:    # 在3016局域网内默认配置，否则依据本机mysql配置

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

# 开启前端服务 (./auxTool-frontEnd-main)
npm run dev

```

## MySQL

```
# 初次开发时，部署数据库
# 开启mySQL
mysql -u root -p

# 创建数据库
CREATE DATABASE auxtool
    DEFAULT CHARACTER SET = 'utf8';
use auxtool

# 导入数据库
use auxtool;
source /xx/xxxx.sql;
exit;

# 导出数据库（可选）
mysqldump -u root -p --default-character-set=utf8 auxtool > xx.sql
```

## Note

```
# 1. 允许其他ip访问mysql
use mysql;
select Host, User,Password from user;
update user set Host='%' where User='root';
flush privileges;
```

## Reference

[gorm](https://gorm.io/docs/)

[View UI Plus](https://www.iviewui.com/view-ui-plus/guide/introduce)
