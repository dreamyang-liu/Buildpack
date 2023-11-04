FROM amazonlinux:2023

ARG cnb_uid=1000
ARG cnb_gid=1000

ENV CNB_USER_ID=${cnb_uid}
ENV CNB_GROUP_ID=${cnb_gid}

RUN yum update -y && \
    yum install shadow-utils -y && \
    groupadd --gid ${cnb_gid} cnb && \
    useradd --uid ${cnb_uid} --gid ${cnb_gid} -m -s /bin/bash cnb && \
    yum clean all && \
    rm -rf /var/cache/yum && \
    rm -rf /var/lib/yum

USER ${cnb_uid}:${cnb_gid}

LABEL io.buildpacks.stack.id="aws-23"
LABEL io.buildpacks.stack.distro.name="Amazon Linux"
LABEL io.buildpacks.stack.distro.version="2023"
LABEL io.buildpacks.stack.version="0.0.1"
LABEL io.buildpacks.stack.maintainer="mengyang"

ENV PORT 8080
EXPOSE 8080