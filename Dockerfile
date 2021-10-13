FROM storezhang/alpine

MAINTAINER storezhang "storezhang@gmail.com"
LABEL architecture="AMD64/x86_64" version="latest" build="2021-10-12"
LABEL Description="Drone持续集成模块化插件，可以修改模块描述文件以及更新依赖"


# 复制文件
COPY mcu /bin


RUN set -ex \
    \
    \
    \
    && apk update \
    && apk --no-cache add go \
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
