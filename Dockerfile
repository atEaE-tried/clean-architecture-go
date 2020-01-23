FROM golang:1.13.0-alpine
ARG approot=/app

# basic setting
ENV LANG C.UTF-8
ENV CGO_ENABLED 0
RUN apk add --update-cache \
        git \
        bash \
        postgresql-client \
        make \
    --no-cache

# make working directory
WORKDIR $approot
ADD . $approot
