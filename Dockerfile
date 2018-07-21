FROM golang:1.10-alpine

LABEL BUILD_IMAGE=true

WORKDIR ${GOPATH}/src/github.com/psanetra/http-db

RUN apk add --no-cache --virtual git && \
    go get -d -u github.com/golang/dep && \
    cd $(go env GOPATH)/src/github.com/golang/dep && \
    export DEP_LATEST=$(git describe --abbrev=0 --tags) && \
    git checkout $DEP_LATEST && \
    go install -ldflags="-X main.version=$DEP_LATEST" ./cmd/dep

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN dep ensure && \
    go install


FROM scratch

WORKDIR /app

COPY --from=0 /go/bin/http-db /app/http-db

EXPOSE 8080

ENTRYPOINT [ "/app/http-db", "serve" ]
