# drone-plugin-mcu

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
