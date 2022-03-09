FROM golang:alpine AS yq


# Yaml修改程序版本
ENV YQ_VERSION 4.21.1
ENV YQ_BINARY yq_linux_amd64


RUN wget https://ghproxy.com/https://github.com/mikefarah/yq/releases/download/v${YQ_VERSION}/${YQ_BINARY} --output-document /usr/bin/yq
RUN chmod +x /usr/bin/yq



# 打包真正的镜像
FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成模块化插件，可以修改模块描述文件以及更新依赖"


# 复制文件
COPY --from=yq /usr/bin/yq /usr/bin/yq
COPY --from=yq /usr/local/go/bin/go /usr/bin/go
# 增加这一步是因为go命令在执行时，需要GOROOT目录，而正常的GOROOT目录是/usr/local/go
COPY --from=yq /usr/local/go/VERSION /usr/local/go/VERSION
COPY mcu /bin



RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/mcu \
    \
    \
    \
    && rm -rf /var/cache/apk/*



ENTRYPOINT /bin/mcu
