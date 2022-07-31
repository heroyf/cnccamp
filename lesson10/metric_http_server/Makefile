APP          := metric_http_server
RELEASE      := target

build:
	gofmt -l -w -s .
	# goimports -l -w .
	go build -a -o ${RELEASE}/bin/${APP}

.PHONY: build
