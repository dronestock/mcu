FROM golang:alpine AS builder


# Yaml修改程序版本
ENV YQ_VERSION 4.13.4
ENV YQ_BINARY yq_linux_amd64


RUN wget https://download.fastgit.org/mikefarah/yq/releases/download/v${YQ_VERSION}/${YQ_BINARY} --output-document /usr/bin/yq
RUN chmod +x /usr/bin/yq



# 打包真正的镜像
FROM storezhang/alpine


MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-10-12"
LABEL Description="Drone持续集成模块化插件，可以修改模块描述文件以及更新依赖"


# 复制文件
COPY --from=builder /usr/bin/yq /usr/bin/yq
COPY --from=builder /usr/local/go/bin/go /usr/bin/go
# 增加这一步是因为go命令在执行时，需要GOROOT目录，而正常的GOROOT目录是/usr/local/go
COPY --from=builder /usr/local/go/VERSION /usr/local/go/VERSION
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
