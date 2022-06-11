FROM golang:1.17.3 AS build
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
WORKDIR /data/cnccamp/project
COPY . /data/cnccamp/project
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -installsuffix cgo  -o http_server .


FROM alpine:latest as prod
WORKDIR /root/
COPY --from=build /data/cnccamp/project/http_server /root/
ENTRYPOINT ["./http_server"]
