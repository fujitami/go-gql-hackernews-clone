# 親イメージの指定 ※alplainだとM1環境に対応できない
FROM golang:1.17-buster
# aptパッケージ更新、git導入
RUN apt update && apt install git

# air(ホットリロード)導入
RUN go install github.com/cosmtrek/air@v1.27.3
# gqlgen導入go getするとバージョン名がついてしまう
RUN go install github.com/99designs/gqlgen@v0.13.0

# コンテナに持っていった.gitを読んでもらうため、起動する場所を明記
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
