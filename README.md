# Kubernetes
[![编译状态](https://github.ruijc.com:20443/api/badges/dronestock/mcu/status.svg)](https://github.ruijc.com:20443/dronestock/mcu)
[![Golang质量](https://goreportcard.com/badge/github.com/dronestock/mcu)](https://goreportcard.com/report/github.com/dronestock/mcu)
![版本](https://img.shields.io/github/go-mod/go-version/dronestock/mcu)
![仓库大小](https://img.shields.io/github/repo-size/dronestock/mcu)
![最后提交](https://img.shields.io/github/last-commit/dronestock/mcu)
![授权协议](https://img.shields.io/github/license/dronestock/mcu)
![语言个数](https://img.shields.io/github/languages/count/dronestock/mcu)
![最佳语言](https://img.shields.io/github/languages/top/dronestock/mcu)
![星星个数](https://img.shields.io/github/stars/dronestock/mcu?style=social)

Drone持续集成模块配置修改插件

## Dart模块带.的问题

正常情况下，`.`用于分隔路径，比如

```yaml
version: 0.0.1
dependencies:
  whiteboard: 0.0.1
```

使用`dependencies.whiteboard`来获取和写入值。假如有如下配置文件

```yaml
version: 0.0.1
dependencies:
  whiteboard.core: 0.0.1
```

那么应该使用`dependencies."whiteboard.core"`来获取和写入值

## 交流

![微信群](https://www.dronestock.tech/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢
[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png)](https://www.jetbrains.com/?from=dronestock/mcu)
