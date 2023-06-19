FROM dockerproxy.com/mikefarah/yq:4.34.1 AS yq
FROM dockerproxy.com/library/golang:1.20.5-alpine AS golang
FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.18.2 AS builder

# 复制文件
COPY --from=yq /usr/bin/yq /docker/usr/bin/yq
COPY --from=golang /usr/local/go/bin/go /docker/usr/bin/go
# 增加这一步是因为go命令在执行时，需要GOROOT目录，而正常的GOROOT目录是/usr/local/go
COPY --from=golang /usr/local/go/VERSION /docker/usr/local/go/VERSION
COPY mcu /docker/usr/local/bin/mcu



# 打包真正的镜像
FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.18.2


LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="Drone持续集成模块化插件，可以修改模块描述文件以及更新依赖"


# 一次性复制所有程序，如果有多个COPY命令需要通过多Builder模式减少COPY登岛
COPY --from=builder /docker /


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /usr/local/bin/mcu \
    \
    \
    \
    && rm -rf /var/cache/apk/*



ENTRYPOINT /usr/local/bin/mcu
