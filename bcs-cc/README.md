# 蓝鲸容器管理平台配置中心模块

## 简介

蓝鲸容器管理平台配置中心是容器集群及节点等的配置中心。

## 主体功能

- 保存项目信息及绑定的蓝鲸CMDB业务信息
- 集群版本信息及集群快照信息
- 生成集群ID，保存master及node IP信息及状态

## 技术栈

- golang
- gin-gonic/gin
- mysql

## 依赖说明
- mysql: 容器管理平台数据库
- bk-iam: 蓝鲸权限中心