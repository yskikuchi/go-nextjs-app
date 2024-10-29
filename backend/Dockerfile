FROM golang:1.23.0-alpine

WORKDIR /go/src

RUN apk update && apk add git

# マイグレーションツール
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# ホットリロード用
RUN go install github.com/air-verse/air@latest

EXPOSE 8080
