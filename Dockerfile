FROM storezhang/alpine

MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-10-12"
LABEL Description="Drone持续集成模块化插件，可以修改模块描述文件以及更新依赖"


# 复制文件
COPY mcu /bin


# Yaml修改程序版本
ENV YQ_VERSION 4.13.4
ENV YQ_BINARY yq_linux_amd64


RUN set -ex \
    \
    \
    \
    && apk update \
    && apk --no-cache add go \
    \
    \
    \
    && wget https://download.fastgit.org/mikefarah/yq/releases/download/v${YQ_VERSION}/${YQ_BINARY} --output-document /usr/bin/yq \
    && chmod +x /usr/bin/yq \
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
