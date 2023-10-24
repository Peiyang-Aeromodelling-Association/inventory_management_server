# 物料管理系统后端

[![unittest](https://github.com/Peiyang-Aeromodelling-Association/inventory_management_server/actions/workflows/test.yml/badge.svg?event=push)](https://github.com/Peiyang-Aeromodelling-Association/inventory_management_server/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/Peiyang-Aeromodelling-Association/inventory_management_server/branch/main/graph/badge.svg)](https://app.codecov.io/gh/Peiyang-Aeromodelling-Association/inventory_management_server)

## 项目简介

一个go语言实现的物料管理系统后端，web框架使用gin，数据库使用postgresql。

## 开发

查看makefile中的命令

```bash
make help
```

运行单元测试

```bash
make test
```

重新生成swagger文档

```bash
make swagger
```

## 部署

通过docker-compose部署：

```bash
make dockercompose
```

该命令其实就是`docker-compose up -d`。

如果出现意外，重启docker-compose也是相同的命令。一般来说，postgresql会自动重启。

重新构建镜像：

```bash
docker-compose build
```