FROM golang:1.16-alpine AS build
WORKDIR /data/cnccamp/project
COPY . /data/cnccamp/project
RUN go build -o /data/cnccamp/bin/httpServer

FROM scratch
COPY --from=build /data/cnccamp/bin/httpServer /data/cnccamp/bin/httpServer
ENTRYPOINT ["/data/cnccamp/bin/httpServer"]