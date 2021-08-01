FROM golang:1.15-alpine

ENV GOPATH /usr/src/app/go
ARG dir=$GOPATH/src/github.com/mrauer
WORKDIR ${dir}

RUN apk add make curl ffmpeg python3 && ln -sf python3 /usr/bin/python

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl && chmod a+rx /usr/local/bin/youtube-dl

WORKDIR $GOPATH/src/github.com/mrauer/grabify
COPY . .
