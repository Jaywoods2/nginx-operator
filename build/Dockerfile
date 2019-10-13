#FROM registry.access.redhat.com/ubi7/ubi-minimal:latest
FROM registry.cn-shenzhen.aliyuncs.com/jaywoods/nginx-operator:base

ENV OPERATOR=/usr/local/bin/nginx-operator \
    USER_UID=1001 \
    USER_NAME=nginx-operator

# install operator binary
COPY build/_output/bin/nginx-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
